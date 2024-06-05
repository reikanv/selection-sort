package main

import (
	"fmt"

	"github.com/reikanv/selection-sort/internal/datagen"
	"github.com/reikanv/selection-sort/internal/sorts"
)

func main() {
	gen := datagen.NewDataGenerator(100000)
	dataset1 := gen.GenIntSlice(26)
	dataset2 := gen.GenIntSlice(52)
	dataset3 := gen.GenIntSlice(104)
	dataset4 := gen.GenIntSlice(208)
	dataset5 := gen.GenIntSlice(416)
	dataset6 := gen.GenIntSlice(832)
	dataset7 := gen.GenIntSlice(1663)

	compareCount1, shiftCount1 := sorts.BubbleSort(&dataset1)
	compareCount2, shiftCount2 := sorts.BubbleSort(&dataset2)
	compareCount3, shiftCount3 := sorts.BubbleSort(&dataset3)
	compareCount4, shiftCount4 := sorts.BubbleSort(&dataset4)
	compareCount5, shiftCount5 := sorts.BubbleSort(&dataset5)
	compareCount6, shiftCount6 := sorts.BubbleSort(&dataset6)
	compareCount7, shiftCount7 := sorts.BubbleSort(&dataset7)

	fmt.Printf(
		"ds1: len %v; compare %v; shift %v\n",
		len(dataset1),
		compareCount1,
		shiftCount1,
	)

	fmt.Printf(
		"ds2: len %v; compare %v; shift %v\n",
		len(dataset2),
		compareCount2,
		shiftCount2,
	)

	fmt.Printf(
		"ds3: len %v; compare %v; shift %v\n",
		len(dataset3),
		compareCount3,
		shiftCount3,
	)

	fmt.Printf(
		"ds4: len %v; compare %v; shift %v\n",
		len(dataset4),
		compareCount4,
		shiftCount4,
	)

	fmt.Printf(
		"ds5: len %v; compare %v; shift %v\n",
		len(dataset5),
		compareCount5,
		shiftCount5,
	)

	fmt.Printf(
		"ds6: len %v; compare %v; shift %v\n",
		len(dataset6),
		compareCount6,
		shiftCount6,
	)

	fmt.Printf(
		"ds7: len %v; compare %v; shift %v\n",
		len(dataset7),
		compareCount7,
		shiftCount7,
	)
}
