package entity_test

import (
	"testing"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/le-michael/cozyworld/clock"
	"github.com/le-michael/cozyworld/entity"
)

func TestCurrentPosition(t *testing.T) {
	testcases := []struct {
		eq        entity.Motion
		advanceMs int64
		expected  mgl64.Vec2
	}{
		{
			eq:        entity.NewMotion(mgl64.Vec2{0, 0}, 0, 100.0, mgl64.Vec2{100.0, 0.0}),
			advanceMs: 0,
			expected:  mgl64.Vec2{0, 0},
		},
		{
			eq:        entity.NewMotion(mgl64.Vec2{0, 0}, 0, 100.0, mgl64.Vec2{100.0, 0.0}),
			advanceMs: 500,
			expected:  mgl64.Vec2{50.0, 0},
		},
		{
			eq:        entity.NewMotion(mgl64.Vec2{0, 0}, 0, 100.0, mgl64.Vec2{100.0, 0.0}),
			advanceMs: 1000,
			expected:  mgl64.Vec2{100.0, 0},
		},
		{
			eq:        entity.NewMotion(mgl64.Vec2{0, 0}, 0, 100.0, mgl64.Vec2{-100.0, 100.0}),
			advanceMs: 500,
			expected:  mgl64.Vec2{-35.35533905932738, 35.3553390593273},
		},
		{
			eq:        entity.NewMotion(mgl64.Vec2{0, 0}, 0, 100.0, mgl64.Vec2{-100.0, 100.0}),
			advanceMs: 1415,
			expected:  mgl64.Vec2{-100, 100},
		},
		{
			eq:        entity.NewMotion(mgl64.Vec2{0, 0}, 0, 100.0, mgl64.Vec2{-23.2, 40.2}),
			advanceMs: 500000,
			expected:  mgl64.Vec2{-23.2, 40.2},
		},
	}

	for _, tc := range testcases {
		fakeClock := clock.NewFakeClock(0)
		fakeClock.AdvanceMs(tc.advanceMs)
		got := tc.eq.CurrentPosition(fakeClock.CurrentTime())
		if !got.ApproxEqual(tc.expected) {
			t.Fatalf("CurrentPosition() = %v; want = %v", got, tc.expected)
		}
	}
}

func TestMoveTo(t *testing.T) {
	testcases := []struct {
		start     mgl64.Vec2
		speed     float64
		advanceMs []int64
		moveTo    []mgl64.Vec2
		expected  []mgl64.Vec2
	}{
		{
			start:     mgl64.Vec2{0.0, 0.0},
			speed:     100.0,
			advanceMs: []int64{0, 500, 1000},
			moveTo: []mgl64.Vec2{
				{100.0, 0.0},
				{100.0, 0.0},
				{150.0, 0.0},
			},
			expected: []mgl64.Vec2{
				{0.0, 0.0},
				{50.0, 0.0},
				{150.0, 0.0},
			},
		},
		{
			start:     mgl64.Vec2{400.0, -200.0},
			speed:     100.0,
			advanceMs: []int64{9999999, 99999999, 9999999, 9999999},
			moveTo: []mgl64.Vec2{
				{100.0, 0.0},
				{100.0, 0.0},
				{150.0, 0.0},
				{500.0, 0.0},
			},
			expected: []mgl64.Vec2{
				{100.0, 0.0},
				{100.0, 0.0},
				{150.0, 0.0},
				{500.0, 0.0},
			},
		},
	}

	for _, tc := range testcases {
		fakeClock := clock.NewFakeClock(0)
		m := entity.NewIdleMotion(tc.start, fakeClock.CurrentTime())
		for i := range tc.advanceMs {
			m = m.MoveTo(fakeClock.CurrentTime(), tc.speed, tc.moveTo[i])
			fakeClock.AdvanceMs(tc.advanceMs[i])
			got := m.CurrentPosition(fakeClock.CurrentTime())
			if !got.ApproxEqual(tc.expected[i]) {
				t.Fatalf("CurrentPosition() = %v; want = %v", got, tc.expected[i])
			}
		}
	}
}
