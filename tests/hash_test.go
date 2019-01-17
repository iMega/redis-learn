package tests

import (
	"reflect"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestHash(t *testing.T) {
	var (
		key      = uuid.New().String()
		expected = map[string]interface{}{
			"0": "4bac27393bdd9777ce02453256c5577cd02275510b2227f473d03f533924f877",
			"1": "4bac27393bdd9777ce02453256c5577cd02275510b2227f473d03f533924f877",
			"3": "sha256",
		}
	)

	t.Run("Insert hash", func(t *testing.T) {
		_, err := GetRedis().HMSet(key, expected).Result()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Getting hash", func(t *testing.T) {
		actual, err := GetRedis().HGetAll(key).Result()
		if err != nil {
			t.Fatal(err)
		}

		if reflect.DeepEqual(expected, actual) {
			t.Fatal("values not equal")
		}
	})

	t.Run("Inspected the object by key", func(t *testing.T) {
		obj, err := GetRedis().DebugObject(key).Result()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(obj)
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
