package main

import (
	"reflect"
	"testing"
)

func TestSideViews(t *testing.T) {
	// Example Tree:
	//      1
	//     / \
	//    2   3
	//     \   \
	//      5   4
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 4},
		},
	}

	tests := []struct {
		name     string
		got      []int
		want     []int
	}{
		{"LeftSideView", LeftSideView(root), []int{1, 2, 5}},
		{"RightSideView", RightSideView(root), []int{1, 3, 4}},
		{"TopView", TopView(root), []int{2, 1, 3, 4}},
		{"BottomView", BottomView(root), []int{2, 5, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("%s got %v, want %v", tt.name, tt.got, tt.want)
			}
		})
	}
}
