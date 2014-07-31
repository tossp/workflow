package workflow

import (
	"time"
)

//流程实例
type FlowInstance struct {
	FlowDefine IWorkflow         //流程定义
	AppData    map[string]string //业务数据
	FlowCase   *FlowCase         //流程实例的数据
	FlowItems  *[]FlowItem       //流程实例的任务数据
}

func New_FlowInstance(flowDef IWorkflow) *FlowInstance {
	fi := &FlowInstance{
		FlowDefine: flowDef,
		AppData:    make(map[string]string),
	}
	return fi
}

func (this *FlowInstance) Run() bool {
	return true
}

//流程实例的数据
type FlowCase struct {
	CaseId      string    //实例id
	FlowId      string    //流程定义id
	Creator     string    //创建人id
	CreatorName string    //创建人名字
	CurrentItem int32     //当前任务项id
	CreateTime  time.Time //创建时间
	EndTime     time.Time //结束时间

}

//流程实例的任务数据
type FlowItem struct {
	ItemId     string
	PartUser   string
	PartName   string
	ItemState  int32
	ActivityId string
}
