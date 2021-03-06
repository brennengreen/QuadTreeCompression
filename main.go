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
	if ((x <= q.Bounds.Max.X && y <= q.Bounds.Max.Y) && (x >= q.Bounds.Min.X && y >= q.Bounds.Min.Y)) {
		if (q.Size < q.Capacity && q.Northwest == nil) {
			q.Points = append(q.Points, image.Point{x, y})
			q.Size++
		} else {
			q = q.Subdivide()
			q.Northeast = q.Northeast.Insert(x, y)
			q.Northwest = q.Northwest.Insert(x, y)
			q.Southeast = q.Southeast.Insert(x, y)
			q.Southwest = q.Southwest.Insert(x, y)
		}
	}
	return q
}

func (q QuadTree) Subdivide() QuadTree {
	q.Northwest = QuadTree{4, 0, image.Rectangle{image.Pt(0,0), image.Pt(q.Bounds.Max.X/2, q.Bounds.Max.Y/2)}, []image.Point{}, nil, nil, nil, nil}
	q.Northeast = QuadTree{4, 0, image.Rectangle{image.Pt(q.Bounds.Max.X/2,0), image.Pt(q.Bounds.Max.X, q.Bounds.Max.Y/2)}, []image.Point{}, nil, nil, nil, nil}
	q.Southwest = QuadTree{4, 0, image.Rectangle{image.Pt(0,q.Bounds.Max.Y/2), image.Pt(q.Bounds.Max.X/2, q.Bounds.Max.Y)}, []image.Point{}, nil, nil, nil, nil}
	q.Southeast = QuadTree{4, 0, image.Rectangle{image.Pt(q.Bounds.Max.X/2, q.Bounds.Max.Y/2), image.Pt(q.Bounds.Max.X, q.Bounds.Max.Y)}, []image.Point{}, nil, nil, nil, nil}

	return q
}

func main() {
	imgFileName := "small.jpg"
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
	
	for x_i := 0; x_i < width; x_i++ {
		for y_i := 0; y_i < height; y_i++ {
			q = q.Insert(x_i, y_i)
		}
	}


}

