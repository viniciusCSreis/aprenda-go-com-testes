package ponteiros_e_erros

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Carteira struct {
	saldo Bitcoin
}

func (c *Carteira) Depositar(valor Bitcoin) {
	c.saldo += valor
}

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

func (c *Carteira) Retirar(valor Bitcoin) error {
	if c.saldo < valor {
		return ErroSaldoInsuficiente
	}
	c.saldo -= valor
	return nil
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}
