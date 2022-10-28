package todoServices

import (
	"tarefas/db"
	"tarefas/entities"
)

// Service reponsável por realizar o read de uma entidade.
func ListSingleTodoService(id int64) (todo entities.Todo, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return // Será tratado no Handler.
	}
	defer connection.Close() // Fecha a conexão após o read.

	row := connection.QueryRow(`SELECT * FROM todos WHERE id = $1`, id)

	// Executa a query de read, atribuindo a todo os valores retornados pela query.
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return // Será tratado no Handler.
}
