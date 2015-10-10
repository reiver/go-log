package flog


var (
	singltonDiscardRouter = DiscardRouter{}
)

func NewDiscardRouter() *DiscardRouter {
	return &singltonDiscardRouter
}


type DiscardRouter struct{}


func (router *DiscardRouter) Route(message string, context map[string]interface{}) error {
	return nil
}
