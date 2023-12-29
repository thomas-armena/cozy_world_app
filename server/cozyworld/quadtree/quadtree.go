package quadtree

import (
	"github.com/go-gl/mathgl/mgl64"
)

const (
	// Maximum number of items in a node before it splits.
	DefaultBucketSize = 32
)

type QuadTreeOpts struct {
	MaxBucketSize int
}

var (
	DefaultBucketOpts = &QuadTreeOpts{MaxBucketSize: DefaultBucketSize}
)

type QuadTreePoint[T any] struct {
	position mgl64.Vec2
	value    T
}

type QuadTreeNode[T any] struct {
	topLeft     *QuadTreeNode[T]
	topRight    *QuadTreeNode[T]
	bottomLeft  *QuadTreeNode[T]
	bottomRight *QuadTreeNode[T]

	bounds Rect
	points []QuadTreePoint[T]

	opts *QuadTreeOpts
}

func NewQuadTreeNode[T any](bounds Rect, opts *QuadTreeOpts) *QuadTreeNode[T] {
	if opts == nil {
		opts = DefaultBucketOpts
	}
	return &QuadTreeNode[T]{
		bounds: bounds,
		points: make([]QuadTreePoint[T], 0),
		opts:   opts,
	}
}

func (q *QuadTreeNode[T]) isLeaf() bool {
	return q.topLeft == nil &&
		q.topRight == nil &&
		q.bottomLeft == nil &&
		q.bottomRight == nil
}

func (q *QuadTreeNode[T]) Bounds() Rect {
	return q.bounds
}

func (q *QuadTreeNode[T]) Insert(position mgl64.Vec2, value T) {
	q.insert(QuadTreePoint[T]{position: position, value: value})
}

func (q *QuadTreeNode[T]) insert(point QuadTreePoint[T]) {
	if !q.bounds.Contains(point.position) {
		return
	}

	if len(q.points) < DefaultBucketSize {
		q.points = append(q.points, point)
		return
	}

	if q.isLeaf() {
		tlBound, trBound, blBound, brBound := q.bounds.Split()
		q.topLeft = NewQuadTreeNode[T](tlBound, q.opts)
		q.topRight = NewQuadTreeNode[T](trBound, q.opts)
		q.bottomLeft = NewQuadTreeNode[T](blBound, q.opts)
		q.bottomRight = NewQuadTreeNode[T](brBound, q.opts)
	}

	// TODO: Add a check to make a point is only added once.
	q.topLeft.insert(point)
	q.topRight.insert(point)
	q.bottomLeft.insert(point)
	q.bottomRight.insert(point)
}

func (q *QuadTreeNode[T]) Query(area Rect) []T {
	res := make([]T, 0)
	q.query(area, &res)
	return res
}

func (q *QuadTreeNode[T]) query(area Rect, res *[]T) {
	if !q.bounds.Overlaps(area) {
		return
	}

	for _, point := range q.points {
		if area.Contains(point.position) {
			*res = append((*res), point.value)
		}
	}

	if q.isLeaf() {
		return
	}

	q.topLeft.query(area, res)
	q.topRight.query(area, res)
	q.bottomLeft.query(area, res)
	q.bottomRight.query(area, res)
}
