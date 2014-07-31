package workflow

type FlowPool struct {
	Pool     map[string]IWorkflow //所有流程的定义
	Provider IFlowProvider        //流程读取器
}

func New_FlowPool() *FlowPool {
	return &FlowPool{
		Pool:     make(map[string]IWorkflow),
		Provider: FlowProviderFactory(),
	}
}

func (this *FlowPool) GetFlowDef(flowid string) (IWorkflow, bool) {
	if v, ok := this.Pool[flowid]; ok {
		return v, true
	}
	return nil, false
}

func (this *FlowPool) RemoveFlowDef(flowid string) {
	delete(this.Pool, flowid)
}

func (this *FlowPool) LoadFlowDef(flowid string) {
	flowXML := this.Provider.GetFlow(flowid)
	f := New_Workflow(flowXML)
	this.Pool[flowid] = f
}

func (this *FlowPool) ReloadFlowDef(flowid string) {
	this.RemoveFlowDef(flowid)
	this.LoadFlowDef(flowid)
}

func (this *FlowPool) ReloadAllFlow() {
	xmlDefs := this.Provider.GetFlowDefines()
	for k, _ := range this.Pool {
		delete(this.Pool, k)
	}
	for k, v := range xmlDefs {
		f := New_Workflow(v)
		this.Pool[k] = f
	}
}
