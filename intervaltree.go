package intervaltree

import (
    "time"
    "errors"
    // "fmt"
)


// wishlist
// polymorphic intervals with an interface and a method to compare them
// a payload within each interval
// AVL tree or red black tree insertion/deletion instead of BST


/* 
    Base IntervalTree class.
    Entry point for the tree (instead of just holding a pointer to root).
    Use *IntervalTree Receiver Type to avoid copying
 */
type IntervalTree struct {
    root *IntervalTreeNode
}

func NewIntervalTree() *IntervalTree {
    // constructor: create tree (empty)
    return &IntervalTree{nil}
}

func (tree *IntervalTree) Empty() bool {
    // Empty: returns boolean if tree is empty
    return tree.root == nil
}

func (tree *IntervalTree) Insert(i *Interval) {
    // Insert: inserts an interval
    if tree.Empty() {
        tree.root = newIntervalTreeNode(i)
    } else {
        tree.root.insert(i)
    }
}

func (tree *IntervalTree) Max() (*Interval, error) {
    // Max: furthest right interval (TODO: consider returning nil, no error)
    if tree.Empty() {
        return nil, errors.New("IntervalTree::Max requires non-empty tree.")
    } else {
        max := tree.root.max()
        return max, nil
    }
}

func (tree *IntervalTree) Min() (*Interval, error) {
    // Min: furthest left interval (TODO: consider returning nil, no error)
    if tree.Empty() {
        return nil, errors.New("IntervalTree::Min requires non-empty tree.")
    } else {
        min := tree.root.min()
        return min, nil
    }
}

func (tree *IntervalTree) FindOverlap(i Interval) []Interval {
    // FindOverlap: find all intervals overlapping with an interval
    if tree.Empty() {
        var overlaps []Interval
        return overlaps
    } else {
        overlaps := tree.root.findOverlap(i)
        return overlaps
    }
}

func (tree *IntervalTree) Overlaps(i Interval) bool {
    // Overlaps: check if any interval in tree overlaps an interval
    if tree.Empty() {
        return false
    } else {
        return tree.root.overlaps(i)
    }
}




/*
    A node in the interval tree, useful for balanced tree insertion
    without having to contaminate interval class.
 */
type IntervalTreeNode struct {
    i *Interval
    subTreeMax time.Time
    left *IntervalTreeNode
    right *IntervalTreeNode
}

func newIntervalTreeNode(i *Interval) *IntervalTreeNode {
    // TODO: return error if i is nil
    node := new(IntervalTreeNode)
    node.i = i
    node.subTreeMax = i.End()
    return node
}

func (node *IntervalTreeNode) insert(i *Interval) *IntervalTreeNode  {
    start := node.i.Start()

    if i.End().Before(start) {
        if node.left == nil {
            node.left = newIntervalTreeNode(i)
        } else {
            node.left.insert(i)
        }
    } else {
        if node.right == nil {
            node.right = newIntervalTreeNode(i)
        } else {
            node.right.insert(i)
        }
    }

    // update max for searching later
    if node.subTreeMax.Before(i.End()) {
        node.subTreeMax = i.End()
    }

    return node
}

func (node *IntervalTreeNode) findOverlap(i Interval) []Interval {
    // TODO: be more efficient with searching, this is just going through every single node
    var overlaps []Interval

    if Overlaps(*node.i, i) {
        overlaps = append(overlaps, *node.i)
    }

    if node.left != nil {
        overlaps = append(overlaps, node.left.findOverlap(i)...)
    }

    if node.right != nil {
        overlaps = append(overlaps, node.right.findOverlap(i)...)
    }

    return overlaps
}

func (node *IntervalTreeNode) max() *Interval {
    if node.right != nil{
        return node.right.max()
    } else {
        return node.i
    }
}

func (node *IntervalTreeNode) min() *Interval {
    if node.left != nil{
        return node.left.min()
    } else {
        return node.i
    }
}

func (node *IntervalTreeNode) overlaps(i Interval) bool {
    // TODO: be more efficient searching
    if Overlaps(*node.i, i) {
        return true
    } else if node.left != nil && node.left.overlaps(i) {
        return true
    } else if node.right != nil && node.right.overlaps(i) {
        return true
    } else {
        return false
    }
}



/*
    Start and End times only. 
    TODO: polymorphic intervals
    Immutable start and end so that intervals are always valid.
 */
type Interval struct {
    start time.Time
    end time.Time
    // Payload PayLoad
}

func NewInterval(start time.Time, end time.Time) (*Interval, error) {
    if start.After(end) {
        return nil, errors.New("Interval::NewInterval end cannot come after start.")
    } else if start.Equal(end) {
        return nil, errors.New("Interval::NewInterval start cannot equal end.")
    } else {
        return &Interval{start, end}, nil
    }
}

func (i *Interval) Start() time.Time {
    return i.start
}

func (i *Interval) End() time.Time {
    return i.end
}

// start_datetime1 <= end_datetime2 and end_datetime1 >= start_datetime2
func Overlaps(i1 Interval, i2 Interval) bool {
    return i1.Start().Before(i2.End()) && i1.End().After(i2.Start())
}