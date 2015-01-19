package bud

// import (
// 	"fmt"
// 	"github.com/tbud/base/encoding/json"
// 	"reflect"
// 	// "encoding/json"
// )

// type CCC struct {
// 	Abc
// 	Name string
// }

// type Abc struct {
// 	AbcEfg string
// 	Tttt   int
// }

// func main() {
// 	// a := Abc{"py", 1}
// 	// a := &Abc{"py", 1}
// 	// a := CCC{{"py", 1}, "ttt"}
// 	// var a CCC
// 	// a.AbcEfg = "py"
// 	// a.Tttt = 1
// 	// a.Name = "tt"

// 	a := &CCC{Abc{"py", 1}, "tt"}

// 	fmt.Printf("%v\n", reflect.ValueOf(a).Type())

// 	b, e := json.Marshal(a)
// 	if e != nil {
// 		return
// 	}

// 	fmt.Printf("Hi, bud!%v\n", string(b))

// 	bb := true
// 	fmt.Printf(``, bb)
// }
