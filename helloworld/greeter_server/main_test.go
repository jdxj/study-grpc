package main

import (
	"fmt"
	"net"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

type myValue struct {
	p1 int64
	p2 int64
	p3 int64
}

func TestServer_SayHello(t *testing.T) {
	value := 3.4
	v := reflect.ValueOf(&value)
	tmp := (*myValue)(unsafe.Pointer(&v))
	fmt.Printf("%d\n", tmp.p2)
	fmt.Printf("%d\n", &value)

	p := v.Elem()
	tmp2 := (*myValue)(unsafe.Pointer(&p))
	fmt.Printf("%d\n", tmp2.p2)
}

func TestPointer(t *testing.T) {
	value := 3.4
	p := unsafe.Pointer(&value)
	fmt.Printf("%x\n", &value)
	fmt.Printf("%x\n", p)
	p2 := &p
	fmt.Printf("%x\n", p2)
}

type MyI interface {
	P()
}

type ab struct {
}

func (aa *ab) P() {

}
func TestInterfaceP(t *testing.T) {
	var i interface{} = (*MyI)(nil)
	//var a *ab
	//var mi MyI = a
	//var i interface{}
	//var i interface{} = MyI(nil)

	v := reflect.TypeOf(i)
	fmt.Printf("v.Kind: %v\n", v.Kind())

	p := v.Elem()
	fmt.Printf("p.Kind: %v\n", p.Kind())
	fmt.Printf("p.Name: %v\n", p.Name())

	p2 := p.Elem()
	fmt.Printf("p2.Kind: %v\n", p2.Kind())
}

func TestTCPDeadline(t *testing.T) {
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		t.Fatalf("%s\n", err)

	}
	go func() {
		con, err := lis.Accept()
		if err != nil {
			t.Fatalf("%s\n", err)
		}

		con.SetDeadline(time.Now().Add(5 * time.Second))

		buf := make([]byte, 1024)
		fmt.Printf("read time: %s\n", time.Now())
		_, err = con.Read(buf)
		if err != nil {
			fmt.Printf("time: %s, err: %s\n", time.Now(), err)
		}
	}()

	con, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	con.SetDeadline(time.Now().Add(10 * time.Second))

	time.Sleep(time.Minute)
}
