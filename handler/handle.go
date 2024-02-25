package handler

import (
	"database/sql"

	"github.com/leonardoksn/rinha-backend-go/config"
)

var db *sql.DB

func InitializeHandler() {
	db = config.GetMysql()
}
