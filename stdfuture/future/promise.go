package future

type Promise interface {
	SetValue(interface{})
	SetException(interface{})
	MakeFuture() Future
}

type PromiseData struct {
	state     string
	result    interface{}
	exception interface{}
}

func NewPromise() *PromiseData {
	return &PromiseData{state: "PENDING"}
}

func (pd *PromiseData) SetValue(result interface{}) {
	pd.result = result
	pd.state = "COMPLETED"
}

func (pd *PromiseData) SetException(error interface{}) {
	pd.exception = error
	pd.state = "FAILED"
}

func (pd *PromiseData) MakeFuture() Future {
	return &FutureData{promise: pd}
}
