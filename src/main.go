package main

import (
	"fmt"
	"math"
	"strings"

	mypackage "curso_golang_platzi/src/mypackage"
)

type car struct {
	brand   string
	year    int
	seating int
	color   string
	owner   string
}

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

	// USO DE FMT
	helloMessage := "Hello"
	worldMessage := "world"

	// Println: Salto de Linea Automatico
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	// Con valores seguros (%s es un string y %d un entero)
	fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos)
	// Con valores inseguros (%v si no se que tipo de dato )
	fmt.Printf("%v tiene mas de %v cursos \n", nombre, cursos)

	// Sprintf
	//genera un string pero no lo imprime en consola lo guarda
	message := fmt.Sprintf("%v tiene mas de %v cursos", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos saber con %T
	fmt.Printf("helloMessage: %T \n", helloMessage)
	fmt.Printf("custos: %T \n", cursos)

	value := returnValue(2)
	fmt.Println(value)

	value1, value2 := dobleReturn(3)
	fmt.Println(value1, value2)
	//si me retorno dos o mas y solo me interesa un valor pongo el _
	//go descarta ese valor y toma el que defina
	value3, _ := dobleReturn(5)
	fmt.Println(value3)

	//area circulo
	fmt.Println("Area del circulo: ", areaCalCirculo(8))

	//bucles
	bucleFor(8)
	bucleForWhile(5)
	//bucleForever(0)

	//operadores
	isPair(3)
	isPair(4)

	if isValidUser("omar", "123pass") {
		fmt.Println("Credenciales correctas")
	} else {
		fmt.Println("credenciales incorrectas")
	}

	dayOfWeek(5)
	trimesterOfYear()
	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("un número par")
	default:
		fmt.Println("un número impar")
	}

	//keywords defer
	deferKeyword()
	keywordsBreakAndContinue()

	//Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	//len cantidad de elementos en un array
	//cap indica capacidad maxima de ese array
	fmt.Println(array, len(array), cap(array))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice, len(slice), cap(slice))
	// Métodos en el slice
	fmt.Println(slice[0])
	//imprima hasta le inidce 3
	fmt.Println(slice[:3])
	//imprima desde : hasta
	fmt.Println(slice[2:4])
	//imprima desde
	fmt.Println(slice[4:])
	//los segundos elementos (:hasta) los excluye por eso no lo muestra

	// Append
	slice = append(slice, 11)
	fmt.Println(slice)
	// Append list
	newSlice := []int{12, 13, 14}
	//descomprimeeses como si hubiese hecho (slice, 12,13,14)
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	//slice recorrido
	slice2 := []string{"hola", "bienvenido", "a", "GO"}
	//indice, valor
	for i, valor := range slice2 {
		fmt.Println(i, valor)
	}
	//escapo el indice no lo quiero ver, valor
	for _, valor := range slice2 {
		fmt.Println(valor)
	}
	//solo el inidce
	for i := range slice2 {
		fmt.Println(i)
	}

	fmt.Println(isPalindrome("anna"))
	fmt.Println(isPalindrome("omar"))

	//Forma recomendada de inicializar el Map: map Con Make
	//corchete va el tipo de dato que va alli para acceder al valor
	//segundo tipo de dato
	m := make(map[string]int)
	m["jose"] = 14
	m["pepito"] = 20
	fmt.Println(m)
	// Recorrer un map
	for i, v := range m {
		fmt.Printf("%s tiene %d años\n", i, v)
	}
	// Encontrar un valor
	valormap := m["jose"]
	fmt.Println(valormap)
	//sino tengo certeza si una llave esta en el map
	valor2, ok := m["jose"]
	fmt.Println(valor2, ok)

	mapInicializado1Paso := map[string]int{"abc": 1}
	fmt.Println("Ejemplo: mapInicializado1Paso=", mapInicializado1Paso)

	myCar := car{brand: "Mazda", year: 2021, seating: 5, color: "azul", owner: "yo"}
	fmt.Println("Los Datos de mi auto son:", myCar)

	var otherCar car
	otherCar.brand = "Renault"
	fmt.Println(otherCar)

	var myCarPublic mypackage.CarPublic
	myCarPublic.Brand = "Ferrari"
	myCarPublic.Year = 2021
	fmt.Println(myCarPublic)

	mypackage.PrintMessage("mi mensaje personalizado")

	//PUNTEWROS
	numberA := 50
	//& sera la direccion de memoria
	numberB := &numberA
	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB)
	//modifico b que tiene 100 pero imprimo a
	*numberB = 100
	fmt.Println(numberA)

	myPC := mypackage.Pc{Ram: 16, Disk: 200, Brand: "msi"}
	fmt.Println(myPC)
	myPC.Ping()
	myPC.DuplicateRam()
	myPC.DuplicateDisk()
	fmt.Println(myPC)

}

func isPalindrome(text string) bool {
	fmt.Println(strings.ToLower(text))
	for i := 0; i < len(text)/2; i++ {
		//text[i] em retorna el ascii con (string)text[i] meda la letra y tner encuenta si es mayus o no
		if text[i] != text[len(text)-i-1] {
			return false
		}
	}
	return true
}

// los parametros asi (a int, b int, c string)
// si a y b son el mismo tipo de dato puedo hacer
//(a, b int, c string)
func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

//para retornar dos o mas valores
func dobleReturn(a int) (x, y int) {
	return a, a * 3
}

func areaCalCirculo(radio float64) float64 {
	return math.Pi * radio
}

func bucleFor(iterator int) {
	for i := iterator; i > 0; i-- {
		fmt.Printf("mi iteracion va en: %d \n", i)
	}
}

func bucleForWhile(counter int) {
	for counter < 10 {
		fmt.Printf("Mi countes va en: %d \n", counter)
		counter++
	}
}

func bucleForever(counterForever int) {
	for {
		fmt.Printf("Output: %d \n", counterForever)
		counterForever++
	}
}

func isPair(valor int) {
	if valor%2 == 0 {
		fmt.Println(valor, "Es par")
	} else {
		fmt.Println(valor, "No es par")
	}
}

func isValidUser(userName, pass string) bool {
	return userName == "Alpha" && pass == "MyPassword"
}

func dayOfWeek(day int) {
	switch day {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Ese no es un día valido de la semana!")
	}
}

func trimesterOfYear() {
	var mes int8 = 10
	switch {
	case mes <= 3:
		fmt.Println("Primer Trimestre")
	case mes > 3 && mes <= 6:
		fmt.Println("Segundo Trimestre")
	case mes > 6 && mes <= 9:
		fmt.Println("Tercer Trimestre")
	case mes > 9 && mes <= 12:
		fmt.Println("Cuarto Trimestre")
	default:
		fmt.Println("Este no es un mes valido")
	}
}

func deferKeyword() {
	//defer ejecuta la ultima funcion antes de que todo termine
	/* ** defer: ** ejecuta la parte de código especificada al final de la ejecución de la función o método. Como buena práctica usar solo una sentencia “defer” dentro de un bloque. */
	defer fmt.Println("Hola")
	fmt.Println("Mundo")
}

func keywordsBreakAndContinue() {
	defer fmt.Println("Aprendiendo Keywords esto se escribe primero y se imprime ultimo por defer")
	// continue y break
	for i := 0; i < 10; i++ {
		// Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//break
		if i == 8 {
			fmt.Println("Es 8 -> Break")
			break
		}
		fmt.Println(i)
	}
}
