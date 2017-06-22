package main

import (
	"fmt"
)

type Controller struct {
	bindings map[rune]func()
}

func newController() *Controller {
	bindings := make(map[rune]func())

	return &Controller{
		bindings: bindings,
	}
}

func (c *Controller) bind(key rune, function func()) error {
	if _, has := c.bindings[key]; has {
		return fmt.Errorf("Key %c is already bound.", key)
	}

	c.bindings[key] = function

	return nil
}

func (c *Controller) processKey(key rune) {
	if function, has := c.bindings[key]; has {
		function()
	}
}
