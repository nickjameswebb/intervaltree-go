package intervaltree

import (
    "time"
    "errors"
)


type IntervalTree struct {
    Root *IntervalTreeNode
}

func NewIntervalTree(inode *IntervalTreeNode) *IntervalTree {
    return &IntervalTree{inode}
}



type IntervalTreeNode struct {
    I *Interval
    Max time.Time
    Left *IntervalTreeNode
    Right *IntervalTreeNode
}

func NewIntervalTreeNode(i *Interval) *IntervalTreeNode {
    inode := new(IntervalTreeNode)
    inode.I = i
    inode.Max = i.End
    inode.Left = nil
    inode.Right = nil
    return inode
}



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