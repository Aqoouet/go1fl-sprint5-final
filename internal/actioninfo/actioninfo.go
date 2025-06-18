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
			return
		}

		msg, err := dp.ActionInfo()

		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(msg)

	}
}
