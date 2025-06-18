// Package daysteps provides functionality for processing and describing daily walking data.

// It includes:
// - A DaySteps struct for storing walk-related data.
// - Exported methods (Parse and ActionInfo) for parsing input strings and generating informational output.

// The package is intended to be used across different modules of the project,
// and therefore uses exported identifiers to ensure accessibility.

package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Aqoouet/go1fl-sprint5-final/internal/common"
	"github.com/Aqoouet/go1fl-sprint5-final/internal/personaldata"
	"github.com/Aqoouet/go1fl-sprint5-final/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse parses a string like "678,0h50m" to number of steps and duration.
// Results of parsing are stored in DaySteps structure.
// Returns errors in case of parsing errors or if expected limits of parameters are exceeded.
func (ds *DaySteps) Parse(datastring string) (err error) {

	dataSl := strings.Split(datastring, ",")
	if len(dataSl) != 2 {
		ds.Steps = 0
		ds.Duration = time.Duration(0)
		err = fmt.Errorf("slice after string (%q) split is equal to %v with length %d, expected length = 2: %w", datastring, dataSl, len(dataSl), common.ErrSliceLen)
		return
	}

	steps, err := strconv.Atoi(dataSl[0])
	if err != nil {
		ds.Steps = 0
		ds.Duration = time.Duration(0)
		err = fmt.Errorf("string representing steps = %q: %w", dataSl[0], common.ErrParseInt)
		return
	}

	if steps <= 0 {
		ds.Steps = 0
		ds.Duration = time.Duration(0)
		err = fmt.Errorf("steps number = %d,  expected non-negative and non-zero steps number: %w", steps, common.ErrParamLimitExceeded)
		return
	}

	duration, err := time.ParseDuration(dataSl[1])
	if err != nil {
		ds.Steps = 0
		ds.Duration = time.Duration(0)
		err = fmt.Errorf("string representing duration = %q: %w", dataSl[1], common.ErrParseDuration)
		return
	}

	if duration <= 0 {
		ds.Steps = 0
		ds.Duration = time.Duration(0)
		err = fmt.Errorf("duration = %v, expected non-negative and non-zero duration: %w", duration, common.ErrParamLimitExceeded)
		return
	}

	ds.Steps = steps
	ds.Duration = duration
	err = nil

	return
}

// ActionInfo returns informative message using information about user weight and height:
// Message contains next parameters:
//   - number od steps
//   - length of a trip
//   - calories
func (ds DaySteps) ActionInfo() (string, error) {

	steps, duration, weight, height := ds.Steps, ds.Duration, ds.Personal.Weight, ds.Personal.Height

	if steps <= 0 {
		return "", fmt.Errorf("number of steps = %v, expected non-negative and non-zero number of steps: %w", steps, common.ErrParamLimitExceeded)
	}

	// distanceM := float64(steps) * common.StepLength
	// distanceKm := distanceM / common.MinKm

	distanceKm := spentenergy.Distance(steps, height)

	calories, err := spentenergy.WalkingSpentCalories(steps, weight, height, duration)

	if err != nil {
		log.Println(err)
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n",
		steps, distanceKm, calories), nil

}
