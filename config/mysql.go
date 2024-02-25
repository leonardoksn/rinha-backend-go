package config

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql" // Driver MySQL
)

func InitializeMySQL() (*sql.DB, error) {
	fmt.Println("Conectando ao MySQL...")
    // String de conexão (substitua os valores)
    db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/transactions")
    if err != nil {
		return nil, err
    }
    // Verificando conexão
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return db, nil
}