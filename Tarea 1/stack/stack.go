package stack

import "fmt"

type Stack struct {
	items []int
}

// Funciones del stack
// Principales
func (s *Stack) Push(dato int) {
	s.items = append(s.items, dato)
}
func (s *Stack) Pop() {
	if s.SkVacio() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

// Complementos o extras
func (s *Stack) SkVacio() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}
func (s *Stack) SkTop() (int, error) {
	if s.SkVacio() {
		return 0, fmt.Errorf("stack vacio")
	}
	return s.items[len(s.items)-1], nil
}
func (s *Stack) Printpila() {
	fmt.Print("La pila tiene: ")
	if s.SkVacio() {
		fmt.Println("0 datos")
		return
	}
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Print("datos")
	fmt.Println()
}
