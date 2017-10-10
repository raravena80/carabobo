# parents

The problem. Given a tree with parent child relationships like this

Suppose we have some input data describing relationships between
parents and children over multiple generations. The data is formatted
as a list of (parent, child) pairs, where each individual is assigned
a unique integer identifier.

For example, in this diagram, 3 is a child of 1 and 2, and 5 is a child of 4:
```
1   2   4
 \ /   / \
  3   5   8
   \ / \   \
    6   7   9
```
Write a function that takes this data as input and returns
two collections:
one containing all individuals with zero known parents,
and one containing all individuals with exactly one known parent.

Sample output:
```
Zero parents: 1, 2, 4
One parent: 5, 7, 8, 9
```

Sample input:
```
parent_child_pairs = [[1, 3], [2, 3], [3, 6], [5, 6], [5, 7], [4, 5], [4, 8], [8, 9]]
```

Solution is in Go:

- parents.go - Solution in Go. Run with: go run parents.go
- parents2.go - Alternativeolution in Go. Run with: go run parents2.go
