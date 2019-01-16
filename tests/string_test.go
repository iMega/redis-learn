package tests

import (
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
}
