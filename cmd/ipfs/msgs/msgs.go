package msgs

import (
	"time"

	"github.com/ipfs/go-ipfs/cmd/ipfs/consts"
	"github.com/ipfs/go-ipfs/cmd/ipfs/tools"

	"github.com/ClarkGuan/jni"
)

//
var (
	msgArr         []byte
	count          int
	orderID        int
	callbackMethod uintptr
	clazzID        uintptr
	envID          jni.Env
)

func init() {
	msgArr = []byte{}
	count = 0
	orderID = 0
	clazzID = 0
	callbackMethod = 0
	envID = 0
	go sendCallback()
}

// AddMessage - adds a message to global message array
func AddMessage(msg []byte) {
	msgArr = append(msgArr, msg...)
	count++
}

// GetMessageArr - returns global message array
func GetMessageArr() []byte {
	msgBytes := []byte{consts.MsgTypeAcceptor, byte(1)}
	msgBytes = append(msgBytes, tools.NumToBytes(time.Now().UTC().Unix())...)
	msgBytes = append(msgBytes, tools.NumToBytes(count)...)
	msgBytes = append(msgBytes, msgArr...)
	msgArr = []byte{}
	count = 0
	return msgBytes
}

// SetEnv - sets the variables used for call back
func SetEnv(env uintptr, clazz uintptr) {
	if envID == 0 {
		envID = jni.Env(env)
	}
	if clazzID == 0 {
		clazzID = clazz
	}
	if callbackMethod == 0 {
		callbackMethod = jni.Env(env).GetStaticMethodID(clazz, "callback", "(I)V")
	}
}

// Sends callback to java
func sendCallback() {
	ticker := time.NewTicker(consts.CallbackInterval)
	for range ticker.C {
		if count > 0 {
			vm, _ := envID.GetJavaVM()
			newEnv, _ := vm.AttachCurrentThread()
			newEnv.CallStaticVoidMethodA(clazzID, callbackMethod, jni.IntValue(consts.PluginToken))
			vm.DetachCurrentThread()
		}
	}
}

// GetOrderID - returns a new orderID for a clientOrderID
func GetOrderID(clientID int) int {
	orderID++
	msgBytes := []byte{consts.MsgTypeGetOrderID, consts.MsgTypeGetOrderID}
	msgBytes = append(msgBytes, tools.NumToBytes(time.Now().UTC().UnixNano())...)
	msgBytes = append(msgBytes, tools.NumToBytes(orderID)...) // orderID
	msgBytes = append(msgBytes, tools.NumToBytes(clientID)...)
	msgBytes = append(tools.NumToBytes(len(msgBytes)), msgBytes...)
	AddMessage(msgBytes)
	return orderID
}

// SendResult - sends result back to java
func SendResult(msgSubType byte, orderID int, result []byte) {
	msgBytes := []byte{consts.MsgTypeEth, msgSubType}
	msgBytes = append(msgBytes, tools.NumToBytes(time.Now().UTC().UnixNano())...)
	msgBytes = append(msgBytes, tools.NumToBytes(orderID)...)
	msgBytes = append(msgBytes, result...)
	msgBytes = append(tools.NumToBytes(len(msgBytes)), msgBytes...)
	AddMessage(msgBytes)
}
