package vivopush

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/bitly/go-simplejson"
)

var authToken *AuthToken = new(AuthToken)

type VivoClient struct {
	AppId     string
	AppKey    string
	AppSecret string
}

type VivoTokenPar struct {
	AppId     string `json:"appId"`
	AppKey    string `json:"appKey"`
	Timestamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
}

type AuthToken struct {
	token      string
	valid_time int64
}

type VivoPush struct {
	host       string
	Auth_token string
}

func NewClient(appId, appKey, appSecret string) (*VivoPush, error) {
	vc := &VivoClient{
		appId,
		appKey,
		appSecret,
	}
	token, err := vc.GetToken()
	if err != nil {
		return nil, err
	}
	return &VivoPush{
		host:       ProductionHost,
		Auth_token: token,
	}, nil
}

//----------------------------------------Token----------------------------------------//
//获取token  返回的expiretime 秒  当过期的时候
func (vc *VivoClient) GetToken() (string, error) {
	now := time.Now().UnixNano() / 1e6
	if authToken != nil {
		if authToken.valid_time > now {
			return authToken.token, nil
		}
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(vc.AppId + vc.AppKey + strconv.FormatInt(now, 10) + vc.AppSecret))
	sign := hex.EncodeToString(md5Ctx.Sum(nil))

	formData, err := json.Marshal(&VivoTokenPar{
		AppId:     vc.AppId,
		AppKey:    vc.AppKey,
		Timestamp: now,
		Sign:      sign,
	})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", ProductionHost+AuthURL, bytes.NewReader(formData))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	result, err := handleResponse(resp)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("network error")
	}
	js, err := simplejson.NewJson(result)
	if err != nil {
		return "", err
	}
	token, err := js.Get("authToken").String()
	if err != nil {
		return "", err
	}
	authToken.token = token
	authToken.valid_time = now + 3600000 //1小时有效
	return token, nil
}

//----------------------------------------Sender----------------------------------------//
// 根据regID，发送消息到指定设备上
func (m *VivoPush) Send(msg *Message, regID string) (*SendResult, error) {
	params := m.assembleSendParams(msg, regID)
	res, err := m.doPost(m.host+RegURL, params)
	if err != nil {
		return nil, err
	}
	var result SendResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	if result.Result != 0 {
		return nil, errors.New(result.Desc)
	}
	fmt.Println(result)
	return &result, nil
}

//----------------------------------------Tracer----------------------------------------//
// 获取指定消息的状态。
func (m *VivoPush) GetMessageStatusByJobKey(jobKey string) (*BatchStatusResult, error) {
	params := m.assembleStatusByJobKeyParams(jobKey)
	res, err := m.doGet(m.host+MessagesStatusURL, params)
	if err != nil {
		return nil, err
	}
	var result BatchStatusResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (m *VivoPush) assembleSendParams(msg *Message, regID string) []byte {
	msg.RegId = regID
	jsondata := msg.JSON()
	return jsondata
}

func (m *VivoPush) assembleStatusByJobKeyParams(jobKey string) string {
	form := url.Values{}
	form.Add("taskIds", jobKey)
	return "?" + form.Encode()
}

func handleResponse(response *http.Response) ([]byte, error) {
	defer func() {
		_ = response.Body.Close()
	}()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *VivoPush) doPost(url string, formData []byte) ([]byte, error) {
	var result []byte
	var req *http.Request
	var resp *http.Response
	var err error

	req, err = http.NewRequest("POST", url, bytes.NewReader(formData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authToken", m.Auth_token)
	client := &http.Client{}
	tryTime := 0
tryAgain:
	resp, err = client.Do(req)
	if err != nil {
		tryTime += 1
		if tryTime < PostRetryTimes {
			goto tryAgain
		}
		return nil, err
	}
	result, err = handleResponse(resp)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("network error")
	}
	return result, nil
}

func (m *VivoPush) doGet(url string, params string) ([]byte, error) {
	var result []byte
	var req *http.Request
	var resp *http.Response
	var err error
	req, err = http.NewRequest("GET", url+params, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authToken", m.Auth_token)

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err = handleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}
