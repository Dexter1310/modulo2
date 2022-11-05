package modul2

import (
	"fmt"
	"strconv"
	"strings"
)

func Sumar(expresion string) int {
	valores := strings.Split(expresion, "+")
	var result int

	for _, valor := range valores {

		num, error := strconv.Atoi(valor) // convierte de string a int
		fmt.Println(error)
		if error != nil { // cuando no hay errores devuelve nil
		
			fmt.Println("Error: Tiene que ingresar un número entero")
			fmt.Println("o !Solo debes realizar con operador +!!")
		} else {
			result += num
		}
	}

	return result
}

func SumaNumero(){
	var expresion string
	var result int
	
	fmt.Print("AÑADE UNA SUMA ejemplo (4+4)=> ")
	fmt.Scanln(&expresion)
	
	result = Sumar(expresion)
	
	fmt.Printf("=> %d\n", result)
}


