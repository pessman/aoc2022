package main

import "testing"

func Test_getParent(t *testing.T) {
	type test struct {
		name   string
		node   int
		parent int
	}

	cases := []test{
		{
			name:   "node 1 - parent node 0",
			node:   1,
			parent: 0,
		},
		{
			name:   "node 2 - parent node 0",
			node:   2,
			parent: 0,
		},
		{
			name:   "node 5 - parent 2",
			node:   5,
			parent: 2,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := getParent(c.node); got != c.parent {
				t.Errorf("got = %d, want = %d", got, c.parent)
			}
		})
	}
}

func Test_leftChild(t *testing.T) {
	type test struct {
		name   string
		parent int
		child  int
	}

	cases := []test{
		{
			name:   "root left child",
			parent: 0,
			child:  1,
		},
		{
			name:   "node 1 left child",
			parent: 1,
			child:  3,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := leftChild(c.parent); got != c.child {
				t.Errorf("got = %d, want = %d", got, c.child)
			}
		})
	}
}

func Test_rightChild(t *testing.T) {
	type test struct {
		name   string
		parent int
		child  int
	}

	cases := []test{
		{
			name:   "root right child",
			parent: 0,
			child:  2,
		},
		{
			name:   "node 2 right child",
			parent: 1,
			child:  4,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := rightChild(c.parent); got != c.child {
				t.Errorf("got = %d, want = %d", got, c.child)
			}
		})
	}
}
