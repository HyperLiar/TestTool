package main

import (
	"sort"
	"testing"
)

func TestSortMap(t *testing.T) {
	m := map[string]int {
		"something": 10,
		"yo": 20,
		"blaj" : 20,
	}

	type kv struct {
		Key string
		Value int
	}

	var ss []kv

	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Key > ss[j].Key
	})

	for _, kv := range ss {
		t.Logf("key_desc||%s, %d\n", kv.Key, kv.Value)
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Key  < ss[j].Key
	})


	for _, kv := range ss {
		t.Logf("key_desc||%s, %d\n", kv.Key, kv.Value)
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Key < ss[j].Key
	})

	for _, kv := range ss {
		t.Logf("key_asc||%s, %d\n", kv.Key, kv.Value)
	}
}