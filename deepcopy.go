package deepcopy

import (
	"fmt"
	. "reflect"
)

type copier = func(old any, oldToNewPtr map[uintptr]any) (any, error)

var kindToCopier map[Kind]copier

func _notSupported(old any, oldToNewPtr map[uintptr]any) (any, error) {
	kind := ValueOf(old).Kind()
	return nil, fmt.Errorf("value %v of kind %v is not supported", old, kind)
}

func _primitive(old any, oldToNewPtr map[uintptr]any) (any, error) {
	return old, nil
}

func _array(old any, oldToNewPtr map[uintptr]any) (any, error) {
	v := ValueOf(old)
	t := TypeOf(old)
	size := v.Len()
	new := New(ArrayOf(size, t.Elem())).Elem()

	for i := range size {
		newi, err := _deepCopy(v.Index(i).Interface(), oldToNewPtr)
		if err != nil {
			return nil, fmt.Errorf("failed to deep copy array element at index %v: %v", i, err)
		}
		newiv := ValueOf(newi)
		if newiv.IsValid() {
			new.Index(i).Set(newiv)
		}
	}

	return new.Interface(), nil
}

func _slice(old any, oldToNewPtr map[uintptr]any) (any, error) {
	v := ValueOf(old)
	t := TypeOf(old)
	size := v.Len()
	new := MakeSlice(t, size, size)
	// new := New(SliceOf(t.Elem())).Elem()
	// new.SetLen(size)
	// new.SetCap(size)

	for i := range size {
		newi, err := _deepCopy(v.Index(i).Interface(), oldToNewPtr)
		if err != nil {
			return nil, fmt.Errorf("failed to deep copy slice element at index %v: %v", i, err)
		}
		newiv := ValueOf(newi)
		if newiv.IsValid() {
			new.Index(i).Set(newiv)
		}
	}

	return new.Interface(), nil
}

func _map(old any, oldToNewPtr map[uintptr]any) (any, error) {
	v := ValueOf(old)
	t := TypeOf(old)
	new := MakeMap(t)

	iter := v.MapRange()

	for iter.Next() {
		newk, err := _deepCopy(iter.Key().Interface(), oldToNewPtr)
		if err != nil {
			return nil, fmt.Errorf("failed to deep copy map key: %v", err)
		}
		newv, err := _deepCopy(iter.Value().Interface(), oldToNewPtr)
		if err != nil {
			return nil, fmt.Errorf("failed to deep copy map value: %v", err)
		}
		newkv := ValueOf(newk)
		newvv := ValueOf(newv)
		if newkv.IsValid() && newvv.IsValid() {
			new.SetMapIndex(newkv, newvv)
		}
	}

	return new.Interface(), nil
}

func _pointer(old any, oldToNewPtr map[uintptr]any) (any, error) {
	v := ValueOf(old)
	t := TypeOf(old)

	if v.IsNil() {
		return Zero(t).Interface(), nil
	}

	if new, ok := oldToNewPtr[v.Pointer()]; ok {
		return new, nil
	}

	new := New(t.Elem())
	oldToNewPtr[v.Pointer()] = new.Interface()

	newv, err := _deepCopy(v.Elem().Interface(), oldToNewPtr)
	if err != nil {
		return nil, fmt.Errorf("failed to deep copy pointer element %v: %v", v.Elem().Interface(), err)
	}
	newvv := ValueOf(newv)
	if newvv.IsValid() {
		new.Elem().Set(newvv)
	}
	return new.Interface(), nil
}

func _struct(old any, oldToNewPtr map[uintptr]any) (any, error) {
	v := ValueOf(old)
	t := TypeOf(old)
	new := New(t).Elem()
	for i := range t.NumField() {
		if t.Field(i).PkgPath != "" {
			continue
		}
		newi, err := _deepCopy(v.Field(i).Interface(), oldToNewPtr)
		if err != nil {
			return nil, fmt.Errorf("failed to deep copy struct field %v with value %v: %v", t.Field(i).Name, v.Field(i).Interface(), err)
		}
		newiv := ValueOf(newi)
		if newiv.IsValid() {
			new.Field(i).Set(newiv)
		}
	}
	return new.Interface(), nil
}

func init() {
	kindToCopier = map[Kind]copier{
		Invalid:       _notSupported,
		Bool:          _primitive,
		Int:           _primitive,
		Int8:          _primitive,
		Int16:         _primitive,
		Int32:         _primitive,
		Int64:         _primitive,
		Uint:          _primitive,
		Uint8:         _primitive,
		Uint16:        _primitive,
		Uint32:        _primitive,
		Uint64:        _primitive,
		Uintptr:       _primitive,
		Float32:       _primitive,
		Float64:       _primitive,
		Complex64:     _primitive,
		Complex128:    _primitive,
		Array:         _array,
		Chan:          _notSupported,
		Func:          _notSupported,
		Interface:     _notSupported,
		Map:           _map,
		Pointer:       _pointer,
		Slice:         _slice,
		String:        _primitive,
		Struct:        _struct,
		UnsafePointer: _notSupported,
	}
}

func _deepCopy(old any, oldToNewPtr map[uintptr]any) (any, error) {
	kind := ValueOf(old).Kind()
	copier := kindToCopier[kind]
	return copier(old, oldToNewPtr)
}

func DeepCopy(src any) (any, error) {
	oldToNewPtr := make(map[uintptr]any)
	return _deepCopy(src, oldToNewPtr)
}
