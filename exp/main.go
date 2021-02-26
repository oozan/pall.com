package main

import (
	"fmt"

	"pall.com/hash"
)

func main() {
	hmac := hash.NewHMAC("secret key")
	fmt.Println(hmac.Hash("secret key 2"))
}
