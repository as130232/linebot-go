package dao

import (
	"gorm.io/gorm"
	"linebot-go/domain/lineUser/entity"
	"linebot-go/global"
	"linebot-go/infrastructure/consts/table"
)

type LineUserDaoGorm struct {
	db *gorm.DB
}

func NewLineUserDaoGorm() *LineUserDaoGorm {
	return &LineUserDaoGorm{db: global.DbGorm}
}

func (r *LineUserDaoGorm) Save(entity entity.LineUser) (*entity.LineUser, error) {
	err := r.db.Table(table.LineUser).Create(&entity).Error
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *LineUserDaoGorm) FindOne(id string) (*entity.LineUser, error) {
	var result entity.LineUser
	err := r.db.Table(table.LineUser).Take(&result, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *LineUserDaoGorm) FindAll() ([]entity.LineUser, error) {
	result := make([]entity.LineUser, 0)
	err := r.db.Table(table.LineUser).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
