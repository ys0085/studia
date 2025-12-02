package solve

/*
Priority queue implementation on a Binary Tree (smallest priority first)
Priority = state.Cost + state.Heuristic
*/

type PriorityQueue []StateType

func (pq *PriorityQueue) enqueue(state StateType) {
	*pq = append(*pq, state)
	pq.upHeap(len(*pq) - 1)
}

func (pq *PriorityQueue) dequeue() StateType {
	if len(*pq) == 0 {
		return StateType{}
	}
	root := (*pq)[0]
	last := len(*pq) - 1
	(*pq)[0] = (*pq)[last]
	*pq = (*pq)[:last]
	pq.downHeap(0)
	return root
}

func (pq *PriorityQueue) upHeap(index int) {
	parent := (index - 1) / 2
	for index > 0 && compare((*pq)[parent], (*pq)[index]) {
		(*pq)[index], (*pq)[parent] = (*pq)[parent], (*pq)[index]
		index = parent
		parent = (index - 1) / 2
	}
}

func (pq *PriorityQueue) downHeap(index int) {
	size := len(*pq)
	smallest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < size && compare((*pq)[smallest], (*pq)[left]) {
		smallest = left
	}
	if right < size && compare((*pq)[smallest], (*pq)[right]) {
		smallest = right
	}
	if smallest != index {
		(*pq)[index], (*pq)[smallest] = (*pq)[smallest], (*pq)[index]
		pq.downHeap(smallest)
	}
}

func compare(a, b StateType) bool {
	//compares states. Returns true if a > b, false otherwise.
	return a.Heuristic+a.Cost > b.Heuristic+b.Cost
}
