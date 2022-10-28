package entities

// Reponsável por armazenar as entidades usadas na aplicação.

type Todo struct {
	ID          int64  `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Done        bool   `json: "done"`
}
