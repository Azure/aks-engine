package symdiff

import (
	"errors"
	"reflect"
	"unsafe"
)

var (
	ErrDifferentArgumentsTypes = errors.New("src and dst must be of same type")
	ErrNilArguments            = errors.New("src and dst must not be nil")
)

// During deepSymmetricDifference, we need to keep track of
// checks that are in progress. The comparison algorithm
// assumes that all checks in progress are true when it
// re-encounters them. Visited comparisons are stored in a
// map indexed by visit.
type visit struct {
	a1  unsafe.Pointer
	a2  unsafe.Pointer
	typ reflect.Type
}

// Diff finds the symmetric difference between two interfaces.
func Diff(dst, src interface{}) error {
	if dst == nil || src == nil {
		return ErrNilArguments
	}

	vDst, vSrc, err := resolveValues(dst, src)
	if err != nil {
		return err
	}

	if vDst.Type() != vSrc.Type() {
		return ErrDifferentArgumentsTypes
	}

	return deepSymmetricDifference(vDst, vSrc, make(map[visit]bool), 0)
}

// Merges the deep symmetric difference using reflected types into dst. The map argument tracks
// comparisons that have already been seen, which allows short circuiting on recursive types.
func deepSymmetricDifference(dst, src reflect.Value, visited map[visit]bool, depth int) error {
	// We want to avoid putting more in the visited map than we need to.
	// For any possible reference cycle that might be encountered,
	// hard(t) needs to return true for at least one of the types in the cycle.
	hard := func(k reflect.Kind) bool {
		switch k {
		case reflect.Map, reflect.Slice, reflect.Ptr, reflect.Interface:
			return true
		}
		return false
	}

	if dst.CanAddr() && src.CanAddr() && hard(dst.Kind()) {
		addr1 := unsafe.Pointer(dst.UnsafeAddr())
		addr2 := unsafe.Pointer(src.UnsafeAddr())
		if uintptr(addr1) > uintptr(addr2) {
			// Canonicalize order to reduce number of entries in visited.
			// Assumes non-moving garbage collector.
			addr1, addr2 = addr2, addr1
		}

		typ := dst.Type()
		v := visit{addr1, addr2, typ}
		// Short circuit if references have already been seen.
		if visited[v] {
			return nil
		}
		// Remember that we've visited this node
		visited[v] = true
	}

	switch dst.Kind() {
	case reflect.Array:
		for i := 0; i < dst.Len(); i++ {
			if err := deepSymmetricDifference(dst.Index(i), src.Index(i), visited, depth+1); err != nil {
				return err
			}
		}
	case reflect.Slice:
		if dst.IsNil() || src.IsNil() {
			break
		}

		if dst.Len() != src.Len() {
			break
		}

		if dst.Pointer() == src.Pointer() && dst.Elem().CanSet() {
			dst.Elem().Set(reflect.Zero(dst.Elem().Type()))
		}

		deepEqual := true
		for i := 0; i < dst.Len(); i++ {
			if !reflect.DeepEqual(dst.Index(i), src.Index(i)) {
				deepEqual = false
			}
		}
		if deepEqual && dst.CanSet() {
			dst.Set(reflect.Zero(dst.Type()))
		}
	case reflect.Interface:
		if dst.IsNil() || src.IsNil() {
			break
		}
		return deepSymmetricDifference(dst.Elem(), src.Elem(), visited, depth+1)
	case reflect.Ptr:
		if dst.Pointer() == src.Pointer() && dst.Elem().CanSet() {
			dst.Elem().Set(reflect.Zero(dst.Elem().Type()))
		}
		return deepSymmetricDifference(dst.Elem(), src.Elem(), visited, depth+1)
	case reflect.Struct:
		if hasExportedField(dst) {
			for i := 0; i < dst.NumField(); i++ {
				if err := deepSymmetricDifference(dst.Field(i), src.Field(i), visited, depth+1); err != nil {
					return err
				}
			}
		} else {
			if dst.CanSet() && dst.Interface() == src.Interface() {
				dst.Set(reflect.Zero(dst.Type()))
			}
		}
	case reflect.Map:
		for _, k := range dst.MapKeys() {
			val1 := dst.MapIndex(k)
			val2 := src.MapIndex(k)
			if val1.IsValid() && val2.IsValid() {
				if err := deepSymmetricDifference(val1, val2, visited, depth+1); err != nil {
					return err
				}
			}
		}
	default:
		if dst.CanSet() && dst.Interface() == src.Interface() {
			dst.Set(reflect.Zero(dst.Type()))
		}
	}
	return nil
}

func hasExportedField(dst reflect.Value) (exported bool) {
	for i, n := 0, dst.NumField(); i < n; i++ {
		field := dst.Type().Field(i)
		if field.Anonymous && dst.Field(i).Kind() == reflect.Struct {
			exported = exported || hasExportedField(dst.Field(i))
		} else {
			exported = exported || len(field.PkgPath) == 0
		}
	}
	return
}

func resolveValues(dst, src interface{}) (vDst, vSrc reflect.Value, err error) {
	if dst == nil || src == nil {
		err = ErrNilArguments
		return
	}
	vDst = reflect.ValueOf(dst).Elem()
	vSrc = reflect.ValueOf(src)
	// We check if vSrc is a pointer to dereference it.
	if vSrc.Kind() == reflect.Ptr {
		vSrc = vSrc.Elem()
	}
	return
}
