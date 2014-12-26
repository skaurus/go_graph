package graphs

import (
	"math/rand"
	"reflect"
	"testing"
	//"fmt"
)

func createRandomGraph(edges int64, undirected bool) (ug *Graph) {
	ug = &Graph{
		VertexEdges: make(map[uint64]map[uint64]bool),
	}

	for i := int64(0); i < edges; i++ {
		from := uint64(rand.Int63() % edges)
		to := uint64(rand.Int63() % edges)
		if from != to {
			if _, ok := ug.VertexEdges[from]; ok {
				ug.VertexEdges[from][to] = true
			} else {
				ug.VertexEdges[from] = map[uint64]bool{to: true}
			}

			if undirected {
				if _, ok := ug.VertexEdges[to]; ok {
					ug.VertexEdges[to][from] = true
				} else {
					ug.VertexEdges[to] = map[uint64]bool{from: true}
				}
			}
		}
	}

	return
}

// This test checks if we can get by DFS the two paths that connects all the
// elements in two separate graphs without any connection between them
func TestUndDFS(t *testing.T) {
	gr := GetUndirected(
		[][2]uint64{
			[2]uint64{0, 1},
			[2]uint64{0, 2},
			[2]uint64{1, 2},
			[2]uint64{2, 3},
			[2]uint64{2, 4},

			[2]uint64{5, 6},
			[2]uint64{6, 7},
			[2]uint64{6, 9},
			[2]uint64{9, 5},
		},
	)

	expectedFromZero := map[uint64]bool{
		0: true,
		1: true,
		2: true,
		3: true,
		4: true,
	}
	expectedFromFive := map[uint64]bool{
		5: true,
		6: true,
		7: true,
		9: true,
	}
	if !reflect.DeepEqual(gr.Dfs(0), expectedFromZero) {
		t.Error("Expeceted path from Zero:", expectedFromZero, "but:", gr.Dfs(0), "obtained.")
	}
	if !reflect.DeepEqual(gr.Dfs(5), expectedFromFive) {
		t.Error("Expeceted path from Five:", expectedFromFive, "but:", gr.Dfs(5), "obtained.")
	}
}

func TestUndConnectedComponents(t *testing.T) {
	gr := GetUndirected(
		[][2]uint64{
			[2]uint64{0, 1},
			[2]uint64{0, 2},
			[2]uint64{1, 2},
			[2]uint64{2, 3},
			[2]uint64{2, 4},

			[2]uint64{5, 6},
			[2]uint64{6, 7},
			[2]uint64{6, 9},
			[2]uint64{9, 5},
		},
	)

	expected := []map[uint64]bool{
		map[uint64]bool{
			0: true,
			1: true,
			2: true,
			3: true,
			4: true,
		},
		map[uint64]bool{
			5: true,
			6: true,
			7: true,
			9: true,
		},
	}

	comps := gr.GetConnectedComponents()
	if len(comps) != len(expected) {
		t.Error("We expected:", len(expected), "components, but:", len(comps), "found")
	}

compLoop:
	for _, c := range comps {
		for _, ec := range expected {
			if reflect.DeepEqual(c, ec) {
				continue compLoop
			}
		}

		t.Error("No component found:", c)
	}
}

func TestUndBFS(t *testing.T) {
	gr := GetUndirected(
		[][2]uint64{
			[2]uint64{0, 1},
			[2]uint64{0, 2},
			[2]uint64{0, 5},
			[2]uint64{1, 2},
			[2]uint64{2, 3},
			[2]uint64{2, 4},
			[2]uint64{4, 3},
			[2]uint64{3, 5},
		},
	)

	expectedDistances := map[uint64]uint64{
		0: 0,
		1: 1,
		2: 1,
		3: 2,
		4: 2,
		5: 1,
	}
	expectedPaths := map[uint64]uint64{
		0: 0,
		1: 0,
		2: 0,
		3: 2,
		4: 2,
		5: 0,
	}
	path, dist := gr.Bfs(0)
	if !reflect.DeepEqual(path, expectedPaths) {
		t.Error("Expeceted paths from Zero:", expectedPaths, "but:", path, "obtained.")
	}

	if !reflect.DeepEqual(dist, expectedDistances) {
		t.Error("Expeceted distances from Zero:", expectedDistances, "but:", dist, "obtained.")
	}
}