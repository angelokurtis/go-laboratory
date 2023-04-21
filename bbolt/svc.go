package main

import bolt "go.etcd.io/bbolt"

type Service struct {
	BoltDB *bolt.DB
}
