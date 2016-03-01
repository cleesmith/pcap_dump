package main

import (
	"fmt"
	"strings"
)

func main() {
	var timeSess, sessReqresp []string
	var s, timestamp, session, requestResponse string

	s = "2016-01-21T00:14:09Z082.165.177.154.00080-192.168.000.005.52751: HTTP/1.1 200 OK"
	// s = "zebraZmuff-ponies: spud"

	timeSess = strings.SplitAfter(s, "Z")
	fmt.Printf("timeSess[%v]=%#v\n", len(timeSess), timeSess)
	if len(timeSess) > 1 {
		timestamp = timeSess[0]

		sessReqresp = strings.Split(timeSess[1], ": ")
		session = sessReqresp[0]
		session = strings.Replace(session, "-", " -> ", 1)
		if len(sessReqresp) > 1 {
			requestResponse = sessReqresp[1]
		}
	} else {
		// wtf? just use it as-is:
		timestamp = s
	}

	fmt.Printf("%s\n", strings.Repeat("_", 79))
	fmt.Printf("%v  %v\n", timestamp, session)
	fmt.Printf("%v\n", requestResponse)
}
