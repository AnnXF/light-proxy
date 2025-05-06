package cmd

import (
	"github.com/up-zero/gotool/netutil"
	"io"
	"log"
	"strings"
)

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

func showLocalIpv4s() {
	ips, err := netutil.Ipv4sLocal()
	if err == nil {
		log.Printf("[sys] local ipv4: %s \n", strings.Join(ips, ";"))
	}
}
