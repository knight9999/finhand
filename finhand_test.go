package finhand

import (
	"fmt"
	"testing"
)

func TestHandlers(t *testing.T) {
	AddHandlerWithNameAndPriority("handle1", 0, func() {
		fmt.Println("OK")
	})
	if _, ok := GetHandler("handle1"); !ok {
		t.Errorf("handle1 is not found")
	}
	if _, ok := GetHandler("handle2"); ok {
		t.Errorf("handle2 is found")
	}
	RemoveHandler("handle1")
	if _, ok := GetHandler("handle1"); ok {
		t.Errorf("handle1 is found")
	}
	result := 0
	AddHandlerWithNameAndPriority("handle3", 0, func() {
		result = 1
	})
	RunHandlers()
	if result != 1 {
		t.Errorf("does not exec handle3")
	}
}

func TestHandlersWithPriority(t *testing.T) {
	result := 0
	AddHandlerWithPriority(10, func() {
		result = result + 1
	})
	AddHandlerWithPriority(5, func() {
		result = result * 2
	})
	RunHandlers()
	if result != 2 {
		t.Errorf("fail in Priority Test")
	}
	result = 0
	AddHandlerWithPriority(5, func() {
		result = result + 1
	})
	AddHandlerWithPriority(10, func() {
		result = result * 2
	})
	RunHandlers()
	if result != 1 {
		t.Errorf("fail inpriority Test")
	}

}

func TestAnonymousHandlers(t *testing.T) {
	result := 0
	name1 := AddHandler(func() {
		result = result + 1
	})
	name2 := AddHandler(func() {
		result = result * 2
	})
	name3 := AddHandler(func() {
		result = result + 10
	})
	RemoveHandler(name3)

	hand1, ok1 := GetHandler(name1)
	if !ok1 {
		t.Errorf("fail in register hand1")
	}
	if name1 != hand1.name {
		t.Errorf("fail in register hand1")
	}
	hand2, ok2 := GetHandler(name2)
	if !ok2 {
		t.Errorf("fail in register hand1")
	}
	if name2 != hand2.name {
		t.Errorf("fail in register hand1")
	}

	RunHandlers()
	if result != 2 {
		t.Errorf("fail in TestAnonymousHandlers")
	}
	result = 0
	RunHandlers()
	if result != 0 {
		t.Errorf("fail in Clearing Handlers")
	}
}
