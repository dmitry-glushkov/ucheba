package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// Investement ...
type Investement struct {
	UID int
	PID int
	Val int
}

// Save ...
func (d *Investement) Save(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(
		ctx,
		`
			INSERT INTO investements
				(uid, pid, val)
				VALUES ($1, $2, $3);
		`,
		d.UID, d.PID, d.Val,
	)
	if err != nil {
		err = fmt.Errorf("...: %w", err)
		return err
	}

	return nil
}

func GetProjectInvestements(ctx context.Context, db *pgx.Conn, pid int) ([]Investement, error) {
	rows, err := db.Query(
		ctx,
		`
		SELECT uid, pid, val
			FROM goals
			WHERE pid = $1;	
		`,
		pid,
	)
	if err != nil {
		err = fmt.Errorf("...: %w", err)
		return nil, err
	}
	defer rows.Close()

	var donates []Investement
	for rows.Next() {
		var donate Investement
		err = rows.Scan()
		if err != nil {
			err = fmt.Errorf("...: %w", err)
			return nil, err
		}

		donates = append(donates, donate)
	}

	return donates, nil
}

func GetUserInvestements(ctx context.Context, db *pgx.Conn, uid int) ([]Investement, error) {
	rows, err := db.Query(
		ctx,
		`
		SELECT uid, pid, val
			FROM goals
			WHERE uid = $1;	
		`,
		uid,
	)
	if err != nil {
		err = fmt.Errorf("...: %w", err)
		return nil, err
	}
	defer rows.Close()

	var donates []Investement
	for rows.Next() {
		var donate Investement
		err = rows.Scan()
		if err != nil {
			err = fmt.Errorf("...: %w", err)
			return nil, err
		}

		donates = append(donates, donate)
	}

	return donates, nil
}