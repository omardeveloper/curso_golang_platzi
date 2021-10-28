package main

import "fmt"

func main() {
	// Declaracion de constantes (valor que no cambiara en el tiempo)
	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("Hola Omar, Bienvenido a GO!!!")
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Declaracion de variables Enteras
	//:= es si no la hemos usado antes, el : crea la variable y le asigna el valor
	base := 12
	var altura int = 14 //aca declaramos que tipo de dato y le asignamos el valor
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	//Go asigna vaalores a variables vacías
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Calcular area del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es: ", areaCuadrado)
}
