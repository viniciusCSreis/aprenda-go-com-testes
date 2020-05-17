package iteracao

import (
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a")
	esperando := "aaaaa"
	if repeticoes != esperando {
		t.Errorf("esperando '%s' mas obteve '%s'", esperando, repeticoes)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a")
	}
}