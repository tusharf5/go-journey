package ds

type MinHeap struct {
	arr []HeapNode
}

type MaxHeap struct {
	arr []HeapNode
}

type HeapNode struct {
	score int
}

func (d *MinHeap) MinHeapInsert(score int) {
	arr := append(d.arr, HeapNode{score})
	d.arr = arr
	d.MinHeapify(len(d.arr) - 1)
}

func (d *MinHeap) MinHeapDel() {
	last := d.arr[len(d.arr)-1]
	d.arr[0] = last
	d.arr = d.arr[:len(d.arr)-1]
	d.MinHeapify(len(d.arr) - 1)
}

func (d *MinHeap) MinHeapTop() int {
	return d.arr[0].score
}

func (d *MinHeap) MinHeapify(currentIndex int) {

	parentIndex := (currentIndex - 1) / 2

	if parentIndex < 0 {
		return
	}

	parentScore := d.arr[parentIndex].score

	if d.arr[currentIndex].score < parentScore {

		d.arr[currentIndex], d.arr[parentIndex] = d.arr[parentIndex], d.arr[currentIndex]
		d.MinHeapify(parentIndex)
	}

}

func GenMinHeap() MinHeap {
	return MinHeap{make([]HeapNode, 0)}
}

func GenMaxHeap() MaxHeap {
	return MaxHeap{make([]HeapNode, 0)}
}
