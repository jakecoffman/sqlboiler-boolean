package main

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"github.com/jakecoffman/sqlboiler-tests/models"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
)

//go:embed schema.sql
var schema string

//go:generate sqlboiler --wipe psql

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	db, err := sql.Open("postgres", "postgres://localhost:5432/sqlboiler-test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// migrate
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	author := &models.Author{
		ID:    1,
		Books: []byte(`[{"title":"1984"}]`),
	}
	err = author.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}

	v, err := models.Authors(
		qm.Select("id"),
		qm.Where("id=?", author.ID),
	).One(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}

	_, err = json.Marshal(v)
	if err != nil {
		// BUG: since Books is not selected it is blank which is invalid JSON.
		log.Fatal("Couldn't marshal", err)
	}
}
