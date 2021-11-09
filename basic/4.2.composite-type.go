// 复合类型

package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// 数组
	arr1 := [3]int8{1, 2, 3}
	arr2 := [...]int{1, 2, 3, 4, 5}
	arr3 := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("array:", arr1[0], arr2[4], arr3[1][0])

	// slice - 引用类型
	silce1 := []int{}
	silce2 := []int{1, 2, 3}
	silce3 := arr2[1:3]
	silce4 := arr2[:]

	silce3[0] = 22

	fmt.Println("slice:", silce1, silce2[2], silce3[0], silce4[1], len(silce4))

	// map - 引用类型
	m1 := map[string]int{"c": 1, "go": 2}
	m2 := make(map[string]int)
	m2 = m1

	delete(m1, "c")
	info, ok := m2["c"]

	// 获取成功
	if ok {
		fmt.Println("get c", info)
	} else {
		fmt.Println("no c info found")
	}

	m3 := make(map[string]int)
	fmt.Println("m3", m3)

	/**
		- make 用于给 slice / map / channel 这几个内建复合类型分配内存, 初始化变量 并返回引用
		- new 用于给其他自定义类型分配内存，未初始化变量，值为 零 值，返回指针
	**/

	varnumpar()
}

// 传递变长参数
func varnumpar() {
	x := arrayLen(1, 3, 2, 0)
	fmt.Printf("arrayLen: %d\n", x)

	a := make([]interface{}, 5)
	x = arrayLen(a...)

	fmt.Printf("arrayLen: %d", x)
}

func arrayLen(a ...interface{}) int {
	return len(a)
}

func convertAssignRows(dest, src interface{}) error {
	// Common cases, without reflect.
	switch s := src.(type) {
	case string:
		switch d := dest.(type) {
		case *string:
			if d == nil {
				return nil
			}
			*d = s
			return nil
		case *[]byte:
			if d == nil {
				return nil
			}
			*d = []byte(s)
			return nil
		}
	case []byte:
		switch d := dest.(type) {
		case *string:
			if d == nil {
				return nil
			}
			*d = string(s)
			return nil
		}
	case nil:
		switch d := dest.(type) {
		case *interface{}:
			if d == nil {
				return nil
			}
			*d = nil
			return nil
		case *[]byte:
			if d == nil {
				return nil
			}
			*d = nil
			return nil
		case *RawBytes:
			if d == nil {
				return nil
			}
			*d = nil
			return nil
		}

	var sv reflect.Value

	switch d := dest.(type) {
	case *string:
		sv = reflect.ValueOf(src)
		switch sv.Kind() {
		case reflect.Bool,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			*d = asString(src)
			return nil
		}
	case *[]byte:
		sv = reflect.ValueOf(src)
		if b, ok := asBytes(nil, sv); ok {
			*d = b
			return nil
		}
	case *RawBytes:
		sv = reflect.ValueOf(src)
		if b, ok := asBytes([]byte(*d)[:0], sv); ok {
			*d = RawBytes(b)
			return nil
		}
	case *bool:
		bv, err := driver.Bool.ConvertValue(src)
		if err == nil {
			*d = bv.(bool)
		}
		return err
	case *interface{}:
		*d = src
		return nil
	}

	if scanner, ok := dest.(Scanner); ok {
		return scanner.Scan(src)
	}

	dpv := reflect.ValueOf(dest)
	if dpv.Kind() != reflect.Ptr {
		return errors.New("destination not a pointer")
	}
	if dpv.IsNil() {
		return nil
	}

	if !sv.IsValid() {
		sv = reflect.ValueOf(src)
	}

	dv := reflect.Indirect(dpv)
	if sv.IsValid() && sv.Type().AssignableTo(dv.Type()) {
		switch b := src.(type) {
		case []byte:
			dv.Set(reflect.ValueOf(cloneBytes(b)))
		default:
			dv.Set(sv)
		}
		return nil
	}

	if dv.Kind() == sv.Kind() && sv.Type().ConvertibleTo(dv.Type()) {
		dv.Set(sv.Convert(dv.Type()))
		return nil
	}

	// The following conversions use a string value as an intermediate representation
	// to convert between various numeric types.
	//
	// This also allows scanning into user defined types such as "type Int int64".
	// For symmetry, also check for string destination types.
	switch dv.Kind() {
	case reflect.Ptr:
		if src == nil {
			dv.Set(reflect.Zero(dv.Type()))
			return nil
		}
		dv.Set(reflect.New(dv.Type().Elem()))
		return convertAssignRows(dv.Interface(), src, rows)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if src == nil {
			return fmt.Errorf("converting NULL to %s is unsupported", dv.Kind())
		}
		s := asString(src)
		i64, err := strconv.ParseInt(s, 10, dv.Type().Bits())
		if err != nil {
			err = strconvErr(err)
			return fmt.Errorf("converting driver.Value type %T (%q) to a %s: %v", src, s, dv.Kind(), err)
		}
		dv.SetInt(i64)
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if src == nil {
			return fmt.Errorf("converting NULL to %s is unsupported", dv.Kind())
		}
		s := asString(src)
		u64, err := strconv.ParseUint(s, 10, dv.Type().Bits())
		if err != nil {
			err = strconvErr(err)
			return fmt.Errorf("converting driver.Value type %T (%q) to a %s: %v", src, s, dv.Kind(), err)
		}
		dv.SetUint(u64)
		return nil
	case reflect.Float32, reflect.Float64:
		if src == nil {
			return fmt.Errorf("converting NULL to %s is unsupported", dv.Kind())
		}
		s := asString(src)
		f64, err := strconv.ParseFloat(s, dv.Type().Bits())
		if err != nil {
			err = strconvErr(err)
			return fmt.Errorf("converting driver.Value type %T (%q) to a %s: %v", src, s, dv.Kind(), err)
		}
		dv.SetFloat(f64)
		return nil
	case reflect.String:
		if src == nil {
			return fmt.Errorf("converting NULL to %s is unsupported", dv.Kind())
		}
		switch v := src.(type) {
		case string:
			dv.SetString(v)
			return nil
		case []byte:
			dv.SetString(string(v))
			return nil
		}
	}

	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", src, dest)
}
