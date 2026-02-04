package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kmaxsoul/demo-project/models"
)

func CreateTodo(pool *pgxpool.Pool, titel string, completed bool, userID string) (*models.Todo, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
        INSERT INTO todos (title, completed, user_id)
        VALUES ($1, $2, $3)
        RETURNING id, title, completed, created_at, updated_at ,user_id
    `

	var todo models.Todo
	err := pool.QueryRow(ctx, query, titel, completed, userID).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
		&todo.UserID,
	)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func GetAllTodos(pool *pgxpool.Pool, userID string) ([]models.Todo, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
        SELECT id, title, completed, created_at, updated_at, user_id
        FROM todos
		WHERE user_id = $1
		ORDER BY created_at DESC
    `

	rows, err := pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
			&todo.UserID,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return todos, nil
}

func GetTodoByID(pool *pgxpool.Pool, id int, userID string) (*models.Todo, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var query string = `
        SELECT id, title, completed, created_at, updated_at, user_id
        FROM todos
        WHERE id = $1 AND user_id = $2
    `
	var todo models.Todo
	err := pool.QueryRow(ctx, query, id, userID).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
		&todo.UserID,
	)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func UpdateTodo(pool *pgxpool.Pool, id int, title string, completed bool, userID string) (*models.Todo, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
        UPDATE todos
        SET title = $1, completed = $2, updated_at = NOW()
        WHERE id = $3 AND user_id = $4
        RETURNING id, title, completed, created_at, updated_at, user_id
    `

	var todo models.Todo
	err := pool.QueryRow(ctx, query, title, completed, id, userID).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
		&todo.UserID,
	)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func DeleteTodo(pool *pgxpool.Pool, id int, userID string) error {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
        DELETE FROM todos
        WHERE id = $1 AND user_id = $2
    `

	result, err := pool.Exec(ctx, query, id, userID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
