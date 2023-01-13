package bookingkamaroperasi

import (
	"errors"
	"log"
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
		log.Println(data.endDate, n.startDate)
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
		log.Println("new", "\t", t.root.data.startDate, t.root.data.endDate)
		return true, nil
	} else {
		ok, err := t.root.insert(data)
		log.Println("BinaryTree.insert ok", ok, err)
		return ok, err
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

	log.Println("node", "\t", node)
	log.Println("diff", "\t", diff, diff < 2*time.Hour)
	log.Println("existing", "\t", n.startDate, n.endDate)
	log.Println("new", "\t", data.startDate, data.endDate)

	if diff < 2*time.Hour {
		log.Println("BookingKamarOperasi_v2.isLessThanTwoHoursApart return 2, nil")
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
	} else if node == 1 {
		if isLess, err := n.data.isLessThanTwoHoursApart(data, 1); isLess == 0 {
			return false, err
		} else if isLess == 2 {
			return false, nil
		} else if isLess == 1 {
			// continue
		} else {
			return false, errors.New("isLess " + strconv.Itoa(isLess) + " is not implemented")
		}

		if n.left == nil {
			n.left = &BinaryNode{left: nil, right: nil, data: data}
		} else {
			return n.left.insert(data)
		}

		return true, nil
	} else if node == 2 {
		if isLess, err := n.data.isLessThanTwoHoursApart(data, 2); isLess == 0 {
			return false, err
		} else if isLess == 2 {
			log.Println("BinaryNode.isLess " + strconv.Itoa(isLess) + " " + strconv.FormatBool(false))
			return false, nil
		} else if isLess == 1 {
			// continue
		} else {
			return false, errors.New("isLess " + strconv.Itoa(isLess) + " is not implemented")
		}

		if n.right == nil {
			log.Println("CREATION")
			n.right = &BinaryNode{left: nil, right: nil, data: data}
		} else {
			log.Println("RECURSIVE")
			return n.right.insert(data)
		}

		log.Println("triggered")
		return true, nil
	} else {
		return false, errors.New("unknown error")
	}
}

type SafeDB_v2 struct {
	mu sync.Mutex
	v  BinaryTree
}

func (c *SafeDB_v2) TryInsert(startDate time.Time, duration time.Duration) (bool, error) {
	c.mu.Lock()
	ok, err := c.v.insert(startDate, startDate.Add(duration))
	log.Println("SafeDB_v2.Insert " + strconv.FormatBool(ok))
	log.Println()
	c.mu.Unlock()

	return ok, err
}

var C_v2 SafeDB_v2

func InitializeDB_v2() {
	C_v2 = SafeDB_v2{v: BinaryTree{}}
}
