package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Brilliant ML Lab"
	app.Usage = "Calculate solutions for Brilliant ML course"
	app.Action = func(clic *cli.Context) error {
		clic.App.Command("noop").Run(clic)
		return nil
	}

	app.Commands = []cli.Command {
		{
			Name: "noop",
			Usage: "Do nothing",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "foo"},
			},
			Action: func(clic *cli.Context) error {
				println("Hello World")
				return nil
			},
		},
		{
			Name: "p-learn",
			Usage: "Find a linear decision boundary",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name: "weight, w",
					Value: 0,
				},
				cli.IntFlag{
					Name: "bias, b",
					Value: 0,
				},
				cli.IntSliceFlag{
					Name: "points, x",
				},
			},
			Action: func(clic *cli.Context) error {
				fmt.Printf("init weight: %d\ninit bias: %d\npoints: %s",
					clic.Int("weight"), clic.Int("bias"), nil)
				return nil
			},
		},
	}

	app.Run(os.Args)
}