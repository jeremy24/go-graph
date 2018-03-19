package GraphMatrix

//import "fmt"
import (
	"math/rand"
//	"fmt"
)

const (
	ON  	uint32 = 0x1
	OFF 	uint32 = 0x0
	ADD 	uint32 = 0x3
	REMOVE	uint32 = 0x4
	MARGIN	float32 = 5.0/100.0 
)


// this is the basic interface that is used for a graph edge graph-matrix
type GraphMatrix interface {
	Has(row int, col int) bool
	//Weight(i int32, j int32) float64
	Connect(rows int, cols int)
	//ConnectMany([]int, []int)
	Dims() (int, int)
	Remove(row int, col int)
	Weight(row int, col int) float32
	AddWeight(row int, col int, weight float32)
	Density() float32
	FillToDensity(density float32)
}

type bitMatrix struct {
	edges [][]uint32
	weights [][]float32
	nedges int
	nverts int
}


func IsWithinMargin(num float32) bool {
	upper := num + MARGIN
	lower := num + MARGIN
	return num < upper && num > lower
}

// make sure x <= y
func Order(x int, y int) (uint, uint) {
	if x < y {
		return uint(x), uint(y)
	}
	return uint(y), uint(x)
}

/*
	This will initialize all of the values in a new matrix for us.
	It only stores the values as a triangle matrix in order to save space
 */
func NewMatrix(nverts int) GraphMatrix {
	edges := make([][]uint32, nverts) // initialize a slice of dy slices

	for i := 0; i < nverts; i++ {
		edges[i] = make([]uint32, nverts - i + 1) // initialize a slice of dx unit8 in each of dy slices
	}

	weights := make([][]float32, nverts) // initialize a slice of dy slices

	for i := 0; i < nverts; i++ {
		weights[i] = make([]float32, nverts - i + 1) // initialize a slice of dx unit8 in each of dy slices
	}

	return &bitMatrix{edges, weights, 0, nverts}
}


func (g bitMatrix) FillToDensity(density float32) {

	for {
		x := rand.Int() % g.nverts
		y := rand.Int() % g.nverts
		g.Connect(x, y)
		if IsWithinMargin( g.Density() ) {
			return
		}
	}
}


func (g bitMatrix) Dims() (int, int) {
	return g.nverts, g.nverts
}

// Check if an edge exists
func (g bitMatrix) Has(i int, j int) bool {
	row, col := Order(i, j)
	bit := g.GetBit(uint(row), uint(col))
	ret := bit == true
	return ret
}

func (g *bitMatrix) Connect(i int, j int) {
	row, col := Order(i, j)

	if g.Has(int(row), int(col)) {
		return
	}
	g.SetBit(uint(row), uint(col), ON)
	g.nedges += 1
	//fmt.Println("Connected:", i, j, g.nedges)
}

func (g bitMatrix) Weight(i int, j int) float32 {
	row, col := Order(i, j)
	return g.weights[row][col]
}

func (g bitMatrix) AddWeight(i int, j int, weight float32) {
	row, col := Order(i, j)
	g.weights[row][col] = weight
}

func (g *bitMatrix) Remove(i int, j int) {
	row, col := Order(i, j)

	if g.Has(int(row), int(col)) {
		g.SetBit(uint(row), uint(col), OFF)
		g.nedges -= 1
	//	fmt.Println("Removed:", i, j, g.nedges)
	}

}




// Below here are internal functions




func (g bitMatrix) SetBit(i uint, j uint, flag uint32) {
	row := i
	col := uint32(j) / uint32(32)
	offset := uint32(j) % uint32(32)
	chunk := g.edges[row][col]

	// set the bit to zero
	if flag == OFF {
		g.edges[row][col] = chunk - (chunk & offset)
		return
	}

	// set the bit to one
	g.edges[row][col] = chunk | offset
}

// the return for this will ALWAYS be 0 or 1
func (g bitMatrix) GetBit(i uint, j uint) bool {
	row := i
	col := uint32(j) / uint32(32)
	offset := uint32(j) % uint32(32)
	chunk := g.edges[row][col]

	// get the bit and mask off the rest
	ret := (chunk & offset) == offset
	return ret
}

func AlterEdges(g *bitMatrix, action uint32) {
	if action == ADD {
		g.nedges += 1
	} else {
		g.nedges -= 1
	}
}

func (g bitMatrix) Density() float32 {
	E := float32(g.nedges)
	V := float32(g.nverts)
	//fmt.Println("E:",E, " V:", V)
	return E / (V * ( V-1))
}

