// Ping.go
package main

import (
	// "bytes"
	"fmt"
	// "io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage :", os.Args[0], "host")
		os.Exit(1)
	}
	addr, err := net.ResolveIPAddr("ip", os.Args[1])
	checkError(err)
	conn, err := net.DialIP("ip4:icmp", addr, addr)
	checkError(err)
	var msg [512]byte
	msg[0] = 8
	msg[2] = 0
	msg[3] = 0
	msg[4] = 0
	msg[5] = 13
	msg[6] = 0
	msg[7] = 37
	len := 8
	check := checkSum(msg[0:len])
	fmt.Println("check : ", check)
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)
	_, err = conn.Write(msg[0:len])
	checkError(err)
	_, err = conn.Read(msg[0:])
	checkError(err)
	fmt.Println("Get response")
	if msg[5] == 13 {
		fmt.Println("identifier matches")
	}
	if msg[7] == 37 {
		fmt.Println("Sequence matches")
	}
}

func checkSum(msg []byte) uint16 {
	sum := 0
	for i := 0; i < len(msg)-1; i += 2 {
		sum += int(msg[i])*256 + int(msg[i+1])
	}
	sum = (sum >> 16) + sum&0xffff
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
