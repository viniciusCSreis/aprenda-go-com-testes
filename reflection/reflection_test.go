package reflection

import (
	"reflect"
	"testing"
)

type Pessoa struct {
	Nome   string
	Perfil Perfil
}

type Perfil struct {
	Idade  int
	Cidade string
}

func TestPercorre(t *testing.T) {

	casos := []struct {
		Nome              string
		Entrada           interface{}
		ChamadasEsperadas []string
	}{
		{
			"Struct com um campo string",
			struct {
				Nome string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct com dois campos tipo string",
			struct {
				Nome   string
				Cidade string
			}{"Chris", "Londres"},
			[]string{"Chris", "Londres"},
		},
		{
			"Struct sem campo tipo string",
			struct {
				Nome  string
				Idade int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Campos aninhados",
			Pessoa{
				"Chris",
				Perfil{33, "Londres"},
			},
			[]string{"Chris", "Londres"},
		},
		{
			"Ponteiros para coisas",
			&Pessoa{
				"Chris",
				Perfil{33, "Londres"},
			},
			[]string{"Chris", "Londres"},
		},
		{
			"Slices",
			[]Perfil{
				{33, "Londres"},
				{34, "Reykjavík"},
			},
			[]string{"Londres", "Reykjavík"},
		},
		{
			"Arrays",
			[2]Perfil{
				{33, "Londres"},
				{34, "Reykjavík"},
			},
			[]string{"Londres", "Reykjavík"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, teste := range casos {
		t.Run(teste.Nome, func(t *testing.T) {
			var resultado []string
			percorre(teste.Entrada, func(entrada string) {
				resultado = append(resultado, entrada)
			})

			if !reflect.DeepEqual(resultado, teste.ChamadasEsperadas) {
				t.Errorf("resultado %v, esperado %v", resultado, teste.ChamadasEsperadas)
			}
		})
	}
}
