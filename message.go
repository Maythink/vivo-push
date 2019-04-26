package vivopush

import (
	"encoding/json"
	"strings"

	gouuid "github.com/satori/go.uuid"
)

//单推
type Message struct {
	RegId           string            `json:"regId"`           // 订阅 PUSH 服务器得到的 id
	NotifyType      int               `json:"notifyType"`      // 通知类型 1:无，2:响铃，3:振动，4:响铃和振动
	Title           string            `json:"title"`           // 通知标题
	Content         string            `json:"content"`         // 通知内容
	TimeToLive      int64             `json:"timeToLive"`      // 可选项。消息保留时长 单位：秒，取值至少60秒，最长7天。当值为空时，默认一天 86400
	SkipType        int               `json:"skipType"`        // 点击跳转类型 1：打开 APP 首页 2：打开链接 3：自定义 4:打开 app 内指定页面
	SkipContent     string            `json:"skipContent"`     // 可选项。跳转内容跳转类型为 2 时，跳转内容最大1000 个字符，跳转类型为 3 或 4 时，跳转内容最大 1024 个字符
	NetworkType     int               `json:"networkType"`     // 可选项。网络方式 -1：不限，1：wifi 下发送，不填默认为-1
	ClientCustomMap map[string]string `json:"clientCustomMap"` // 可选项。客户端自定义键值对自定义key和Value键值对个数不能超过 10 个，且长度不能超过1024 字符, key 和 Value 键值对总长度不能超过 1024 字符。
	Extra           map[string]string `json:"extra"`           // 可选项。高级特性
	RequestId       string            `json:"requestId"`       // 用户请求唯一标识
}

// 保存群推消息
type MessagePayload struct {
	Title           string            `json:"title"`           // 通知标题
	Content         string            `json:"content"`         // 通知内容
	NotifyType      int               `json:"notifyType"`      // 通知类型 1:无，2:响铃，3:振动，4:响铃和振动
	TimeToLive      int64             `json:"timeToLive"`      // 可选项。消息保留时长 单位：秒，取值至少60秒，最长7天。当值为空时，默认一天 86400
	SkipType        int               `json:"skipType"`        // 点击跳转类型 1：打开 APP 首页 2：打开链接 3：自定义 4:打开 app 内指定页面
	SkipContent     string            `json:"skipContent"`     // 可选项。跳转内容跳转类型为 2 时，跳转内容最大1000 个字符，跳转类型为 3 或 4 时，跳转内容最大 1024 个字符
	NetworkType     int               `json:"networkType"`     // 可选项。网络方式 -1：不限，1：wifi 下发送，不填默认为-1
	ClientCustomMap map[string]string `json:"clientCustomMap"` // 可选项。客户端自定义键值对自定义key和Value键值对个数不能超过 10 个，且长度不能超过1024 字符, key 和 Value 键值对总长度不能超过 1024 字符。
	RequestId       string            `json:"requestId"`       // 用户请求唯一标识
}

//群推
type MessageList struct {
	RegIds    []string `json:"regIds"`    // regId 列表 个数大于等于 2，小于等于 1000， regId 长度 23 个字符(regIds，aliases 两者需 一个不为空，两个不为空，取 regIds)
	TaskId    string   `json:"taskId"`    // 公共消息任务号，取 saveListPayload 返回的 taskId
	RequestId string   `json:"requestId"` // 用户请求唯一标识
}

func (m *Message) SetNotifyType(notifyType int) *Message {
	m.NotifyType = notifyType
	return m
}

// 添加自定义字段, 客户端使用
func (m *Message) AddExtra(key, value string) *Message {
	m.Extra[key] = value
	return m
}

func (m *Message) AddCustomMap(key, value string) *Message {
	m.ClientCustomMap[key] = value
	return m
}

func (m *Message) JSON() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return bytes
}

func (m *MessagePayload) JSON() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return bytes
}

//-----------------------------------------------------------------------------------//
// 发送给设备的Message对象
func NewVivoMessage(title, content string) *Message {
	return &Message{
		NotifyType:      2,
		Title:           title,
		Content:         content,
		TimeToLive:      MaxTimeToLive,
		SkipType:        1,
		SkipContent:     "",
		NetworkType:     -1,
		ClientCustomMap: make(map[string]string),
		Extra:           make(map[string]string),
		RequestId:       strings.ToUpper(gouuid.Must(gouuid.NewV4()).String()),
	}
}

// 保存群推消息
func NewListPayloadMessage(title, content string) *MessagePayload {
	return &MessagePayload{
		Title:           title,
		Content:         content,
		NotifyType:      2,
		TimeToLive:      MaxTimeToLive,
		SkipType:        1,
		SkipContent:     "",
		NetworkType:     -1,
		ClientCustomMap: make(map[string]string),
		RequestId:       strings.ToUpper(gouuid.Must(gouuid.NewV4()).String()),
	}
}

// 发送群推消息给regIds
func NewListMessage(regIds []string, taskId string) *MessageList {
	return &MessageList{
		RegIds:    regIds,
		TaskId:    taskId,
		RequestId: strings.ToUpper(gouuid.Must(gouuid.NewV4()).String()),
	}
}

// 打开当前app首页
func (m *Message) SetLauncherActivity() *Message {
	m.SkipType = 1
	return m
}

// 打开网页
func (m *Message) SetJumpWebURL(value string) *Message {
	m.SkipType = 2
	m.SkipContent = value
	return m
}

// 打开自定义
func (m *Message) SetJumpCustom(value string) *Message {
	m.SkipType = 3
	m.SkipContent = value
	return m
}

// 打开当前app内的任意一个Activity。
func (m *Message) SetJumpActivity(value string) *Message {
	m.SkipType = 4
	m.SkipContent = value
	return m
}

//-----------------------------------------广播------------------------------------------//
// 设置通知类型
func (m *MessagePayload) SetPayloadNotifyType(notifyType int) *MessagePayload {
	m.NotifyType = notifyType
	return m
}

// 客户端自定义键值对
func (m *MessagePayload) PayloadAddCustomMap(key, value string) *MessagePayload {
	m.ClientCustomMap[key] = value
	return m
}

// 打开当前app首页
func (m *MessagePayload) SetPayloadLauncherActivity() *MessagePayload {
	m.SkipType = 1
	return m
}

// 打开网页
func (m *MessagePayload) SetPayloadJumpWebURL(value string) *MessagePayload {
	m.SkipType = 2
	m.SkipContent = value
	return m
}

// 打开自定义
func (m *MessagePayload) SetPayloadJumpCustom(value string) *MessagePayload {
	m.SkipType = 3
	m.SkipContent = value
	return m
}

// 打开当前app内的任意一个Activity。
func (m *MessagePayload) SetPayloadJumpActivity(value string) *MessagePayload {
	m.SkipType = 4
	m.SkipContent = value
	return m
}

//-----------------------------------------------------------------------------------//
// TargetedMessage封装了VivoPush推送服务系统中的消息Message对象，和该Message对象所要发送到的目标。

type TargetType int32

const (
	TargetTypeRegID   TargetType = 1
	TargetTypeReAlias TargetType = 2
	TargetTypeAccount TargetType = 3
)

type TargetedMessage struct {
	message    *Message
	targetType TargetType
	target     string
}

func NewTargetedMessage(m *Message, target string, targetType TargetType) *TargetedMessage {
	return &TargetedMessage{
		message:    m,
		targetType: targetType,
		target:     target,
	}
}

func (tm *TargetedMessage) SetTargetType(targetType TargetType) *TargetedMessage {
	tm.targetType = targetType
	return tm
}

func (tm *TargetedMessage) SetTarget(target string) *TargetedMessage {
	tm.target = target
	return tm
}

func (tm *TargetedMessage) JSON() []byte {
	bytes, err := json.Marshal(tm)
	if err != nil {
		panic(err)
	}
	return bytes
}
