
package main

import (
	"fmt"
	"GraphMatrix"
)



func main() {
	mat := GraphMatrix.NewMatrix(10)
	
	mat.Connect(1, 1)
	fmt.Println("Has (1, 1):", mat.Has(1, 1))
}
