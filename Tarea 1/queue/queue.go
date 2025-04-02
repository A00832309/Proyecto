package queue

import "fmt"

type Queue struct {
	items []int
}

// Funciones del Queue
// Principales

func (q *Queue) Inqueue(dato int) {
	q.items = append(q.items, dato)
}

func (q *Queue) Enqueue() {
	q.items = q.items[1:]
}
func (q *Queue) QuVacio() bool {
	if len(q.items) == 0 {
		return true
	}
	return false
}

// Funciones extras
func (q *Queue) Frente() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("Queue vacia")
	}
	return q.items[0], nil
}
func (q *Queue) PrintQueue() {
	fmt.Print("La fila tiene: ")
	if q.QuVacio() {
		fmt.Print("0 Datos")
		return
	}
	for _, item := range q.items {
		fmt.Print(item, " ")
	}
	fmt.Print("datos ")
	fmt.Println()
}
