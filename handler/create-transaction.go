package handler

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/leonardoksn/rinha-backend-go/schemas"
)

func CreateTransactionHandler(c *gin.Context) {

	var requestBody CreateTransactionRequestBody
	var requestParams CreateTransactionRequestParams

	if err := c.ShouldBindUri(&requestParams); err != nil {
		c.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	rows := db.QueryRow("SELECT id FROM clientes WHERE id = ?", requestParams.Id)

	var user schemas.User

	if err := rows.Scan(&user.Id); err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := requestBody.Validate(); err != nil {
		c.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	tx, err := db.Begin();

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	data, err := updateClientBalance(tx, requestParams.Id, requestBody.Valor, requestBody.Tipo)

	if err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = insertTransaction(tx, requestParams.Id, requestBody.Valor, requestBody.Descricao, requestBody.Tipo)

	if err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	tx.Commit()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, data)
}

func updateClientBalance(db *sql.Tx, clientID int, amount int, transactionType string) (schemas.ReturnLimit, error) {

	var sql string
	var binds []interface{}

	if transactionType == "c" {
		sql = `
		UPDATE clientes
		SET
			saldo = saldo + ?
		WHERE
			id = ?
		`

		binds = append(binds, amount, clientID)

	} else if transactionType == "d" {
		sql = `
		UPDATE clientes
		SET saldo = saldo - ?
		WHERE id = ?
		AND abs(saldo - ?) <= limite`

		binds = append(binds, amount, clientID, amount)

	}

	var result schemas.ReturnLimit

	// Execute the balance update in the database
	_, err := db.Exec(sql, binds...)
	if err != nil {
		return result, err
	}

	// Retrieve the new limit and balance for the client
	err = db.QueryRow("SELECT limite, saldo FROM clientes WHERE id = ?", clientID).Scan(&result.Limite, &result.Saldo)
	if err != nil {
		return result, err
	}

	return result, nil
}

func insertTransaction(db *sql.Tx, clientID int, amount int, description string, transactionType string) (error) {

	_,err := db.Exec("INSERT INTO transacoes (valor, tipo, descricao, cliente_id) VALUES (?, ?, ?, ?)", amount, transactionType, description, clientID)
	return err
}
