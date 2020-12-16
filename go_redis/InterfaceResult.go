package go_redis

type InterfaceResult struct {
	Result interface{}
	Err error
}

func NewInterfaceResult(result interface{}, err error) *InterfaceResult {
	return &InterfaceResult{Result: result, Err: err}
}

func(this *InterfaceResult) Unwrap() interface{} {
	if this.Err!=nil{
		panic(this.Err)
	}
	return this.Result
}

func(this *InterfaceResult) Unwrap_Or(v interface{}) interface{} {
	if this.Err!=nil{
		return v
	}
	return this.Result
}
