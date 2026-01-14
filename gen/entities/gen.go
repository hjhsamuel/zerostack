package entities

import "github.com/hjhsamuel/zerostack/pkg/parser/api"

type BaseInfo struct {
	Module  string
	Service string
	Home    string
}

type ApiFile struct {
	Path string
	Api  *api.APIDefinition
}
