package intervaltree

import "testing"
import "time"

func TestTreeCreation(t *testing.T) {
    var tree *IntervalTree = NewIntervalTree()
    if tree == nil {
        t.Error("NewIntervalTree should return a non-null pointer to a tree.")
    }
}

func TestInsertEmptyTree(t *testing.T) {
    tree := NewIntervalTree()
    start := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2017, 10, 2, 0, 0, 0, 0, time.UTC)
    interval, _ := NewInterval(start, end)

    tree.Insert(interval)

    if tree.root == nil || tree.root.i != interval {
        t.Error("IntervalTree root should contain interval.")
    } else if tree.root.left != nil || tree.root.right != nil {
        t.Error("IntervalTree root should be only node in tree.")
    } else if tree.root.subTreeMax != end {
        t.Error("IntervalTree root has incorrect subTreeMax.")
    }
}

func TestInsertLeftOfRoot(t *testing.T) {
    // insert an interval to left of root (root only node in tree)
    tree := NewIntervalTree()

    // root node
    start := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2017, 10, 2, 0, 0, 0, 0, time.UTC)
    interval, _ := NewInterval(start, end)
    tree.Insert(interval)

    // left node
    start2 := time.Date(2016, 10, 1, 0, 0, 0, 0, time.UTC)
    end2 := time.Date(2016, 10, 2, 0, 0, 0, 0, time.UTC)
    interval2, _ := NewInterval(start2, end2)
    tree.Insert(interval2)

    testNode := newIntervalTreeNode(interval2)

    if tree.root.left == nil || *tree.root.left != *testNode {
        t.Error("Node was not inserted correctly.")
    }
}

func TestInsertRightOfRoot(t *testing.T) {
    // insert an interval to right of root (root only node in tree)
    tree := NewIntervalTree()

    // root node
    start := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2017, 10, 2, 0, 0, 0, 0, time.UTC)
    interval, _ := NewInterval(start, end)
    tree.Insert(interval)

    // left node
    start2 := time.Date(2018, 10, 1, 0, 0, 0, 0, time.UTC)
    end2 := time.Date(2018, 10, 2, 0, 0, 0, 0, time.UTC)
    interval2, _ := NewInterval(start2, end2)
    tree.Insert(interval2)

    testNode := newIntervalTreeNode(interval2)

    if tree.root.right == nil || *tree.root.right != *testNode {
        t.Error("Node was not inserted correctly.")
    }
}

func TestInsertOverlappingRootStart(t *testing.T) {
    // insert an interval that overlaps start of root (root only node in tree)
    // should be inserted to the right (based on insertion algorithm)
    tree := NewIntervalTree()

    // root node
    start := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2017, 10, 2, 0, 0, 0, 0, time.UTC)
    interval, _ := NewInterval(start, end)
    tree.Insert(interval)

    // left node
    start2 := time.Date(2017, 9, 1, 0, 0, 0, 0, time.UTC)
    end2 := time.Date(2017, 10, 1, 2, 0, 0, 0, time.UTC)
    interval2, _ := NewInterval(start2, end2)
    tree.Insert(interval2)

    testNode := newIntervalTreeNode(interval2)

    if tree.root.right == nil || *tree.root.right != *testNode {
        t.Error("Node was not inserted correctly.")
    }
}

func TestInsertOverlappingRootEnd(t *testing.T) {
    // insert an interval that overlaps end of root (root only node in tree)
    // should be inserted to the right (based on insertion algorithm)
    tree := NewIntervalTree()

    // root node
    start := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2017, 10, 2, 0, 0, 0, 0, time.UTC)
    interval, _ := NewInterval(start, end)
    tree.Insert(interval)

    // left node
    start2 := time.Date(2017, 10, 1, 2, 0, 0, 0, time.UTC)
    end2 := time.Date(2017, 10, 3, 0, 0, 0, 0, time.UTC)
    interval2, _ := NewInterval(start2, end2)
    tree.Insert(interval2)

    testNode := newIntervalTreeNode(interval2)

    if tree.root.right == nil || *tree.root.right != *testNode {
        t.Error("Node was not inserted correctly.")
    }
}

func TestInsertWithinRoot(t *testing.T) {
    // insert an interval that is encapsulated by root (root only node in tree)
    // should be inserted to the right (based on insertion algorithm)
    tree := NewIntervalTree()

    // root node
    start := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2017, 10, 2, 0, 0, 0, 0, time.UTC)
    interval, _ := NewInterval(start, end)
    tree.Insert(interval)

    // left node
    start2 := time.Date(2017, 10, 1, 2, 0, 0, 0, time.UTC)
    end2 := time.Date(2017, 10, 1, 4, 0, 0, 0, time.UTC)
    interval2, _ := NewInterval(start2, end2)
    tree.Insert(interval2)

    testNode := newIntervalTreeNode(interval2)

    if tree.root.right == nil || *tree.root.right != *testNode {
        t.Error("Node was not inserted correctly.")
    }
}

func TestInsertEncapsulatingRoot(t *testing.T) {
    // insert an interval that encapsulates root (root only node in tree)
    // should be inserted to the right (based on insertion algorithm)
    tree := NewIntervalTree()

    // root node
    start := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2017, 10, 2, 0, 0, 0, 0, time.UTC)
    interval, _ := NewInterval(start, end)
    tree.Insert(interval)

    // left node
    start2 := time.Date(2016, 10, 1, 2, 0, 0, 0, time.UTC)
    end2 := time.Date(2018, 10, 1, 4, 0, 0, 0, time.UTC)
    interval2, _ := NewInterval(start2, end2)
    tree.Insert(interval2)

    testNode := newIntervalTreeNode(interval2)

    if tree.root.right == nil || *tree.root.right != *testNode {
        t.Error("Node was not inserted correctly.")
    }
}

func TestOverlapsEmptyTree(t *testing.T) {
    tree := NewIntervalTree()

    start := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2017, 10, 2, 0, 0, 0, 0, time.UTC)
    interval, _ := NewInterval(start, end)

    if tree.Overlaps(interval) {
        t.Error("Tree should not detect overlap when empty.")
    }
}

func TestOverlapsPopulatedTree(t *testing.T) {
    tree := NewIntervalTree()

    start := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2017, 10, 2, 0, 0, 0, 0, time.UTC)
    interval, _ := NewInterval(start, end)

    start2 := time.Date(2017, 10, 3, 0, 0, 0, 0, time.UTC)
    end2 := time.Date(2017, 10, 4, 0, 0, 0, 0, time.UTC)
    interval2, _ := NewInterval(start2, end2)

    start3 := time.Date(2015, 10, 1, 0, 0, 0, 0, time.UTC)
    end3 := time.Date(2017, 4, 2, 0, 0, 0, 0, time.UTC)
    interval3, _ := NewInterval(start3, end3)

    start4 := time.Date(2016, 10, 1, 0, 0, 0, 0, time.UTC)
    end4 := time.Date(2016, 10, 30, 0, 0, 0, 0, time.UTC)
    interval4, _ := NewInterval(start4, end4)

    tree.Insert(interval)
    tree.Insert(interval2)
    tree.Insert(interval3)
    tree.Insert(interval4)

    // overlaps nothing
    start5 := time.Date(2017, 4, 10, 0, 0, 0, 0, time.UTC)
    end5 := time.Date(2017, 4, 11, 0, 0, 0, 0, time.UTC)
    testInterval1, _ := NewInterval(start5, end5)

    // overlaps all intervals
    start6 := time.Date(2014, 10, 1, 0, 0, 0, 0, time.UTC)
    end6 := time.Date(2019, 10, 2, 0, 0, 0, 0, time.UTC)
    testInterval2, _ := NewInterval(start6, end6)

    if tree.Overlaps(testInterval1) {
        t.Error("Tree should not detect overlap with testInterval1.")
    }

    if !tree.Overlaps(testInterval2) {
        t.Error("Tree should detect overlap with testInterval2.")
    }
}

func TestFindOverlapEmptyTree(t *testing.T) {
    tree := NewIntervalTree()
    interval, _ := NewInterval(
            time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC), 
            time.Date(2017, 10, 2, 0, 0, 0, 0, time.UTC),
        )
    overlaps := tree.FindOverlap(interval)
    if len(overlaps) != 0 {
        t.Error("Tree found a phantom overlapping interval.")
    }
}

func TestFindOverlapNothingReturned(t *testing.T) {
    tree := NewIntervalTree()

    start := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2017, 10, 2, 0, 0, 0, 0, time.UTC)
    interval, _ := NewInterval(start, end)

    start2 := time.Date(2017, 10, 3, 0, 0, 0, 0, time.UTC)
    end2 := time.Date(2017, 10, 4, 0, 0, 0, 0, time.UTC)
    interval2, _ := NewInterval(start2, end2)

    start3 := time.Date(2015, 10, 1, 0, 0, 0, 0, time.UTC)
    end3 := time.Date(2017, 4, 2, 0, 0, 0, 0, time.UTC)
    interval3, _ := NewInterval(start3, end3)

    start4 := time.Date(2016, 10, 1, 0, 0, 0, 0, time.UTC)
    end4 := time.Date(2016, 10, 30, 0, 0, 0, 0, time.UTC)
    interval4, _ := NewInterval(start4, end4)

    tree.Insert(interval)
    tree.Insert(interval2)
    tree.Insert(interval3)
    tree.Insert(interval4)

    start5 := time.Date(2017, 4, 10, 0, 0, 0, 0, time.UTC)
    end5 := time.Date(2017, 4, 11, 0, 0, 0, 0, time.UTC)
    testInterval1, _ := NewInterval(start5, end5)

    overlaps := tree.FindOverlap(testInterval1)
    if len(overlaps) != 0 {
        t.Error("Tree found phantom overlapping interval(s) with testInterval1.")
    }
}

func TestFindOverlapMultipleReturned(t *testing.T) {
    tree := NewIntervalTree()

    start := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2017, 10, 2, 0, 0, 0, 0, time.UTC)
    interval, _ := NewInterval(start, end)

    start2 := time.Date(2017, 10, 3, 0, 0, 0, 0, time.UTC)
    end2 := time.Date(2017, 10, 4, 0, 0, 0, 0, time.UTC)
    interval2, _ := NewInterval(start2, end2)

    start3 := time.Date(2015, 10, 1, 0, 0, 0, 0, time.UTC)
    end3 := time.Date(2017, 4, 2, 0, 0, 0, 0, time.UTC)
    interval3, _ := NewInterval(start3, end3)

    start4 := time.Date(2016, 10, 1, 0, 0, 0, 0, time.UTC)
    end4 := time.Date(2016, 10, 30, 0, 0, 0, 0, time.UTC)
    interval4, _ := NewInterval(start4, end4)

    tree.Insert(interval)
    tree.Insert(interval2)
    tree.Insert(interval3)
    tree.Insert(interval4)

    start5 := time.Date(2014, 4, 10, 0, 0, 0, 0, time.UTC)
    end5 := time.Date(2018, 4, 11, 0, 0, 0, 0, time.UTC)
    testInterval1, _ := NewInterval(start5, end5)

    overlaps := tree.FindOverlap(testInterval1)
    numOverlaps := len(overlaps)
    expectedNumOverlaps := 4
    if numOverlaps != expectedNumOverlaps {
        t.Error("Tree did not find correct overlaps. Expected", expectedNumOverlaps, "found", numOverlaps, "instead.")
    }
}

func TestTreeEmpty(t *testing.T) {
    tree := NewIntervalTree()
    if tree.Empty() == false {
        t.Error("NewIntervalTree should return an empty tree.")
    }

    start := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2017, 10, 2, 0, 0, 0, 0, time.UTC)
    interval, _ := NewInterval(start, end)
    tree.Insert(interval)

    if tree.Empty() == true {
        t.Error("IntervalTree should not be empty after insertion.")
    }
}
