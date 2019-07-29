package snippers

type IntHeap []int

func (a IntHeap) Len() int           { return len(a) }
func (a IntHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntHeap) Less(i, j int) bool { return a[i] < a[j] }

func (h *IntHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func (h *IntHeap) Pop() interface{} {
	n := h.Len()
	r := (*h)[n-1]
	*h = (*h)[:n-1]
	return r
}
