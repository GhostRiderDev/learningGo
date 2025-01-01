package dbmysql

import (
	"context"
	"database/sql"

	"github.com/ghostriderdev/movies/metadata/internal/repository"
	"github.com/ghostriderdev/movies/metadata/pkg/model"
	_ "github.com/go-sql-driver/mysql"
)

// Repository defines a MYSQL-based movie metadata repository
type Repository struct {
	db *sql.DB
}

// New creates a new instance of MYSQL based repository
func New() (*Repository, error) {
	db, err := sql.Open("mysql", "root:root@/main")
	if err != nil {
		return nil, err
	}
	return &Repository{db}, nil
}

// Get retrieves movie metadata for by movie id
func (r *Repository) Get(ctx context.Context, id string) (*model.Metadata, error) {
	var title, description, director string

	row := r.db.QueryRowContext(ctx, "SELECT title, description, director FROM Movie WHERE id = ?", id)

	if err := row.Scan(&title, &description, &director); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return &model.Metadata{
		ID:          id,
		Title:       title,
		Description: description,
		Director:    director,
	}, nil
}

// Put adds movie metadata for a given movie id.
func (r *Repository) Put(ctx context.Context, id string,
	metadata *model.Metadata) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO Movie (id, title, description, director) VALUES (?, ?, ?, ?)",
		id, metadata.Title, metadata.Description, metadata.
			Director)
	return err
}
