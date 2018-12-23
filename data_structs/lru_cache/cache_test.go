package lru_cache

import "testing"
import "reflect"

func check(t *testing.T, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		t.Errorf("Expected: %v, Actual: %v", exp, act)
	}
}

func TestLRUGetAndSet(t *testing.T) {
	cache := New(5)
	cache.Set(1, "a")
	cache.Set(2, "b")
	cache.Set(3, "c")
	cache.Set(4, "d")
	cache.Set(5, "e")
	check(t, "a", cache.Get(1))
	check(t, "b", cache.Get(2))
	check(t, "c", cache.Get(3))
	check(t, "d", cache.Get(4))
	check(t, "e", cache.Get(5))
	cache.Set(6, "f")
	check(t, "f", cache.Get(6))
	check(t, nil, cache.Get(1))
	cache.Set(7, "g")
	check(t, "g", cache.Get(7))
	check(t, nil, cache.Get(1))
	check(t, nil, cache.Get(2))
}

func TestLRUGetSetWithReplace(t *testing.T) {
	cache := New(3)
	cache.Set(1, "a")
	cache.Set(2, "b")
	cache.Set(3, "c")
	check(t, "a", cache.Get(1))
	check(t, "b", cache.Get(2))
	check(t, "c", cache.Get(3))
	cache.Set(1, "aa")
	cache.Set(4, "d")
	check(t, "aa", cache.Get(1))
	check(t, nil, cache.Get(2))
	check(t, "c", cache.Get(3))
	check(t, "d", cache.Get(4))
	cache.Set(1, "aaa")
	cache.Set(5, "e")
	check(t, "aaa", cache.Get(1))
	check(t, nil, cache.Get(2))
	check(t, nil, cache.Get(3))
	check(t, "d", cache.Get(4))
	check(t, "e", cache.Get(5))
}
