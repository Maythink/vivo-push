package vivopush

import (
	"testing"
)

var appId string = "11984"
var appKey string = "e6611c0e-851a-46df-87da-149f9cc50c93"
var appSecret string = "6fdb2a63-6c7a-4e98-bbf7-c89c47b789ac"

var msg1 *Message = NewVivoMessage("hi baby1", "hi1")

var regID1 string = "15559958796931198455743"

func TestMiPush_Send(t *testing.T) {
	client, err := NewClient(appId, appKey, appSecret)
	if err != nil {
		t.Errorf("TestMiPush_Send failed :%v\n", err)
	}
	result, err := client.Send(msg1, regID1)
	if err != nil {
		t.Errorf("TestMiPush_Send failed :%v\n", err)
	}
	t.Logf("result=%#v\n", result)
}

func TestMiPush_GetMessageStatusByJobKey(t *testing.T) {
	client, err := NewClient(appId, appKey, appSecret)
	if err != nil {
		t.Errorf("TestMiPush_Send failed :%v\n", err)
	}
	result, err := client.GetMessageStatusByJobKey("570247239105613824")
	if err != nil {
		t.Errorf("TestMiPush_GetMessageStatusByJobKey failed :%v\n", err)
	}
	t.Logf("result=%#v\n", result)
}
