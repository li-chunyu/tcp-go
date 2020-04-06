package tcpip

import (
	"net"
)

const (
	ETH_P_ARP uint16 = 0x0806
	ETH_P_IP  uint16 = 0x0800
)

type eth_hdr struct {
	dmac      net.HardwareAddr // 6 bytes alias for []byte
	smac      net.HardwareAddr // 6 bytes
	ethertype uint16
	playload  []byte
}

func (hdr *eth_hdr) encode() []byte {
	b := make([]byte, 0)
	b = append(b, hdr.dmac...)
	b = append(b, hdr.smac...)
	b = append(b, writeUint16ToByteBigEndian(hdr.ethertype)...)
	b = append(b, hdr.playload...)
	return b
}

func (hdr *eth_hdr) decode(b []byte) {
	hdr.dmac = b[0:6]
	hdr.smac = b[6:12]
	hdr.ethertype = readUint16FromByteBigEndian(b[12:14])
	hdr.playload = b[14:]
}
