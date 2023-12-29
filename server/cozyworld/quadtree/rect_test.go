package quadtree_test

import (
	"testing"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/le-michael/cozyworld/quadtree"
)

func TestClosedRectContains(t *testing.T) {
	closed := quadtree.NewClosedRect(0, 0, 5, 5)

	validPoints := []mgl64.Vec2{
		{0, 0}, {5, 0}, {5, 5},
		{0, 5}, {2.5, 3.3}, {1.2, 5},
	}
	for _, point := range validPoints {
		if !closed.Contains(point) {
			t.Fatalf("closed.Contains(%v) = false; want = true", point)
		}
	}

	invalidPoints := []mgl64.Vec2{
		{-0.1, 0}, {-0.1, 2}, {0, -0.1},
		{-0.1, -20}, {2, 5000}, {-0.0001, -0.00001},
	}
	for _, point := range invalidPoints {
		if closed.Contains(point) {
			t.Fatalf("closed.Contains(%v) = true; want = false", point)
		}
	}

}

func TestSplit(t *testing.T) {
	closed := quadtree.NewClosedRect(0, 0, 5, 5)
	tl, tr, bl, br := closed.Split()

	testcases := []struct {
		point      mgl64.Vec2
		tlExpected bool
		trExpected bool
		blExpected bool
		brExpected bool
	}{
		{
			point:      mgl64.Vec2{0.0, 0.0},
			tlExpected: true, trExpected: false, blExpected: false, brExpected: false,
		},
		{
			point:      mgl64.Vec2{2.5, 0.0},
			tlExpected: false, trExpected: true, blExpected: false, brExpected: false,
		},
		{
			point:      mgl64.Vec2{5.0, 0.0},
			tlExpected: false, trExpected: true, blExpected: false, brExpected: false,
		},
		{
			point:      mgl64.Vec2{0, 2.5},
			tlExpected: false, trExpected: false, blExpected: true, brExpected: false,
		},
		{
			point:      mgl64.Vec2{2.5, 2.5},
			tlExpected: false, trExpected: false, blExpected: false, brExpected: true,
		},
		{
			point:      mgl64.Vec2{5.0, 5.0},
			tlExpected: false, trExpected: false, blExpected: false, brExpected: true,
		},
	}

	for _, tc := range testcases {
		if tl.Contains(tc.point) != tc.tlExpected {
			t.Fatalf(
				"tl.Contains(%v) = %v; want = %v", tc.point, tl.Contains(tc.point), tc.tlExpected)
		}
		if tr.Contains(tc.point) != tc.trExpected {
			t.Fatalf(
				"tr.Contains(%v) = %v; want = %v", tc.point, tr.Contains(tc.point), tc.trExpected)
		}
		if bl.Contains(tc.point) != tc.blExpected {
			t.Fatalf(
				"bl.Contains(%v) = %v; want = %v", tc.point, bl.Contains(tc.point), tc.blExpected)
		}
		if br.Contains(tc.point) != tc.brExpected {
			t.Fatalf(
				"br.Contains(%v) = %v; want = %v", tc.point, br.Contains(tc.point), tc.brExpected)
		}
	}
}
