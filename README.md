# ds

This project identifies the most commonly used data structures in software engineering. It is an open source project created by [@Farischt](https://github.com/farischt)

This project is also a go package via the following command :

```bash
go get github.com/farischt/ds
```

## Packages

The project currently contains 6 different packages, one for each implemented data structure.

Here is a list of the different packages :

- `ds/queue` for queues.
- `ds/stack` for stacks.
- `ds/linked_list` for linked lists.
- `ds/tree` for trees.
- `ds/graph` for binary graphs.
- `ds/heap` for heaps.

`ds/tree` and `ds/linked_list` have each a separated `Node` structure.

## How to use the different data structures

### Queue

In order to create a queue:

```go
    import (
        "github.com/farischt/ds/queue"
    )

    func main() {
        capacity := 1
        q := queue.New[int](capacity)
        // You can now use any method implementend in IQueue as seen in docs.
    }
```

### Stack

In order to create a stack:

```go
    import (
        "github.com/farischt/ds/stack"
    )

    func main() {
        capacity := 1
        s := stack.New[int](capacity)
        // You can now use any method implementend in IStack as seen in docs.
    }
```

### Linked list

In order to create a linked list:

```go
    import (
        "github.com/farischt/ds/linked_list"
    )

    func main() {
        l := ll.New[int]()
        // You can now use any method implementend in ILinkedList as seen in docs.
    }
```

### Binary search tree

In order to create a binary search tree:

```go
    import (
        "github.com/farischt/ds/tree"
    )

    func main() {
        rootData := 10
        root := tree.NewNode(rootData)
        l := tree.New(root)
        // You can now use any method implementend in IBinarySearchTree as seen in docs.
    }
```

### Graph

In order to create a graph:

```go
    import (
        "github.com/farischt/ds/graph"
    )

    func main() {
        g := graph.New[int]()
        // You can now use any method implementend in IGraph as seen in docs.

        // Before adding any edge, make sure to create nodes.
        src := 10
        dst := 20
        g.Add(src)
        g.Add(dst)
        g.AddUndirectedEdge(src, dst)
    }
```

### Heap

In order to create a heap:

```go
    import (
        "github.com/farischt/ds/heap"
    )

    func main() {
        minHeap := heap.New[int](heap.MinHeap)
        maxHeap := heap.New[int](heap.MaxHeap)
        // You can now use any method implementend in IGraph as seen in docs.

        // In some case a heap could only contain a native number type (int, uint...).
        data := 10 // This is the value used to compute the heap.
        minHeap.Push(data, nil)

        // On another hand, a heap could be use to store more than just a number. For example, in the diksjtra algorithm, we use a heap to store the value of a node and the current distance.
        node := 10 // This is the value used to compute the heap.
        distance := 30 // It could be any data type.
        minHeap.Push(node, distance)


        // If you pop the top element it will return an pointer of heap.Item.
        item, _ := minHeap.Pop()

        // This item has to fields Value (the node value) and Information (in the previous case, it will be the distance 30).
        // Refer to the heap.Item docs for more informations.
    }
```

For more information about the various methods please refere to the [package documentation](https://pkg.go.dev/github.com/farischt/ds).
