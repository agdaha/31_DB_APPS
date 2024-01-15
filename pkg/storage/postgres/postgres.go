package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

type Task struct {
	ID         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

func (s *Storage) NewTask(t Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO tasks (author_id, assigned_id, title, content)
		VALUES ($1, $2, $3, $4) RETURNING id;
		`,
		t.Title,
		t.Content,
		t.AuthorID,
		t.AssignedID,
	).Scan(&id)
	return id, err
}

func (s *Storage) Tasks() ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT
			id,
			opened,
			closed,
			author_id,
			assigned_id,
			title,
			content
		FROM tasks
		ORDER BY id;
	`)
	return query(err, rows)
}

func query(err error, rows pgx.Rows) ([]Task, error) {
	if err != nil {
		return nil, err
	}
	var tasks []Task

	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

func (s *Storage) TasksWithAuthor(authorId int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
			SELECT
				id,
				opened,
				closed,
				author_id,
				assigned_id,
				title,
				content
			FROM tasks
			WHERE
				($1 = 0 OR author_id = $1)
			ORDER BY id;
		`,
		authorId,
	)
	return query(err, rows)
}

func (s *Storage) TasksWithLabel(taskId int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
			SELECT
				id,
				opened,
				closed,
				author_id,
				assigned_id,
				title,
				content
			FROM tasks
			WHERE
				id in (SELECT task_id FROM public.task_labels WHERE label_id=$1)
			ORDER BY id;
		`,
		taskId,
	)
	return query(err, rows)
}

func (s *Storage) UpdateTask(id int, t Task) error {
	_, err := s.db.Exec(context.Background(), `
	UPDATE public.tasks
	SET closed=$1, author_id=$2, assigned_id=$3, title=$4, content=$5
	WHERE id=$6;
	`,
		t.Closed,
		t.AuthorID,
		t.AssignedID,
		t.Title,
		t.Content,
		id,
	)
	return err
}

func (s *Storage) DeleteTask(id int) error {
	_, err := s.db.Exec(context.Background(), `
	DELETE FROM public.task_labels
	WHERE task_id=$1;
	`,
		id,
	)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(context.Background(), `
	DELETE FROM public.tasks
	WHERE id=1;
	`,
		id,
	)
	return err
}

func (s *Storage) Close() {
	s.db.Close()
}
