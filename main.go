package main

import patterns "go_project/design_patterns"

func main() {

	var h1 patterns.Handler = patterns.Invoke1{}
	var h2 patterns.Handler = patterns.Invoke2{}

	chain := patterns.HandlerChain{Handlers: []*patterns.Handler{&h1, &h2}}

	chain.DoHandler("1")
}
