package workflow

type IFlowBuilder interface {
	Name() string
	Descript() string
	Activities() *map[string]IActivity
	Transitions() *[]ITransition
}
