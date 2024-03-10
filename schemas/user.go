package schemas

type User struct {
	Id int
}

type ReturnLimit struct {
	Limite int `json:"limite"`
	Saldo  int `json:"saldo"`
}