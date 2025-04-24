package repository

import "linebot-go/domain/lineUser/entity"

type ILineUserRepository interface {
	Save(entity entity.LineUser) (*entity.LineUser, error)
	FindOne(id string) (*entity.LineUser, error)
	FindAll() ([]entity.LineUser, error)
}
