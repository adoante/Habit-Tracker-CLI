package main

import (
	"os"
	"strconv"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Habbit"
	app.Usage = "To Track Habits"
	app.Version = "0.0.1"

	app.EnableBashCompletion = true

	var addHabit string
	var palomaHabit string
	var markHabit int
	var removeHabit int

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "add,a",
			Usage:       "add habit to track",
			Value:       "name",
			Destination: &addHabit,
		},
		cli.StringFlag{
			Name:        "paloma,p",
			Usage:       "set check mark character",
			Value:       "character",
			Destination: &palomaHabit,
		},
		cli.IntFlag{
			Name:        "mark,m",
			Usage:       "mark a habit complete",
			Value:       0,
			Destination: &markHabit,
		},
		cli.IntFlag{
			Name:        "remove,r",
			Usage:       "remove a habit from tracker",
			Value:       0,
			Destination: &removeHabit,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:        "stats",
			Aliases:     []string{"s"},
			Usage:       "Shows habit tracker statistics.",
			Description: "idk",
			Action: func(c *cli.Context) error {
				println("stats command was used")
				return nil
			},
		},
		{
			Name:        "check",
			Aliases:     []string{"c"},
			Usage:       "Shows habit tracker card.",
			Description: "idk",
			Action: func(c *cli.Context) error {
				println("check command was used")
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		if addHabit != "name" {
			println("Added '" + addHabit + "' to habit tracker.")
		} else if palomaHabit != "character" {
			println("'" + palomaHabit + "' is your new check mark.")
		} else if markHabit != 0 {
			println("Habit: " + strconv.Itoa(markHabit) + "has been marked!")
		} else if removeHabit != 0 {
			println("Removed '" + strconv.Itoa(removeHabit) + "' from habit tracker.")
		}
		return nil
	}

	app.UsageText = "habbit [command] [arg]"
	app.Author = "Adolfo Gante - https://github.com/adoante"
	app.Run(os.Args)
}
