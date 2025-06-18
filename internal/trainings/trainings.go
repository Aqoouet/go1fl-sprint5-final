package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Aqoouet/go1fl-sprint5-final/internal/common"
	"github.com/Aqoouet/go1fl-sprint5-final/internal/personaldata"
	"github.com/Aqoouet/go1fl-sprint5-final/internal/spentenergy"
)

// Structure for storing info about training.
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse parses a string like "3456,Ходьба,3h00m" to number of steps, activity type and duration.
// It store results to an input structure Training.
// Returns errors in case of parsing errors or if expected limits of parameters are exceeded.
func (t *Training) Parse(datastring string) (err error) {
	dataSl := strings.Split(datastring, ",")
	if len(dataSl) != 3 {
		t.Steps = 0
		t.TrainingType = ""
		t.Duration = time.Duration(0)
		err = fmt.Errorf("slice after string (%q) split is equal to %v with length %d, expected length = 3: %w", datastring, dataSl, len(dataSl), common.ErrSliceLen)
		return
	}

	steps, err := strconv.Atoi(dataSl[0])
	if err != nil {
		t.Steps = 0
		t.TrainingType = ""
		t.Duration = time.Duration(0)
		err = fmt.Errorf("string representing steps = %q: %w", dataSl[0], common.ErrParseInt)
		return
	}

	if steps <= 0 {
		t.Steps = 0
		t.TrainingType = ""
		t.Duration = time.Duration(0)
		err = fmt.Errorf("steps number = %d,  expected non-negative and non-zero steps number: %w", steps, common.ErrParamLimitExceeded)
		return
	}

	activity := strings.TrimSpace(dataSl[1])

	if len(activity) == 0 {
		t.Steps = 0
		t.TrainingType = ""
		t.Duration = time.Duration(0)
		err = fmt.Errorf("activity must be set: %w", common.ErrEmptyString)
		return
	}

	duration, err := time.ParseDuration(dataSl[2])
	if err != nil {
		t.Steps = 0
		t.TrainingType = ""
		t.Duration = time.Duration(0)
		err = fmt.Errorf("string representing duration = %q: %w", dataSl[2], common.ErrParseDuration)
		return
	}

	if duration <= time.Duration(0) {
		t.Steps = 0
		t.TrainingType = ""
		t.Duration = time.Duration(0)
		err = fmt.Errorf("duration = %v, expected non-negative and non-zero duration: %w", duration, common.ErrParamLimitExceeded)
		return
	}

	t.Steps = steps
	t.TrainingType = activity
	t.Duration = duration
	err = nil
	return
}

// ActionInfo generates informative message about training using:
//   - data string ("3456,Ходьба,3h00m")
//   - user's weight and height
//
// # All data is provided to method using structure Training
//
// Returns formatted string or error if parsing or calculation failed.
func (t Training) ActionInfo() (string, error) {

	var calories float64
	var err error

	steps, weight, height, activity, duration := t.Steps, t.Personal.Weight, t.Personal.Height, t.TrainingType, t.Duration

	switch activity {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(steps, weight, height, duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(steps, weight, height, duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: activity = %q, expected activity = ['Ходьба', 'Бег']: %w", activity, common.ErrParamLimitExceeded)
	}

	if err != nil {
		return "", err
	}

	dist := spentenergy.Distance(steps, height)
	mSpeed := spentenergy.MeanSpeed(steps, height, duration)

	var b strings.Builder

	b.WriteString(fmt.Sprintf("Тип тренировки: %s\n", activity))
	b.WriteString(fmt.Sprintf("Длительность: %.2f ч.\n", duration.Hours()))
	b.WriteString(fmt.Sprintf("Дистанция: %.2f км.\n", dist))
	b.WriteString(fmt.Sprintf("Скорость: %.2f км/ч\n", mSpeed))
	b.WriteString(fmt.Sprintf("Сожгли калорий: %.2f\n", calories))

	return b.String(), nil

}
