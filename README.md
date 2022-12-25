# ds

This project identifies the most commonly used data structures in software engineering. It is an open source project created by [@Farischt](https://github.com/farischt)

This project is also a go package via the following command :

## **Currently not a package yet**

```bash
go get github.com/farischt/ds
```

## Packages

The project currently contains 4 different packages, one for each implemented data structure.

Here is a list of the different packages :

- `ds/queue` for queues.
- `ds/stack` for stacks.
- `ds/linked_list` for linked lists.
- `ds/tree` for trees.

`ds/tree` and `ds/linked_list` have each a separated `Node` structure.

For more information about the various methods please refere to the package documentation.

## Todo

- [ ] Add parent field to bst. (Currently added as a public field, but should be private. Done for insert and findInorderSuccessor, but not remove method yet).
  - [x] Insert
  - [x] FindInorderSuccessor
  - [ ] IsFromRight
  - [ ] RemoveNodeWithTwoChild
  - [ ] RemoveNodeWithOneChild
  - [ ] RemoveLeafNode
