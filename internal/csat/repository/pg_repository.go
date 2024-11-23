package repository

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/pkg/errors"
)

type CSATRepository struct {
	db *sql.DB
}

func NewCSATPGRepository(db *sql.DB) *CSATRepository {
	return &CSATRepository{db: db}
}

func (r *CSATRepository) GetStatistics(ctx context.Context) ([]*models.CSATStat, error) {
	rows, err := r.db.QueryContext(ctx, getStatistics)
	if err != nil {
		return nil, errors.Wrap(err, "GetStatistics.Query")
	}
	defer rows.Close()

	var stats []*models.CSATStat
	for rows.Next() {
		stat := &models.CSATStat{}
		err := rows.Scan(
			&stat.ID,
			&stat.Topic,
			&stat.QuestionID,
			&stat.Question,
			&stat.AverageScore,
		)
		if err != nil {
			return nil, errors.Wrap(err, "GetStatistics.Query")
		}
		stats = append(stats, stat)
	}

	return stats, nil
}
