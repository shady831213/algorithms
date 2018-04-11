package disjointSetTree

type disjointSet struct {
	p     *disjointSet
	rank  int
	Value interface{}
}

func findSet(e *disjointSet) *disjointSet {
	if e.p != e {
		e.p = findSet(e.p)
	}
	return e.p
}

func makeSet(value interface{}) *disjointSet {
	t := new(disjointSet)
	t.Value = value
	t.p = t
	t.rank = 0
	return t
}

func union(e1, e2 *disjointSet) *disjointSet {
	return link(findSet(e1), findSet(e2))
}

func link(s1, s2 *disjointSet) *disjointSet {
	if s1 != s2 {
		if s1.rank < s2.rank {
			s1.p = s2
			return s2
		}
		s2.p = s1
		if s1.rank == s2.rank {
			s1.rank++
		}
		return s1
	}
	return s1
}
