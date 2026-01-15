package entities

import "github.com/hjhsamuel/zerostack/pkg/parser/api"

type BaseInfo struct {
	Module    string
	Service   string
	Workspace string
	SrvHome   string
}

type ApiFile struct {
	Path string
	Api  *api.APIDefinition
}

type RouteInfo struct {
	Syntax   string
	FileName string
	Params   map[string]any
}
