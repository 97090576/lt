package main

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}
func (h *MinHeap) Pop() int {
	old := *h
	n := len(old)
	x := old[0]
	*h = old[1:n]
	h.down(0)
	return x
}
func (h *MinHeap) up(i int) {
	for {
		parent := (i - 1) / 2
		if parent == i || !h.Less(i, parent) {
			break
		}
		h.Swap(i, parent)
		i = parent
	}
}
func (h *MinHeap) down(i int) {
	n := len(*h)
	for {
		left := 2*i + 1
		if left >= n || left < 0 { // left < 0 after int overflow
			break
		}
		child := left
		if right := left + 1; right < n && h.Less(right, left) {
			child = right
		}
		if !h.Less(child, i) {
			break
		}
		h.Swap(i, child)
		i = child
	}
}
