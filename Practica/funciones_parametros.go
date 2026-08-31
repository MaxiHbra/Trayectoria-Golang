package main
import "fmt"

func hola (s string) string {
	return "Hola " + s
}
func resta_dos (a int) int {
	return (a-2)
}

func main() {
	fmt.Println(hola("Masi"))
	fmt.Println("Si restamos 2 a 5= ", resta_dos(5))
}
