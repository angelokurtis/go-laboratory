package main

import (
	"fmt"
)

func main() {
	svc, err := initialize()
	dieIfErr(err)

	d := svc.Diskv

	// Write three bytes to the key "gamma".
	key := "gamma"
	err = d.Write(key, []byte{'1', '2', '3'})
	dieIfErr(err)

	// Read the value back out of the store.
	value, err := d.Read(key)
	dieIfErr(err)
	fmt.Printf("%v\n", value)

	// Erase the key+value from the store (and the disk).
	err = d.Erase(key)
	dieIfErr(err)
}
