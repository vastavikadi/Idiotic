package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/urfave/cli/v2"
)

var pool *pgxpool.Pool
var ctx = context.Background()


func Init(){
	var err error
	pool, err = pgxpool.New(ctx,"postgres://admin:secret@localhost:5432/idiotic")
	if err !=nil{
		log.Fatal("Unable to connect to database:", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Unable to ping database:", err)
	}

	fmt.Println("Connected to database successfully!")

}

func main(){
	app := &cli.App{
		Name: "Idiotic",
		Usage: "A simple CLI program to execute operations related to the Idiotic application",
        Commands: []*cli.Command{
            // We'll add commands here
        },
	}
	err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)
	}
}