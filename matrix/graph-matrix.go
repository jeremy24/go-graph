package main

import "fmt"

// this is the basic interface that is used for a graph edge matrix
type GraphMatrix interface {
	Has(i int, j int) bool
	//Weight(i int32, j int32) float64
	Connect(r int, c int)
	Dims() (int, int)
}

type bitMatrix struct {
	edges [][]uint32
	rows  int
	cols  int
}

func NewMatrix(rows int, cols int) GraphMatrix {
	edges := make([][]uint32, rows) // initialize a slice of dy slices

	for i := 0; i < rows; i++ {
		edges[i] = make([]uint32, cols) // initialize a slice of dx unit8 in each of dy slices
	}

	return bitMatrix{edges, rows, cols}
}

func (g bitMatrix) Dims() (int, int) {
	return g.rows, g.cols
}

// Check if an edge exists
func (g bitMatrix) Has(i int, j int) bool {
	bit := g.GetBit(uint(i), uint(j))
	ret := bit == true
	return ret
}

func (g bitMatrix) Connect(i int, j int) {
	g.SetBit(uint(i), uint(j))
}

func (g bitMatrix) SetBit(i uint, j uint) {
	row := i
	var col uint32 = uint32(j) / uint32(32)
	var offset uint32 = uint32(j) % uint32(32)
	chunk := g.edges[row][col]

	// set the offset bit to one
	g.edges[row][col] = chunk | offset
}

// the return for this will ALWAYS be 0 or 1
func (g bitMatrix) GetBit(i uint, j uint) bool {
	row := i
	var col uint32 = uint32(j) / uint32(32)
	var offset uint32 = uint32(j) % uint32(32)
	chunk := g.edges[row][col]

	// get the bit and mask off the rest
	ret := (chunk & offset) == offset
	return ret
}

func main() {
	g := NewMatrix(100, 100)
	r, c := g.Dims()
	fmt.Printf("Dims: (%d. %d)\n", r, c)
	g.Connect(1, 1)
	fmt.Println("E(1, 1): ", g.Has(1, 1))
	fmt.Println("E(2, 1): ", g.Has(2, 1))

	g.Connect(40, 40)
	fmt.Println("E(40, 40): ", g.Has(40, 40))
}
