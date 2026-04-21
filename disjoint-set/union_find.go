package main

import "fmt"

// UnionFind implements the Disjoint Set Union (DSU) data structure.
type UnionFind struct {
	parent []int
	rank   []int
}

// NewUnionFind initializes a Union-Find structure with n elements.
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent: parent, rank: rank}
}

// Find returns the representative of the set containing i, with path compression.
func (uf *UnionFind) Find(i int) int {
	if uf.parent[i] == i {
		return i
	}
	uf.parent[i] = uf.Find(uf.parent[i])
	return uf.parent[i]
}

// Union merges the sets containing i and j using union by rank.
func (uf *UnionFind) Union(i, j int) {
	rootI := uf.Find(i)
	rootJ := uf.Find(j)
	if rootI != rootJ {
		if uf.rank[rootI] < uf.rank[rootJ] {
			uf.parent[rootI] = rootJ
		} else if uf.rank[rootI] > uf.rank[rootJ] {
			uf.parent[rootJ] = rootI
		} else {
			uf.parent[rootI] = rootJ
			uf.rank[rootJ]++
		}
	}
}

func main() {
	uf := NewUnionFind(10)
	uf.Union(1, 2)
	uf.Union(2, 3)
	fmt.Printf("Find(1): %d\n", uf.Find(1))
	fmt.Printf("Find(3): %d\n", uf.Find(3))
	fmt.Printf("Are 1 and 3 connected? %v\n", uf.Find(1) == uf.Find(3))
	fmt.Printf("Are 1 and 4 connected? %v\n", uf.Find(1) == uf.Find(4))
}
