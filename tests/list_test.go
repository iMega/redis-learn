package tests

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

var email = "user@example.com"

func EmailToHash(value string) string {
	h := sha256.New()

	h.Write([]byte(value))

	ret := h.Sum(nil)

	fmt.Println(base64.URLEncoding.EncodeToString(h.Sum(nil)))

	return string(ret)
}

func EmailToHash2(value string) string {
	ret := sha256.Sum256([]byte(value))

	fmt.Printf("%x\n", ret)

	return fmt.Sprintf("%x", ret)
}

func EmailToHash3(value string) string {
	return "117" +
		"115" +
		"101" +
		"114" +
		"50" +
		"64" +
		"101" +
		"120" +
		"97" +
		"109" +
		"112" +
		"108" +
		"101" +
		"46" +
		"99" +
		"111" +
		"109"
}

func TestList_key_email(t *testing.T) {
	var (
		key      = email
		expected = []string{
			"4bac27393bdd9777ce02453256c5577cd02275510b2227f473d03f533924f877",
			"4bac27393bdd9777ce02453256c5577cd02275510b2227f473d03f533924f877",
			"sha256",
		}
	)

	t.Run("Insert list", func(t *testing.T) {
		_, err := GetRedis().LPush(key, expected).Result()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Getting list", func(t *testing.T) {
		actual, err := GetRedis().LRange(key, 0, -1).Result()
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

func TestList_key_email_ints(t *testing.T) {
	var (
		key      = EmailToHash3(email)
		expected = []string{
			"4bac27393bdd9777ce02453256c5577cd02275510b2227f473d03f533924f877",
			"4bac27393bdd9777ce02453256c5577cd02275510b2227f473d03f533924f877",
			"sha256",
		}
	)

	t.Run("Insert list", func(t *testing.T) {
		_, err := GetRedis().LPush(key, expected).Result()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Getting list", func(t *testing.T) {
		actual, err := GetRedis().LRange(key, 0, -1).Result()
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

func TestList_key_email_hash_string(t *testing.T) {
	var (
		key      = EmailToHash(email)
		expected = []string{
			"4bac27393bdd9777ce02453256c5577cd02275510b2227f473d03f533924f877",
			"4bac27393bdd9777ce02453256c5577cd02275510b2227f473d03f533924f877",
			"sha256",
		}
	)

	t.Run("Insert list", func(t *testing.T) {
		_, err := GetRedis().LPush(key, expected).Result()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Getting list", func(t *testing.T) {
		actual, err := GetRedis().LRange(key, 0, -1).Result()
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

func TestList_key_email_hash_byte(t *testing.T) {
	var (
		key      = EmailToHash2(email)
		expected = []string{
			"4bac27393bdd9777ce02453256c5577cd02275510b2227f473d03f533924f877",
			"4bac27393bdd9777ce02453256c5577cd02275510b2227f473d03f533924f877",
			"sha256",
		}
	)

	t.Run("Insert list", func(t *testing.T) {
		_, err := GetRedis().LPush(key, expected).Result()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Getting list", func(t *testing.T) {
		actual, err := GetRedis().LRange(key, 0, -1).Result()
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
