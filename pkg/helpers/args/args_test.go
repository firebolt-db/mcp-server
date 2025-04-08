package args

import (
	"testing"
)

func TestString(t *testing.T) {
	args := map[string]any{
		"key1": "value1",
	}

	val, err := String(args, "key1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != "value1" {
		t.Errorf("expected 'value1', got '%s'", val)
	}

	_, err = String(args, "key2")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestMaybeString(t *testing.T) {
	args := map[string]any{
		"key1": "value1",
	}

	val, err := MaybeString(args, "key1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val == nil || *val != "value1" {
		t.Errorf("expected 'value1', got '%v'", val)
	}

	val, err = MaybeString(args, "key2")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != nil {
		t.Errorf("expected nil, got '%v'", val)
	}
}

func TestStrings(t *testing.T) {
	args := map[string]any{
		"key1": "value1",
		"key2": "value2",
	}

	vals, err := Strings(args, "key1", "key2")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(vals) != 2 || vals[0] != "value1" || vals[1] != "value2" {
		t.Errorf("expected ['value1', 'value2'], got '%v'", vals)
	}

	_, err = Strings(args, "key1", "key3")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestMaybeStrings(t *testing.T) {
	args := map[string]any{
		"key1": "value1",
		"key2": "value2",
	}

	vals, err := MaybeStrings(args, "key1", "key2")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(vals) != 2 || *vals[0] != "value1" || *vals[1] != "value2" {
		t.Errorf("expected ['value1', 'value2'], got '%v'", vals)
	}

	vals, err = MaybeStrings(args, "key1", "key3")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(vals) != 2 || *vals[0] != "value1" || vals[1] != nil {
		t.Errorf("expected ['value1', nil], got '%v'", vals)
	}
}
