package graph

// disjoint set implementation

type DisjointSet struct {
	father []int
}

func BuildDisjointSet(n int) *DisjointSet {
	father := make([]int, n)
	for i := 0; i < n; i++ {
		father[i] = i
	}
	return &DisjointSet{father: father}
}

// Find the v root, contains path compress
func (d *DisjointSet) Find(v int) int {
	u := d.father[v]
	if u == v {
		return u
	} else {
		// path compress
		u = d.Find(u)
		d.father[v] = u
		return u
	}
}

// Union w and v, not implement rank optimization
func (d *DisjointSet) Union(w, v int) {
	w = d.Find(w)
	v = d.Find(v)
	if w == v {
		return
	}
	// w -> v
	d.father[v] = w
}

// IsSame w is same v
func (d *DisjointSet) IsSame(w, v int) bool {
	wRoot := d.Find(w)
	vRoot := d.Find(v)
	return wRoot == vRoot
}
