//go:build windows

package cmd

import (
	"golang.org/x/sys/windows/registry"
	"log"
	"net/url"
)

func configClear() {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.SET_VALUE)
	if err != nil {
		log.Fatalf("[sys] registry open error:%s \n", err.Error())
	}
	defer key.Close()

	if err = key.SetDWordValue("ProxyEnable", 0x0); err != nil {
		log.Fatalf("[sys] registry set error:%s \n", err.Error())
	}

	println("[sys] proxy config clear successfully")
}

func configSet(domain string) {
	host := ""
	u, err := url.Parse(domain)
	if err != nil {
		host = domain
	} else {
		host = u.Host
	}

	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.SET_VALUE)
	if err != nil {
		log.Fatalf("[sys] registry open error:%s \n", err.Error())
	}
	defer key.Close()

	if err = key.SetDWordValue("ProxyEnable", 0x1); err != nil {
		log.Fatalf("[sys] registry set proxy enable error:%s \n", err.Error())
	}
	if err = key.SetStringValue("ProxyServer", host); err != nil {
		log.Fatalf("[sys] registry set proxy server error:%s \n", err.Error())
	}
	if err = key.SetStringValue("ProxyOverride", `localhost;127.*;10.*;172.16.*;172.17.*;172.18.*;172.19.*;172.20.*;172.21.*;172.22.*;172.23.*;172.24.*;172.25.*;172.26.*;172.27.*;172.28.*;172.29.*;172.30.*;172.31.*;192.168.*;127.0.0.1;<local>`); err != nil {
		log.Fatalf("[sys] registry set proxy override error:%s \n", err.Error())
	}

	println("[sys] proxy config set successfully")
}
