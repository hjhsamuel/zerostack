package gofile

import (
	"go/ast"
	"strings"
)

type FuncSignature struct {
	Recv    string   // Api, *Api
	Name    string   // GetUser
	Params  []string // *gp.Context, *schema.GetUserReq
	Results []string // *schema.GetUserRsp, error
}

func BuildFunSignature(fn *ast.FuncDecl) FuncSignature {
	return FuncSignature{
		Recv:    extractReceiver(fn),
		Name:    fn.Name.Name,
		Params:  extractFieldTypes(fn.Type.Params),
		Results: extractFieldTypes(fn.Type.Results),
	}
}

func exprString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name

	case *ast.StarExpr:
		return "*" + exprString(t.X)

	case *ast.SelectorExpr:
		return exprString(t.X) + "." + t.Sel.Name

	case *ast.ArrayType:
		return "[]" + exprString(t.Elt)

	case *ast.MapType:
		return "map[" + exprString(t.Key) + "]" + exprString(t.Value)

	case *ast.Ellipsis:
		return "..." + exprString(t.Elt)

	case *ast.InterfaceType:
		return "interface{}"

	default:
		return "<unknown>"
	}
}

func extractReceiver(f *ast.FuncDecl) string {
	if f.Recv == nil || len(f.Recv.List) == 0 {
		return ""
	}
	return exprString(f.Recv.List[0].Type)
}

func extractFieldTypes(fields *ast.FieldList) []string {
	if fields == nil {
		return nil
	}

	var res []string
	for _, f := range fields.List {
		typ := exprString(f.Type)

		// 多个参数共用一个类型
		count := len(f.Names)
		if count == 0 {
			count = 1
		}

		for i := 0; i < count; i++ {
			res = append(res, typ)
		}
	}
	return res
}

func IsSameFuncSignature(a, b FuncSignature) bool {
	if a.Recv != b.Recv {
		return false
	}
	if a.Name != b.Name {
		return false
	}

	if len(a.Params) != len(b.Params) {
		return false
	}
	for index, item := range a.Params {
		if item != b.Params[index] {
			return false
		}
	}

	if len(a.Results) != len(b.Results) {
		return false
	}
	for index, item := range a.Results {
		if item != b.Results[index] {
			return false
		}
	}

	return true
}

type SwaggerSignature struct {
	FuncName string
	Summary  string
	Tag      string
	Req      string
	Rsp      string
	Path     string
	Method   string
}

func BuildSwaggerSignature(fn *ast.FuncDecl) SwaggerSignature {
	meta := SwaggerSignature{}
	if fn.Doc == nil {
		return meta
	}
	for index, comment := range fn.Doc.List {
		line := strings.TrimSpace(strings.TrimPrefix(comment.Text, "//"))
		if strings.HasPrefix(line, "@") {
			parseSwaggerLine(&meta, line)
		} else {
			if index == 0 {
				meta.FuncName = line
			}
		}
	}
	return meta
}

func parseSwaggerLine(meta *SwaggerSignature, line string) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "@Summary":
		meta.Summary = strings.Join(parts[1:], " ")
	case "@Tags":
		meta.Tag = strings.Join(parts[1:], " ")
	case "@Param":
		if len(parts) >= 4 {
			meta.Req = parts[3]
		}
	case "@Success":
		meta.Rsp = strings.TrimSuffix(strings.TrimPrefix(parts[len(parts)-1], "gp.Response{data="), "}")
	case "@Router":
		if len(parts) != 3 {
			return
		}
		meta.Path = parts[1]
		meta.Method = strings.TrimSuffix(strings.TrimPrefix(parts[2], "["), "]")
	default:
		return
	}
}

func IsSameSwaggerSignature(a, b SwaggerSignature) bool {
	if a.FuncName != b.FuncName {
		return false
	}
	if a.Summary != b.Summary {
		return false
	}
	if a.Tag != b.Tag {
		return false
	}
	if a.Req != b.Req {
		return false
	}
	if a.Rsp != b.Rsp {
		return false
	}
	if a.Path != b.Path {
		return false
	}
	if a.Method != b.Method {
		return false
	}
	return true
}
