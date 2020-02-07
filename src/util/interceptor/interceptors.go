package interceptor

type Interceptor struct{}

func InterceptorHandler() *Interceptor {
	return &Interceptor{}
}

type InterceptorInterface interface {
	InterceptorMessage()
}

func InterceptorMessage() {

}
