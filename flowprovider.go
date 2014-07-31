package workflow

//--------------------流程定义提供者类----------------------------------------
type IFlowProvider interface {
	GetFlowDefines() map[string]string
	GetFlow(flowid string) string
}

func FlowProviderFactory() IFlowProvider {
	//todo: 根据配置文件, 选择不同的provider
	var p IFlowProvider
	p = new(PgFlowProvider)
	return p
}

type PgFlowProvider struct {
}

func (this *PgFlowProvider) GetFlowDefines() map[string]string {
	//sql := ""
	//todo: load from postgresql
	flows := make(map[string]string)
	return flows
}

func (this *PgFlowProvider) GetFlow(flowid string) string {
	//sql := ""
	//todo: load from postgresql
	flow := ``
	return flow
}

//------------------流程实例提供者类---------------------------------------
type ICaseProvider interface {
	GetCase(caseid string) *FlowCase
	GetFlowItems(caseid string) *[]FlowItem
}

func CaseProviderFactory() ICaseProvider {
	//todo: 根据配置文件, 选择不同的provider
	var p ICaseProvider
	p = new(PgCaseProvider)
	return p
}

type PgCaseProvider struct {
}

func (this *PgCaseProvider) GetCase(caseid string) *FlowCase {
	return new(FlowCase)
}
func (this *PgCaseProvider) GetFlowItems(caseid string) *[]FlowItem {
	its := make([]FlowItem, 10)
	return &its
}
