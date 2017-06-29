// Copyright 2017 Martin Planer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

var (
	black = color.RGBA{0, 0, 0, 255}
	red   = color.RGBA{255, 0, 0, 255}
	green = color.RGBA{0, 255, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
)

const imageSize = 500
const tileSize = imageSize / boardSize

var col color.Color

func exportBoard(b Board, filename string) {
	img := drawBoard(b)

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func drawBoard(b Board) image.Image {
	var img = image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))

	col = red
	drawGrid(img)

	for _, t := range b {
		drawTile(img, t)
	}

	return img
}

func drawGrid(img *image.RGBA) {
	for l := 0; l <= 10; l++ {
		ll := max(0, min(imageSize-1, l*tileSize))
		for i := 0; i < imageSize; i++ {
			img.SetRGBA(ll, i, red)
			img.SetRGBA(i, ll, red)
		}
	}
}

func drawTile(img *image.RGBA, t Tile) {
	x1 := t.x * tileSize
	x2 := x1 + t.w*tileSize
	y1 := t.y * tileSize
	y2 := y1 + t.h*tileSize

	col = blue
	drawRectFill(img, x1, y1, x2, y2)
	col = black
	drawRectOutline(img, x1, y1, x2, y2)
}

func drawRectOutline(img *image.RGBA, x1, y1, x2, y2 int) {
	drawHLine(img, y1, x1, x2)
	drawHLine(img, y2, x1, x2)
	drawVLine(img, x1, y1, y2)
	drawVLine(img, x2, y1, y2)
}

func drawRectFill(img *image.RGBA, x1, y1, x2, y2 int) {
	for y := y1; y <= y2; y++ {
		drawHLine(img, y, x1, x2)
	}
}

func drawHLine(img *image.RGBA, y, x1, x2 int) {
	for x := x1; x <= x2; x++ {
		img.Set(x, y, col)
	}
}

func drawVLine(img *image.RGBA, x, y1, y2 int) {
	for y := y1; y <= y2; y++ {
		img.Set(x, y, col)
	}
}
