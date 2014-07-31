package workflow

type WorkFlowFactory struct {
	FlowPool     *FlowPool     //流程定义提供者
	CaseProvider ICaseProvider //流程实例提供者
}

func New_WorkFlowFactory() *WorkFlowFactory {
	return &WorkFlowFactory{
		FlowPool:     New_FlowPool(),
		CaseProvider: CaseProviderFactory(),
	}
}

//获取一个流程实例
//流程id, 实例id
func (this *WorkFlowFactory) GetWorkFlow(caseid string) *FlowInstance {
	c := this.CaseProvider.GetCase(caseid)
	ci := this.CaseProvider.GetFlowItems(caseid)

	fd, _ := this.FlowPool.GetFlowDef(c.FlowId)
	return &FlowInstance{
		FlowDefine: fd,
		FlowCase:   c,
		FlowItems:  ci,
	}
}
