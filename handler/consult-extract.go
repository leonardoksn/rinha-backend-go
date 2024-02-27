package handler

import "github.com/gin-gonic/gin"

func ConsultTransactionHandler(c *gin.Context) {

	rows, err := db.Query("SELECT * FROM clientes")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []Users
	for rows.Next() {
		var user Users
		if err := rows.Scan(&user.Id, &user.Limite, &user.Saldo, &user.Nome); err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}
