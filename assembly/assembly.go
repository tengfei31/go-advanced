package assembly

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"
)

var Id int

var NameData [8]byte
var Name string

var helloWorld = "你好，世界"

func PrintHelloWorld()

var Num [2]int

var (
	boolValue  bool
	trueValue  bool
	falseValue bool
)

var int32Value int32
var uint32Value uint32

var (
	float32Value float32
	float64Value float64
)

var helloworld string

var sliceValue []byte

var mapValue map[string]int
var chanValue chan int

//go:nosplit
func Swap(int, int) (int, int)

//go:nosplit
func Foo(bool, int16) []byte

//go:nosplit
func Foo1()

//go:nosplit
func If(bool, int, int) int

//go:nosplit
func LoopAdd(int, int, int) int

//go:nosplit
func SyscallWrite_Darwin(fd int, msg string) int

//go:nosplit
func CopySlice_AVX2(dst, src []byte, len int)

//go:nosplit
func getG() unsafe.Pointer


// Test 测试汇编代码
func Test() {
	fmt.Println(Id)
	fmt.Println(Name)
	NameData[0] += 1
	fmt.Println(Name)
	//PrintHelloWorld()
	fmt.Println(Num)
	fmt.Println(boolValue)
	fmt.Println(trueValue)
	fmt.Println(falseValue)
	fmt.Println("int32Value=", int32Value)
	fmt.Println("uint32Value=", uint32Value)
	fmt.Println("float32Value=", float32Value)
	fmt.Println("float64Value=", float64Value)
	fmt.Println("helloworld=", helloworld)
	fmt.Println("sliceValue=", string(sliceValue))
	fmt.Println("mapValue=", mapValue)
	fmt.Println("chanValue=", chanValue)
	var a, b = Swap(1, 2)
	fmt.Printf("Swap(1, 2)=%d, %d\n", a, b)
	//var c = Foo(false, 8)
	//fmt.Println("Foo(false, 8)=", string(c))
	fmt.Printf("If(false, %d, %d)=%d\n", 1, 2, If(false, 1, 2))
	fmt.Printf("LoopAdd(%d, %d, %d)=%d\n", 3, 1, 1, LoopAdd(3, 1, 1))
	/*type Tree struct {
		Node string `json:"node"`
	}
	var tree Tree = Tree{
		Node: "wtf",
	}
	rv := reflect.ValueOf(tree)
	fmt.Println(rv.Kind(), rv.IsNil())*/
	var base int64 = 10
	fmt.Printf("1 << %d = %v\n", base, 1<<base)

	if runtime.GOOS == "darwin" {
		SyscallWrite_Darwin(1, "hello syscall!\n")
	}

	var dst, src []byte
	src = []byte("hello,hello,hello,hello,hello,hello,hello,hello,hello,hello,hello,hello")
	dst = make([]byte, len(src))
	CopySlice_AVX2(dst, src, len(src))
	fmt.Printf("CopySlice_AVX2(dst, []byte(\"hello\"), len(src)), dst=%s\n", string(dst))

	//打印goroutine
	fmt.Printf("GetGoroutineId()=%d\n", GetGoroutineId())
}

var offsetDictMap = map[string]int64{
	"go1.16": 152,
	"go1.10": 152,
	"go1.9":  152,
	"go1.8":  192,
}
var goidOffset = func() int64 {
	goversion := runtime.Version()
	for key, off := range offsetDictMap {
		if goversion == key || strings.HasPrefix(goversion, key) {
			return off
		}
	}
	panic("unsupport go verion:"+goversion)
}()
// GetGoroutineId 打印goroutine
func GetGoroutineId() int64 {
	g := getG()
	p := (*int64)(unsafe.Pointer(uintptr(g) + uintptr(goidOffset)))
	return *p
}
