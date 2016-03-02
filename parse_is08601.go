package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	s := "\x1b[0;34m2016-01-21T00:14:09Z082.165.177.154.00080-192.168.000.005.52751: HTTP/1.1 200 OK"
	s = strings.Replace(s, "\x1b[0;34m", "", 1)
	//t, err := time.Parse(time.RFC3339, "2016-01-21T00:14:09Z")
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		parseError, ok := err.(*time.ParseError)
		fmt.Printf("ok=%v\nparseError.ValueElem=%#v\n", ok, parseError.ValueElem)
		fmt.Printf("t=%v\nerr=%T=%#v\n", t, err, err)
		//fmt.Printf("err.ValueElem=%T=%#v\n", err, err.ValueElem)
		return
	}
	fmt.Println(t.UTC().Format(time.RFC3339))
}
