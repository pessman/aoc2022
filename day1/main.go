package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var file = flag.String("file", "day1/example.txt", "input file path from root directory")
	flag.Parse()
	input := loadData(*file)
	sum := 0
	for i := 0; i < 3; i++ {
		v, err := input.Pop()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(v)
		sum += v
	}
	log.Println(sum)
}

func loadData(filename string) *Heap[int] {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	h := NewHeap(less)
	sum := 0
	for scanner.Scan() {
		val := scanner.Text()
		if val == "" {
			h.Add(sum)
			sum = 0
			continue
		}

		i, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		sum += i
	}
	h.Add(sum)
	return h
}

type LessFn[T any] func(a, b T) bool

func less(a, b int) bool {
	return a < b
}

type Heap[T any] struct {
	heap []T
	less LessFn[T]
}

func NewHeap[T any](fn LessFn[T]) *Heap[T] {
	return &Heap[T]{
		heap: make([]T, 0),
		less: fn,
	}
}

func (h *Heap[T]) Add(val T) {
	h.heap = append(h.heap, val)
	h.heapifyUp(h.size() - 1)
}

func (h *Heap[T]) Pop() (T, error) {
	var ret T
	if h.size() == 0 {
		return ret, fmt.Errorf("heap is empty, nothing to pop")
	}

	ret = h.heap[0]
	tail := h.size() - 1
	h.heap[0] = h.heap[tail]
	h.heap = h.heap[:tail]
	h.heapifyDown(0)
	return ret, nil
}

func (h *Heap[T]) heapifyDown(parentIndex int) {
	lc, rc := leftChild(parentIndex), rightChild(parentIndex)
	size := h.size()
	if lc >= size {
		return // reached end of heap
	}

	child := lc
	// right child in bounds and rc is the bigger of the two
	if rc < size && h.less(h.heap[lc], h.heap[rc]) {
		child = rc
	}

	if !h.less(h.heap[parentIndex], h.heap[child]) {
		return // no longer need to check lower children as parent is larger than biggest child
	}

	h.heap[parentIndex], h.heap[child] = h.heap[child], h.heap[parentIndex]
	h.heapifyDown(child)
}

func (h *Heap[T]) heapifyUp(tailIndex int) {
	parentIndex := getParent(tailIndex)
	if h.less(h.heap[parentIndex], h.heap[tailIndex]) {
		h.heap[tailIndex], h.heap[parentIndex] = h.heap[parentIndex], h.heap[tailIndex]
		h.heapifyUp(parentIndex)
	}
}

func (h *Heap[T]) size() int {
	return len(h.heap)
}

// getParent returns parent index postion for input child
func getParent(i int) int {
	return (i - 1) / 2
}

// leftChild returns the left child node index for input parent
func leftChild(i int) int {
	return (i * 2) + 1
}

// rightChild returns the right child node index for input parent
func rightChild(i int) int {
	return (i * 2) + 2
}
