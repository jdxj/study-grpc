package non_varint

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestEncoding(t *testing.T) {
	num := uint32(2855731)
	ct := make([]byte, 4)
	binary.BigEndian.PutUint32(ct, num)
	fmt.Printf("ct: %b\n", ct) // 00000000 00101011 10010011 00110011

	req := &Test1{
		A: &num,
	}
	d, err := proto.Marshal(req)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("d: %0 x\n", d)
	fmt.Printf("len(d): %d\n", len(d))

	// 1. 8位一组 00000000 00101011 10010011 00110011
	// 2. 逆序   00110011 10010011 00101011 00000000
	// 3. 补序号和类型 00001101 00110011 10010011 00101011 00000000
	// 4. 结果             0d       33       93       2b       00
}
