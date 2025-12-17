package bidi

import (
	"fmt"
)

// PerformActions executes a sequence of input actions.
func (c *Client) PerformActions(context string, actions []map[string]interface{}) error {
	// If no context provided, get the first one from the tree
	if context == "" {
		tree, err := c.GetTree()
		if err != nil {
			return fmt.Errorf("failed to get browsing context: %w", err)
		}
		if len(tree.Contexts) == 0 {
			return fmt.Errorf("no browsing contexts available")
		}
		context = tree.Contexts[0].Context
	}

	params := map[string]interface{}{
		"context": context,
		"actions": actions,
	}

	_, err := c.SendCommand("input.performActions", params)
	return err
}

// Click performs a mouse click at the specified coordinates.
func (c *Client) Click(context string, x, y float64) error {
	actions := []map[string]interface{}{
		{
			"type": "pointer",
			"id":   "mouse",
			"parameters": map[string]interface{}{
				"pointerType": "mouse",
			},
			"actions": []map[string]interface{}{
				{
					"type":     "pointerMove",
					"x":        int(x),
					"y":        int(y),
					"duration": 0,
				},
				{
					"type":   "pointerDown",
					"button": 0,
				},
				{
					"type":   "pointerUp",
					"button": 0,
				},
			},
		},
	}

	return c.PerformActions(context, actions)
}

// ClickElement finds an element and clicks its center.
func (c *Client) ClickElement(context, selector string) error {
	info, err := c.FindElement(context, selector)
	if err != nil {
		return err
	}

	x, y := info.GetCenter()
	return c.Click(context, x, y)
}

// DoubleClick performs a double-click at the specified coordinates.
func (c *Client) DoubleClick(context string, x, y float64) error {
	actions := []map[string]interface{}{
		{
			"type": "pointer",
			"id":   "mouse",
			"parameters": map[string]interface{}{
				"pointerType": "mouse",
			},
			"actions": []map[string]interface{}{
				{
					"type":     "pointerMove",
					"x":        int(x),
					"y":        int(y),
					"duration": 0,
				},
				{
					"type":   "pointerDown",
					"button": 0,
				},
				{
					"type":   "pointerUp",
					"button": 0,
				},
				{
					"type":   "pointerDown",
					"button": 0,
				},
				{
					"type":   "pointerUp",
					"button": 0,
				},
			},
		},
	}

	return c.PerformActions(context, actions)
}

// MoveMouse moves the mouse to the specified coordinates.
func (c *Client) MoveMouse(context string, x, y float64) error {
	actions := []map[string]interface{}{
		{
			"type": "pointer",
			"id":   "mouse",
			"parameters": map[string]interface{}{
				"pointerType": "mouse",
			},
			"actions": []map[string]interface{}{
				{
					"type":     "pointerMove",
					"x":        int(x),
					"y":        int(y),
					"duration": 0,
				},
			},
		},
	}

	return c.PerformActions(context, actions)
}
