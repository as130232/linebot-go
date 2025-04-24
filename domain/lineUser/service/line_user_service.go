package service

import (
	"linebot-go/domain/lineUser/entity"
	"linebot-go/domain/lineUser/repository"
)

type LineUserService struct {
	lineUserRepo repository.ILineUserRepository
}

func NewLineUserService(lineUserRepo repository.ILineUserRepository) *LineUserService {
	return &LineUserService{lineUserRepo: lineUserRepo}
}

func (s *LineUserService) GetById(id string) (*entity.LineUser, error) {
	return s.lineUserRepo.FindOne(id)
}

func (s *LineUserService) Save(lineUser entity.LineUser) (*entity.LineUser, error) {
	return s.lineUserRepo.Save(lineUser)
}
