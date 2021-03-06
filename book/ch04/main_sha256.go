package main

import (
	"github.com/kbl/gopher_exercises/book/ch04/my_sha256"
	"os"
)

func main() {
	var param string
	if len(os.Args) > 1 {
		param = os.Args[1]
	}
	hash_type := my_sha256.SHA256
	if param == "--384" {
		hash_type = my_sha256.SHA384
	} else if param == "--512" {
		hash_type = my_sha256.SHA512
	}

	my_sha256.Hash(os.Stdin, hash_type)
}
