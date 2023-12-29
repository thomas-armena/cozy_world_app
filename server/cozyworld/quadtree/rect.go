package quadtree

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Rect struct {
	top    float64
	left   float64
	width  float64
	height float64

	// Closed referes to whether or not we allow points to exist on the
	// far edge of a rect.
	//
	// For example, if closedRight is true then a point can be placed on the right edge.
	closedRight  bool
	closedBottom bool
}

func NewClosedRect(top, left, width, height float64) Rect {
	return NewRect(top, left, width, height, true, true)
}

func NewRect(top, left, width, height float64, closedRight, closedBottom bool) Rect {
	return Rect{
		top:    top,
		left:   left,
		width:  width,
		height: height,

		closedRight:  closedRight,
		closedBottom: closedBottom,
	}
}

// Split divides a Rect into 4 sections.
func (r Rect) Split() (topLeft, topRight, bottomLeft, bottomRight Rect) {
	topLeft = NewRect(r.top, r.left, r.width/2, r.height/2, false, false)
	topRight = NewRect(r.top, r.left+(r.width/2), r.width/2, r.height/2, true, false)
	bottomLeft = NewRect(r.top+(r.height/2), r.left, r.width/2, r.height/2, false, true)
	bottomRight = NewRect(r.top+(r.height/2), r.left+(r.width/2), r.width/2, r.height/2, true, true)
	return topLeft, topRight, bottomLeft, bottomRight
}

// Contains returns whether a point is within a Rect.
func (r Rect) Contains(point mgl64.Vec2) bool {
	ok := true
	if r.closedBottom {
		ok = ok &&
			r.top <= point.Y() &&
			point.Y() <= r.top+r.height

	} else {
		ok = ok &&
			r.top <= point.Y() &&
			point.Y() < r.top+r.height
	}

	if r.closedRight {
		ok = ok &&
			r.left <= point.X() && point.X() <= r.left+r.width
	} else {
		ok = ok &&
			r.left <= point.X() &&
			point.X() < r.left+r.width
	}

	return ok
}

// Overlaps return whether two Rects overlap.
func (r Rect) Overlaps(o Rect) bool {
	return !(r.left+r.width < o.left ||
		r.left > o.left+o.width ||
		r.top+r.height < o.top ||
		r.top > o.top+o.height)
}
