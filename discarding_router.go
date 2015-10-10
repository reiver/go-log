package flog


var (
	singltonDiscardRouter = DiscardRouter{}
)


// NewDiscardRouter returns an initialized DiscardRouter.
func NewDiscardRouter() *DiscardRouter {
	return &singltonDiscardRouter
}


// DiscardRouter is a Router that discards any message (and its context)
// it is asked to route.
//
// Conceptually it is similar to /dev/null
type DiscardRouter struct{}


func (router *DiscardRouter) Route(message string, context map[string]interface{}) error {
	return nil
}
