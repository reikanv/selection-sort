package datagen

import (
	"math/rand"
	"time"
)

type dataGenerator struct {
	randEngine *rand.Rand
	rangeCeil  int
}

func (dg *dataGenerator) genRandomInt() int {
	return dg.randEngine.Intn(dg.rangeCeil)
}

func (dg *dataGenerator) GenIntSlice(length int) []int {
	res := make([]int, length)

	for i := 0; i < length; i++ {
		res[i] = dg.genRandomInt()
	}

	return res
}

func NewDataGenerator(rangeCeil int) *dataGenerator {
	seed := time.Now().UnixNano()
	randEngine := rand.New(rand.NewSource(seed))

	return &dataGenerator{randEngine, rangeCeil}
}
