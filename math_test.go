package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)
	if total != 30 {
		t.Errorf("Esperado 30, resultado %d", total)
	}
}
