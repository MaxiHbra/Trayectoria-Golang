package main
import "fmt"

func hola (s string) string {
	return "Hola " + s
}
func resta_dos (a int) int {
	return (a-2)
}
func esPar(x int) string{
	if x==0 {
		return "Es par"
}	else{
		return esImpar(x-1)
}}
func esImpar(x int) string{
	if x==0 {
		return "Es impar"
}	else{
		return esPar(x-1)
}}
func main() {
	fmt.Println(hola("Masi"))
	fmt.Println("Si restamos 2 a 5= ", resta_dos(5))
	fmt.Println("El numero 8 ",esPar(8))
}

