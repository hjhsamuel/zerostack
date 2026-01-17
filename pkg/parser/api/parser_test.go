package api

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseAPIFile(t *testing.T) {
	var apiFilePath string = "../../../example/example.api"

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
	//fmt.Println("============ Type ==============")
	//for _, item := range rsp.Types {
	//	fmt.Println(item.Name)
	//	fmt.Println("========> Fields <=========")
	//	for _, field := range item.Fields {
	//		fmt.Printf("%#v\n", field)
	//		for _, tag := range field.Tags {
	//			fmt.Printf("=======> Tag: %#v\n", tag)
	//		}
	//	}
	//}
	fmt.Println("============ Group ===============")
	fmt.Println(rsp.Group.Name)
	fmt.Printf("Meta: %#v\n", rsp.Group.RouteMeta)
	fmt.Println("========> Handlers <=========")
	for _, item := range rsp.Group.Handlers {
		content, err := json.Marshal(item)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(content))
	}
}
