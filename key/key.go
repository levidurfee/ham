package main

import (
	"encoding/base64"
	"fmt"

	"github.com/gorilla/securecookie"
)

func main() {
	fmt.Println("Starting")

	k := securecookie.GenerateRandomKey(32)
	ke := base64.StdEncoding.EncodeToString(k)
	fmt.Println(ke)
}
