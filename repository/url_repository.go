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

func (r *URLRepository) Save(longURL, shortCode string) error {
	query := `INSERT INTO urls (long_url,short_code) VALUES ($1,$2)`
	_, err := r.DB.Exec(query, longURL, shortCode)
	return err
}

func (r *URLRepository) Exists(shortcode string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(
		SELECT 1 
		FROM urls
		WHERE short_code = $1
	)`
	err := r.DB.QueryRow(query, shortcode).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
