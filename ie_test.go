package brdoc

import (
	"testing"
)

func TestIsIE(t *testing.T) {
	for i, tc := range []struct {
		name  string
		doc   string
		uf    UF
		valid bool
	}{
		{"NotImplementedUF", "030123459", AP, false},

		{"AC_InvalidData", "", AC, false},
		{"AC_InvalidData", "AAAAAAAAAAAAA", AC, false},

		{"AC_InvalidFormat", "010048230011", AC, false},
		{"AC_InvalidFormat", "01004823001123", AC, false},
		{"AC_InvalidFormat", "02.004.823/001-12", AC, false},
		{"AC_InvalidFormat", "01-004-823.001/12", AC, false},

		{"AC_InvalidDigit", "01.004.823/001-02", AC, false},
		{"AC_InvalidDigit", "01.004.823/001-13", AC, false},

		// Example from the official rules.
		{"AC_Valid", "01.004.823/001-12", AC, true},
		{"AC_Valid", "0100482300112", AC, true},

		// Public IEs of companies.
		{"AC_Valid", "01.008.267/001-08", AC, true},

		{"AL_InvalidData", "", AL, false},
		{"AL_InvalidData", "AAAAAAAAA", AL, false},

		{"AL_InvalidFormat", "24000004", AL, false},
		{"AL_InvalidFormat", "2400000480", AL, false},
		{"AL_InvalidFormat", "250000048", AL, false},

		{"AL_InvalidDigit", "240000049", AL, false},

		// Example from the official rules.
		{"AL_Valid", "240000048", AL, true},

		// Public IEs of companies.
		{"AL_Valid", "245008403", AL, true},

		// Types of company not listed by older rules.
		{"AL_Valid", "241000009", AL, true},

		{"AM_InvalidData", "", AM, false},
		{"AM_InvalidFormat", "04.900.976", AM, false},
		{"AM_InvalidDigit", "04.900.976-2", AM, false},
		// Examples from the official rules.
		{"AM_Valid", "04.900.976-1", AM, true},
		{"AM_Valid", "04.150.272-8", AM, true},

		{"CE_InvalidData", "", CE, false},
		{"CE_InvalidFormat", "0600001-5", CE, false},
		{"CE_InvalidDigit", "06000001-4", CE, false},
		// Examples from the official rules.
		{"CE_Valid", "06000001-5", CE, true},
		// Public IEs of companies.
		{"CE_Valid", "06863259-2", CE, true},

		{"DF_InvalidData", "", DF, false},
		{"DF_InvalidFormat", "0730000100109-", DF, false},
		{"DF_InvalidDigit", "07300001001-08", DF, false},
		// Examples from the official rules.
		{"DF_Valid", "07300001001-09", DF, true},
		// Public IEs of companies.
		{"DF_Valid", "07656443030-76", DF, true},

		{"ES_InvalidData", "", ES, false},
		{"ES_InvalidFormat", "99999999", ES, false},
		{"ES_InvalidDigit", "999999991", ES, false},
		// Examples from the official rules.
		{"ES_Valid", "999999990", ES, true},
		// Public IEs of companies.
		{"ES_Valid", "000002976", ES, true},

		{"GO_InvalidData", "", GO, false},
		{"GO_InvalidFormat", "12.987.654-7", GO, false},
		{"GO_InvalidDigit", "10.987.654-8", GO, false},
		// Examples from the official rules.
		{"GO_Valid", "10.987.654-7", GO, true},
		// Public IEs of companies.
		{"GO_Valid", "10.277.380-7", GO, true},

		{"MA_InvalidData", "", MA, false},
		{"MA_InvalidFormat", "130000385", MA, false},
		{"MA_InvalidDigit", "120000386", MA, false},
		// Examples from the official rules.
		{"MA_Valid", "120000385", MA, true},
		// Public IEs of companies.
		{"MA_Valid", "126001049", MA, true},

		{"MS_InvalidData", "", MS, false},
		{"MS_InvalidFormat", "292909575", MS, false},
		{"MS_InvalidDigit", "282909576", MS, false},
		// Public IEs of companies.
		{"MS_Valid", "282909575", MS, true},
		{"MS_Valid", "283242809", MS, true},

		{"MT_InvalidData", "", MT, false},
		{"MT_InvalidFormat", "013000001-9", MT, false},
		{"MT_InvalidDigit", "0013000001-8", MT, false},
		// Examples from the official rules.
		{"MT_Valid", "0013000001-9", MT, true},
		// Public IEs of companies.
		{"MT_Valid", "0013144158-2", MT, true},

		{"PA_InvalidData", "", PA, false},
		{"PA_InvalidFormat", "16999999-5", PA, false},
		{"PA_InvalidDigit", "15999999-4", PA, false},
		// Examples from the official rules.
		{"PA_Valid", "15999999-5", PA, true},
		{"PA_Valid", "75000002-3", PA, true},
		// Public IEs of companies.
		{"PA_Valid", "15177432-3", PA, true},
		{"PA_Valid", "15771637-6", PA, true},

		{"PB_InvalidData", "", PB, false},
		{"PB_InvalidFormat", "0600001-5", PB, false},
		{"PB_InvalidDigit", "06000001-6", PB, false},
		// Examples from the official rules.
		{"PB_Valid", "06000001-5", PB, true},
		// Public IEs of companies.
		{"PB_Valid", "16999351-5", PB, true},

		{"PI_InvalidData", "", PI, false},
		{"PI_InvalidFormat", "01234567", PI, false},
		{"PI_InvalidDigit", "012345678", PI, false},
		// Examples from the official rules.
		{"PI_Valid", "012345679", PI, true},
		// Public IEs of companies.
		{"PI_Valid", "194155722", PI, true},
		{"PI_Valid", "195195337", PI, true},

		{"RJ_InvalidData", "", RJ, false},
		{"RJ_InvalidFormat", "91.018.0-39", RJ, false},
		{"RJ_InvalidDigit", "91.018.03-8", RJ, false},
		// Public IEs of companies.
		{"RJ_Valid", "91.018.03-9", RJ, true},

		{"RN_InvalidData", "", RN, false},
		{"RN_InvalidFormat", "21.040.040-1", RN, false},
		{"RN_InvalidDigit", "20.040.040-2", RN, false},
		// Examples from the official rules.
		{"RN_Valid", "20.040.040-1", RN, true},
		{"RN_Valid", "20.0.040.040-0", RN, true},
		// Public IEs of companies.
		{"RN_Valid", "20.300.936-3", RN, true},

		{"RR_InvalidData", "", RR, false},
		{"RR_InvalidFormat", "25006628-1", RR, false},
		{"RR_InvalidDigit", "24006628-2", RR, false},
		// Examples from the official rules.
		{"RR_Valid", "24006628-1", RR, true},
		{"RR_Valid", "24001755-6", RR, true},
		{"RR_Valid", "24003429-0", RR, true},
		// Public IEs of companies.
		{"RR_Valid", "24002153-4", RR, true},

		{"RS_InvalidData", "", RS, false},
		{"RS_InvalidFormat", "224/365879", RS, false},
		{"RS_InvalidDigit", "224/3658793", RS, false},
		// Examples from the official rules.
		{"RS_Valid", "224/3658792", RS, true},
		// Public IEs of companies.
		{"RS_Valid", "900/0000802", RS, true},

		{"SC_InvalidData", "", SC, false},
		{"SC_InvalidFormat", "251.040.85", SC, false},
		{"SC_InvalidDigit", "251.040.853", SC, false},
		// Examples from the official rules.
		{"SC_Valid", "251.040.852", SC, true},
		// Public IEs of companies.
		{"SC_Valid", "252.085.442", SC, true},

		{"SE_InvalidData", "", SE, false},
		{"SE_InvalidFormat", "2712345-3", SE, false},
		{"SE_InvalidDigit", "27123456-4", SE, false},
		// Examples from the official rules.
		{"SE_Valid", "27123456-3", SE, true},
		// Public IEs of companies.
		{"SE_Valid", "27077995-7", SE, true},

		{"PR_InvalidData", "", PR, false},
		{"PR_InvalidData", "AAAAAAAAAA", PR, false},

		{"PR_InvalidFormat", "123456785", PR, false},
		{"PR_InvalidFormat", "12345678501", PR, false},
		{"PR_InvalidFormat", "123.456.785-0", PR, false},
		{"PR_InvalidFormat", "12.345678-50", PR, false},

		{"PR_InvalidDigit", "123.45678-40", PR, false},
		{"PR_InvalidDigit", "123.45678-51", PR, false},

		// Example from the official rules.
		{"PR_Valid", "123.45678-50", PR, true},
		{"PR_Valid", "12345678-50", PR, true},
		{"PR_Valid", "1234567850", PR, true},

		// Public IEs of companies.
		{"PR_Valid", "099.02241-88", PR, true},
		{"PR_Valid", "101.79579-92", PR, true},

		{"SP_InvalidData", "", SP, false},
		{"SP_InvalidData", "AAAAAAAAAAAA", SP, false},

		{"SP_InvalidFormat", "11004249011", SP, false},
		{"SP_InvalidFormat", "1100424901145", SP, false},
		{"SP_InvalidFormat", "110-042-490-114", SP, false},
		{"SP_InvalidFormat", "p-01100424.3/002", SP, false},
		{"SP_InvalidFormat", "P-0110042.43/002", SP, false},
		{"SP_InvalidFormat", "P-01100424.3/0021", SP, false},

		{"SP_InvalidDigit", "110.042.491.114", SP, false},
		{"SP_InvalidDigit", "110.042.490.115", SP, false},
		{"SP_InvalidDigit", "P-01100424.4/002", SP, false},

		// Examples from the official rules.
		{"SP_Valid", "110.042.490.114", SP, true},
		{"SP_Valid", "110042490114", SP, true},
		{"SP_Valid", "P-01100424.3/002", SP, true},
		{"SP_Valid", "P011004243002", SP, true},
		{"SP_Valid", "P-01100424.3/999", SP, true},

		// Public IEs of companies.
		{"SP_Valid", "310.035.324.119", SP, true},
		{"SP_Valid", "108.354.656.114", SP, true},
		{"SP_Valid", "142.270.790.110", SP, true},
		{"SP_Valid", "142.484.958.110", SP, true},
		{"SP_Valid", "102.654.009.110", SP, true},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			assertEq(t, tc.valid, IsIE(tc.doc, tc.uf))
		})
	}
}
