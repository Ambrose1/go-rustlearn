package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	// 设置参数
	device := "en0"
	snapshotLen := int32(65535)
	promiscuous := false
	timeout := pcap.BlockForever

	// 打开网络接口
	handle, err := pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// 设置过滤器
	filter := "tcp and port 80"
	if err := handle.SetBPFFilter(filter); err != nil {
		log.Fatal(err)
	}

	// 开始抓包
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
