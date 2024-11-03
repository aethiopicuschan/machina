package service

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/aethiopicuschan/machina/internal/service/common"
	"github.com/aethiopicuschan/machina/internal/service/python"
)

type Service interface {
	Load(r io.ReadCloser) error
	GenerateBundles() ([]common.Bundle, error)
	Close() error
}

// Return a new Service
func New(fp string) (s Service, err error) {
	// Check if source file exists
	if _, err = os.Stat(fp); errors.Is(err, os.ErrNotExist) {
		return
	}

	// Check if it is a supported language
	switch filepath.Ext(fp) {
	case ".py":
		s = &python.PythonService{}
	default:
		err = fmt.Errorf("unsupported language: %s", filepath.Ext(fp))
	}
	if err != nil {
		return
	}

	// Read the source file
	file, err := os.Open(fp)
	if err != nil {
		return
	}

	// Load the source file
	err = s.Load(file)

	return
}
