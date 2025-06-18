<<<<<<< HEAD
package personaldata

type Personal struct {
	// TODO: добавить поля
}

func (p Personal) Print() {
	// TODO: реализовать функцию
=======
// Package for storing of structure Personal used across the project.

package personaldata

import "fmt"

// Structure for storing user data
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Method for user data printing
func (p Personal) Print() {
	fmt.Printf(`Имя: %s
Вес: %.2f кг.
Рост: %.2f м.

`, p.Name, p.Weight, p.Height)
>>>>>>> modif
}
