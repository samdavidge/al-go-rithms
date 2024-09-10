package recursion

import (
	"reflect"
	"testing"
)

func TestItCanSolve4by4MazeWithOnePath(t *testing.T) {
	maze := []string{
		"xxex",
		"x  x",
		"x xx",
		"xsxx",
	}

	start := Point{X: 1, Y: 3}
	end := Point{X: 2, Y: 0}

	solution := []Point{
		{1, 3},
		{1, 2},
		{1, 1},
		{2, 1},
	}

	path := Solve(maze, "x", start, end)

	if !reflect.DeepEqual(solution, path) {
		t.Fatal(path)
	}
}

func TestItCanSolve4by4MazeWithMultiplePaths(t *testing.T) {
	maze := []string{
		"xxex",
		"x  x",
		"x  x",
		"xs x",
	}

	start := Point{X: 1, Y: 3}
	end := Point{X: 2, Y: 0}

	solution := []Point{
		{1, 3},
		{2, 3},
		{2, 2},
		{1, 2},
		{1, 1},
		{2, 1},
	}

	path := Solve(maze, "x", start, end)

	if !reflect.DeepEqual(solution, path) {
		t.Fatal(path)
	}
}

func TestItCanSolve10By10MazeWithSinglePath(t *testing.T) {
	maze := []string{
		"xxxxxxexxx",
		"xxxxxx xxx",
		"xxxxxx xxx",
		"xxx    xxx",
		"xxx xxxxxx",
		"xxx xxxxxx",
		"xxx   xxxx",
		"xxxxx xxxx",
		"xxxxx xxxx",
		"xxxxxsxxxx",
	}

	start := Point{X: 5, Y: 9}
	end := Point{X: 6, Y: 0}

	solution := []Point{
		{5, 9},
		{5, 8},
		{5, 7},
		{5, 6},
		{4, 6},
		{3, 6},
		{3, 5},
		{3, 4},
		{3, 3},
		{4, 3},
		{5, 3},
		{6, 3},
		{6, 2},
		{6, 1},
	}

	path := Solve(maze, "x", start, end)

	if !reflect.DeepEqual(solution, path) {
		t.Fatal(path)
	}
}
