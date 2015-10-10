package flog


func NewDiscardRouter() *DiscardRouter {
	router := DiscardRouter{}

	return &router
}


type DiscardRouter struct{}


func (router *DiscardRouter) Route(message string, context map[string]interface{}) error {
	return nil
}
