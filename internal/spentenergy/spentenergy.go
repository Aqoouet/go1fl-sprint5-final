<<<<<<< HEAD
package spentenergy

import (
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
=======
// Package spentenergy provides functions for calculating calories burned during physical activities
// such as walking and running. It also includes utilities for calculating average speed and distance.

// These functions are intended to be used across different packages within the project,
// so all public methods are exported and follow Go's naming conventions.

package spentenergy

import (
	"fmt"
	"time"

	"github.com/Aqoouet/go1fl-sprint5-final/internal/common"
)

// distance calculates distance using number of steps and height
func Distance(steps int, height float64) float64 {
	stepLength := common.StepLengthCoefficient * height
	distanceM := stepLength * float64(steps)
	distanceKm := distanceM / common.MinKm
	return distanceKm
}

// meanSpeed takes number of steps, height and activity duration.
// Returns mean speed.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 {
		return 0.
	}
	d := Distance(steps, height)
	hours := duration.Hours()

	if hours <= 0 {
		return 0.0
	}

	return d / hours
}

// RunningSpentCalories calculates spent calories number for running activity.
// Returns an error if parameters exceeds allowable limits.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("steps number = %d,  expected non-negative and non-zero steps number: %w", steps, common.ErrParamLimitExceeded)
	}

	if weight <= 0. {
		return 0, fmt.Errorf("weight = %.2f,  expected non-negative and non-zero weight: %w", weight, common.ErrParamLimitExceeded)
	}

	if height <= 0 {
		return 0, fmt.Errorf("height = %.2f,  expected non-negative and non-zero height: %w", height, common.ErrParamLimitExceeded)
	}

	if duration <= 0 {
		return 0, fmt.Errorf("duration = %v,  expected non-negative and non-zero duration: %w", duration, common.ErrParamLimitExceeded)
	}

	mSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	return (weight * mSpeed * durationInMinutes) / common.MinInH, nil
}

// WalkingSpentCalories calculates spent calories number for walking activity.
// Returns an error if parameters exceeds allowable limits.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	c, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return c * common.WalkingCaloriesCoefficient, nil

>>>>>>> modif
}
