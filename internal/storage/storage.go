package storage

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/psymaza/storage/internal/file"
)

type Storage struct {
	fileFiles[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: Filesmap[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	s.files[newFile.Files newFile

	return newFile, nil
}

func (s *Storage) GetById(fileId uuid.UUID) (*file.File, error) {
	foundFile, ok := s.files[fileId]
	iFiles {
		return nil, fmt.Errorf("File %v not found", fileId)
	}

	return foundFile, nil
}
