package buffer

import (
	"reflect"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	rb := NewRingBuffer(3)

	rb.Push(1)
	rb.Push(2)
	rb.Push(3)

	got := rb.Pop()
	if !reflect.DeepEqual(got, 1) {
		t.Errorf("want: %v, got: %v", 1, got)
	}

	got = rb.Pop()
	if !reflect.DeepEqual(got, 2) {
		t.Errorf("want: %v, got: %v", 2, got)
	}

	got = rb.Pop()
	if !reflect.DeepEqual(got, 3) {
		t.Errorf("want: %v, got: %v", 3, got)
	}

	got = rb.Pop()
	if !reflect.DeepEqual(got, -1) {
		t.Errorf("want: %v, got: %v", -1, got)
	}

	rb.Push(1)
	rb.Push(2)
	rb.Push(3)

	got = rb.Pop()
	if !reflect.DeepEqual(got, 1) {
		t.Errorf("want: %v, got: %v", 1, got)
	}

	got = rb.Pop()
	if !reflect.DeepEqual(got, 2) {
		t.Errorf("want: %v, got: %v", 2, got)
	}

	got = rb.Pop()
	if !reflect.DeepEqual(got, 3) {
		t.Errorf("want: %v, got: %v", 3, got)
	}

	got = rb.Pop()
	if !reflect.DeepEqual(got, -1) {
		t.Errorf("want: %v, got: %v", -1, got)
	}

	rb.Push(1)
	rb.Push(2)
	rb.Push(3)
	rb.Push(4)

	got = rb.Pop()
	if !reflect.DeepEqual(got, 2) {
		t.Errorf("want: %v, got: %v", 2, got)
	}

	got = rb.Pop()
	if !reflect.DeepEqual(got, 3) {
		t.Errorf("want: %v, got: %v", 3, got)
	}

	got = rb.Pop()
	if !reflect.DeepEqual(got, 4) {
		t.Errorf("want: %v, got: %v", 4, got)
	}

	got = rb.Pop()
	if !reflect.DeepEqual(got, -1) {
		t.Errorf("want: %v, got: %v", -1, got)
	}

	rb.Push(5)
	rb.Push(6)
	rb.Push(7)
	rb.Push(9)

	got = rb.Pop()
	if !reflect.DeepEqual(got, 6) {
		t.Errorf("want: %v, got: %v", 6, got)
	}

	got = rb.Pop()
	if !reflect.DeepEqual(got, 7) {
		t.Errorf("want: %v, got: %v", 7, got)
	}

	got = rb.Pop()
	if !reflect.DeepEqual(got, 9) {
		t.Errorf("want: %v, got: %v", 9, got)
	}

	got = rb.Pop()
	if !reflect.DeepEqual(got, -1) {
		t.Errorf("want: %v, got: %v", -1, got)
	}

	rb.Push(100)

	got = rb.Pop()
	if !reflect.DeepEqual(got, 100) {
		t.Errorf("want: %v, got: %v", 100, got)
	}

}
