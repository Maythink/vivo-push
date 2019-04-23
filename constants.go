package vivopush

const (
	ProductionHost = "https://api-push.vivo.com.cn"
)

const (
	AuthURL           = "/message/auth"          // 推送鉴权接口
	RegURL            = "/message/send"          // 单推接口
	MessagesStatusURL = "/report/getStatistics " // 获取消息推送的统计值接口
)

var (
	PostRetryTimes = 3 //重试次数
)
