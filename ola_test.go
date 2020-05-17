package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Vinicius")
	esperando := "Olá, Vinicius"
	if resultado != esperando {
		t.Errorf("resultado '%s', esperando '%s'", resultado, esperando)
	}
}
