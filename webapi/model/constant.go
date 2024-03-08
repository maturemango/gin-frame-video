package model

// 密钥格式要对
const PubKey = `-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAN/XLyUx1DwoWE3QgtJr7vYGk0mN0XsQ
lR7yG7WSEQj44QsxA2ue0ph83EUURvZ/n9Wca6rd7R6IKwEWIZYIQOsCAwEAAQ==
-----END PUBLIC KEY-----`

const PriKey = `-----BEGIN PRIVATE KEY-----
MIIBVgIBADANBgkqhkiG9w0BAQEFAASCAUAwggE8AgEAAkEA39cvJTHUPChYTdCC
0mvu9gaTSY3RexCVHvIbtZIRCPjhCzEDa57SmHzcRRRG9n+f1Zxrqt3tHogrARYh
lghA6wIDAQABAkEAn65CU6ZYYRHm7Jvyt2mH7rqCF9azubb6qjjMy5qHzH1o6e8X
k1eW3tz11eMEMYsTGjP+TUCNSL3NQeiNdjabAQIhAP384PzmCpt9MpqFvjdYKhQ6
xV/6R7VgYCRzysm4oE/lAiEA4Z0ppLv9d9DjHNj1MgZVK2D6TnFWd3kO2SvATsKT
II8CIDX7TjJSDkUX6e5vqIsIwQDFsPeCMUV6c1SsC5iuFdyFAiEAi01v1gAg66b1
Y+1tz7prQgJ56o8+VTxQ97R04+xtzW8CIQCn/bOpl5/XQAtMLClLf917801jwrN5
g0EFYzOncAHsfQ==
-----END PRIVATE KEY-----`

// 视频弹幕用缓存键
var VideoInputTimeKey string = "danmu:Video:No.%s;Node:%d;User:no.%s;Time:%d" // 视频某个时间点的弹幕键
var GetVideoInputTimeKey string = "danmu:Video:No.%s;Node:%d;User:"
var VideoInputKey string = "danmu:video:no.%s;node:%d;user:no.%s;time:%d" // 视频弹幕根据视频编号和用户编号以及发送的弹幕时间戳来设置键
var GetUserVideoInputKey string = "danmu:video:no.%s;user:no.%s;time:"  // 查询用户在同一视频下的所有弹幕

// 视频操作相关
var QueryTripletSQL = `select * from gf_video_triplet where video_no = ? and user_id = ?`

// 操作类型
type UserServiceDesc string

const (
	Login          UserServiceDesc = "用户登录"
	UpdateData     UserServiceDesc = "修改用户信息"
	UploadVideo    UserServiceDesc = "用户上传视频"
	VideoManage    UserServiceDesc = "用户视频管理"
)

// 操作详情
type OperateDetail string

const (
	LoginSuccess      OperateDetail = "登录成功"
)

// 日志等级
type LogLevel string

const (
	Debug   LogLevel = "DEBUG"
	Info    LogLevel = "INFO"
	Warn    LogLevel = "WARN"
	Error   LogLevel = "ERROR"
)

// 角色id
const (
	SuperUser = iota + 1   // 超管角色id
	Admin                  // 管理员角色id
	User                   // 普通用户角色id
)