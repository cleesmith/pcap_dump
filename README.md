***

#### Feb 27, 2016:
```
https://github.com/cleesmith/pcap_dump
Go 1.6
go get github.com/google/gopacket
godep save
  - what does godep do?
    - looks at the "import"s in each ".go" file
    - makes a copy of the "go get" packages found
    - our "vendor" folder
    - it removes the ".git" folder from each package
    - creates the "Godeps/Godeps.json" listing

go run cls_readpcap.go
go run httpassembly.go -r example.com-4.pcap
```

***

#### Feb 4, 2016:
```
tcpflow is kind of like CapMe or CapAnalysis
? limited to tcp sessions ?

brew install tcpflow

tcpflow -h
tcpflow -hh ... more details
http://www.circlemud.org/jelson/software/tcpflow/tcpflow.1.html
https://www.mankier.com/1/tcpflow
https://tournasdimitrios1.wordpress.com/2011/02/21/tcpflow-a-tcp-flow-recorder/

... pcaps:
http://chrissanders.org/packet-captures/

... usages:
http://kalilinuxtutorials.com/tcpflow/
https://sesblog.amazon.com/blog/tag/TCP+Flow
http://simson.net/ref/2013/2013-12-05_tcpflow-and-BE-update.pdf

tcpflow -c -g -T %t_%A.%a-%B.%b%V%v%C%c_ -X /dev/null -r http_espn.pcap
... timestamp epochSeconds_srcIP.srcPort-dstIP.dstPortvlanNumconnectionCount
tcpflow -c -g -T %t_%A.%a-%B.%b%V%v%C%c_ -X /dev/null -r http_espn.pcap > tcp.out
... show http requests to host:
tcpflow -c -r http_espn.pcap | grep -oE '(GET|POST|HEAD) .* HTTP/1.[01]|Host: .*'
or with a specified host expression:
tcpflow -c -g -r http_espn.pcap 'host 192.168.146.131' | grep -oE '(GET|POST|HEAD) .* HTTP/1.[01]|Host: .*'

... this yielded an empty html page:
tcpflow -a -o tcpflows -r http_espn.pcap

... suppress report.xml use -X /dev/null
tcpflow -c -g -FT -X /dev/null -r http_espn.pcap > tcpflow.out
```

***
***
