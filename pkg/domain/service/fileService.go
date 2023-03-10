package service

import (
	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/repository"
)

type FileService struct {
	repository repository.FileRepository
}

func NewFileService(repo repository.FileRepository) *FileService {
	return &FileService{repository: repo}
}

func (s *FileService) Exists(f domain.File) bool {
	res, err := s.repository.FindByID(f.GetID())
	if err != nil || res == nil {
		return false
	}

	return true
}
