package gofile

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

func ParseGoFile(path string) (*token.FileSet, *ast.File, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}
	return fset, f, nil
}

func AstToGoCode(fset *token.FileSet, f any) (string, error) {
	buf := new(bytes.Buffer)
	if err := format.Node(buf, fset, f); err != nil {
		return "", err
	}
	return buf.String(), nil
}
