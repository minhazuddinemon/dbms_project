package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
	RegisterStudentTx(ctx context.Context, stParams CreateStudentParams, mobiles []string) (Student, error)
}

type SQLStore struct {
	*Queries
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) Store {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}

func (st *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := st.db.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	return tx.Commit(ctx)
}

func (store *SQLStore) RegisterStudentTx(ctx context.Context, stParams CreateStudentParams, mobiles []string) (Student, error) {
	var result Student
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result, err = q.CreateStudent(ctx, stParams)
		if err != nil {
			return err
		}

		for _, mobile := range mobiles {
			err = q.AddStudentMobile(ctx, AddStudentMobileParams{
				StudentID: result.StudentID,
				Mobile:    mobile,
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	return result, err
}
