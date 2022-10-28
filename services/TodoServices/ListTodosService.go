package todoServices

import (
	"tarefas/db"
	"tarefas/entities"
)

// Service reponsável por realizar o read de uma entidade.
func ListTodosService() (todos []entities.Todo, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return // Será tratado no Handler.
	}
	defer connection.Close() // Fecha a conexão após o read.

	rows, err := connection.Query(`SELECT * FROM todos`)
	if err != nil {
		return // Será tratado no Handler.
	}

	for rows.Next() {
		var todo entities.Todo

		// Executa a query de read, atribuindo a entidade os valores retornados pela query.
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue // Apenas ignoramos o erro.
		}

		todos = append(todos, todo)
	}

	return // Será tratado no Handler.
}
