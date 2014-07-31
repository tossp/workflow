package workflow

//流程参数
type EventArgs struct {
	AppData    *map[string]string //业务数据
	CurrStatus string             //当前状态
	CurrUser   string             //当前用户
	UserChoice string             //用户选择结果
}

//流程环节接口对象
type IActivity interface {
	GetName() string
	GetDescript() string
	CanEnter(args *EventArgs) bool //是否可以进入
	CanExit(args *EventArgs) bool  //是否可以退出
	Enter(args *EventArgs)         //进入
	Execute(args *EventArgs)       //执行
	Exit(args *EventArgs)          //退出
	IsRoot() bool                  //是否跟节点
}

//流程环节
type Activity struct {
	name          string                        //环节名称
	descript      string                        //描述
	isRoot        bool                          //是否根节点
	EnterMode     ActivityRunMode               //进入模式
	ExitMode      ActivityRunMode               //退出模式
	JoinMode      TransitJoinMode               //迁入模式
	SplitMode     TransitSplitMode              //迁出模式
	CheckEnter    func(*map[string]string) bool //检查是否可以进入的函数
	CheckExit     func(*map[string]string) bool //检查是否可以退出的函数
	EnterHandler  func(*EventArgs)              //进入事件处理器
	ExcuteHandler func(*EventArgs)              //执行事件处理
	ExitHandler   func(*EventArgs)              //退出事件处理器
}

func New_Activity(name string) *Activity {
	act := &Activity{
		name:      name,
		EnterMode: ActivityRunMode_Auto,
		ExitMode:  ActivityRunMode_Auto,
		JoinMode:  TransitJoinMode_XOR,
		SplitMode: TransitSplitMode_XOR,
	}
	return act
}

//-------IActivity interface----------
func (this *Activity) GetName() string {
	return this.name
}
func (this *Activity) GetDescript() string {
	return this.descript
}

func (this *Activity) CanEnter(args *EventArgs) bool {
	if this.CheckEnter == nil {
		return true
	}
	return this.CheckEnter(args.AppData)
}

func (this *Activity) CanExit(args *EventArgs) bool {
	if this.CanExit == nil {
		return true
	}
	return this.CheckExit(args.AppData)
}

func (this *Activity) Enter(args *EventArgs) {
	if this.EnterHandler != nil {
		this.EnterHandler(args)
	}
}

func (this *Activity) Excute(args *EventArgs) {
	if this.ExcuteHandler != nil {
		this.ExcuteHandler(args)
	}
}

func (this *Activity) Exit(args *EventArgs) {
	if this.ExitHandler != nil {
		this.ExitHandler(args)
	}
}

func (this *Activity) IsRoot() bool {
	return this.isRoot
}

//-------IActivity interface----------
