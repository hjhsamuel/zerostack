package entities

import "github.com/hjhsamuel/zerostack/pkg/parser/api"

type BaseInfo struct {
	Module    string
	Service   string
	Workspace string
	SrvHome   string
}

type ConfigInfo struct {
	Server ServerInfo `json:"server"`
	Log    LogInfo    `json:"log"`
}

type ServerInfo struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Salt string `json:"salt"`
}

type LogInfo struct {
	Level    string `json:"level"`
	MaxRolls int    `json:"rolls"`
}

type ApiFile struct {
	Path string
	Api  *api.APIDefinition
}

type SyntaxInfo struct {
	Syntax string
	Groups []*RouteInfo
}

type RouteInfo struct {
	GroupName string
	Params    *GenInfo
}

type GenInfo struct {
	Module  string     `json:"module"`
	Service string     `json:"service"`
	Syntax  string     `json:"syntax"`
	Types   []*GenType `json:"types"`
	Group   *GenGroup  `json:"group"`
}

type GenType struct {
	Name   string      `json:"name"`
	Fields []*GenField `json:"fields"`
}

type GenField struct {
	Name    string    `json:"name"`
	Type    string    `json:"type"`
	Tags    []*GenTag `json:"tags"`
	Comment string    `json:"comment"`
}

type GenTag struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

type GenGroup struct {
	Name    string                 `json:"name"`
	Route   *GenRoute              `json:"route"`
	Handler map[string]*GenHandler `json:"handler"`
}

type GenRoute struct {
	Tag  string `json:"tag"`
	Auth bool   `json:"auth"`
}

type GenHandler struct {
	Method  string  `json:"method"`
	Path    string  `json:"path"`
	Req     string  `json:"req"`
	Rsp     *GenRsp `json:"rsp"`
	Handler string  `json:"handler"`
	Doc     *GenDoc `json:"doc"`
}

type GenRsp struct {
	Base bool   `json:"base"`
	Type string `json:"type"`
}

type GenDoc struct {
	Summary string `json:"summary"`
}
