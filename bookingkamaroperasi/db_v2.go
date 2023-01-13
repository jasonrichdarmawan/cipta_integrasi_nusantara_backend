package bookingkamaroperasi

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

type BookingKamarOperasi_v2 struct {
	startDate time.Time
	endDate   time.Time
}

// Return
// 0 = the startDate overlaps with an existing booking
// 1 = left
// 2 = right
func (n *BookingKamarOperasi_v2) leftOrRight(data BookingKamarOperasi_v2) (int, error) {
	if data.endDate.Before(n.startDate) {
		// scenario 1: endDate n.startDate
		return 1, nil
	} else {
		// scenario 2: n.startDate endDate

		if data.startDate.Before(n.endDate) {
			// scenario 3: n.startDate startDate n.endDate
			return 0, errors.New("startDate overlaps with an existing booking")
		}

		return 2, nil
	}
}

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode

	data BookingKamarOperasi_v2
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(startDate, endDate time.Time) (bool, error) {
	data := BookingKamarOperasi_v2{startDate: startDate, endDate: endDate}

	if t.root == nil {
		t.root = &BinaryNode{
			left: nil, right: nil,
			data: data}
		return true, nil
	} else {
		return t.root.insert(data)
	}
}

// Return
// 0 = this method is not to be used directly
// 1 = false
// 2 = true
func (n *BookingKamarOperasi_v2) isLessThanTwoHoursApart(data BookingKamarOperasi_v2, node int) (int, error) {
	var diff time.Duration
	if node == 1 {
		// endDate n.startDate
		diff = n.startDate.Sub(data.endDate)
	} else if node == 2 {
		// n.startDate n.endDate startDate
		diff = data.startDate.Sub(n.endDate)
	} else {
		return 0, errors.New("this method is not to be used directly")
	}

	if diff < 2*time.Hour {
		return 2, nil
	} else {
		return 1, nil
	}

}

// This method is not to be used directly.
func (n *BinaryNode) insert(data BookingKamarOperasi_v2) (bool, error) {
	// error.
	if n == nil {
		return false, errors.New("this method is not to be used directly")
	}

	node, err := n.data.leftOrRight(data)
	if node == 0 {
		return false, err
	}

	if isLess, err := n.data.isLessThanTwoHoursApart(data, node); isLess == 0 {
		return false, err
	} else if isLess == 2 {
		return false, nil
	} else if isLess == 1 {
		// continue
	} else {
		return false, errors.New("isLess " + strconv.Itoa(isLess) + " is not implemented")
	}

	if node == 1 {
		if n.left == nil {
			n.left = &BinaryNode{left: nil, right: nil, data: data}
		} else {
			return n.left.insert(data)
		}
	} else if node == 2 {
		if n.right == nil {
			n.right = &BinaryNode{left: nil, right: nil, data: data}
		} else {
			return n.right.insert(data)
		}
	} else {
		return false, err
	}

	return true, nil
}

type SafeDB_v2 struct {
	mu sync.Mutex
	v  BinaryTree
}

func (c *SafeDB_v2) TryInsert(startDate time.Time, duration time.Duration) (bool, error) {
	c.mu.Lock()
	ok, err := c.v.insert(startDate, startDate.Add(duration))
	c.mu.Unlock()

	return ok, err
}

var C_v2 SafeDB_v2

func InitializeDB_v2() {
	C_v2 = SafeDB_v2{v: BinaryTree{}}
}
