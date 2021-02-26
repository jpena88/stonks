package main

import (
	"log"
	"os"

	"github.com/jpena88/stonks/api"
	"github.com/urfave/cli/v2"
)

// HotList is a slice of all of the popular stonks
var HotList = []string{
	"tsla",
	"aapl",
	"msft",
	"twtr",
	"nflx",
}

func main() {
	var symbols []string
	app := &cli.App{
		Name:  "stonks",
		Usage: "a CLI tool for all of your investments",
	}
	app.EnableBashCompletion = true
	app.Action = func(c *cli.Context) error {
		symbols = append(symbols, HotList...)
		s := api.Stonk{SymbolList: symbols}
		s.Print()
		return nil
	}
	app.Commands = []*cli.Command{
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Get details of particular stonk by symbol",
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return nil
				}
				symbolInput := c.Args().First()
				symbols = append(symbols, symbolInput)
				s := api.Stonk{SymbolList: symbols}
				s.Print()
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
