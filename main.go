package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"main.go/core"
)

func main() {

	app := &cli.App{
		Name:  "Elephant cache",
		Usage: "Run the server with TTL to cache everything pass through",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "ttl",
				Value:       "",
				Usage:       "Time to live for the cache",
				Destination: &core.Options.TTL,
			},
		},
		Action: func(cCtx *cli.Context) error {

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
