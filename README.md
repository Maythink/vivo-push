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
	msg1 := vv.NewVivoMessage("hi baby1", "hi1")
	_, err = client.Send(msg1, "regID")
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

```

### Sender APIs

- [x] Send(msg *Message, regID string)

### Stats APIs

- [x] GetMessageStatusByJobKey(jobKey string) (*BatchStatusResult, error) 
