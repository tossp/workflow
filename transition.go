package workflow

//迁移接口
type ITransition interface {
	Name() string             //名字
	Condition() ICondition    //条件
	FromActivity() IActivity  //起始环节
	ToActivity() IActivity    //截至环节
	Validate() bool           //验证环节是否正确
	Eval(args EventArgs) bool //计算是否通过
}

//迁移
type Transition struct {
	name         string     //名字
	condition    ICondition //条件
	fromActivity IActivity  //起始环节
	toActivity   IActivity  //截至环节
}

func New_Transition(name string, condition ICondition, fromAct, toAct IActivity) *Transition {
	return &Transition{
		name:         name,
		condition:    condition,
		fromActivity: fromAct,
		toActivity:   toAct,
	}
}

//-------ITransition interface----------
func (this *Transition) Name() string {
	return this.name
}

func (this *Transition) Condition() ICondition {
	return this.condition
}

func (this *Transition) FromActivity() IActivity {
	return this.fromActivity
}

func (this *Transition) ToActivity() IActivity {
	return this.toActivity
}

func (this *Transition) Validate() bool {
	return true
}

func (this *Transition) Eval(args EventArgs) bool {
	if this.condition == nil {
		return true
	}
	return this.condition.Eval()
}

//-------ITransition interface----------
