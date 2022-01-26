package base128_varints

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestEncoding(t *testing.T) {
	num := int32(2855731) // b: 101011 10010011 00110011
	fmt.Printf("%b\n", num)
	req := &Test2{
		A: &num,
	}
	d, err := proto.Marshal(req)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	// 1. 使用 ZigZag, 公式 (n << 1) ^ (n >> 31), 算术右移
	fmt.Printf("zigzag: %b\n", (num<<1)^(num>>31)) // b: 10101110010011001100110
	// 2. 7位一组 0000010 1011100 1001100 1100110
	// 3. 逆序  1100110  1001100  1011100  0000010
	// 4. 补位 11100110 11001100 11011100 00000010
	// 5. 补序号和类型 00001000 11100110 11001100 11011100 00000010
	// 6. 结果 0000100011100110110011001101110000000010
	fmt.Printf("d: %0 x\n", d)
}
