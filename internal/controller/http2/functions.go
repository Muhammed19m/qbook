package http2

func AdaptToRW(handle HandleFunc) HandleFuncRW {
	return func(ctx ContextRW) (any, error) {
		return handle(ctx)
	}
}

func WrapHandlerInMidllewares(handler HandleFunc, midllewares ...Middleware) HandleFuncRW {
	hmw := AdaptToRW(handler)

	for _, mw := range midllewares {
		hmw = mw(hmw)
	}
	return hmw
}
