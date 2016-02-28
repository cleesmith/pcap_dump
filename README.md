### Goals
* explore [osquery](https://osquery.io/)
* deeper understanding of gopacket and OSI layers; especially network, transport, and session layers
* fix ```gopacket```'s packet dump (hex/human)
  * so [unifiedbeat](https://github.com/cleesmith/unifiedbeat) works with Go 1.6 and older
* replace Sguil's and CapME's transcript (session) feature
  * do this using gopacket
    * instead of external programs such as tcpdump and tcpflow
    * so there is just one binary to copy(install) onto servers

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

##### on a virtualbox vm named Go do:
```
lsb_release -a
    No LSB modules are available.
    Distributor ID: Ubuntu
    Description:  Ubuntu 14.04.3 LTS
    Release:  14.04
    Codename: trusty
sudo apt-get install libpcap-dev
    - http://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket
    - http://askubuntu.com/questions/436203/how-do-i-get-the-pcap-lib-on-ubuntu
    - compiling gopacket needs pcap dev headers like "pcap.h"
go version --> go version go1.5.1 linux/amd64
    - go versions prior to 1.6 work with gopacket, coz
       there's no reflect/struct/interface bug fix, see:
        https://golang.org/doc/go1.6#reflect
        https://github.com/golang/go/issues/12367
git clone https://github.com/cleesmith/golang_learning.git
cd /home/cleesmith/go/golang_learning/gopacket
go get github.com/google/gopacket
go run pcapdump.go -r test_ethernet.pcap
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
