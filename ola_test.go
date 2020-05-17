package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Vinicius","")
		esperando := "Olá, Vinicius"
		verificaMensagemCorreta(t, resultado, esperando)
	})

	t.Run("diz 'olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("","")
		esperando := "Olá, Mundo"
		verificaMensagemCorreta(t, resultado, esperando)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "Espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})
	t.Run("em frances", func(t *testing.T) {
		resultado := Ola("Elodie", "Francês")
		esperado := "Bonjour, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

}
