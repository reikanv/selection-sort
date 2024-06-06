package main

import (
	"flag"
	"fmt"

	"github.com/reikanv/selection-sort/internal/datagen"
	"github.com/reikanv/selection-sort/internal/sorts"
)

type flags struct {
	alg  string
	size int
	ceil int
}

func main() {
	f := flags{}
	flag.IntVar(&f.size, "size", 1, "size of initial unsorted slice")
	flag.IntVar(&f.ceil, "ceil", 1, "maximum value of generated int")
	flag.StringVar(&f.alg, "alg", "bubble", "algorhitm to use")
	flag.Parse()

	if f.size < 1 {
		msg := fmt.Sprintf("invalid size %v can't be more less than 1", f.size)
		panic(msg)
	}

	if f.alg != "bubble" && f.alg != "selection" {
		msg := fmt.Sprintf("invalid algorhitm %v available values: bubble, selection. bubble is default", f.size)
		panic(msg)
	}

	gen := datagen.NewDataGenerator(f.ceil)
	dataset := gen.GenIntSlice(f.size)

	fmt.Println("original slice:")
	fmt.Println(dataset)

	switch f.alg {
	case "bubble":
		sorts.BubbleSort(&dataset)
	case "selection":
		sorts.SelectionSort(&dataset)
	default:
		sorts.BubbleSort(&dataset)
	}

	fmt.Println("sorted slice:")
	fmt.Println(dataset)
}
