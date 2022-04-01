package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "Brilliant ML Lab"
	app.Usage = "Calculate solutions for Brilliant ML course"
	app.Action = func(clic *cli.Context) error {
		clic.App.Command("noop").Run(clic)
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "noop",
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
			Name:  "p-learn",
			Usage: "Find a linear decision boundary",
			Flags: []cli.Flag{
				// cli.Float64SliceFlag{
				// 	Name:  "weight, w",
				// 	Value: zeros,
				// },
				cli.Float64Flag{
					Name:  "bias, b",
					Value: 0.0,
				},
				cli.StringFlag{
					Name: "points, x",
				},
			},
			Action: func(clic *cli.Context) error {
				points := ReadInputPoints(clic.String("points"))
				weights := make([]float64, len(points[0].x))
				bias := clic.Float64("bias")
				FindLinearBoundary(&points, &weights, &bias)
				fmt.Printf("Weights: %f\nBias: %f\n", weights, bias)
				return nil
			},
		},
	}

	app.Run(os.Args)
}

func FindLinearBoundary(points *[]point, weights *[]float64, bias *float64) {

}
