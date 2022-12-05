package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	st := stack[string]{}

	if !st.isEmpty() {
		t.Log("The stack is new, it must be empty!")
		t.Fail()
	}

	st.push("hi")

	if st.peek() != "hi" {
		t.Log("Pushing to the stack must not be working")
		t.Fail()
	}

	if x := st.pop(); x != "hi" {
		t.Log("Unexpected value popped from stack. Expeded `hi`, got ", x)
		t.Fail()
	}

	if !st.isEmpty() {
		t.Log("Stack was just emptied, isEmpty must be true!")
		t.Fail()
	}
}
