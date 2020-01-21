// tutorialesprogrammacionya.com go T11

package main

import "fmt"

func main() {
	var numTriangulos, superior int
	fmt.Println("Triángulos a calcular:")
	fmt.Scan(&numTriangulos)
	for f := 0; f < numTriangulos; f++ {
		var base, altura float32
		fmt.Println("Introduzca la base: ")
		fmt.Scan(&base)
		fmt.Println("Introduzca la altura: ")
		fmt.Scan(&altura)
		fmt.Println("El triángulo ", (f + 1), " tiene una base de ", base, " , una altura de ", altura, " y una superficie de ", ((base * altura) / 2))
		if ((base * altura) / 2) > 12 {
			superior++
		}
	}
	fmt.Println("Hay ", superior, " triángulos mayores a 12")

}
