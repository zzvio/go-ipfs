package consts

import (
	"time"
)

// Plugin Vars -
var (
	PluginToken int
)

// Plugin consts -
const (
	CallbackInterval   = 100 * time.Millisecond
	MatchInterval      = 1000 * time.Millisecond
	MsgRateLimitPerMin = 60
)

// Wrapper MsgType
const (
	MsgTypeInitiator = byte(1)
	MsgTypeAcceptor  = byte(2)
)

// MsgTypes
const (
	MsgTypeStartPlugin = byte(0)
	MsgTypeGetOrderID  = byte(1)
	MsgTypeEth         = byte(100)
	MsgTypeReceive     = byte(255)
)

// MsgSubType for MsgTypeEth
const (
	MsgSubTypeMode = byte(0)
)
