package main

import (
	"Tarea_1/queue"
	"Tarea_1/stack"
	"fmt"
)

func main() {
	//Prueba con la pila
	pila := stack.Stack{}
	pila.Push(5)
	pila.Push(10)
	pila.Push(15)
	pila.Push(20)
	//Datos
	pila.Printpila()

	top, err := pila.SkTop()
	//Pruebas para stak
	if err == nil {
		fmt.Println("El item Top es: ", top)
	}
	pila.Pop()
	pila.Printpila()
	pila.Pop()
	pila.Pop()
	pila.Pop()
	pila.Printpila()
	//Prueba para la queue
	queque := queue.Queue{}
	//Datos
	queque.Inqueue(1)
	queque.Inqueue(2)
	queque.Inqueue(3)
	queque.Inqueue(4)
	queque.Inqueue(5)
	//Checar datos
	queque.PrintQueue()
	//Pruebas
	queque.Enqueue()
	queque.Enqueue()
	queque.Enqueue()

	queque.PrintQueue()

	queque.Enqueue()
	queque.Enqueue()
	queque.PrintQueue()
	//Pruebas para Diccionarios4

	dicc := make(map[string]int)
	dicc["Uno"] = 1
	dicc["Dos"] = 2
	dicc["Tres"] = 3

	//diccionario := map[string]int{
	//	"Uno":  1,
	//	"Dos":  2,
	//	"Tres": 3,
	//}

	fmt.Println("")
	fmt.Println(dicc)
	delete(dicc, "Dos")
	fmt.Println(dicc)
	//fmt.Println(diccionario)

}
