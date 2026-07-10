package repository

import (
	"database/sql"
)

// db *sql.DB represents the database connection pool
type URLRepository struct {
	DB *sql.DB
}

// constructor function to create a new instance of URLRepository
func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{
		DB: db,
	}
}

func (r *URLRepository) Save(longURL string) (int64, error) {
	query := `
		INSERT INTO urls (long_url)
		VALUES ($1)
		RETURNING id
	`

	var id int64

	err := r.DB.QueryRow(query, longURL).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *URLRepository) GetByID(id int64) (string, error) {
	query := `
		SELECT long_url
		FROM urls
		WHERE id = $1
	`

	var url string
	
	err := r.DB.QueryRow(query, id).Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}