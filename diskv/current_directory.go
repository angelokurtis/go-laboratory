package main

import (
	"os"

	"github.com/pkg/errors"
)

type CurrentDirectory string

func NewCurrentDirectory() (CurrentDirectory, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", errors.WithStack(err)
	}

	return CurrentDirectory(path), err
}
