// Reponsável por estabelecer a conexão com banco de dados.

package db

import (
	"database/sql"
	"fmt"
	"tarefas/configs"

	_ "github.com/lib/pq" // Precisamos do _ antes para informar que mesmo não ultilizando o import, precisamos dele.
)

// Abre uma conexão com banco de dados, para nos services usarmos ela.
func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	connection, err := sql.Open("postgres", stringConnection)
	if err != nil {
		panic(err) // Encerra a aplicação.
	}

	// Realiza um ping no banco, para realmente ter certeza que a conexão foi estabelecida.
	err = connection.Ping()

	return connection, err
}
