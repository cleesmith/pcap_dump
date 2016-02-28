package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

func main() {
	var b bytes.Buffer
	res := []byte{0x0f, 0x25, 0x18, 0xd7, 0xa9, 0x40, 0xad, 0xe2}
	fmt.Printf("res=%T\n", res)
	fmt.Fprint(&b, hex.Dump(res))
	fmt.Println(b.String())
}
