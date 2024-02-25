package main

import (
	"fmt"

	"github.com/leonardoksn/rinha-backend-go/config"
	"github.com/leonardoksn/rinha-backend-go/router"
)

func main() {
	err := config.Init()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	
	router.InitRouter()
}
