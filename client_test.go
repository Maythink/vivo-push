package vivopush

import (
	"testing"
	// "time"
)

var packageName string = "sbkssbkssbkssbkssbkssbkssbkssbks"

var client = NewClient("sbkssbkssbkssbkssbkssbkssbkssbks")

var msg1 *Message = NewVivoMessage("hi baby1", "hi1")
var msg2 *Message = NewVivoMessage("hi baby2", "hi2 ")

var regID1 string = "WFioJi0fiIco7vOrI4dnxxjeKAUqR7fjugoGkHUgxeo="
var regID2 string = "52Pe7fPIRXWsXhzn4eYJ1njYhBhN8Lcp8IJPOMjThdk="

func TestMiPush_Send(t *testing.T) {
	result, err := client.Send(msg1, regID1)
	if err != nil {
		t.Errorf("TestMiPush_Send failed :%v\n", err)
	}
	t.Logf("result=%#v\n", result)
}

func TestMiPush_GetMessageStatusByJobKey(t *testing.T) {
	result, err := client.GetMessageStatusByJobKey("key111")
	if err != nil {
		t.Errorf("TestMiPush_GetMessageStatusByJobKey failed :%v\n", err)
	}
	t.Logf("result=%#v\n", result)
}
