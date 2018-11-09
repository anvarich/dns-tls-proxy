package main

import (
	"crypto/tls"
	"fmt"
	"log"

	//"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/miekg/dns"
)

//
//func dnsUdpHandler(w dns.ResponseWriter, m *dns.Msg) {
//	dnsHandler(w, m, "udp")
//}
//
//func dnsTcpHandler(w dns.ResponseWriter, m *dns.Msg) {
//	dnsHandler(w, m, "tcp")
//}

func main() {
	//dns.HandleFunc()
	go serve("tcp", dnsHandler)
	go serve("udp", dnsHandler)
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig
	fmt.Printf("Signal (%s) received, stopping\n", s)
}

func serve(net string, handle dns.HandlerFunc) {
	server := &dns.Server{Addr: ":8053", Net: net, Handler: dns.HandlerFunc(handle)}
	log.Printf("Listening on " + net + "8053")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Failed to setup the "+net+" server: %s\n", err.Error())
	}
}
func dnsHandler(w dns.ResponseWriter, m *dns.Msg) {

	fmt.Printf("%s\n", m.Question[0].Name)
	log.Printf("Question recieved " + m.Question[0].Name + ".")
	m.Question[0].Name = strings.ToUpper(m.Question[0].Name)
	c := new(dns.Client)
	c.Net = "tcp-tls"
	c.TLSConfig = &tls.Config{
		InsecureSkipVerify: false,
	}
	r, _, err := c.Exchange(m, "1.1.1.1:853")
	if err != nil {
		fmt.Printf("failed to exchange adreess: %s\n", err.Error())
	}
	if r == nil {
		fmt.Printf("no response from server ")
	}
	r.Question[0].Name = strings.ToLower(r.Question[0].Name)
	for i := 0; i < len(r.Answer); i++ {
		r.Answer[i].Header().Name = strings.ToLower(r.Answer[i].Header().Name)
	}
	//log.Printf("%s\n", r)
	w.WriteMsg(r)
}
