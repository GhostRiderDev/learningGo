package dbmysql

import (
	"context"
	"database/sql"

	"github.com/ghostriderdev/movies/rating/internal/repository"
	model "github.com/ghostriderdev/movies/rating/pkg"
	_ "github.com/go-sql-driver/mysql"
)

// Repository defines MYSQL based rating repository
type Repository struct {
	db *sql.DB
}

// New creates a new MySQL-based rating repository.
func New() (*Repository, error) {
	db, err := sql.Open("mysql", "root:root@/main")
	if err != nil {
		return nil, err
	}
	return &Repository{db}, nil
}

// Get retrieves all ratings for a given record.
func (r *Repository) Get(ctx context.Context, id model.RecordID, recordType model.RecordType) (*[]model.Rating, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT user_id, value FROM Rating WHERE record_id = ? AND record_type = ?", id, recordType)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var res []model.Rating

	for rows.Next() {
		var userID string
		var value int32

		if err := rows.Scan(&userID, value); err != nil {
			return nil, err
		}

		res = append(res, model.Rating{
			UserID: model.UserID(userID),
			Value:  model.RatingValue(value),
		})
	}

	if len(res) == 0 {
		return nil, repository.ErrNotFound
	}

	return &res, nil
}

// Put adds a rating for a given record.
func (r *Repository) Put(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO Rating (record_id, record_type, user_id, value) VALUES (?, ?, ?, ?)", recordID, recordType, rating.UserID, rating.Value)
	return err
}
