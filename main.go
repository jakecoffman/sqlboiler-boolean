package main

import (
	"context"
	"database/sql"
	_ "embed"
	"github.com/jakecoffman/sqlboiler-boolean/models"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
)

//go:embed schema.sql
var schema string

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

	blah := &models.Blah{
		ID:        1,
		IsDeleted: false,
	}
	err = blah.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}

	_, err = models.Blahs(
		qm.Where("id=$1", blah.ID),
	).UpdateAll(context.Background(), db, models.M{"is_deleted": true})
	if err != nil {
		// bug?
		//models: unable to update all for blah: pq: column "is_deleted" is of type boolean but expression is of type integer
		log.Fatal(err)
	}
}
