package io

import (
	"errors"
	"io"
	"testing"
	"time"
)

func TestWriterTimeout_Write(t *testing.T) {
	timeout := time.Second
	w := NewTimeoutWriter(timeout)
	data := "Hello, World"

	start := time.Now()
	n, err := w.Write([]byte (data))
	end:= time.Now()
	// to do
	if err != nil {
		t.Fatal("the writer should be slow, but shouldn't fail")
	}
	if n != len(data) {
		t.Fatal("the writer should write the whole data")
	}
	if end.Sub(start) < timeout {
		t.Fatal("the writer should write in more than", timeout)
	}
}

func TestByte(t *testing.T)  {
	Byte := 
}

func TestSlowWriter_Write(t *testing.T) {
	speed := 2

	w := NewSlowWriter(speed)
	data := "Hello, World"

	for {
		n, err := w.Write([]byte(data))
		if err != nil {
			if err != io.EOF {
				break
			}
			t.Fatal("the writer should be slow, but shouldn't fail")
		}

		if n != speed {
			t.Fatal("the writer doesn't write at the specified speed")
		}
		data = data[:n]
	}
}
