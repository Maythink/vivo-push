package vivopush

const (
	ProductionHost = "https://api-push.vivo.com.cn"
)

const (
	AuthURL            = "/message/auth"            // 推送鉴权接口
	SendURL            = "/message/send"            // 单推接口
	SaveListPayloadURL = "/message/saveListPayload" // 保存群推消息公共体接口
	PushToListURL      = "/message/pushToList"      // 批量推送用户接口
	PushToAllURL       = "/message/all"             // 全量发送接口
	MessagesStatusURL  = "/report/getStatistics "   // 获取消息推送的统计值接口
)

var (
	PostRetryTimes       = 3         //重试次数
	MaxTimeToLive  int64 = 3600 * 24 //消息保留时长
)
