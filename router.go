package log


type Router interface {
	Route(message string, context map[string]interface{}) error
}
