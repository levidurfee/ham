package main

import (
	"encoding/base64"
	"fmt"

	"github.com/gorilla/securecookie"
)

func main() {
	k := securecookie.GenerateRandomKey(32)
	b := base64.StdEncoding.EncodeToString(k)
	fmt.Println(b)
}
