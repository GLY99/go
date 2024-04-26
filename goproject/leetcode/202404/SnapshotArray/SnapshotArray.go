package main

type SnapshotArray struct {
	idx int
	arr [][][2]int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{idx: 0, arr: make([][][2]int, length)}
}

func (snapArr *SnapshotArray) Set(index int, val int) {
	snapArr.arr[index] = append(snapArr.arr[index], [2]int{snapArr.idx, val})
}

func (snapArr *SnapshotArray) Snap() int {
	snapArr.idx++
	return snapArr.idx - 1
}

func (snapArr *SnapshotArray) Get(index int, snap_id int) int {
	length := len(snapArr.arr[index])
	left, right := 0, length
	for left < right {
		mid := (left + right) >> 1
		if snapArr.arr[index][mid][0] <= snap_id {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left <= 0 {
		return 0
	}
	return snapArr.arr[index][left-1][1]
}
