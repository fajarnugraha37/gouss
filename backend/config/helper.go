package config

import (
	"fmt"
	"reflect"
	"strings"
)

func get(path string) (any, error) {
	parts := strings.Split(path, ".")
	var value reflect.Value = reflect.
        ValueOf(&AppProperty).
        Elem()

	for _, part := range parts {
		if value.Kind() != reflect.Struct {
			return nil, fmt.Errorf("cannot access field %s in non-struct value", part)
        }

        value = value.FieldByName(part)
        if !value.IsValid() {
            return nil, fmt.Errorf("invalid field: %s", part)
        }
	}

	return value.Interface(), nil
}

func GetOrNil(path string) any {
	value, err := get(path)
	if err != nil {
		return nil
	}

	return value
}

func GetOrDefault(path string, def any) any {
	value, err := get(path)
	if err != nil {
		return def
	}

	return value
}

func Set(path string, value any) error {
    parts := strings.Split(path, ".")
    v := reflect.
        ValueOf(&AppProperty).
        Elem()

    for i, part := range parts {
        if v.Kind() != reflect.Struct {
            return fmt.Errorf("invalid path: %s is not a struct", strings.Join(parts[:i], "."))
        }
        
        if i == len(parts)-1 {
            field := v.FieldByName(part)
            if !field.IsValid() {
                return fmt.Errorf("field not found: %s", part)
            }
            if !field.CanSet() {
                return fmt.Errorf("cannot set field: %s", part)
            }
            
            newValue := reflect.ValueOf(value)
            if !newValue.Type().AssignableTo(field.Type()) {
                return fmt.Errorf("cannot assign %v to %s (type %v)", value, part, field.Type())
            }
            field.Set(newValue)
         
            return nil
        }
        
        v = v.FieldByName(part)
        if !v.IsValid() {
            return fmt.Errorf("invalid field: %s", part)
        }
    }

    return fmt.Errorf("path does not lead to a field")
}
