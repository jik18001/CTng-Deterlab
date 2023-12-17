package main

import (
        "fmt"
        "math/rand"
        "os"
        "time"
        "encoding/json"

        CA "github.com/jik18001/CTngV2/CA"
        Logger "github.com/jik18001/CTngV2/Logger"
        gossiper "github.com/jik18001/CTngV2/gossiper"
        monitor "github.com/jik18001/CTngV2/monitor"
        util "github.com/jik18001/CTngV2/util"
)

type GossiperConn struct {
        ID                 int
        ConnectedGossipers []string
}
type MonitorConn struct {
        ID               int
        ConnectedCAs     []string
        ConnectedLoggers []string
}


func OverwriteOrginatorMonitorConnection(ctx_m *monitor.MonitorContext, MID string) {
        path := "Topology/monitorconn_" + MID + ".json"
        var monitorconn MonitorConn = MonitorConn{ID: 0}
        monitorconn_json, _ := util.ReadByte(path)
        _ = json.Unmarshal(monitorconn_json, &monitorconn)
        ctx_m.Monitor_private_config.CA_URLs = monitorconn.ConnectedCAs
        ctx_m.Monitor_private_config.Logger_URLs = monitorconn.ConnectedLoggers
        fmt.Println(monitorconn)
        fmt.Println("Overwrite the monitor connection for monitor ", MID)
        fmt.Println("CA_URLs: ", ctx_m.Monitor_private_config.CA_URLs)
        fmt.Println("Logger_URLs: ", ctx_m.Monitor_private_config.Logger_URLs)
}

/*
func OverwriteOrginatorMonitorConnectionD1(ctx_m *monitor.MonitorContext, MID string) {
        path := "CLD1/monitorconn_" + MID + ".json"
        var monitorconn MonitorConn = MonitorConn{ID: 0}
        monitorconn_json, _ := util.ReadByte(path)
        _ = json.Unmarshal(monitorconn_json, &monitorconn)
        ctx_m.Monitor_private_config.CA_URLs = monitorconn.ConnectedCAs
        ctx_m.Monitor_private_config.Logger_URLs = monitorconn.ConnectedLoggers
        fmt.Println(monitorconn)
        fmt.Println("Overwrite the monitor connection for monitor ", MID)
        fmt.Println("CA_URLs: ", ctx_m.Monitor_private_config.CA_URLs)
        fmt.Println("Logger_URLs: ", ctx_m.Monitor_private_config.Logger_URLs)
}

func OverwriteOrginatorMonitorConnectionD2(ctx_m *monitor.MonitorContext, MID string) {
        path := "CLD2/monitorconn_" + MID + ".json"
        var monitorconn MonitorConn = MonitorConn{ID: 0}
        monitorconn_json, _ := util.ReadByte(path)
        _ = json.Unmarshal(monitorconn_json, &monitorconn)
        ctx_m.Monitor_private_config.CA_URLs = monitorconn.ConnectedCAs
        ctx_m.Monitor_private_config.Logger_URLs = monitorconn.ConnectedLoggers
        fmt.Println(monitorconn)
        fmt.Println("Overwrite the monitor connection for monitor ", MID)
        fmt.Println("CA_URLs: ", ctx_m.Monitor_private_config.CA_URLs)
        fmt.Println("Logger_URLs: ", ctx_m.Monitor_private_config.Logger_URLs)
}

func OverwriteOrginatorMonitorConnectionD3(ctx_m *monitor.MonitorContext, MID string) {
        path := "CLD3/monitorconn_" + MID + ".json"
        var monitorconn MonitorConn = MonitorConn{ID: 0}
        monitorconn_json, _ := util.ReadByte(path)
        _ = json.Unmarshal(monitorconn_json, &monitorconn)
        ctx_m.Monitor_private_config.CA_URLs = monitorconn.ConnectedCAs
        ctx_m.Monitor_private_config.Logger_URLs = monitorconn.ConnectedLoggers
        fmt.Println(monitorconn)
        fmt.Println("Overwrite the monitor connection for monitor ", MID)
        fmt.Println("CA_URLs: ", ctx_m.Monitor_private_config.CA_URLs)
        fmt.Println("Logger_URLs: ", ctx_m.Monitor_private_config.Logger_URLs)
}
*/
func OverwriteGossiperConnection(ctx_g *gossiper.GossiperContext, GID string) {
        path := "DMLCG/g_" + GID + ".json"
        var gossiperconn GossiperConn
        gossiperconn_json, _ := util.ReadByte(path)
        json.Unmarshal(gossiperconn_json, &gossiperconn)
        ctx_g.Gossiper_private_config.Connected_Gossipers = gossiperconn.ConnectedGossipers
        fmt.Println("Overwrite the gossiper connection for gossiper ", GID)
        fmt.Println("ConnectedGossipers: ", ctx_g.Gossiper_private_config.Connected_Gossipers)
}

func StartCA(CID string) {
        path_prefix := "ca_testconfig/" + CID
        path_1 := path_prefix + "/CA_public_config.json"
        path_2 := path_prefix + "/CA_private_config.json"
        path_3 := path_prefix + "/CA_crypto_config.json"
        ctx_ca := CA.InitializeCAContext(path_1, path_2, path_3)
        ctx_ca.RevocationRatio = 0.02
        ctx_ca.Max_latency = 100
        CA.StartCA(ctx_ca)
}

func StartLogger(LID string) {
        path_prefix := "logger_testconfig/" + LID
        path_1 := path_prefix + "/Logger_public_config.json"
        path_2 := path_prefix + "/Logger_private_config.json"
        path_3 := path_prefix + "/Logger_crypto_config.json"
        ctx_logger := Logger.InitializeLoggerContext(path_1, path_2, path_3)
        ctx_logger.Max_latency = 100
        Logger.StartLogger(ctx_logger)
}

func StartMonitor(MID string) {
        path_prefix := "monitor_testconfig/" + MID
        path_1 := path_prefix + "/Monitor_public_config.json"
        path_2 := path_prefix + "/Monitor_private_config.json"
        path_3 := path_prefix + "/Monitor_crypto_config.json"
        ctx_monitor := monitor.InitializeMonitorContext(path_1, path_2, path_3, MID)
        // clean up the storage
        ctx_monitor.InitializeMonitorStorage("monitor_testdata/")
        // delete all the files in the storage
        ctx_monitor.CleanUpMonitorStorage()
        //ctx_monitor.Mode = 0
        //wait for 60 seconds
        fmt.Println("Delay 60 seconds to start monitor server")
        time.Sleep(60 * time.Second)
        ctx_monitor.Period_Offset = util.GetCurrentPeriod()
        ctx_monitor.Maxdrift_miliseconds = 1000
        ctx_monitor.Clockdrift_miliseconds = rand.Intn(1000)
        OverwriteOrginatorMonitorConnection(ctx_monitor, MID)
        ctx_monitor.Max_latency = 100
        monitor.StartMonitorServer(ctx_monitor)
}

func StartGossiper(GID string) {
        path_prefix := "gossiper_testconfig/" + GID
        path_1 := path_prefix + "/Gossiper_public_config.json"
        path_2 := path_prefix + "/Gossiper_private_config.json"
        path_3 := path_prefix + "/Gossiper_crypto_config.json"
        ctx_gossiper := gossiper.InitializeGossiperContext(path_1, path_2, path_3, GID)
        ctx_gossiper.StorageDirectory = "gossiper_testdata/" + ctx_gossiper.StorageID + "/"
        ctx_gossiper.StorageFile = "gossiper_testdata.json"
        ctx_gossiper.CleanUpGossiperStorage()
        ctx_gossiper.Total_Logger = 20
        ctx_gossiper.Total_CA = 100
        // create the storage directory if not exist
        util.CreateDir(ctx_gossiper.StorageDirectory)
        // create the storage file if not exist
        util.CreateFile(ctx_gossiper.StorageDirectory + ctx_gossiper.StorageFile)
        //ctx_gossiper.Optimization_threshold = 999999
        //ctx_gossiper.Optimization_mode = false
        OverwriteGossiperConnection(ctx_gossiper,GID)
        ctx_gossiper.Max_latency = 100
        gossiper.StartGossiperServer(ctx_gossiper)
}

func main() {
        if len(os.Args) < 2 {
                fmt.Println("Usage: go run Test1.go <CA|Logger|Monitor|Gossiper> <ID>")
                os.Exit(1)
        }
        time.AfterFunc(3*time.Minute, func() {
                fmt.Println("Terminating the program after 3 minutes.")
                os.Exit(0)
        })
        switch os.Args[1] {
        case "CA":
                StartCA(os.Args[2])
        case "Logger":
                StartLogger(os.Args[2])
        case "Monitor":
                StartMonitor(os.Args[2])
        case "Gossiper":
                StartGossiper(os.Args[2])
        default:
                fmt.Println("Usage: go run Test1.go <CA|Logger|Monitor|Gossiper> <ID>")
                os.Exit(1)
        }
}
