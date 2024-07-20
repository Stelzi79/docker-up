package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	fmt.Println("📢 Hallo!!!") // Code here
	app := &cli.App{
		Name:  "docker-up",
		Usage: "A simplification of commands surrounding Docker Compose up.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Enable debug mode",
				Value:   false,
			},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.Bool("debug") {
				fmt.Println("🐞 Debugging enabled")
			}
			fmt.Printf("Hello %q", cCtx.Args().Get(0))
			fmt.Println("🚀")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
