package sha256

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"log"
	"math/bits"
)

type HashType int

const (
	SHA256 HashType = iota
	SHA384
	SHA512
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

func Hash(in io.Reader, t HashType) {
	var b bytes.Buffer
	_, err := b.ReadFrom(in)
	if err != nil {
		log.Fatalf("Reading from %q failed with %q", in, err)
	}
	if t == SHA256 {
		fmt.Printf("%x\n", sha256.Sum256(b.Bytes()))
	} else if t == SHA384 {
		fmt.Printf("%x\n", sha512.Sum384(b.Bytes()))
	} else {
		fmt.Printf("%x\n", sha512.Sum512(b.Bytes()))
	}
}
