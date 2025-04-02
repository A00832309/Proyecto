package main

import (
	"Tarea_1/queue"
	"Tarea_1/stack"
	"fmt"
)

func main() {
	//Prueba con la pila
	//Para estos casos solamente se probo agregar valores dentro de la pila
	pila := stack.Stack{}
	pila.Push(5)
	pila.Push(10)
	pila.Push(15)
	pila.Push(20)
	//Datos
	pila.Printpila()

	top, err := pila.SkTop()

	//Para este caso solamente se probo cual es el top para ver que si se agrego bien
	if err == nil {
		fmt.Println("El item Top es: ", top)
	}
	//Quitamos los valores eh imprimomos el resultado hasya que quede vacio
	pila.Pop()
	pila.Printpila()
	pila.Pop()
	pila.Pop()
	pila.Pop()
	pila.Printpila()
	//Prueba para la queue creamos y agregamos los valores
	queque := queue.Queue{}
	//Datos
	queque.Inqueue(1)
	queque.Inqueue(2)
	queque.Inqueue(3)
	queque.Inqueue(4)
	queque.Inqueue(5)
	//Checar datos si entraron correctamete
	queque.PrintQueue()
	//sacamos los 3 valores
	queque.Enqueue()
	queque.Enqueue()
	queque.Enqueue()
	// comprobamo que salieron bien del queue
	queque.PrintQueue()

	queque.Enqueue()
	queque.Enqueue()
	queque.PrintQueue()
	//Pruebas para Diccionarios4

	dicc := make(map[string]int)
	//incertamos datos
	dicc["Uno"] = 1
	dicc["Dos"] = 2
	dicc["Tres"] = 3

	//diccionario := map[string]int{
	//	"Uno":  1,
	//	"Dos":  2,
	//	"Tres": 3,
	//}
	//imprimimos los resultados
	fmt.Println("")
	fmt.Println(dicc)
	//eliminamos el segundo valor del dicionario
	delete(dicc, "Dos")
	fmt.Println(dicc)
	//no se pudo vaciar el diccionario dado a que mandaba error
	//dicc := make(map[string]int)
	fmt.Println(dicc)
	//fmt.Println(diccionario)

}
