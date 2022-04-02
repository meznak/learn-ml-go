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
	point_count := len(*points)
	points_since_change := 0

	/*
		while points_since_change < point_count:
			check sign of result vs class
		    if result != class:
				update weights & bias
				reset points_since_change
			else:
				increment points_since_change

			if index == point_count:
				reset index
	*/

	for {
		for _, point := range *points {
			result := dot_product(&point.x, weights) + *bias
			if result*point.y <= 0 {
				update_wb(point, weights, bias)
				points_since_change = 0
			} else {
				points_since_change++
				if points_since_change == point_count {
					return
				}
			}
		}
	}
}

func dot_product(a, b *[]float64) float64 {
	total := 0.0

	for i, x := range *a {
		total += x * (*b)[i]
	}

	return total
}

func scalar_product(coeff float64, vec []float64) []float64 {
	output := make([]float64, len(vec))

	for i, v := range vec {
		output[i] = v * coeff
	}

	return output
}

func vec_sum(a, b []float64) []float64 {
	output := make([]float64, len(a))

	for i, v := range a {
		output[i] = v + b[i]
	}

	return output
}

func update_wb(p point, weights *[]float64, bias *float64) {
	scaled_x := scalar_product(p.y, p.x)

	*weights = vec_sum(*weights, scaled_x)
	*bias = *bias + p.y
}
