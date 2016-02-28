package main

import (
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	for _, file := range []struct {
		filename       string
		num            int
		expectedLayers []gopacket.LayerType
	}{
		{"test_loopback.pcap",
			24,
			[]gopacket.LayerType{
				layers.LayerTypeLoopback,
				layers.LayerTypeIPv6,
				layers.LayerTypeTCP,
			},
		},
		{"test_ethernet.pcap",
			16,
			[]gopacket.LayerType{
				layers.LayerTypeEthernet,
				layers.LayerTypeIPv4,
				layers.LayerTypeTCP,
			},
		},
		{"test_dns.pcap",
			10,
			[]gopacket.LayerType{
				layers.LayerTypeEthernet,
				layers.LayerTypeIPv4,
				layers.LayerTypeUDP,
				layers.LayerTypeDNS,
			},
		},
	} {
		log.Printf("Processing file %s\n", file.filename)

		packets := []gopacket.Packet{}
		handle, err := pcap.OpenOffline(file.filename)
		if err != nil {
			log.Fatal(err)
		} else {
			packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
			for packet := range packetSource.Packets() {
				packets = append(packets, packet)
			}
		}
		log.Printf("packets: got=%v want=%v\n", len(packets), file.num)
		if len(packets) != file.num {
			log.Fatal("Incorrect number of packets, want", file.num, "got", len(packets))
		}

		for i, p := range packets {

			// see "type PacketMetadata struct" and "type CaptureInfo struct"
			// in github.com/google/gopacket/packet.go:
			log.Printf("\n%v. packet=%v:\n", i+1, p.Metadata().CaptureInfo.Timestamp.UTC())
			// log.Printf("\n%v. packet:\n%#v\n", i+1, p)
			// log.Printf("file.expectedLayers=%T=%#v\n", file.expectedLayers, file.expectedLayers)
			// see issue: https://github.com/google/gopacket/issues/175
			// log.Printf(p.Dump()) // fails in Go 1.6, see https://golang.org/doc/go1.6#reflect
			// log.Printf(p.String()) // fails in Go 1.6, see https://golang.org/doc/go1.6#reflect

			// Iterate over all layers, printing out each layer type
			log.Println("All packet layers:")
			for _, layer := range p.Layers() {
				log.Println("- ", layer.LayerType())
			}

			// for _, layertype := range file.expectedLayers {
			// 	if p.Layer(layertype) == nil {
			// 		log.Fatal("Packet", i, "has no layer type\n%s", layertype, p.Dump())
			// 	}
			// }

		} // for range packets
	}
}
