package simple_message

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestEncoding(t *testing.T) {
	num := int32(2855731) // b: 101011 10010011 00110011
	fmt.Printf("%b\n", num)
	req := &Test1{
		A: &num,
	}
	d, err := proto.Marshal(req)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	// 1. 7位一组 0000001  0101110  0100110  0110011
	// 2. 逆序    0110011  0100110  0101110  0000001
	// 3. 补位   10110011 10100110 10101110 00000001
	// 4. 补序号和类型 00001000 10110011 10100110 10101110 00000001
	// 5. 结果 0000100010110011101001101010111000000001
	fmt.Printf("d: %0 x\n", d)
	fmt.Printf("len(d): %d\n", len(d))
}

func TestEncodingNegative(t *testing.T) {
	num := int64(-2855731)
	ct := make([]byte, 8)
	binary.BigEndian.PutUint64(ct, uint64(num)) // 11111111 11111111 11111111 11111111 11111111 11010100 01101100 11001101
	fmt.Printf("ct: %b\n", ct)

	req := &Test2{
		B: &num,
	}
	d, err := proto.Marshal(req)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	// 1. 7位一组 0000001 1111111 1111111 1111111 1111111 1111111 1111110 1010001 1011001 1001101
	// 2. 逆序    1001101  1011001  1010001  1111110  1111111  1111111  1111111  1111111  1111111  0000001
	// 3. 补位   11001101 11011001 11010001 11111110 11111111 11111111 11111111 11111111 11111111 00000001
	// 4. 补序号和类型 00001000 11001101 11011001 11010001 11111110 11111111 11111111 11111111 11111111 11111111 00000001
	//                     08       cd       d9       d1       fe       ff       ff       ff       ff       ff       01
	fmt.Printf("d: %0 x\n", d)
	fmt.Printf("len(d): %d\n", len(d))
}
