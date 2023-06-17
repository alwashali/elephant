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
		Usage: "Run the server with TTL to cache everything passes through",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "ttl",
				Value:       "",
				Usage:       "Time to live for the cache",
				Required:    true,
				Destination: &core.Options.TTL,
			},
		},
		Action: func(Ctx *cli.Context) error {
			core.Run()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
