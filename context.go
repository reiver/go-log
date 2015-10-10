package flog


func newContext(cascade ...interface{}) map[string]interface{} {

	context := make(map[string]interface{})

	for _,x := range cascade {
		switch xx := x.(type) {
		case map[string]string:
			for key, value := range xx {
				context[key] = value
			}
		case map[string]interface{}:
			for key, value := range xx {
				context[key] = value
			}
		case string:
			context["text"] = xx
		}
	}

	return context
}
