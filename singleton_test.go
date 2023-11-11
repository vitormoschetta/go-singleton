package teste

import "testing"

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
