# Run the Go solutions locally

Each folder here is one NeetCode problem. Its `main.go` contains:

1. a copy of the latest Go submission from `Data Structures & Algorithms/<problem>/`
2. the types NeetCode normally provides (`ListNode`, `TreeNode`, ...)
3. a `main()` with commented test cases that prints `PASS` / `FAIL`

The synced `submission-*.go` files are not touched, because NeetCode's GitHub Sync owns them.

## Run one problem

```sh
cd local
go run ./two-integer-sum
```

## Run every problem

```sh
cd local
for d in */; do echo "== $d"; go run "./$d" | tail -1; done
```

## Add your own test case

Open the problem's `main.go`, go to `func main()` and add a line:

```go
// what this case is testing
check("my case", twoSum([]int{1, 2, 3}, 5), []int{1, 2})
```

`check(name, got, want)` prints `PASS` when the two values are equal.

Note: if you submit a new solution on NeetCode, the copy in `local/` does not update by itself — paste the new code over the "Solution" part of `main.go`.
