package main

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func createWiFi(ssid, password string) {
	cmd := exec.Command("netsh", "wlan", "set", "hostednetwork", "mode=allow", "ssid="+ssid, "key="+password)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("netsh", "wlan", "start", "hostednetwork")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Wi-Fi 热点已创建。")
}

func Test_wifeMock(*testing.T) {
	ssid := "YourSSID"
	password := "YourPassword"
	createWiFi(ssid, password)
}
