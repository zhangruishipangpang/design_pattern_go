package design_patterns

import "testing"

func TestHandler(t *testing.T) {
	var h1 Handler = Invoke1{}
	var h2 Handler = Invoke2{}

	chain := HandlerChain{Handlers: []*Handler{&h1, &h2}}

	chain.DoHandler("1")
}
