package iteracao

func Repetir(texto string) string{
	resultado := ""
	for i := 0 ; i < 5 ; i++ {
		resultado += texto
	}
	return resultado
}
