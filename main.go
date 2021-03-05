package main

import (
	"fmt"
	"log"
	"os"
	"image"
	//"image/color"
	_ "image/jpeg"
)

type QuadTreer interface {
	Insert(x int, y int)
	Subdivide()
}

type QuadTree struct {
	Capacity int
	Size int
	Bounds image.Rectangle
	Points []image.Point
	Northeast *QuadTree
	Northwest *QuadTree
	Southeast *QuadTree
	Southwest *QuadTree
}


func (q QuadTree) Insert(x int, y int) QuadTree {
	if (x <= q.Bounds.Max.X && y <= q.Bounds.Max.Y) {
		fmt.Println("Can Insert!")
		fmt.Println(q.Points, q.Size, q.Capacity)
		if (q.Size < q.Capacity) {
			q.Points = append(q.Points, image.Point{x, y})
			q.Size++
		} else {
			q.Subdivide()
		}
	} else {
		fmt.Println("Can't Insert!")
	}

	return q
}

func (q QuadTree) Subdivide() {
	fmt.Println("Subdividing!")
}

func main() {
	imgFileName := "test.jpg"
	reader, err := os.Open(imgFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	width := m.Bounds().Max.X
	height := m.Bounds().Max.Y
	fmt.Println(width, height)

	q := QuadTree{4, 0, image.Rectangle{image.Pt(0,0), image.Pt(width, height)}, []image.Point{}, nil, nil, nil, nil}
	q = q.Insert(100, 150)
	q = q.Insert(500, 600)
	q = q.Insert(300, 750)
	q = q.Insert(500, 600)
	q = q.Insert(1, 1)

}

