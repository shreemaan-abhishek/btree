package main

import "testing"

func TestTreeDepth(t *testing.T) {
	v := 3
	tree := &Node{
		value: &v,
	}
	if TreeDepth(tree) != 1 {
		t.Errorf("depth not 1")
	}

	l := tree
	r := tree
	anotherTree := &Node{
		value: &v,
		left:  l,
		right: r,
	}
	twodep := TreeDepth(anotherTree)
	if twodep != 2 {
		t.Errorf("depth not 2 %d", twodep)
	}

	l2 := anotherTree
	r2 := anotherTree
	three := &Node{
		value: &v,
		left:  l2,
		right: r2,
	}
	threedep := TreeDepth(three)
	if threedep != 3 {
		t.Errorf("wrong depth found %d", twodep)
	}
}
