package brdoc

import (
	. "testing"
)

func TestIsCPF(t *T) {
    t.Run("Invalid CPF format", func(t *T) {
        v := IsCPF("3467875434578764345789654")
        assert(t, false, v)
        
        v = IsCPF("123")
        assert(t, false, v)
        
        v = IsCPF("#$%¨&*(ABCDEF")
        assert(t, false, v)
    })
    t.Run("Invalid digits in CPF", func(t *T) {
        v := IsCPF("000.000.000-11")
        assert(t, false, v)
        
        v = IsCPF("111.111.111-22")
        assert(t, false, v)
        
        v = IsCPF("638.190.204-83")
        assert(t, false, v)
    })
    t.Run("Valid CPF", func(t *T) {
        v := IsCPF("000.000.000-00")
        assert(t, false, v)
        
        v = IsCPF("111.111.111-11")
        assert(t, false, v)
        
        v = IsCPF("638.190.204-38")
        assert(t, false, v)
    })
}

func TestIsCNPJ(t *T) {
    t.Run("Invalid CNPJ format", func(t *T) {
        v := IsCNPJ("3467875434578764345789654")
        assert(t, false, v)
        
        v = IsCNPJ("123")
        assert(t, false, v)
        
        v = IsCNPJ("#$%¨&*(ABCDEF")
        assert(t, false, v)
    })
    t.Run("Invalid digits in CNPJ", func(t *T) {
        v := IsCNPJ("00.000.000/0000-11")
        assert(t, false, v)
        
        v = IsCNPJ("11.111.111/1111-00")
        assert(t, false, v)
        
        v = IsCNPJ("28.637.456/1000-95")
        assert(t, false, v)
    })
    t.Run("Valid CNPJ", func(t *T) {
        v := IsCNPJ("00.000.000/0000-00")
        assert(t, false, v)
        
        v = IsCNPJ("11.111.111/1111-80")
        assert(t, false, v)
        
        v = IsCNPJ("28.637.456/1000-59")
        assert(t, false, v)
    })
}

func assert(t *T, expected bool, value bool) {
	if value == expected {
		t.Logf("Result should be \"%v\": ja!", expected)
	} else {
		t.Errorf("Result should be \"%v\": nein! %s", expected, value)
	}
}