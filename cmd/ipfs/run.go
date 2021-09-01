package main

import (
	"os"
	"strings"

	"github.com/ipfs/go-ipfs/cmd/ipfs/consts"
	"github.com/ipfs/go-ipfs/cmd/ipfs/tools"
	"golang.org/x/sync/syncmap"
)

var (
	orderCount         int
	clientToOrderIDMap *syncmap.Map
	sessionChan        chan bool
)

func init() {
	clientToOrderIDMap = new(syncmap.Map)
	orderCount = 0
	sessionChan = make(chan bool)
}

// call - Runs the contract, It receive data as parsed byte and returns back a parsed byte array
func call(arr []byte) []byte {
	if len(arr) >= 6 && arr[0] == consts.MsgTypeStartPlugin && arr[1] == consts.MsgTypeStartPlugin {
		consts.PluginToken = tools.BytesToInt(arr[2:6])

		args := strings.Split(string(arr[6:]), " ")
		args = append([]string{"ipfs"}, args...)
		os.Exit(mainRet(args))
		return []byte{}
	}
	return []byte{}
}
