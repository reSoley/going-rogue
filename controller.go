package main

import ()

type Controller struct {
	bindings map[rune]func()
}

func newController() *Controller {
	bindings := make(map[rune]func())

	return &Controller{
		bindings: bindings,
	}
}

func (c *Controller) bind(key rune, function func()) {
	c.bindings[key] = function
}

func (c *Controller) processKey(key rune) {
	if function, has := c.bindings[key]; has {
		function()
	}
}
