package cmd

import _ "embed"

//go:embed example.api
var apiExample string

var (
	// VarBoolInit initialization service or not
	VarBoolInit bool
	// VarStringApiFile the api file path
	VarStringApiFile string
	// VarBoolSwagger create swagger doc or not
	VarBoolSwagger bool
	// VarStringDB the type of database
	VarStringDB string
	// VarStringModule the module name in go.mod
	VarStringModule string
)
