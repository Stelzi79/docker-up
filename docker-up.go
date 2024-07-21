package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var debug bool = false
var tpls bool = false

func main() {

	fmt.Println("📢 Hallo!!!", "asfd") // Code here

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
			&cli.BoolFlag{
				Name:        "tpls",
				Usage:       "Saves Templates to 💾. Needs flag '--debug' to work",
				Value:       false,
				Destination: &tpls,
			},
		},
		Action: func(cCtx *cli.Context) error {
			chkDebug()

			fmt.Printf("Hello %q\n", cCtx.Args().Get(0))

			fmt.Println(templates.colorsTest)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func chkDebug() {
	if debug {
		fmt.Println(templates.debugEnabled)
		// fmt.Println(cli.AppHelpTemplate)
		if tpls {
			os.WriteFile("tpls/AppHelpTemplate.txt", []byte(cli.AppHelpTemplate), 0644)
			fmt.Println("AppHelpTemplate.txt ... saved")
			os.WriteFile("tpls/CommandHelpTemplate.txt", []byte(cli.CommandHelpTemplate), 0644)
			fmt.Println("CommandHelpTemplate.txt ... saved")
			os.WriteFile("tpls/SubcommandHelpTemplate.txt", []byte(cli.SubcommandHelpTemplate), 0644)
			fmt.Println("SubcommandHelpTemplate.txt ... saved")
		}
	}
}
