package workflow

//流程对象接口
type IWorkflow interface {
	GetName() string                           //名字
	GetDescript() string                       //描述
	GetRoot() IActivity                        //根节点
	GetActivity(name string) IActivity         //获取指定名字的节点
	GetTransition(from, to string) ITransition //获取两个节点之间的定向迁移
	Activities() *map[string]IActivity         //所有环节
	Transitions() *[]ITransition               //所有迁移
	//AppData() *map[string]string               //业务数据
}

//流程对象
type Workflow struct {
	name        string
	descript    string
	activities  map[string]IActivity
	transitions []ITransition
	//appData     map[string]string
}

func New_Workflow(flowDef string) *Workflow {
	//todo: flowDef 是xml的流程定义, 这里进行解析, 进行赋值
	flow := &Workflow{
		name:     "",
		descript: "",
	}
	return flow
}

//-------IWorkflow interface----------
func (this *Workflow) GetName() string {
	return this.name
}

func (this *Workflow) GetDescript() string {
	return this.descript
}

func (this *Workflow) GetRoot() IActivity {
	for _, v := range this.activities {
		if v.IsRoot() {
			return v
		}
	}
	return nil
}

func (this *Workflow) GetActivity(name string) IActivity {
	for k, v := range this.activities {
		if k == name {
			return v
		}
	}
	return nil
}

func (this *Workflow) GetTransition(from, to string) ITransition {
	for _, v := range this.transitions {
		if v.FromActivity().GetName() == from && v.ToActivity().GetName() == to {
			return v
		}
	}
	return nil
}

func (this *Workflow) Activities() *map[string]IActivity {
	return &this.activities
}

func (this *Workflow) Transitions() *[]ITransition {
	return &this.transitions
}

//func (this *Workflow) AppData() *map[string]string {
//	return &this.appData
//}

//-------IWorkflow interface----------
