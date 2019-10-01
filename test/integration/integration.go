package main

type (
	plugin string

	auth    plugin
	secrets plugin
	storage plugin
)

type node struct {
	kind     plugin
	config   map[string]interface{}
	children []*node
}
