package simulation

import (
	"image/color"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/benoitmasson/plotters/piechart"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
)

// Results holds the data obtained from the simulation
type Results struct {
	Susceptible, Infected, Resistant int
}

// Simulate runs the actual simulation for the program to display
// Returns the number of timesteps used during the actual simulation
func Simulate(pop, initI, until int, S2I, I2R, S2R float64) int {
	result := Results{
		pop - initI,
		initI,
		0,
	}

	//max time steps
	if until < 0 || until > 2000 {
		until = 2000
	}

	t := 0
	for ; t < until; t++ {
		simStep(&result, S2I, I2R, S2R)
		generatePieChart(result, pop, t)
		if result.Infected >= pop || result.Susceptible+result.Resistant >= pop {
			break
		}
	}
	return t + 1
}

func simStep(res *Results, S2I, I2R, S2R float64) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	sus := res.Susceptible
	for i := 0; i < sus; i++ {
		rand := rng.Float64()
		if rand < S2I {
			res.Susceptible--
			res.Infected++
		}
	}
	sus = res.Susceptible
	for i := 0; i < sus; i++ {
		rand := rng.Float64()
		if rand < S2R {
			res.Susceptible--
			res.Resistant++
		}
	}
	inf := res.Infected
	for i := 0; i < inf; i++ {
		rand := rng.Float64()
		if rand < I2R {
			res.Infected--
			res.Resistant++
		}
	}
}

func generatePieChart(r Results, pop, i int) {
	var err error
	var p *plot.Plot
	if p, err = plot.New(); err != nil {
		log.Fatalf("Could not generate pie chart. Error: %v", err.Error())
	}
	p.HideAxes()
	p.BackgroundColor = color.RGBA{0, 0, 0, 0}
	p.Legend.Top = true
	p.Legend.Color = color.RGBA{42, 159, 214, 255}

	var susPie *piechart.PieChart
	if susPie, err = piechart.NewPieChart(plotter.Values{float64(r.Susceptible)}); err != nil {
		log.Fatalf("Susceptible pie not created! Error: %v", err.Error())
	}
	susPie.Color = color.RGBA{255, 0, 0, 255}
	susPie.Total = float64(pop)
	p.Add(susPie)
	p.Legend.Add("Susceptible", susPie)

	var infPie *piechart.PieChart
	if infPie, err = piechart.NewPieChart(plotter.Values{float64(r.Infected)}); err != nil {
		log.Fatalf("Infected pie not created! Error: %v", err.Error())
	}
	infPie.Color = color.RGBA{0, 255, 0, 255}
	infPie.Total = float64(pop)
	infPie.Offset.Value = float64(r.Susceptible)
	p.Add(infPie)
	p.Legend.Add("Infected", infPie)

	var resPie *piechart.PieChart
	if resPie, err = piechart.NewPieChart(plotter.Values{float64(r.Resistant)}); err != nil {
		log.Fatalf("Resistant pie not created! Error: %v", err.Error())
	}
	resPie.Color = color.RGBA{0, 0, 255, 255}
	resPie.Total = float64(pop)
	resPie.Offset.Value = float64(r.Susceptible + r.Infected)
	p.Add(resPie)
	p.Legend.Add("Resistant", resPie)

	p.Save(4*vg.Inch, 4*vg.Inch, "astiServer/static/pieCharts/pieChart"+strconv.Itoa(i)+".png")
}
