package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alwashali/elephant/core"
	opts "github.com/alwashali/elephant/options"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "Elephant cache",
		Usage: "Run the server with TTL to cache everything pass through",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "ttl",
				Value:       "",
				Usage:       "Time to live for the cache, example: 24h",
				Destination: &opts.TTL,
			},
		},
		Action: func(cCtx *cli.Context) error {
			if opts.TTL != "" {
				fmt.Printf("\n\n")
				cacheDuration, _ := time.ParseDuration(opts.TTL)
				fmt.Println("Caching Duration: ", cacheDuration.String())
				core.Run()

			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
