package task2

import (
	"log"
	"strconv"
	"strings"
)

// task2 - Put in main and run
// contentPoints := task1.ReadFromFile(task2File2)
// contentRectangular := task1.ReadFromFile(task2File1)
// var points = task2.ParsePoints(contentPoints)
// var rect = task2.ParseNewRectangular(contentRectangular)
// for _, p := range points {
// 	if rect.CheckPoint(p) != -1 {
// 		fmt.Println(p, rect.CheckPoint(p))
// 	}
// }

const (
	// os.Args[1]
	task2File1 = "./files/task2_1.txt"
	// os.Args[2]
	task2File2 = "./files/task2_2.txt"
)

type point struct {
	X float64
	Y float64
}

type rectangular struct {
	points []point
}

func (rect rectangular) CheckPoint(p point) int {
	switch {
	case rect.isVertex(p):
		return 0
	case rect.isOnSide(p):
		return 1
	case rect.isInside(p):
		return 2
	case rect.isOutside(p):
		return 3
	}
	return -1
}

func (rect rectangular) isVertex(p point) bool {
	for _, vertex := range rect.points {
		//fmt.Printf("Comparing vertex - %v and point - %v \n", vertex, p)
		if vertex == p {
			return true
		}
	}
	return false
}

func (rect rectangular) isOnSide(p point) bool {
	condSide1 := p.Y > rect.points[0].Y && p.Y < rect.points[1].Y &&
		p.X == rect.points[0].X
	condSide2 := p.X > rect.points[1].X && p.X < rect.points[2].X &&
		p.Y == rect.points[2].Y
	condSide3 := p.Y < rect.points[2].Y && p.Y > rect.points[3].Y &&
		p.X == rect.points[2].X
	condSide4 := p.X > rect.points[0].X && p.X < rect.points[3].X &&
		p.Y == rect.points[0].Y
	if condSide1 || condSide2 || condSide3 || condSide4 {
		return true
	}
	return false
}

func (rect rectangular) isInside(p point) bool {
	condX := p.X > rect.points[0].X && p.X < rect.points[3].X
	condY := p.Y > rect.points[0].Y && p.Y < rect.points[1].Y
	if condX && condY {
		return true
	}
	return false
}

func (rect rectangular) isOutside(p point) bool {
	condX := p.X < rect.points[0].X || p.X > rect.points[2].X
	condY := p.Y < rect.points[0].Y || p.Y > rect.points[2].Y
	if condX || condY {
		return true
	}
	return false
}

func ParsePoints(input string) []point {
	var points []point
	pointsRaw := strings.Split(strings.TrimSpace(input), "\n")
	for _, pointsPair := range pointsRaw {
		sliceXY := strings.Split(pointsPair, " ")
		points = append(points, newPoint(sliceXY))
	}
	return points
}

func ParseNewRectangular(input string) rectangular {
	return rectangular{ParsePoints(input)}
}

func newPoint(sliceXY []string) point {
	x, err := strconv.ParseFloat(strings.TrimSpace(sliceXY[0]), 64)
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.ParseFloat(strings.TrimSpace(sliceXY[1]), 64)
	if err != nil {
		log.Fatal(err)
	}

	return point{x, y}
}
