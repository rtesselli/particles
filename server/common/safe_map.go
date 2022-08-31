package common

import "sync"

type SafeMap[T comparable, Y any] struct {
	mutex  sync.RWMutex
	values map[T]Y
}

func NewSafeMap[T comparable, Y any]() SafeMap[T, Y] {
	return SafeMap[T, Y]{values: make(map[T]Y)}
}

func (m *SafeMap[T, Y]) AddValue(key T, value Y) {
	m.mutex.Lock()
	m.values[key] = value
	m.mutex.Unlock()
}

func (m *SafeMap[T, Y]) RemoveValue(key T) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.values, key)
}

func (m *SafeMap[T, Y]) GetValue(key T) Y {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.values[key]
}

func (m *SafeMap[T, Y]) GetMap() map[T]Y {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	newMap := make(map[T]Y)
	for k, v := range m.values {
		newMap[k] = v
	}
	return newMap
}

func (m *SafeMap[T, Y]) Reset() {
	m.mutex.Lock()
	m.values = make(map[T]Y)
	m.mutex.Unlock()
}
