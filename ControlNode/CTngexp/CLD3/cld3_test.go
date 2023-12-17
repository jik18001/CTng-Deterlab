package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/jik18001/CTngV2/crypto"
	"github.com/jik18001/CTngV2/gossiper"
	"github.com/jik18001/CTngV2/monitor"
	"github.com/jik18001/CTngV2/util"
)

type all_urls struct {
	CA_URLs     []string
	Logger_URLs []string
}

type MonitorConn struct {
	ID               int
	ConnectedCAs     []string
	ConnectedLoggers []string
}

func Containindex(arr []int, index int) bool {
	for _, i := range arr {
		if i == index {
			return true
		}
	}
	return false
}

func connectCAsAndLoggersToMonitors(monitors []MonitorConn, numMonitors, T int, CA_urls []string, Logger_urls []string) {
	// divide monitors into 4 groups
	// each group has numMonitors/4 monitors
	// should always be divisible by 4
	numeachgroup := T
	fmt.Println(numeachgroup)
	connectedMonitors := make([]int, 0)
	for j := 0; j < numeachgroup; j++ {
		connectedMonitors = append(connectedMonitors, j)
	}
	fmt.Println(connectedMonitors)
	// Connect CAs to Monitors
	for _, caURL := range CA_urls {
		for _, monitorIndex := range connectedMonitors {
			monitors[monitorIndex].ConnectedCAs = append(monitors[monitorIndex].ConnectedCAs, caURL)
		}
	}

	// Connect Loggers to Monitors
	for _, loggerURL := range Logger_urls {
		for _, monitorIndex := range connectedMonitors {
			monitors[monitorIndex].ConnectedLoggers = append(monitors[monitorIndex].ConnectedLoggers, loggerURL)
		}
	}
	// set the rest of monitors to be connected to nothing (empty list)
	for i := 0; i < numMonitors; i++ {
		if !Containindex(connectedMonitors, i) {
			monitors[i].ConnectedCAs = make([]string, 0)
			monitors[i].ConnectedLoggers = make([]string, 0)
		}
	}
}

func TestGenMonitorConn(t *testing.T) {
	gpath := "../gossiper_testconfig/1/Gossiper_public_config.json"
	var gconf gossiper.Gossiper_public_config
	bytes, _ := util.ReadByte(gpath)
	json.Unmarshal(bytes, &gconf)
	mpath := "../monitor_testconfig/1/Monitor_public_config.json"
	var mconf monitor.Monitor_public_config
	bytes, _ = util.ReadByte(mpath)
	json.Unmarshal(bytes, &mconf)
	numMonitors := len(gconf.Gossiper_URLs)
	gcpath := "../gossiper_testconfig/1/Gossiper_crypto_config.json"
	var gcconf *crypto.CryptoConfig
	gcconf, _ = crypto.ReadCryptoConfig(gcpath)
	Threshold := gcconf.Threshold
	monitors := make([]MonitorConn, numMonitors)
	for i := 0; i < numMonitors; i++ {
		monitors[i].ID = i + 1
	}
	connectCAsAndLoggersToMonitors(monitors, numMonitors, Threshold, mconf.All_CA_URLs, mconf.All_Logger_URLs)
	for _, monitor := range monitors {
		util.WriteData("monitorconn_"+strconv.Itoa(monitor.ID)+".json", monitor)
	}
}
