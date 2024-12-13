package utils

import (
	"os"
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	// Test integer queue
	intQueue := Queue[int]{}
	if !intQueue.IsEmpty() {
		t.Errorf("Expected queue to be empty")
	}

	intQueue.Enqueue(1)
	intQueue.Enqueue(2)
	intQueue.Enqueue(3)

	if intQueue.IsEmpty() {
		t.Errorf("Expected queue to not be empty")
	}

	val, ok := intQueue.Dequeue()
	if !ok || val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}

	val, ok = intQueue.Dequeue()
	if !ok || val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}

	val, ok = intQueue.Dequeue()
	if !ok || val != 3 {
		t.Errorf("Expected 3, got %v", val)
	}

	if !intQueue.IsEmpty() {
		t.Errorf("Expected queue to be empty after dequeuing all elements")
	}

	// Test string queue
	stringQueue := Queue[string]{}
	stringQueue.Enqueue("hello")
	stringQueue.Enqueue("world")

	valStr, ok := stringQueue.Dequeue()
	if !ok || valStr != "hello" {
		t.Errorf("Expected 'hello', got %v", valStr)
	}

	valStr, ok = stringQueue.Dequeue()
	if !ok || valStr != "world" {
		t.Errorf("Expected 'world', got %v", valStr)
	}

	if !stringQueue.IsEmpty() {
		t.Errorf("Expected string queue to be empty after dequeuing all elements")
	}
}

func TestLoadFile(t *testing.T) {
	// Create a temporary test file
	tempFileName := "test_input.txt"
	testContent := "line1\r\nline2\r\nline3"
	err := os.WriteFile(tempFileName, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tempFileName) // Cleanup after test

	lines := LoadFile(tempFileName)
	expected := []string{"line1", "line2", "line3"}

	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("Expected %v, got %v", expected, lines)
	}

	// Test file not found
	nonExistentFile := "non_existent.txt"
	lines = LoadFile(nonExistentFile)
	if lines != nil {
		t.Errorf("Expected nil for non-existent file, got %v", lines)
	}
}
