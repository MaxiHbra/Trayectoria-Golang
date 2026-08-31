package main
import "fmt"

func main(){
	//var edad int
	//var nombre string
	//var valor bool
	var ciudad string = "Salta"
	var temperatura float32 = 23.5
	var temperatura2 float64 = 25.67
	var numeros [5]int
	numeros= [5]int{1 ,2, 3, 4, 5}
	//Pero gracias a la inferencia de tipos puedo escribir una variable sin incluir su tipo ya que la estoy asignando
	//pais := "Argentina"
	fmt.Println(ciudad)
	fmt.Println(temperatura)
	fmt.Println(temperatura2)
	fmt.Println(numeros[2])
}
