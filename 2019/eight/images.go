package eight

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

//GetLayers return is [layer]pixel
func GetLayers(in string, w, h int) [][]string {
	layers := make([][]string, 0)

	currentH := 0
	layerIndex := 0
	layers = append(layers, make([]string, 0))

	tempStr := ""
	for i := 1; i <= len(in); i++ {

		if currentH == h {
			layers = append(layers, make([]string, 0))
			layerIndex++
			currentH = 0
		}

		if i%w != 0 {
			tempStr += string(in[i-1])
		} else {

			tempStr += string(in[i-1])
			layers[layerIndex] = append(layers[layerIndex], tempStr)
			tempStr = ""
			currentH++
		}
	}
	// layers[layerIndex] = append(layers[layerIndex], tempStr)

	return layers
}

func ImageCheckSum(in string, w, h int) int {
	layers := GetLayers(in, w, h)
	layerWithfewest0 := LayerWithMin(layers, "0")

	countOfOneInLayer := CountNumberLayer(layers[layerWithfewest0], "1")
	countOfTwoInLayer := CountNumberLayer(layers[layerWithfewest0], "2")

	return countOfOneInLayer * countOfTwoInLayer
}
func LayerWithMax(in [][]string, cut string) int {
	max := -1
	index := -1
	for i, l := range in {
		c := CountNumberLayer(l, cut)
		if c > max {
			max = c
			index = i
		}
	}

	return index
}

func RenderImage(in string, w, h int) string {
	layers := GetLayers(in, w, h)

	rendered := make([]string, h)
	for i := 0; i < h; i++ {
		str := ""
		for j := 0; j < w; j++ {
			str += "2"
		}
		rendered[i] = str
	}

	// Crap. . . Im n^3 thats not good.
	for _, l := range layers {
		for i, r := range l {
			for j, c := range r {
				cStr := string(c)
				if cStr != "2" && string(rendered[i][j]) == "2" {
					o := []rune(rendered[i])
					o[j] = c
					rendered[i] = string(o)
				}
			}
		}
	}
	WriteImage(rendered, w, h)
	return strings.Join(rendered, "")
}

func WriteImage(input []string, w, h int) {
	rect := image.Rect(0, 0, w*20, h*20)
	img := image.NewRGBA(rect)
	green := color.RGBA{14, 184, 14, 0xff}
	for r, row := range input {
		for c, p := range row {
			for y := r * 20; y < (r+1)*20; y++ {
				for x := c * 20; x < (c+1)*20; x++ {
					switch string(p) {
					case "0":
						img.Set(x, y, color.Black)
					case "1":
						img.Set(x, y, green)
					}
				}
			}
		}
	}
	f, _ := os.Create(fmt.Sprintf("img-%d-%d.png", w, h))
	png.Encode(f, img)
}
func LayerWithMin(in [][]string, cut string) int {
	min := -1
	index := -1
	for i, l := range in {
		c := CountNumberLayer(l, cut)
		if c < min || min == -1 {
			min = c
			index = i
		}
	}

	return index
}
func CountNumberLayer(layer []string, lookingFor string) int {
	count := 0

	for _, r := range layer {
		for _, c := range r {
			if string(c) == lookingFor {
				count++
			}
		}
	}

	return count
}
