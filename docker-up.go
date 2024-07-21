package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	colorsTest = "\n{{.Colors.Green}}Hello sa\tdf adf {{.Colors.Underline}}asd{{.Colors.Bold}} fas df{{.Colors.Reset}}\n"
)

var debug bool

func main() {
	fmt.Println("📢 Hallo!!!") // Code here
	app := &cli.App{
		Name:  "docker-up",
		Usage: "A simplification of commands surrounding Docker Compose up.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				Usage:       "Enable debug mode",
				Value:       false,
				Destination: &debug,
			},
		},
		Action: func(cCtx *cli.Context) error {
			chkDebug()
			//  {{"\"output\""}}
			fmt.Printf("Hello %q", cCtx.Args().Get(0))
			t1 := template.New("t1")
			t1, err := t1.Parse(colorsTest)
			if err != nil {
				panic(err)
			}

			t1.Execute(os.Stdout, map[string]Col{"Colors": Colors})

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func chkDebug() {
	if debug {
		fmt.Println("🐛 Debugging enabled")
	}
}
