package hashtable_test

import (
	"fmt"
	"testing"

	"github.com/HotPotatoC/hashtable"
)

func populate(n int) hashtable.HashTable {
	ht := hashtable.New()
	for i := 0; i < n; i++ {
		ht.Set(fmt.Sprintf("k%d", i+1), fmt.Sprintf("v%d", i+1))
	}
	return ht
}

func TestSet(t *testing.T) {
	ht := populate(4)
	if ht.Len() != 4 {
		t.Errorf("Failed TestSet -> Expected Size: %d | Got: %d", 4, ht.Len())
	}

	ht.Set("my-key", "value")
	if ht.Len() != 5 {
		t.Errorf("Failed TestSet -> Expected Size: %d | Got: %d", 5, ht.Len())
	}
}

func TestRemove(t *testing.T) {
	ht := populate(4)

	ht.Remove("k1")

	if ht.Len() != 3 {
		t.Errorf("Failed TestRemove -> Expected Size: %d | Got: %d", 3, ht.Len())
	}
}

func TestGet(t *testing.T) {
	ht := populate(5)

	value, ok := ht.Get("k2")
	if !ok {
		t.Errorf("Failed TestGet -> Expected value: %v | Got: %v", true, ok)
	}
	expected := "v2"
	if value != expected {
		t.Errorf("Failed TestGet -> Expected value: %s | Got: %s", expected, value)
	}
}

func TestIter(t *testing.T) {
	ht := populate(5)

	kv := make([]*hashtable.Entry, 0)
	for entry := range ht.Iter() {
		kv = append(kv, entry)
	}

	if len(kv) != 5 {
		t.Errorf("Failed TestIter -> Expected size: %d | Got: %d", 5, len(kv))
	}
}

func TestPopulate_100(t *testing.T) {
	ht := populate(100)
	if ht.Len() != 100 {
		t.Errorf("Failed TestPopulate100 -> Expected Size: %d | Got: %d", 100, ht.Len())
	}

	for i := 0; i < 100; i++ {
		if !ht.Exist(fmt.Sprintf("k%d", i+1)) {
			t.Errorf("Failed TestPopulate100 -> Expected: %v | Got: %v", true, ht.Exist(fmt.Sprintf("k%d", i)))
		}
	}
}

func TestPopulate_1000(t *testing.T) {
	ht := populate(1000)
	if ht.Len() != 1000 {
		t.Errorf("Failed TestPopulate100 -> Expected Size: %d | Got: %d", 1000, ht.Len())
	}

	for i := 0; i < 1000; i++ {
		if !ht.Exist(fmt.Sprintf("k%d", i+1)) {
			t.Errorf("Failed TestPopulate100 -> Expected: %v | Got: %v", true, ht.Exist(fmt.Sprintf("k%d", i)))
		}
	}
}

func TestPopulate_10000(t *testing.T) {
	ht := populate(10000)
	if ht.Len() != 10000 {
		t.Errorf("Failed TestPopulate10000 -> Expected Size: %d | Got: %d", 10000, ht.Len())
	}

	for i := 0; i < 10000; i++ {
		if !ht.Exist(fmt.Sprintf("k%d", i+1)) {
			t.Errorf("Failed TestPopulate100 -> Expected: %v | Got: %v", true, ht.Exist(fmt.Sprintf("k%d", i)))
		}
	}
}

func TestPopulate_100000(t *testing.T) {
	ht := populate(100000)
	if ht.Len() != 100000 {
		t.Errorf("Failed TestPopulate100000 -> Expected Size: %d | Got: %d", 100000, ht.Len())
	}

	for i := 0; i < 100000; i++ {
		if !ht.Exist(fmt.Sprintf("k%d", i+1)) {
			t.Errorf("Failed TestPopulate100 -> Expected: %v | Got: %v", true, ht.Exist(fmt.Sprintf("k%d", i)))
		}
	}
}
