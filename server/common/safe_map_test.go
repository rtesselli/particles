package common

import "testing"

func TestNewSafeMap(t *testing.T) {
	safe_map := NewSafeMap[int, int]()
	if len(safe_map.values) != 0 {
		t.Errorf("Wrong size")
	}
}

func TestAddValue(t *testing.T) {
	safe_map := SafeMap[int, int]{values: make(map[int]int)}
	safe_map.AddValue(1, 10)
	safe_map.AddValue(1, 10)
	if safe_map.values[1] != 10 {
		t.Errorf("Value error")
	}
}

func TestRemoveValue(t *testing.T) {
	safe_map := SafeMap[int, int]{values: make(map[int]int)}
	safe_map.AddValue(1, 10)
	safe_map.RemoveValue(1)
	if len(safe_map.values) != 0 {
		t.Errorf("Wrong size")
	}
}

func TestGetValue(t *testing.T) {
	safe_map := SafeMap[int, int]{values: make(map[int]int)}
	safe_map.AddValue(1, 10)
	if safe_map.GetValue(1) != 10 {
		t.Errorf("Key error")
	}
}

func TestReset(t *testing.T) {
	safe_map := NewSafeMap[int, int]()
	safe_map.AddValue(1, 1)
	safe_map.AddValue(2, 2)
	safe_map.Reset()
	if len(safe_map.values) != 0 {
		t.Errorf("Wrong size")
	}
}
