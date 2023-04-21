package main

func main() {
	svc, err := initialize()
	dieIfErr(err)

	db := svc.BoltDB
	tx, err := db.Begin(true)
	dieIfErr(err)

	defer func() { _ = tx.Rollback() }()

	_, err = tx.CreateBucket([]byte("MyBucket"))
	dieIfErr(err)

	err = tx.Commit()
	dieIfErr(err)
}
