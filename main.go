package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	"github.com/aarondl/opt/omit"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/types"

	"github.com/d0zer11st/bobexample/internal/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	connectionString := "postgres://test_user:test_user@127.0.0.1:5432/test_db?sslmode=disable"
	ctx := context.Background()

	db, openErr := sql.Open("pgx", connectionString)
	if openErr != nil {
		log.Fatal(openErr)
	}
	defer db.Close()

	bobDB := bob.NewDB(db)

	macToTest := "1b:1b:63:84:45:e6"
	hwAddr, parseMacErr := net.ParseMAC(macToTest)
	if parseMacErr != nil {
		log.Fatal(parseMacErr)
	}
	log.Printf("parsed mac: %s", hwAddr)
	omitMacValue := omit.From(
		types.Stringer[net.HardwareAddr]{
			Val: hwAddr,
		},
	)
	exampleSetter := &models.ExampleSetter{
		Mac: omitMacValue,
	}
	_, err := models.Examples.InsertMany(ctx, bobDB, exampleSetter)
	if err != nil {
		log.Fatal(err)
	}
}
