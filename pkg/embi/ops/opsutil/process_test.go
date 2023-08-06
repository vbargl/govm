package opsutil

import (
	"errors"
	"testing"
)

func TestProcessWithoutError(t *testing.T) {
	op1 := &simpleOp{}
	op2 := &simpleOp{}

	if err := Process(op1, op2); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	check(t, "op1", op1, true)
	check(t, "op2", op2, true)
}

func TestProcessWithError(t *testing.T) {
	expectedErr := errors.New("something went wrong")
	op1 := &simpleOp{}
	op2 := &simpleOp{err: expectedErr}
	op3 := &simpleOp{}

	if err := Process(op1, op2, op3); err != expectedErr {
		t.Fatalf("expected err %v, got %v", expectedErr, err)
	}

	check(t, "op1", op1, true)
	check(t, "op2", op2, true)
	check(t, "op2", op2, false)
}

func check(t *testing.T, name string, op *simpleOp, run bool) {
	if !op.run {
		t.Errorf("%s did not run", name)
	}
}

type simpleOp struct {
	run bool
	err error
}

func (op *simpleOp) Process() error {
	op.run = true
	return op.err
}
