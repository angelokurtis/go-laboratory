package main

import (
	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
)

func NewBoltDB(directory CurrentDirectory) (*bolt.DB, error) {
	db, err := bolt.Open(string(directory)+"/demo.db", 0o666, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, nil
}
