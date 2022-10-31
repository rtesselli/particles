package common

import "testing"

func TestNewSafeMap(t *testing.T) {
	safeMap := NewSafeMap[int, int]()
	if len(safeMap.values) != 0 {
		t.Errorf("Wrong size")
	}
}

func TestAddValue(t *testing.T) {
	safeMap := SafeMap[int, int]{values: make(map[int]int)}
	safeMap.AddValue(1, 10)
	safeMap.AddValue(1, 10)
	if safeMap.values[1] != 10 {
		t.Errorf("Value error")
	}
}

func TestRemoveValue(t *testing.T) {
	safeMap := SafeMap[int, int]{values: make(map[int]int)}
	safeMap.AddValue(1, 10)
	safeMap.RemoveValue(1)
	if len(safeMap.values) != 0 {
		t.Errorf("Wrong size")
	}
}

func TestGetValue(t *testing.T) {
	safeMap := SafeMap[int, int]{values: make(map[int]int)}
	safeMap.AddValue(1, 10)
	if safeMap.GetValue(1) != 10 {
		t.Errorf("Key error")
	}
}

func TestReset(t *testing.T) {
	safeMap := NewSafeMap[int, int]()
	safeMap.AddValue(1, 1)
	safeMap.AddValue(2, 2)
	safeMap.Reset()
	if len(safeMap.values) != 0 {
		t.Errorf("Wrong size")
	}
}

func TestSize(t *testing.T) {
	safeMap := NewSafeMap[int, int]()
	safeMap.AddValue(1, 1)
	safeMap.AddValue(2, 2)
	if safeMap.Size() != 2 {
		t.Errorf("Wrong size")
	}
}
