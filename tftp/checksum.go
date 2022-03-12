package main

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
)

func checksum(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("%x", sha512.Sum512_256(b))
}
