# Run the Go solutions locally

Each folder here is one NeetCode problem (all of NeetCode 150). Its `main.go` contains:

0. at the top, the problem described in plain words with example input/output, and a link to the original
1. a copy of the latest Go submission from `Data Structures & Algorithms/<problem>/`
2. the types NeetCode normally provides (`ListNode`, `TreeNode`, ...)
3. a `main()` with commented test cases. For every case it prints `PASS` / `FAIL`, the input, the result and the expected result

The synced `submission-*.go` files are not touched, because NeetCode's GitHub Sync owns them.

The folders are numbered from easy to hard, using the difficulty shown on neetcode.io:
`001`-`028` Easy, `029`-`129` Medium, `130`-`150` Hard. Inside one difficulty they follow the NeetCode 150 list order.

## Run one problem

```sh
cd local
go run ./003-two-integer-sum
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
check("my case", `twoSum([]int{1, 2, 3}, 5)`,
	twoSum([]int{1, 2, 3}, 5), []int{1, 2})
```

`check(name, input, got, want)`:

- `name`: a short description
- `input`: the call written as text between backticks, only used for printing
- `got`: the real call
- `want`: the answer you expect

Example output:

```
PASS  my case
      input:    twoSum([]int{1, 2, 3}, 5)
      result:   [1 2]
      expected: [1 2]
```

Note: if you submit a new solution on NeetCode, the copy in `local/` does not update by itself — paste the new code over the "Solution" part of `main.go`.
