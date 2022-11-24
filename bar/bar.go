package bar

import (
	"fmt"
	"reflect"
)

type t struct{}

func PackagePath() {
	var typ t
	fmt.Println(reflect.TypeOf(typ).PkgPath())
}
