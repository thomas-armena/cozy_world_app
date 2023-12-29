package quadtree_test

import (
	"testing"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/le-michael/cozyworld/quadtree"
)

func TestQuery(t *testing.T) {
	top := -50.0
	left := -50.0
	width := 100.0
	height := 100.0
	q := quadtree.NewQuadTreeNode[int](quadtree.NewClosedRect(top, left, width, height), nil)

	total := 0
	for i := 0; i <= int(width); i++ {
		for j := 0; j <= int(height); j++ {
			q.Insert(mgl64.Vec2{top + float64(i), left + float64(j)}, total)
			total++
		}
	}

	all := q.Query(q.Bounds())
	if len(all) != total {
		t.Fatalf("len(q.Query(q.Bounds())) = %v: want = %v", len(all), total)
	}

	// Should be five points between [(0.5, 0.5), (5.5, 1.5)]
	singleRow := q.Query(quadtree.NewClosedRect(0.5, 0.5, 5.0, 1.0))
	if len(singleRow) != 5 {
		t.Fatalf("len(singleRow) = %v: want = %v", len(singleRow), 5)
	}

	// Should be 22 points between [(35, 35), (45, 36)]
	twoRows := q.Query(quadtree.NewClosedRect(35, 35, 10, 1))
	if len(twoRows) != 22 {
		t.Fatalf("len(twoRows) = %v: want = %v", len(twoRows), 22)
	}

	singlePoint := q.Query(quadtree.NewClosedRect(49.5, 49.5, 1, 1))
	if len(singlePoint) != 1 {
		t.Fatalf("len(singlePoint) = %v: want = %v", len(singleRow), 1)
	}
	if singlePoint[0] != total-1 {
		t.Fatalf("singlePoint[0] = %v: want = %v", singlePoint[0], total-1)
	}
}
