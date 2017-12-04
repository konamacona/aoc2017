package main

type Side int

const (
	EAST Side = iota
	NORTH
	WEST
	SOUTH
)

type Spiral struct {
	m          map[int]map[int]int
	x          int
	y          int
	index      int
	layer      int
	side       Side
	sideIndex  int
	sideLength int
}

func NewSpiral() *Spiral {
	s := new(Spiral)
	s.m = make(map[int]map[int]int)
	s.x = 0
	s.y = 0
	s.layer = 1
	s.index = 0
	s.side = SOUTH
	s.sideIndex = 1
	s.sideLength = 1
	return s
}

func (s *Spiral) Next() {
	s.index++

	if s.sideIndex == s.sideLength {
		s._NextSide()
	}

	s.sideIndex++
	switch s.side {
	case EAST:
		s.y++
	case NORTH:
		s.x--
	case WEST:
		s.y--
	case SOUTH:
		s.x++
	}
}

func (s *Spiral) _NextSide() {
	s.sideIndex = 1
	switch s.side {
	case EAST:
		s.side = NORTH
	case NORTH:
		s.side = WEST
	case WEST:
		s.side = SOUTH
	case SOUTH:
		// Go to the next layer and the next side
		s.layer++
		s.x++
		s.y--
		s.side = EAST
		s.sideLength = (2 * s.layer) - 1
	}
}

func (s *Spiral) Get(x int, y int) int {
	return s.m[x][y]
}

func (s *Spiral) Set(x int, y int, val int) {
	yMap, hasX := s.m[x]
	if !hasX {
		s.m[x] = make(map[int]int)
		yMap = s.m[x]
	}
	yMap[y] = val
}

func (s *Spiral) SetCurrent(val int) {
	s.Set(s.x, s.y, val)
}
