package io

import (
	"errors"
	"github.com/De-Rune/jane/package/jane"
	"os"
	"path/filepath"
)

func GetJn(path string) (*FILE, error) {
	if filepath.Ext(path) != jane.Extension {
		return nil, errors.New(jane.Errors[`file_not_jane`] + path)
	}
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	f := new(FILE)
	f.Path = path
	f.Content = []rune(string(bytes))
	return f, nil
}