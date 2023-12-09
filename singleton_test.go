package main

import (
	"fmt"
	"testing"
	"time"
)

func TestInstance(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()

	if instance1 != instance2 {
		t.Error("Expected the same instance in both calls")
	}
}

func TestInstance2(t *testing.T) {
	instance1 := GetInstance2()
	instance2 := GetInstance2()

	if instance1 != instance2 {
		t.Error("Expected the same instance in both calls")
	}
}

func TestSelect(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "hello"
	}()

	go func() {
		time.Sleep(time.Second * 1)
		ch2 <- "world"
	}()

	select {
	case <-ch1:
		fmt.Println("ch1 was ready")
	case <-ch2:
		fmt.Println("ch2 was ready")
	}
}

func TestSelect2(t *testing.T) {
	a := 3

	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	default:
		fmt.Println("a is not 1 or 2")
	}
}
