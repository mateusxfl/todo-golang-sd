package todoServices

import (
	"tarefas/db"
	"tarefas/entities"
)

// Service reponsável por realizar o insert de uma entidade.
func StoreTodoService(todo entities.Todo) (id int64, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return // Será tratado no Handler.
	}
	defer connection.Close() // Fecha a conexão após o insert.

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	// Executa a query de insert, atribuindo a id o valor retornado pela query.
	err = connection.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return // Será tratado no Handler.
}
