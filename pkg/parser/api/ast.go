package api

type APIDefinition struct {
	Syntax string
	Types  []*TypeDef
	Group  *Group
}

type TypeDef struct {
	Name   string
	Fields []*Field
	Embed  []string
}

type Field struct {
	Name    string
	Type    string
	Tags    []*Tag
	Comment string
}

type Tag struct {
	Key string
	Val string
}

type Group struct {
	Name      string
	RouteMeta *RouteAnnotation
	Handlers  []*Handler
}

type RouteAnnotation struct {
	Tag  string
	Auth bool
}

type DocAnnotation struct {
	Summary string
}

type Handler struct {
	Method   string
	Path     string
	ReqType  string
	RspType  *Param
	FuncName string
	Doc      *DocAnnotation
}

type Param struct {
	Base bool
	Type string
}
