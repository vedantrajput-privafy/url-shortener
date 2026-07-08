package handler

import(
	"database/sql"
)

// store the database connection in Handler struct
type Handler struct {
	DB *sql.DB
}
