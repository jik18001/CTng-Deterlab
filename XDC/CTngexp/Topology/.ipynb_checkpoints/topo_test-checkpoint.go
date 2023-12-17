package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

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

func connectCAsAndLoggersToMonitors(monitors []MonitorConn, numMonitors, T int, CA_urls []string, Logger_urls []string) {
	rand.Seed(time.Now().UnixNano())

	// Connect CAs to Monitors
	for _, caURL := range CA_urls {
		connectedMonitors := rand.Perm(numMonitors)[:T]
		for _, monitorIndex := range connectedMonitors {
			monitors[monitorIndex].ConnectedCAs = append(monitors[monitorIndex].ConnectedCAs, caURL)
		}
	}

	// Connect Loggers to Monitors
	for _, loggerURL := range Logger_urls {
		connectedMonitors := rand.Perm(numMonitors)[:T]
		for _, monitorIndex := range connectedMonitors {
			monitors[monitorIndex].ConnectedLoggers = append(monitors[monitorIndex].ConnectedLoggers, loggerURL)
		}
	}
}

func TestGenMonitorConn(t *testing.T) {
	var allURLs all_urls
	urls_json, _ := util.ReadByte("urls.json")
	json.Unmarshal(urls_json, &allURLs)
	fmt.Println(allURLs)
	numMonitors := 16
	Threshold := 4
	monitors := make([]MonitorConn, numMonitors)
	for i := 0; i < numMonitors; i++ {
		monitors[i].ID = i + 1
	}
	connectCAsAndLoggersToMonitors(monitors, numMonitors, Threshold, allURLs.CA_URLs, allURLs.Logger_URLs)
	for _, monitor := range monitors {
		util.WriteData("monitorconn_"+strconv.Itoa(monitor.ID)+".json", monitor)
	}
}
