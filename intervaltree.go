package intervaltree

import (
    "time"
    "errors"
    "fmt"
)


/* 
    Base IntervalTree class.
    Entry point for the tree (instead of just holding a pointer to root).
 */
type IntervalTree struct {
    root *IntervalTreeNode
}

func NewIntervalTree(node *IntervalTreeNode) *IntervalTree {
    return &IntervalTree{node}
}

func (tree *IntervalTree) Insert(i *Interval) error {
    tree.root = insertInterval(tree.root, i)
    return nil
}

func (tree *IntervalTree) PrintTree() {
    printNode(tree.root)
}

func (tree *IntervalTree) OverlapExists(i *Interval) bool {
    return overlapExists(tree.root, i)
}





/*
    A node in the interval tree, useful for balanced tree insertion
    without having to contaminate interval class.
 */
type IntervalTreeNode struct {
    I *Interval
    Max time.Time
    Left *IntervalTreeNode
    Right *IntervalTreeNode
}

func NewIntervalTreeNode(i *Interval) *IntervalTreeNode {
    node := new(IntervalTreeNode)
    node.I = i
    node.Max = i.End
    node.Left = nil
    node.Right = nil
    return node
}

func insertInterval(node *IntervalTreeNode, i *Interval) *IntervalTreeNode  {
    // base case: create new node out of interval + return
    if node == nil {
        return NewIntervalTreeNode(i)
    }

    // recursive case (non null node): pass either left or right
    start := node.I.Start
    if i.Start.Before(start) {
        node.Left = insertInterval(node.Left, i)
    } else {
        node.Right = insertInterval(node.Right, i)
    }

    // update max for searching later
    if node.Max.Before(i.End) {
        node.Max = i.End
    }

    return node
}

func printNode(node *IntervalTreeNode) {
    if node == nil {
        fmt.Println("end\n")
    } else {
        fmt.Println(*node.I)
        fmt.Println("left")
        printNode(node.Left)
        fmt.Println("right")
        printNode(node.Right)
    }
}

func overlapExists(node *IntervalTreeNode, i *Interval) bool {
    if node == nil {
        return false
    } else {
        if Overlaps(*node.I, *i) {
            return true
        } else if overlapExists(node.Left, i) {
            return true
        } else {
            return overlapExists(node.Right, i)
        }
    }
}



/*
    Start and End times only. 
    TODO: polymorphic intervals
 */
type Interval struct {
    Start time.Time
    End time.Time
}

func NewInterval(start time.Time, end time.Time) (*Interval, error) {
    if start.After(end) {
        return nil, errors.New("interval: end cannot come after start.")
    } else if start.Equal(end) {
        return nil, errors.New("interval: start cannot equal end.")
    } else {
        return &Interval{start, end}, nil
    }
}

// start_datetime1 <= end_datetime2 and end_datetime1 >= start_datetime2
func Overlaps(i1 Interval, i2 Interval) bool {
    return i1.Start.Before(i2.End) && i1.End.After(i2.Start)
}