package helper

import (
	"fmt"
	"reflect"
)

func ExpectType[T any](r any) T {
	expectedType := reflect.TypeOf((*T)(nil)).Elem()
	receivedType := reflect.TypeOf(r)

	fmt.Printf("Expected type: %v, Received type: %v\n", expectedType, receivedType)

	if expectedType == receivedType {
		return r.(T)
	}
	panic(fmt.Sprintf("Expected %v but received %v instead", expectedType, receivedType))
}
