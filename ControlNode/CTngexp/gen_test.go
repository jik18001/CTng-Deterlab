package main

import (
	"ctngexp/Gen"
	"testing"
)

func TestGen(t *testing.T) {
	Gen.Generateall(100, 20,8, 2, 2, 60, 60, "")
	newtemp := Gen.Generate_IP_Json_template(100, 20, 8, "172.30.0.", 11, "172.30.0.", 22, "172.30.0.", 42)
	Gen.Write_IP_Json_to_files("IPLIST.json", newtemp)
	IPLIST := Gen.Read_IP_Json_from_files("IPLIST.json")
	Gen.Map_all(100, 20, 8, 2, IPLIST)
}
