func shortestPath(n int, edges [][]int, src int) map[int]int {
    adj := map[int][][]int{}
    dist := map[int]int{}
    dist[src] = 0

    for _, edge := range edges {
        adj[edge[0]] = append(adj[edge[0]], []int{edge[1], edge[2]})
    }

    h := &MinHeap{
        Item{
            base: src,
            cost: 0,
        },
    }
    heap.Init(h)

    for h.Len() > 0 {
        current := heap.Pop(h).(Item)

        if current.cost > dist[current.base] {
            continue
        }

        for _, neighbor := range adj[current.base] {
            nextBase, nextDist := neighbor[0], neighbor[1]
            newCost := current.cost + nextDist

            existing, ok := dist[nextBase]
            if !ok || newCost < existing {
                dist[nextBase] = newCost
                heap.Push(h, Item{
                    base: nextBase,
                    cost: newCost,
                })
            }
        }
    }

    for i := 0; i < n; i++ {
        if _, ok := dist[i]; !ok {
            dist[i] = -1
        }
    }

    return dist
}

type MinHeap []Item
type Item struct {
    base int
    cost int
}

func (h MinHeap) Len() int {
    return len(h)
}
func (h MinHeap) Less(i, j int) bool {
    return h[i].cost < h[j].cost
}
func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(Item))
}
func (h *MinHeap) Pop() any {
    old := *h
    item := old[old.Len()-1]
    *h = old[:old.Len()-1]
    return item
}
