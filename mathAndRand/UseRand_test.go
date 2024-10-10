package mathAndRand

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_rand(t *testing.T) {
	rand.Intn(12)
	fmt.Println()
}
