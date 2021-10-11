package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/sin392/todo-back/model"
)

type TODOService struct {
	db *sql.DB
}

func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

func (s *TODOService) CreateTODO(ctx context.Context, opts *model.CreateTODORequest) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)
	var todo model.TODO

	res, err := s.db.ExecContext(ctx, insert, opts.Subject, opts.Description)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	todo.ID = int(id)

	err = s.db.QueryRowContext(ctx, confirm, id).Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &todo, err
}

// ReadTODO reads TODO on DB.
func (s *TODOService) ReadTODO(ctx context.Context, opts *model.ReadTODORequest) (*model.TODO, error) {
	const (
		read = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)
	todo := model.TODO{}
	err := s.db.QueryRowContext(ctx, read, opts.ID).Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.CreatedAt)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &todo, err
}

// ReadTODOs reads TODOs on DB.
func (s *TODOService) ReadTODOs(ctx context.Context, opts *model.ReadTODOsRequest) (*[]model.TODO, error) {
	read := `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?, ?`
	var todos []model.TODO

	skip, limit := "0", "1000000"
	// 引数用の構造体をポインタ型にしているがゼロ値が代入されてしまっている？
	// if opts.Skip != nil {
	if *opts.Skip != "" {
		skip = *opts.Skip
	}
	// if opts.Limit != nil {
	if *opts.Limit != "" {
		limit = *opts.Limit
	}

	rows, err := s.db.QueryContext(ctx, read, skip, limit)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	todo := model.TODO{}
	for rows.Next() {
		err := rows.Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.CreatedAt)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		todos = append(todos, todo)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &todos, err
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, opts *model.UpdateTODORequest) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)
	var todo model.TODO

	_, err := s.db.ExecContext(ctx, update, opts.Subject, opts.Description, opts.ID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = s.db.QueryRowContext(ctx, confirm, opts.ID).Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &todo, err
}

// DeleteTODO deletes TODO on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, opts *model.DeleteTODORequest) error {
	const deleteFmt = `DELETE FROM todos WHERE id = ?`
	_, err := s.db.ExecContext(ctx, deleteFmt, opts.ID)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

// DeleteTODOs deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODOs(ctx context.Context, opts *model.DeleteTODOsRequest) error {
	// ExecContextの引数としてidsをわたして埋め込みをするとTruncated incorrect INTEGER valueが発生する？
	deleteFmt := fmt.Sprintf(`DELETE FROM todos WHERE id IN (%s)`, strings.Join(opts.IDs, ","))
	_, err := s.db.ExecContext(ctx, deleteFmt)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
