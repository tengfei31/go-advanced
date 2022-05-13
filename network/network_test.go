package network

import (
	"log"
	"testing"
)

func TestTcp(t *testing.T) {
	err := startTcp()
	if err != nil {
		log.Fatal(err)
	}
}


func TestUdp(t *testing.T) {
	err := startUdp()
	if err != nil {
		log.Fatal(err)
	}
}
