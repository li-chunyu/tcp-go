package tcpip

import (
	"net"
)

type arpstate uint32

const (
	ARP_FREE     arpstate = 0 // never used
	ARP_WAITING  arpstate = 1
	ARP_RESOLVED arpstate = 2

	ARP_REQUEST uint16 = 0x0001
	ARP_REPLY   uint16 = 0x0002

	ARP_ETHERNET uint16 = 0x0001
	ARP_IPV4     uint16 = 0x0800
)

type arp_hdr struct {
	hwtype uint16 // data link protocol type.
	protype uint16 // net layer type.
	hwsize byte // hardware address length.
	prosize byte // net layer address length.
	opcode uint16
	data []byte
}

type arp_ipv4 struct {
	smac net.HardwareAddr
	sip net.IP
	dmac net.HardwareAddr
	dip net.IP
}

func (hdr *arp_hdr) encode() []byte {
	b := make([]byte, 0)
	b = append(b, writeUint16ToByteBigEndian(hdr.hwtype)...)
	b = append(b, writeUint16ToByteBigEndian(hdr.protype)...)
	b = append(b, hdr.hwsize)
	b = append(b, hdr.prosize)
	b = append(b, writeUint16ToByteBigEndian(hdr.opcode)...)
	b = append(b, hdr.data...)
	return b
}

func (hdr *arp_hdr) decode(b []byte) {
	hdr.hwtype = readUint16FromByteBigEndian(b[0:2])
	hdr.protype = readUint16FromByteBigEndian(b[2:4])
	hdr.hwsize = b[4]
	hdr.prosize = b[5]
	hdr.opcode = readUint16FromByteBigEndian(b[6:8])
	hdr.data = b[8:]
}

