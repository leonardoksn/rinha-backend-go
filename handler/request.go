package handler

import "fmt"

type CreateTransactionRequestBody struct {
	Valor     int    `json:"valor"`
	Descricao string `json:"descricao"`
	Tipo      string `json:"tipo"`
}

func (r *CreateTransactionRequestBody) Validate() error {
	if r.Valor <= 0 {
		return fmt.Errorf("o valor deve ser maior que zero")
	}

	if len(r.Descricao) < 1 || len(r.Descricao) > 10 {
		return fmt.Errorf("a descrição deve ter entre 1 e 10 caracteres")
	}

	if r.Tipo != "c" && r.Tipo != "d" {
		return fmt.Errorf("o tipo deve ser 'c' para crédito ou 'd' para débito")
	}
	
	return nil
}

type CreateTransactionRequestParams struct {
	Id int `uri:"id"`
}

