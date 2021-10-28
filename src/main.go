package main

import (
	"fmt"
	"math"
)

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

	// Valores:
	x := 1000
	y := 500
	//Suma:
	sumaResult := x + y
	//Resta:
	restaResult := x - y

	fmt.Println("La suma de los 2 valores es ", sumaResult, ".Mientras que su resta equivale a ", restaResult, ".")

	//Multiplicacion:
	multiResult := x * y
	// División:
	diviResult := x / y

	fmt.Println("La multiplicación de los dos valores es", multiResult, " y su divisón es igual a", diviResult, ".")

	// Modulo:
	modulResult := x % y

	fmt.Println("El modulo de los dos valores es: ", modulResult)

	// Incremental:
	x++
	fmt.Println("El incremental de x es: ", x)
	// Decremental:
	x--
	x--
	fmt.Println("El decremental de x es: ", x)

	// area del circulo
	radio := 2
	areaCirculo := math.Pi * math.Pow(float64(radio), 2)
	println("areaCirculo", areaCirculo)
}
