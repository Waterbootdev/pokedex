package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.first.com", []byte("first.testdata"))

	_, ok := cache.Get("https://example.first.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(baseTime)

	cache.Add("https://example.second.com", []byte("second.testdata"))

	time.Sleep(baseTime)

	cache.Add("https://example.third.com", []byte("third.testdata"))

	_, ok = cache.Get("https://example.first.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
	_, ok = cache.Get("https://example.second.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	_, ok = cache.Get("https://example.third.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
}
