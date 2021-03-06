package tests

import (
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestString(t *testing.T) {
	var (
		key      = uuid.New().String()
		expected = "bar"
	)

	t.Run("Insert string", func(t *testing.T) {
		_, err := GetRedis().Append(key, expected).Result()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Getting string", func(t *testing.T) {
		actual, err := GetRedis().Get(key).Result()
		if err != nil {
			t.Fatal(err)
		}

		if expected != actual {
			t.Fatal("values not equal")
		}
	})

	t.Run("Inspected the object by key", func(t *testing.T) {
		obj, err := GetRedis().DebugObject(key).Result()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(obj)

		val, err := GetRedis().Do("debug", "sdslen", key).Result()
		if err != nil {
			t.Fatal(err)
		}

		sdslen, ok := val.(string)
		if !ok {
			t.Fatal("failed to convert value to string")
		}

		t.Log(sdslen)
	})

	t.Run("Memory usage", func(t *testing.T) {
		b, err := GetRedis().MemoryUsage(key).Result()
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("Memory usage: %d", b)
	})

	t.Run("Memory info", func(t *testing.T) {
		s, err := GetRedis().Info("memory").Result()
		if err != nil {
			t.Fatal(err)
		}

		mi := strings.Split(s, "\n")
		b := strings.Split(mi[1], ":")

		t.Logf("Memory info: %s", b[1])
	})
}
