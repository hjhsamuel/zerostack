package parser

type APIDefinition struct {
	Syntax string
	Types  []*TypeDef
	Groups []*Group
}

type TypeDef struct {
	Name   string
	Fields []*Field
	Embed  []string
}

type Field struct {
	Name    string
	Type    string
	Tag     string
	Comment string
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
	RspType  string
	FuncName string
	Doc      *DocAnnotation
}
