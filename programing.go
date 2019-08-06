package main

import "fmt"
import "time"

func main(){
	fmt.Println("***** MI PROGRAMA CON GO ****")


	momento := time.Now()
	hoy := momento.Weekday();

	switch hoy {
		case 0:
			fmt.Println("Hoy es domingo")

		case 1:
			fmt.Println("Hoy es Lunes")

		case 2:
			fmt.Println("Hoy es Martes")

		case 3:
			fmt.Println("Hoy es Miercoles")

		default:
			fmt.Println("Hoy es otro dia de la semana")
	}

}
