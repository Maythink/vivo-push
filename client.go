package vivopush

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	// "strconv"
	// "strings"
)

var authToken *AuthToken

type AuthToken struct {
	token      string
	valid_time int64
}

type VivoPush struct {
	host       string
	auth_token string
}

func NewClient(auth_token string) *VivoPush {
	return &VivoPush{
		host:       ProductionHost,
		auth_token: auth_token,
	}
}

//----------------------------------------Sender----------------------------------------//
// 根据registrationId，发送消息到指定设备上
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

func (m *VivoPush) handleResponse(response *http.Response) ([]byte, error) {
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
	req.Header.Set("authToken", m.auth_token)
	client := &http.Client{}
	tryTime := 0
tryAgain:
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("vivo push post err", err, tryTime)
		tryTime += 1
		if tryTime < PostRetryTimes {
			goto tryAgain
		}
		return nil, err
	}
	result, err = m.handleResponse(resp)
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
	req.Header.Set("authToken", m.auth_token)

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err = m.handleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}
