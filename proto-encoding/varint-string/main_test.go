package varint_string

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestEncoding(t *testing.T) {
	str := "testing"
	fmt.Printf("str: %0 x\n", str)

	req := &Test1{
		A: &str,
	}

	d, err := proto.Marshal(req)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("d: %0 x\n", d)

	// 序号和类型 00001010
}
