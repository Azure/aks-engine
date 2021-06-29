// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.
package armhelpers

import (
	"fmt"
	"reflect"
)

// DeepCopy dst and src should be the same type in different API version
// dst should be pointer type
func DeepCopy(dst, src interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("fail to copy object %v", r)
		}
	}()
	dstValue := reflect.ValueOf(dst)
	srcValue := reflect.ValueOf(src)
	if dstValue.Kind() != reflect.Ptr {
		return fmt.Errorf("The dst must be Ptr")
	}
	dstValue = dstValue.Elem()
	if dstValue.Type().String() != srcValue.Type().String() {
		return fmt.Errorf("the dst type (%q) and src type (%q) are not the same", dstValue.Type().String(), srcValue.Type().String())
	}
	deepCopyInternal(dstValue, srcValue, 0)
	return err
}

func deepCopyInternal(dstValue, srcValue reflect.Value, depth int) {
	if dstValue.CanSet() {
		switch srcValue.Kind() {
		case reflect.Bool:
			dstValue.SetBool(srcValue.Bool())
		case reflect.String:
			dstValue.SetString(srcValue.String())
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
			dstValue.SetInt(srcValue.Int())
		case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
			dstValue.SetUint(srcValue.Uint())
		case reflect.Float64, reflect.Float32:
			dstValue.SetFloat(srcValue.Float())
		case reflect.Complex64, reflect.Complex128:
			dstValue.SetComplex(srcValue.Complex())
		case reflect.Ptr:
			if !srcValue.IsNil() {
				d := reflect.New(dstValue.Type().Elem())
				dstValue.Set(d)
				deepCopyInternal(dstValue.Elem(), srcValue.Elem(), depth+1)
			}
		case reflect.Slice:
			if !srcValue.IsNil() {
				d := reflect.MakeSlice(dstValue.Type(), srcValue.Len(), srcValue.Cap())
				dstValue.Set(d)
				for i := 0; i < srcValue.Len(); i++ {
					v := dstValue.Index(i)
					deepCopyInternal(v, srcValue.Index(i), depth+1)
					v.Set(v)
				}
			}
		case reflect.Array:
			d := reflect.New(dstValue.Type()).Elem()
			for i := 0; i < srcValue.Len(); i++ {
				v := reflect.New(srcValue.Index(i).Type()).Elem()
				deepCopyInternal(v, srcValue.Index(i), depth+1)
				d.Index(i).Set(v)
			}
			dstValue.Set(d)

		case reflect.Map:
			if !srcValue.IsNil() {
				d := reflect.MakeMap(dstValue.Type())
				for _, key := range srcValue.MapKeys() {
					v := reflect.New(srcValue.MapIndex(key).Type()).Elem()
					deepCopyInternal(v, srcValue.MapIndex(key), depth+1)
					d.SetMapIndex(key, v)
				}
				dstValue.Set(d)
			}
		case reflect.Struct:
			for i := 0; i < srcValue.NumField(); i++ {
				srcField := srcValue.Field(i)
				dstField := dstValue.FieldByName(srcValue.Type().Field(i).Name)
				if dstField.IsValid() && dstField.CanAddr() && dstField.CanSet() {
					deepCopyInternal(dstField, srcField, depth+1)
				}
			}
		default:
		}
	}
}
