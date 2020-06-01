package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SleeperConfiguravel struct {
	duracao time.Duration
	sleep   func(time.Duration)
}

func (s *SleeperConfiguravel) Sleep() {
	s.sleep(s.duracao)
}

func main()  {
	sleeper := &SleeperConfiguravel{1 * time.Second, time.Sleep}
	Contagem(os.Stdout, sleeper)
}

const ultimaPalavra = "VAI!"
const inicioContagem = 3

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem ; i > 0 ; i-- {
		sleeper.Sleep()
		fmt.Fprintf(saida, "%d\n", i)
	}
	sleeper.Sleep()
	fmt.Fprint(saida,ultimaPalavra)

}
