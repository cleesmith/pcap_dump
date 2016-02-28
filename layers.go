package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: go run layers.go file.pcap\n")
	os.Exit(1)
}

func main() {
	flag.Parse()
	args := flag.Args()
	if flag.NArg() != 1 {
		usage()
	}

	log.Printf("layers: processing file '%s'\n", args[0])
	if len(args[0]) <= 0 {
		log.Printf("layers: invalid args[0] in: '-r %s'\n", args[0])
	}

	handle, err := pcap.OpenOffline(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packets := []gopacket.Packet{}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		packets = append(packets, packet)
	}

	for i, p := range packets {

		// see "type PacketMetadata struct" and "type CaptureInfo struct"
		// in github.com/google/gopacket/packet.go:
		log.Printf("\n%v. packet=%v:\n", i+1, p.Metadata().CaptureInfo.Timestamp.UTC())
		// log.Printf("\n%v. packet:\n%#v\n", i+1, p)

		// see issue: https://github.com/google/gopacket/issues/175
		log.Println(p.Dump()) // fails in Go 1.6, see https://golang.org/doc/go1.6#reflect
		// note: p.Dump output is the same as "hexdump -C some.file"
		// log.Println(p.String()) // fails in Go 1.6, see https://golang.org/doc/go1.6#reflect

		// Iterate over all layers, printing out each layer type
		// log.Println("All packet layers:")
		for _, layer := range p.Layers() {
			log.Println("- ", layer.LayerType())
		}

	} // for range packets

	log.Printf("packets count=%v\n", len(packets))
}
