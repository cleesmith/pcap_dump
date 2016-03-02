package main

import (
	"fmt"
	"strings"
)

func main() {
	var timeSess, sessReqresp []string
	var s, timestamp, session, requestResponse string

	// https://github.com/simsong/tcpflow/blob/3cd14c8e350917a2f31733d037472aeebe74c7a8/src/tcpip.cpp#L237
	// tcpflow colors (ANSI escape codes/sequences):
	//                           green=?       blue=client   red=server
	// const char *color[3] = { "\033[0;32m", "\033[0;34m", "\033[0;31m" };
	// terminate color:
	// if (demux.opt.use_color) printf("\033[0m");

	s = "\x1b[0;34m2016-01-21T00:14:09Z082.165.177.154.00080-192.168.000.005.52751: HTTP/1.1 200 OK"
	// s = "zebraZmuff-ponies: spud"
	fmt.Printf("s[%v]=%#v\n", len(s), s)

	s = strings.Replace(s, "\x1b[0;34m", "____________ client/src:\n", 1) // blue=client
	s = strings.Replace(s, "\x1b[0;31m", "", 1)                           // red=server
	s = strings.Replace(s, "\x1b[0;32m", "", 1)                           // green=?
	s = strings.Replace(s, "\x1b[0m", "", 1)                              // terminate

	timeSess = strings.SplitAfterN(s, "Z", 2) // split after 1st "Z" only
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

	// fmt.Printf("%s\n", strings.Repeat("_", 79))
	fmt.Printf("%v  %v\n", timestamp, session)
	fmt.Printf("%v\n", requestResponse)
}
