package dao

import (
	"context"
	"github.com/jackc/pgx/v5"
	"linebot-go/domain/lineUser/entity"
	"linebot-go/global"
)

type LineUserDaoPgx struct {
	db *pgx.Conn
}

func NewLineUserDaoPgx() *LineUserDaoPgx {
	return &LineUserDaoPgx{db: global.DbPgx}
}

func (r *LineUserDaoPgx) FindAll() ([]entity.LineUser, error) {
	rows, err := r.db.Query(context.Background(), "SELECT id, name, femas_token, create_time FROM public.line_user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.LineUser
	for rows.Next() {
		var user entity.LineUser
		if err := rows.Scan(&user.ID, &user.Name, &user.FemasToken, &user.CreateTime); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
