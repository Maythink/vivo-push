# vivo-push
Vivo推送服务 Golang SDK

full golang implementation of Vivo Push API (https://dev.vivo.com.cn/documentCenter/doc/155)

```Go
import (
    "fmt"

    vv "github.com/Maythink/vivo-push"
)

func main() {
	client, err := vv.NewClient("your appId", "your appKey", "your appSecret")
	if err != nil {
		return
	}

	// 单推
	msg1 := vv.NewVivoMessage("hi baby", "hi")
	_, err = client.Send(msg1, "regID")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 群推
	msg2 := vv.NewListPayloadMessage("hello baby", "hello")
	_, err = client.SendList(msg2, []string{"regID1", "regID2"})
	if err != nil {
		fmt.Println(err)
		return
	}

	//全量推送
	msg3 := vv.NewListPayloadMessage("hi all baby", "hi all")
	_, err = client.SendAll(msg3)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

```

### Sender APIs

- [x] Send(msg *Message, regID string)
- [x] SendList(msg *MessagePayload, regIds []string)
- [x] SendAll(msg *MessagePayload)

### Stats APIs

- [x] GetMessageStatusByJobKey(jobKey string) (*BatchStatusResult, error) 
