package main

var dirs = [2][4][2]int{
	{{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
	{{-1, 1}, {1, 1}, {-1, -1}, {1, -1}},
}

type NeighborSum struct {
	gird [][]int
	pos  map[int][2]int
}

func Constructor(grid [][]int) NeighborSum {
	this := NeighborSum{
		gird: grid,
		pos:  make(map[int][2]int),
	}
	for i := range grid {
		for j := range grid[i] {
			this.pos[grid[i][j]] = [2]int{i, j}
		}
	}
	return this
}

func (this *NeighborSum) AdjacentSum(value int) int {
	return this.GetSum(value, 0)
}

func (this *NeighborSum) DiagonalSum(value int) int {
	return this.GetSum(value, 1)
}

func (this *NeighborSum) GetSum(value, action int) int {
	sum := 0
	pos := this.pos[value]
	x, y := pos[0], pos[1]
	for _, dir := range dirs[action] {
		nx, ny := dir[0]+x, dir[1]+y
		if nx >= 0 && nx < len(this.gird) && ny >= 0 && ny < len(this.gird[0]) {
			sum += this.gird[nx][ny]
		}
	}
	return sum
}

/**
 * Your NeighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */
