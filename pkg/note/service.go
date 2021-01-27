package note

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

)

type Service struct {
	pool *pgxpool.Pool
}

type Note struct {
	ID      int32     `json:"id"`
	Name    string    `json:"name"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

func (s *Service) GetNote(ctx context.Context) (items []*Note, err error) {

	rows, err := s.pool.Query(ctx, `
		SELECT * FROM note
	`)

	for rows.Next() {
		item := &Note{}
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Content,
			&item.Created)
		if err != nil {
			log.Print(err)
		}

		items = append(items, item)
	}

	return items, nil

}
