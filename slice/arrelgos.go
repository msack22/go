package slice

import "fmt"

var tabla [10]int = [10]int{10, 0, 59}

func MuestroArreglos() {
	tabla[7] = 44
	tabla[9] = 33
	fmt.Println(tabla)

	tabla2 := [10]string{"pablo", "pepe"}
	fmt.Println(tabla2)

}
