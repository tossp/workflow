package workflow

//迁移条件
type ICondition interface {
	Name() string
	NotOp() bool
	Eval() bool
}

//条件
type Condition struct {
	name        string      //名字
	notOp       bool        //是否取反
	EvalHandler func() bool //条件计算函数
}

func New_Condition(name string, notOp bool) *Condition {
	return &Condition{
		name:  name,
		notOp: notOp,
	}
}

//-------ICondition interface----------
func (this *Condition) Name() string {
	return this.name
}

func (this *Condition) NotOp() bool {
	return this.notOp
}

func (this *Condition) Eval() bool {
	result := true
	if this.EvalHandler != nil {
		result = this.EvalHandler()
	}
	return this.notOpEval(result)
}

func (this *Condition) notOpEval(result bool) bool {
	if this.notOp {
		return !result
	}
	return result
}

//-------ICondition interface----------
