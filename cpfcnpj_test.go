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

func assert(t *T, expected bool, value bool) {
	if value == expected {
		t.Logf("Result should be \"%v\": ja!", expected)
	} else {
		t.Errorf("Result should be \"%v\": nein! %s", expected, value)
	}
}