// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package persistence

import ()

type Account struct {
	ID       int64
	Username string
	Balance  string
	Version  int32
}
