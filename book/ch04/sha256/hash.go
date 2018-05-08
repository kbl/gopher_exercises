package sha256

import (
	"crypto/sha256"
	"fmt"
	"math/bits"
)

func DifferentBits(s1, s2 string) int {
	sum1 := sha256.Sum256([]byte(s1))
	sum2 := sha256.Sum256([]byte(s2))
	var commonBits int
	for i, b1 := range sum1 {
		b2 := sum2[i]
		xorBits := b1 ^ b2
		commonBits += bits.OnesCount(uint(xorBits))
	}
	return commonBits
}
