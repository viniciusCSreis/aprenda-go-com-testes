package sync

import (
	"sync"
	"testing"
)

func TestContador(t *testing.T) {
	t.Run("Incrementar o contato 3 vezes resulta no valor 3", func(t *testing.T) {
		contador := Contador{}
		contador.Incrementa()
		contador.Incrementa()
		contador.Incrementa()

		verificaContador(t, &contador, 3)
	})
	t.Run("Roda concorrentemente em segurança", func(t *testing.T) {
		contagemEsperada := 1000
		contador := Contador{}

		var wg sync.WaitGroup
		wg.Add(contagemEsperada)

		for i := 0; i < contagemEsperada; i++ {
			go func(w *sync.WaitGroup) {
				contador.Incrementa()
				w.Done()
			}(&wg)
		}
		wg.Wait()
		verificaContador(t, &contador, contagemEsperada)
	})
}

func verificaContador(t *testing.T, resultado *Contador, esperado int) {
	t.Helper()
	if resultado.Valor() != esperado {
		t.Errorf("resultado %d, esperado %d", resultado.Valor(), esperado)
	}
}
