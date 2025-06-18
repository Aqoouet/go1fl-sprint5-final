// Package actioninfo provides functionality for summarizing and displaying general information
// about various training sessions and daily walks.

// It defines:
//   - An interface DataParser with methods Parse() and ActionInfo()
//   - A function Info() that accepts data slices and struct instances (e.g., Training and DaySteps)
//     to generate a consolidated output

// The package supports modular design by working with interfaces,
// making it flexible for use with different data types and structures.

package actioninfo

import (
	"fmt"
	"log"

	"github.com/Aqoouet/go1fl-sprint5-final/internal/daysteps"
	"github.com/Aqoouet/go1fl-sprint5-final/internal/trainings"
)

type DataParser interface {
	Parse(a string) error
	ActionInfo() (string, error)
}

var _ DataParser = new(daysteps.DaySteps)
var _ DataParser = new(trainings.Training)

func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		err := dp.Parse(v)

		if err != nil {
			log.Println(err)
			continue
		}

		msg, err := dp.ActionInfo()

		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(msg)

	}
}
