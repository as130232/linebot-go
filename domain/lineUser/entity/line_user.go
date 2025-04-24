package entity

import "time"

type LineUser struct {
	ID         string    `db:"id"`
	Name       string    `db:"name"`
	FemasToken string    `db:"femas_token"`
	IsValid    int       `db:"is_valid"`
	CreateTime time.Time `db:"create_time"`
}
