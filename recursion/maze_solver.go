package recursion

type Point struct {
	X int
	Y int
}

func Solve(maze []string, wall string, start Point, end Point) []Point {
	var path []Point
	var seen [][]bool

	seen = makeEmptySeen(len(maze))

	walk(maze, wall, start, end, &seen, &path)

	return path
}

func walk(maze []string, wall string, current Point, end Point, seen *[][]bool, path *[]Point) bool {
	directions := directions()

	if current.X < 0 || current.X >= len(maze[0]) ||
		current.Y < 0 || current.Y >= len(maze) {
		return false
	}

	if string(maze[current.Y][current.X]) == wall {
		return false
	}

	if current.X == end.X && current.Y == end.Y {
		return true
	}

	if (*seen)[current.Y][current.X] {
		return false
	}

	(*seen)[current.Y][current.X] = true
	*path = append(*path, current)

	for i := 0; i < len(directions); i++ {
		x, y := directions[i][0], directions[i][1]
		if walk(maze, wall, Point{
			X: current.X + x,
			Y: current.Y + y,
		}, end, seen, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]

	return false
}

func directions() [][]int {
	return [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
}

func makeEmptySeen(length int) [][]bool {
	seen := make([][]bool, length)
	for i := range length {
		seen[i] = make([]bool, length)
	}
	return seen
}
