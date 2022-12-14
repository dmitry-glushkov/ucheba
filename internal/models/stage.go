package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
)

// Stage ...
type Stage struct {
	ID      int    `json:"id"`
	PID     int    `json:"pid"`
	Target  int    `json:"target"`
	DueDate string `json:"due_date"`
}

// Save ...
func (g *Stage) Save(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(
		ctx,
		`
			INSERT INTO stages
				(pid, target, due_date)
				VALUES ($1, $2, $3);	
		`,
		g.PID, g.Target, g.DueDate,
	)
	return err
}

func (g *Stage) SaveMock(ctx context.Context, db *pgx.Conn) error {
	return nil
}

// GetStages ...
func GetStages(ctx context.Context, db *pgx.Conn, pid int) ([]Stage, error) {
	rows, err := db.Query(
		ctx,
		`
		SELECT pid, target, due_date
			FROM stages
			WHERE pid = $1
			order by id desc;
		`,
		pid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stages []Stage
	for rows.Next() {
		var stage Stage
		err = rows.Scan(
			&stage.PID,
			&stage.Target,
			&stage.DueDate,
		)
		if err != nil {
			return nil, err
		}

		stages = append(stages, stage)
	}

	return stages, nil
}

func GetStagesMock(ctx context.Context, db *pgx.Conn, pid int) ([]Stage, error) {
	return []Stage{
		{
			ID:      0,
			Target:  50000,
			DueDate: time.Now().AddDate(0, 0, -7).Format("02.01.2006"),
		},
		{
			ID:      1,
			Target:  150000,
			DueDate: time.Now().AddDate(0, 0, 7).Format("02.01.2006"),
		},
		{
			ID:      2,
			Target:  400000,
			DueDate: time.Now().AddDate(0, 1, 0).Format("02.01.2006"),
		},
	}, nil
}
