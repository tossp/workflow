package workflow

import (
	"github.com/nu7hatch/gouuid"
)

//流程中的事件
type FlowEvent struct {
	//流程实列Id
	CaseId uuid.UUID
	//流程任务Id
	TaskId int32
	//事件Id
	EventId uuid.UUID
	//事件名称
	EventName string
	//处理人类型
	HandlerType EventHandleType
}

//Email事件
type EmailEvent struct {
	FlowEvent
	//标题
	Subject string
	//邮件内容
	Body string
	//收件人列表
	MailList []string
}

//短信事件
type SMSEvent struct {
	FlowEvent
	//短信内容
	Content string
	//电话号码
	Mobile []string
}

//消息事件
type MessageEvent struct {
	FlowEvent
	//消息内容
	Content string
	//接收人
	UserList []string
}

//数据接口调用
type DatasourceEvent struct {
	FlowEvent
	//数据源id
	DatasourceId string
	//参数列表
	Params map[string]string
}

//转发事件
type TransferEvent struct {
	FlowEvent
	//下一个处理人
	NextUser string
}
