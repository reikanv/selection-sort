package main

import (
	"fmt"
	"math"

	"github.com/reikanv/selection-sort/internal/datagen"
	"github.com/reikanv/selection-sort/internal/sorts"
	strpad "github.com/reikanv/selection-sort/pkg"
)

type result struct {
	datasetNum    int
	datasetLength int
	compareCount  int
	shiftCount    int
	compareLog    float64
	compareNorm   float64
	shiftLog      float64
	shiftNorm     float64
}

func run(datasets [][]int, sorterFunc func(data *[]int) (int, int)) {
	res := make([]result, len(datasets))
	for i := 0; i < len(datasets); i++ {
		datasetLength := len(datasets[i])
		datasetLengthF64 := float64(datasetLength)
		item := make([]int, datasetLength)
		copy(item, datasets[i])
		compareCount, shiftCount := sorterFunc(&item)
		compareCountF64 := float64(compareCount)
		shiftCountF64 := float64(shiftCount)

		res[i] = result{
			datasetNum:    i + 1,
			compareCount:  compareCount,
			compareLog:    math.Log(compareCountF64),
			compareNorm:   compareCountF64 / datasetLengthF64,
			shiftCount:    shiftCount,
			shiftLog:      math.Log(shiftCountF64),
			shiftNorm:     shiftCountF64 / datasetLengthF64,
			datasetLength: datasetLength,
		}
	}

	fmt.Print("dataset # | ")
	for _, r := range res {
		fmt.Printf(strpad.Left("%v |", 14), r.datasetNum)
	}
	fmt.Print("\n")

	fmt.Print("length    | ")
	for _, r := range res {
		fmt.Printf(strpad.Left("%v |", 14-len(fmt.Sprint(r.datasetLength))+1), r.datasetLength)
	}
	fmt.Print("\n")

	fmt.Print("compare   | ")
	for _, r := range res {
		fmt.Printf(strpad.Left("%v |", 14-len(fmt.Sprint(r.compareCount))+1), r.compareCount)
	}
	fmt.Print("\n")

	fmt.Print("shift     | ")
	for _, r := range res {
		fmt.Printf(strpad.Left("%v |", 14-len(fmt.Sprint(r.shiftCount))+1), r.shiftCount)
	}
	fmt.Print("\n")

	fmt.Println("norm values")

	fmt.Print("compare   | ")
	for _, r := range res {
		fmt.Printf(strpad.Left("%v |", 14-len(strpad.TrimFloat(r.compareNorm))+1), strpad.TrimFloat(r.compareNorm))
	}
	fmt.Print("\n")

	fmt.Print("shift     | ")
	for _, r := range res {
		fmt.Printf(strpad.Left("%v |", 14-len(strpad.TrimFloat(r.shiftNorm))+1), strpad.TrimFloat(r.shiftNorm))
	}
	fmt.Print("\n")

	fmt.Println("log values")

	fmt.Print("compare   | ")
	for _, r := range res {
		fmt.Printf(strpad.Left("%v |", 14-len(strpad.TrimFloat(r.compareLog))+1), strpad.TrimFloat(r.compareLog))
	}
	fmt.Print("\n")

	fmt.Print("shift     | ")
	for _, r := range res {
		fmt.Printf(strpad.Left("%v |", 14-len(strpad.TrimFloat(r.shiftLog))+1), strpad.TrimFloat(r.shiftLog))
	}
	fmt.Print("\n")
}

func main() {
	gen := datagen.NewDataGenerator(100000)
	dataset1 := gen.GenIntSlice(26)
	dataset2 := gen.GenIntSlice(52)
	dataset3 := gen.GenIntSlice(104)
	dataset4 := gen.GenIntSlice(208)
	dataset5 := gen.GenIntSlice(416)
	dataset6 := gen.GenIntSlice(832)
	dataset7 := gen.GenIntSlice(1664)

	datasets := [][]int{dataset1, dataset2, dataset3, dataset4, dataset5, dataset6, dataset7}

	fmt.Println("bubble sort:")
	run(datasets, sorts.BubbleSort)

	fmt.Print("\n")

	fmt.Println("selection sort:")
	run(datasets, sorts.SelectionSort)
}
