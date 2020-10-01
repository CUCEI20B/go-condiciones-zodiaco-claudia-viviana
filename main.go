package main

import "fmt"

func main() {
	var dia int32
	var mes int32
	var signo string

	fmt.Scan(&dia)
	fmt.Scan(&mes)

	switch {
	// ENERO
	case mes == 1:
		if dia >= 21 && dia <= 31 {
			signo = "Acuario"
		} else if dia >= 1 && dia < 21 {
			signo = "Capricornio"
		} else {
			signo = "Día inválido"
		}
	// FEBRERO
	case mes == 2:
		if dia >= 20 && dia <= 29 {
			signo = "Pisis"
		} else if dia >= 1 && dia < 20 {
			signo = "Acuario"
		} else {
			signo = "Día inválido"
		}
	//MARZO
	case mes == 3:
		if dia >= 21 && dia <= 31 {
			signo = "Aries"
		} else if dia >= 1 && dia < 21 {
			signo = "Pisis"
		} else {
			signo = "Día inválido"
		}
	// ABRIL
	case mes == 4:
		if dia >= 21 && dia <= 30 {
			signo = "Tauro"
		} else if dia >= 1 && dia < 21 {
			signo = "Aries"
		} else {
			signo = "Día inválido"
		}
	// MAYO
	case mes == 5:
		if dia >= 22 && dia <= 31 {
			signo = "Géminis"
		} else if dia >= 1 && dia < 22 {
			signo = "Tauro"
		} else {
			signo = "Día inválido"
		}
	// JUNIO
	case mes == 6:
		if dia >= 22 && dia <= 30 {
			signo = "Cáncer"
		} else if dia >= 1 && dia < 22 {
			signo = "Géminis"
		} else {
			signo = "Día inválido"
		}
	// JULIO
	case mes == 5:
		if dia >= 24 && dia <= 31 {
			signo = "Leo"
		} else if dia >= 1 && dia < 24 {
			signo = "Cáncer"
		} else {
			signo = "Día inválido"
		}
	// AGOSTO
	case mes == 5:
		if dia >= 24 && dia <= 31 {
			signo = "Virgo"
		} else if dia >= 1 && dia < 24 {
			signo = "Leo"
		} else {
			signo = "Día inválido"
		}
	// SEPTIEMBRE
	case mes == 5:
		if dia >= 24 && dia <= 30 {
			signo = "Libra"
		} else if dia >= 1 && dia < 24 {
			signo = "Virgo"
		} else {
			signo = "Día inválido"
		}
	// OCTUBRE
	case mes == 5:
		if dia >= 24 && dia <= 31 {
			signo = "Escorpio"
		} else if dia >= 1 && dia < 24 {
			signo = "Libra"
		} else {
			signo = "Día inválido"
		}
	// NOVIEMBRE
	case mes == 5:
		if dia >= 23 && dia <= 30 {
			signo = "Sagitario"
		} else if dia >= 1 && dia < 23 {
			signo = "Escorpio"
		} else {
			signo = "Día inválido"
		}
	// DICIEMBRE
	case mes == 5:
		if dia >= 22 && dia <= 31 {
			signo = "Capricornio"
		} else if dia >= 1 && dia < 22 {
			signo = "Sagitario"
		} else {
			signo = "Día inválido"
		}

	default:
		signo = "Mes inválido"
	}

	fmt.Println(signo)
}
