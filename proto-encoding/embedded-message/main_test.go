package embedded_message

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestEncoding(t *testing.T) {
	num := int32(150)
	req1 := &Test1{
		A: &num,
	}
	req2 := &Test2{
		B: req1,
	}

	d, err := proto.Marshal(req2)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("d: %0 x\n", d)
	fmt.Printf("len(d): %d\n", len(d))
	// 补序号和类型 00010010
	// 结果             12
}
