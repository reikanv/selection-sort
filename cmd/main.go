package main

import (
	"fmt"

	"github.com/reikanv/selection-sort/internal/datagen"
	"github.com/reikanv/selection-sort/internal/sorts"
)

func run(datasets [][]int, sorterFunc func(data *[]int) (int, int)) {
	for i := 0; i < len(datasets); i++ {
		item := make([]int, len(datasets[i]))
		copy(item, datasets[i])
		compareCount, shiftCount := sorterFunc(&item)
		fmt.Printf("dataset %v: len %v; compare %v; shift %v\n", i+1, len(item), compareCount, shiftCount)
	}
}

func main() {
	gen := datagen.NewDataGenerator(100000)
	dataset1 := gen.GenIntSlice(26)
	dataset2 := gen.GenIntSlice(52)
	dataset3 := gen.GenIntSlice(104)
	dataset4 := gen.GenIntSlice(208)
	dataset5 := gen.GenIntSlice(416)
	dataset6 := gen.GenIntSlice(832)
	dataset7 := gen.GenIntSlice(1663)

	datasets := [][]int{dataset1, dataset2, dataset3, dataset4, dataset5, dataset6, dataset7}

	fmt.Println("bubble sort")
	run(datasets, sorts.BubbleSort)

	fmt.Println("selection sort")
	run(datasets, sorts.SelectionSort)
}
