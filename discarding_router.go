package log


var (
	singltonDiscardingRouter = DiscardingRouter{}
)


// NewDiscardingRouter returns an initialized DiscardingRouter.
func NewDiscardingRouter() *DiscardingRouter {
	return &singltonDiscardingRouter
}


// DiscardingRouter is a Router that discards any message (and its context)
// it is asked to route.
//
// Conceptually it is similar to /dev/null
type DiscardingRouter struct{}


func (router *DiscardingRouter) Route(message string, context map[string]interface{}) error {
	return nil
}
