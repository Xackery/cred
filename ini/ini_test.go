package ini

import (
	"os"
	"testing"
	"time"
)

func TestIni(t *testing.T) {
	path, err := os.Getwd()

	p, err := Load(path, "test")
	if err != nil {
		t.Fatalf("Failed to load test: %s", err.Error())
	}
	if p.String("test_string") != "String" {
		t.Fatalf("test_string key returned %s, not %s", p.String("test_string"), "String")
	}
	if p.Duration("test_duration") != 10*time.Second {
		t.Fatalf("test_duration key returned %s, not 10s", p.Duration("test_duration"))
	}

	if p.Int("test_int") != 123 {
		t.Fatalf("test_int key returned %s, not 123", p.Int("test_int"))
	}

	if p.Int64("test_int64") != 123123 {
		t.Fatalf("test_int64 key returned %s, not 123123", p.Int64("test_int64"))
	}

}
