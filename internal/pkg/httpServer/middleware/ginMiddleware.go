package middleware

type GinMiddleware struct {
	RequestLogging *RequestLogging
}

func NewGinMiddleware(requestLogging *RequestLogging) *GinMiddleware {
	return &GinMiddleware{RequestLogging: requestLogging}
}
