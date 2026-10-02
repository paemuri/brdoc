package brdoc

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const pixGUI = "br.gov.bcb.pix"

var (
	pixMCCRegexp  = regexp.MustCompile(`^\d{4}$`)
	pixTxIDRegexp = regexp.MustCompile(`^([0-9A-Za-z]{1,25}|\*\*\*)$`)
	pixISPBRegexp = regexp.MustCompile(`^[0-9A-Za-z]{8}$`)
	// Not official logic: the format of the amount, which is only shown by
	// examples.
	pixAmountRegexp = regexp.MustCompile(`^\d{1,10}(\.\d{1,2})?$`)
)

// emvObject is a data object of an EMV QR Code: an ID and its value.
type emvObject struct {
	id    int
	value string
}

// IsPixBRCode verifies if the given string is a valid BR Code of Pix, the
// content of its QR Code, also used as "Pix copia e cola". It can be static
// (with a Pix key), dynamic (with a URL) or composite (with a URL of the
// recurrence parameters).
//
// The structure, the mandatory fields and the CRC are defined by the
// [Manual do BR Code], which follows the EMV QR Code standard, and the fields
// of Pix by section 2 of the [Manual de Padrões para Iniciação do Pix]. The Pix
// key is validated by [IsPixKey].
//
// [Manual do BR Code]: https://www.bcb.gov.br/content/estabilidadefinanceira/spb_docs/ManualBRCode.pdf
// [Manual de Padrões para Iniciação do Pix]: https://www.bcb.gov.br/content/estabilidadefinanceira/pix/Regulamento_Pix/II_ManualdePadroesparaIniciacaodoPix.pdf
func IsPixBRCode(doc string) bool {
	objects, ok := parseEMVObjects(doc)
	if !ok || len(objects) < 2 {
		return false
	}

	// The payload format indicator is the first object, and the CRC the last
	// one, calculated over everything before its value.
	first, last := objects[0], objects[len(objects)-1]
	if first.id != 0 ||
		first.value != "01" ||
		last.id != 63 ||
		len(last.value) != 4 ||
		!strings.EqualFold(last.value, calcEMVCRC(doc[:len(doc)-4])) {
		return false
	}

	root := map[int]string{}
	for _, o := range objects {
		root[o.id] = o.value
	}

	if v, ok := root[1]; ok && v != "11" && v != "12" {
		return false
	}
	if v, ok := root[54]; ok && !pixAmountRegexp.MatchString(v) {
		return false
	}
	if !pixMCCRegexp.MatchString(root[52]) || root[53] != "986" ||
		root[58] != "BR" || !validEMVLength(root[59], 25) ||
		!validEMVLength(root[60], 15) {
		return false
	}

	additional, ok := parseEMVTemplate(root[62])
	if !ok || !pixTxIDRegexp.MatchString(additional[5]) {
		return false
	}

	account, recurrence, ok := findPixTemplates(objects)
	if !ok {
		return false
	}

	return validPixAccount(account, recurrence)
}

// findPixTemplates returns the merchant account information (IDs 26 to 51) and
// the unreserved template (IDs 80 to 99) of Pix, the latter only if present.
// All the templates of these ranges must have a GUI.
func findPixTemplates(objects []emvObject) (account, recurrence map[int]string, ok bool) {
	for _, o := range objects {
		if (o.id < 26 || o.id > 51) && o.id < 80 {
			continue
		}

		template, ok := parseEMVTemplate(o.value)
		if !ok || template[0] == "" {
			return nil, nil, false
		}
		if !strings.EqualFold(template[0], pixGUI) {
			continue
		}

		if o.id <= 51 && account == nil {
			account = template
		} else if o.id >= 80 && recurrence == nil {
			recurrence = template
		}
	}

	return account, recurrence, account != nil
}

// validPixAccount verifies the merchant account information of Pix. A static BR
// Code has a Pix key, a dynamic one has a URL, and a composite one may have
// none of them, as long as it has the URL of the recurrence parameters.
func validPixAccount(account, recurrence map[int]string) bool {
	key, hasKey := account[1]
	url, hasURL := account[25]

	switch {
	case hasKey && hasURL:
		return false
	case hasKey && !IsPixKey(key):
		return false
	case hasURL && !validPixURL(url):
		return false
	}

	if v, ok := account[2]; ok && !validEMVLength(v, 72) {
		return false
	}
	if v, ok := account[3]; ok && !pixISPBRegexp.MatchString(v) {
		return false
	}

	if recurrence != nil {
		return validPixURL(recurrence[25])
	}

	return hasKey || hasURL
}

// validPixURL verifies a URL of a BR Code, which has up to 77 characters and no
// protocol prefix.
func validPixURL(url string) bool {
	// Not official logic: the URL is not fully validated.
	return validEMVLength(url, 77) && !strings.Contains(url, "://")
}

// validEMVLength verifies if `value` has from 1 to `limit` characters.
func validEMVLength(value string, limit int) bool {
	n := len([]rune(value))
	return n >= 1 && n <= limit
}

// parseEMVTemplate returns the data objects of a template by their IDs, or
// false if it is malformed or has repeated IDs.
func parseEMVTemplate(value string) (map[int]string, bool) {
	objects, ok := parseEMVObjects(value)
	if !ok {
		return nil, false
	}

	template := map[int]string{}
	for _, o := range objects {
		template[o.id] = o.value
	}

	return template, true
}

// parseEMVObjects splits `doc` in its data objects, each one with an ID of 2
// digits, a length of 2 digits and a value of that many characters. It returns
// false if `doc` is malformed or has repeated IDs.
func parseEMVObjects(doc string) ([]emvObject, bool) {
	r := []rune(doc)
	seen := map[int]bool{}

	var objects []emvObject
	for i := 0; i < len(r); {
		if i+4 > len(r) || !allDigit(string(r[i:i+4])) {
			return nil, false
		}

		id, _ := strconv.Atoi(string(r[i : i+2]))
		size, _ := strconv.Atoi(string(r[i+2 : i+4]))
		i += 4
		if size == 0 || i+size > len(r) || seen[id] {
			return nil, false
		}

		seen[id] = true
		objects = append(objects, emvObject{id, string(r[i : i+size])})
		i += size
	}

	return objects, len(objects) > 0
}

// calcEMVCRC returns the CRC-16/CCITT-FALSE (polynomial 0x1021, initial value
// 0xFFFF) of `doc`, as 4 hexadecimal digits.
func calcEMVCRC(doc string) string {
	crc := 0xFFFF
	for i := 0; i < len(doc); i++ {
		crc ^= int(doc[i]) << 8
		for j := 0; j < 8; j++ {
			if crc&0x8000 != 0 {
				crc = crc<<1 ^ 0x1021
			} else {
				crc <<= 1
			}
		}
		crc &= 0xFFFF
	}

	return fmt.Sprintf("%04X", crc)
}
