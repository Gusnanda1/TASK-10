package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func DB_CONN() {

	connURL := "postgres://postgres:Gusnanda123@localhost:5000/db_project"
	var err error

	Conn, err = pgx.Connect(context.Background(), connURL)

	if err != nil {
		fmt.Println("Undefined Database", err)
		os.Exit(1)
	}

	fmt.Println("Connection to database successfully!")

}
