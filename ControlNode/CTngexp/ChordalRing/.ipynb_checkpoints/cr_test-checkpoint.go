package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/jik18001/CTngV2/util"
)

type GossiperConn struct {
	ID                 int
	ConnectedGossipers []string
}

func generateChordalRingNeighbors(ips []string) []GossiperConn {
	if len(ips) != 16 {
		fmt.Println("Error: Input list must contain exactly 16 IP addresses")
		return nil
	}

	neighbors := make([]GossiperConn, 16)
	for i := 0; i < 16; i++ {
		neighbors[i].ID = i + 1 // Set the ID for each Gossiper
		// Each node will have 3 neighbors based on CR16(1,3,9)
		neighbors[i].ConnectedGossipers = append(neighbors[i].ConnectedGossipers, ips[(i+1)%16]) // Neighbor 1 step away
		neighbors[i].ConnectedGossipers = append(neighbors[i].ConnectedGossipers, ips[(i+3)%16]) // Neighbor 3 steps away
		neighbors[i].ConnectedGossipers = append(neighbors[i].ConnectedGossipers, ips[(i+9)%16]) // Neighbor 9 steps away
	}

	return neighbors
}

func TestCR(t *testing.T) {
	ips := []string{
		"172.30.0.42:9000",
		"172.30.0.43:9001",
		"172.30.0.44:9002",
		"172.30.0.45:9003",
		"172.30.0.46:9004",
		"172.30.0.47:9005",
		"172.30.0.48:9006",
		"172.30.0.49:9007",
		"172.30.0.50:9008",
		"172.30.0.51:9009",
		"172.30.0.52:9010",
		"172.30.0.53:9011",
		"172.30.0.54:9012",
		"172.30.0.55:9013",
		"172.30.0.56:9014",
		"172.30.0.57:9015",
	}
	neighbors := generateChordalRingNeighbors(ips)
	for _, monitor := range neighbors {
		util.WriteData("g_"+strconv.Itoa(monitor.ID)+".json", monitor)
	}
}
