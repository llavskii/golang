package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(10)
		c.Set("123", 123)
		val, ok := c.Get("123")
		require.True(t, ok)
		require.Equal(t, 123, val)
		c.Clear()
		require.True(t, c.IsEmpty())
	})
	t.Run("cache overflow", func(t *testing.T) {
		c := NewCache(3)
		c.Set("1", 1)
		c.Set("2", 2)
		c.Set("3", 3)
		c.Set("4", 4)
		_, ok := c.Get("1")
		require.False(t, ok)
	})
	t.Run("element obsolescence", func(t *testing.T) {
		c := NewCache(3)
		c.Set("1", 1)
		c.Set("2", 2)
		c.Set("3", 3)
		c.Get("1")
		_, ok := c.Get("1") // after that, value for key '1' is the oldest cache element
		require.True(t, ok)
		c.Set("4", 4)
		_, ok = c.Get("2")
		require.False(t, ok)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
