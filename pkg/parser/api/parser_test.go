package api

import (
	"fmt"
	"testing"
)

func TestParseAPIFile(t *testing.T) {
	var apiFilePath string = "../../cmd/example.api"

	p, err := NewParser(apiFilePath)
	if err != nil {
		t.Fatal(err)
	}
	//for _, item := range p.ts.tokens {
	//	fmt.Printf("%#v\n", item)
	//}
	rsp, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(rsp.Syntax)
	fmt.Println("============ Type ==============")
	for _, item := range rsp.Types {
		fmt.Println(item.Name)
		fmt.Println("========> Fields <=========")
		for _, field := range item.Fields {
			fmt.Printf("%#v\n", field)
		}
	}
	fmt.Println("============ Group ===============")
	for _, item := range rsp.Groups {
		fmt.Println(item.Name)
		fmt.Printf("Meta: %#v\n", item.RouteMeta)
		fmt.Println("========> Handlers <=========")
		for _, handler := range item.Handlers {
			fmt.Printf("%#v\n", handler)
		}
	}
}
