package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	config := tls.Config{InsecureSkipVerify: false}
	for _, site := range os.Args[1:] {
		expireCheck(site, config)
	}
}

func expireCheck(site string, config tls.Config) {
	conn, err := tls.Dial("tcp", site, &config)
	if err != nil {
		log.Fatalf("%s: %s", site, err)
	}

	defer conn.Close()
	state := conn.ConnectionState()
	for _, cert := range state.PeerCertificates {
		diff := cert.NotAfter.Unix() - time.Now().Unix()
		diff /= 24 * 60 * 60
		if diff < 100 {
			fmt.Printf("%s Certificate expires in %d days at %s\n", site, diff, cert.NotAfter)
		}
	}
}
