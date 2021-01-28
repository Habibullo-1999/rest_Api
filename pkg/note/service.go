package note

import (
	"context"
	"errors"
	"log"

	"github.com/Habibullo-1999/rest_Api/pkg/types"
	"github.com/jackc/pgx/v4/pgxpool"

)

type Service struct {
	pool *pgxpool.Pool
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}
//Get All Notes
func (s *Service) GetNotes(ctx context.Context) (items []*types.Note, err error) {

	rows, err := s.pool.Query(ctx, `
		SELECT * FROM note
	`)

	for rows.Next() {
		item := &types.Note{}
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
//Get By ID
func (s *Service) GetById(ctx context.Context, id int64) (item *types.Note, err error) {

	items := &types.Note{}
	sqlQuery := `Select * From note Where id=$1`
	err = s.pool.QueryRow(ctx, sqlQuery, id).Scan(
		&items.ID,
		&items.Name,
		&items.Content,
		&items.Created)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return items, nil
}

//Save 
func (s *Service) SaveNote(ctx context.Context, itemOld *types.Note) (newNote *types.Note, err error) {
	item := &types.Note{}
	if itemOld.ID == 0 {
		sqlQuery := `Insert into note(name,content) values ($1,$2) returning *`
		err = s.pool.QueryRow(ctx, sqlQuery, itemOld.Name, itemOld.Content).Scan(
			&item.ID,
			&item.Name,
			&item.Content,
			&item.Created)
	} else {
		log.Print(err)
		return nil, errors.New("internal error")
	}
	if err != nil {
		log.Print(err)
		return nil, errors.New("internal error")
	}
	return item, nil
}
//Update 
func (s *Service) UpdateNote(ctx context.Context, itemOld *types.Note) (newNote *types.Note, err error) {
	item := &types.Note{}

		sqlQuery := `UPDATE note Set name=$2, content=$3 WHERE id=$1 returning * `
		err = s.pool.QueryRow(ctx, sqlQuery, itemOld.ID, itemOld.Name, itemOld.Content).Scan(
			&item.ID,
			&item.Name,
			&item.Content,
			&item.Created)
	if err != nil {
		log.Print(err)
		return nil, errors.New("internal error")
	}		

	
	return item, nil
}
