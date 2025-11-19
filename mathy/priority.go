package mathy

type PriorityItem[T comparable] struct {
	Value    T
	Priority int
	index    int
}

type PriorityQueue[T comparable] []*PriorityItem[T]

func (pq PriorityQueue[T]) Len() int           { return len(pq) }
func (pq PriorityQueue[T]) Less(i, j int) bool { return pq[i].Priority < pq[j].Priority }

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*PriorityItem[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
