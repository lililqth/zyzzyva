package conf

import (
	"crypto/rand"
	"fmt"
	"time"
)

var F = 3
var N = 3*F + 1
var M = 1

var KeySize = 2048
var ClientTimeout = 10 * time.Second
var CPInterval = 100

var Extra = []byte("extra")

var UdpMulticastAddr = "224.0.0.1:10000"
var UdpMulticastInterface = ""

func GetReqAddr(id int) string {
	p := 0
	if id < N {
		p = 20000 + id
	} else {
		p = 30000 + id
	}
	return fmt.Sprintf("127.0.0.1:%d", p)
}

func GetListenAddr(id int) string {
	p := 0
	if id < N {
		p = 20000 + id
	} else {
		p = 30000 + id
	}
	return fmt.Sprintf(":%d", p)
}

var RandInputSize = 64

func GetRandInput() ([]byte, error) {
	in := make([]byte, RandInputSize)
	_, err := rand.Read(in)
	if err != nil {
		return nil, err
	}

	return in, nil
}
