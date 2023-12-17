package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/jik18001/CTngV2/gossiper"
	"github.com/jik18001/CTngV2/util"
)

// Directed Multi-Layer Cyclic Graph (DMLCG) Topology Generator
// This graph has the following properties:
// 1. Each node has N/4 outgoing edges pointing to nodes in the next group.
// 2. Each node has N/4 incoming edges pointing from nodes in the previous group.
// 3. The Number of nodes in the graph is N, where N is a multiple of 4 and N >= 8.
// 4. f+1 = N/4, where f is the number of faulty nodes allowed in the graph.

type GossiperConn struct {
	ID                 int
	ConnectedGossipers []string
}

// validateMonitorCount checks if the total number of monitors is valid (8, 16, 24, or 32).
func validateMonitorCount(count int) bool {
	return count%8 == 0 && count <= 32
}

// generateGroupedTopology generates a topology based on the specified requirements.
func generateMLCFG(ips []string) ([]GossiperConn, error) {
	if !validateMonitorCount(len(ips)) {
		return nil, fmt.Errorf("invalid number of monitors: %d. Must be 8, 16, 24, or 32", len(ips))
	}

	groupSize := len(ips) / 4
	gossipers := make([]GossiperConn, len(ips))

	for i := 0; i < len(ips); i++ {
		gossipers[i].ID = i + 1
		nextGroupStart := ((i / groupSize) + 1) % 4 * groupSize
		for j := 0; j < groupSize; j++ {
			gossipers[i].ConnectedGossipers = append(gossipers[i].ConnectedGossipers, ips[nextGroupStart+j])
		}
	}

	return gossipers, nil
}

func TestDMLCG(t *testing.T) {
	path := "../gossiper_testconfig/1/Gossiper_public_config.json"
	var conf gossiper.Gossiper_public_config
	bytes, _ := util.ReadByte(path)
	json.Unmarshal(bytes, &conf)
	ips := conf.Gossiper_URLs
	neighbors, _ := generateMLCFG(ips)
	for _, gossiper := range neighbors {
		util.WriteData("g_"+strconv.Itoa(gossiper.ID)+".json", gossiper)
	}
}
