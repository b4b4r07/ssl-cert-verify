package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"time"
)

var (
	showDays = flag.Bool("show-days", false, "Show only remaining days")
)

func main() {
	flag.Parse()
	config := tls.Config{InsecureSkipVerify: false}
	for _, site := range flag.Args() {
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
		days := cert.NotAfter.Unix() - time.Now().Unix()
		days /= 24 * 60 * 60
		if days < 100 {
			if *showDays {
				fmt.Println(days)
			} else {
				fmt.Printf("%s Certificate expires in %d days at %s\n", site, days, cert.NotAfter)
			}
		}
	}
}
