package handler

type Users struct {
	Id int `json:"id"`
	Limite int  `json:"limite"`
	Saldo int `json:"saldo"`
	Nome string `json:"nome"`
}