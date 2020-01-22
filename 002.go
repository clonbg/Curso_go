package main

import "fmt"

// tutorialesprogrammacionya.com go T11

func main() {
	var edad, promedioManana, edadManana, promedioTarde, edadTarde, promedioNoche, edadNoche float32
	var texto string
	fmt.Println("Cinco edades de turno de mañana: ")
	for f := 0; f < 5; f++ {
		fmt.Scan(&edad)
		edadManana = edadManana + edad
	}
	promedioManana = (edadManana / 5.00)
	fmt.Println("Mañanas:", promedioManana)

	fmt.Println("Seis edades de turno de tarde: ")
	edad = 0
	for f := 0; f < 6; f++ {
		fmt.Scan(&edad)
		edadTarde = edadTarde + edad
	}
	promedioTarde = (edadTarde / 6.00)
	fmt.Println("Tardes: ", promedioTarde)

	fmt.Println("Seis edades de turno de noche: ")
	edad = 0
	for f := 0; f < 11; f++ {
		fmt.Scan(&edad)
		edadNoche = edadNoche + edad
	}
	promedioNoche = (edadNoche / 11.00)
	fmt.Println("Noches: ", promedioNoche)

	if promedioManana > promedioTarde && promedioManana > promedioNoche {
		texto = "Mañanas"
	} else if promedioTarde > promedioManana && promedioTarde > promedioNoche {
		texto = "Tardes"
	} else {
		texto = "Noches"
	}

	fmt.Println("El promedio mayor es el de: ", texto)
}
