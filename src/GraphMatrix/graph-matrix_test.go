package GraphMatrix

import (
	"testing"
	"math/rand"
	"math"
//	"fmt"
)


func GetMatrix(nverts int) GraphMatrix {
	mat := NewMatrix(nverts)
	return mat
}


func TestConnect(t *testing.T) {
	n := 15000
	mat := GetMatrix(n)


	for i := 0 ; i < int(math.Sqrt(float64(n))) ; i++ {
		x := rand.Int() % n
		y := rand.Int() % n
		mat.Connect(x, y)
		if mat.Has(x, y) != true {
			t.Fatalf("(%d, %d) was not connected right\n", x, y)
		}
	}
}

// func TestRemove(t *testing.T) {
// 	n := 10
// 	mat := GetMatrix(n)

// 	for i := 0 ; i < int(math.Sqrt(float64(n))) ; i++ {
// 		x := rand.Int() % n
// 		y := rand.Int() % n
// 		mat.Connect(x, y)
// 		if mat.Has(x, y) != true {
// 			t.Fatalf("(%d, %d) was not connected right\n", x, y)
// 		}
// 		mat.Remove(x, y)
// 		if mat.Has(x, y) != false {
// 			t.Fatalf("(%d, %d) on iter %d was not removed right\n", x, y, i)
// 		}
// 	}
// }


func TestDensity( t *testing.T) {
	n := 2
	mat := GetMatrix(n)
	mat.Connect(0,1)
	if mat.Density() != 0.5 {
		t.Fatalf("Density is not right, want .5 got:", mat.Density())
	}
}

// func TestWeight(t *testing.T) {
// 	n := 15000
// 	mat := GetMatrix(n)

// 	for i := 0 ; i < int(math.Sqrt(float64(n))) ; i++ {
// 		x := rand.Int() % n
// 		y := rand.Int() % n
// 		w := rand.Float32()
// 		mat.Connect(x, y)
// 		mat.AddWeight(x, y, w)
// 		if mat.Has(x, y) != true {
// 			t.Fatalf("(%d, %d) was not connected right\n", x, y)
// 		}
// 		//if mat.Weight(x, y) != w {
// 		//	t.Fatalf("(%d, %d) was not weighted right\n", x, y)
// 		//}
// 	}
// }