package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	// "io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whitelndex = 0 // Первый цвет палитры
	blacklndex = 1 // Следующий цвет палитры
)

func main() {

	server := "localhost:8000"
	http.HandleFunc("/", master)
	http.HandleFunc("/lissajous", liss)
	log.Fatal(http.ListenAndServe(server, nil))

}

func liss(writer http.ResponseWriter, req *http.Request) {
	lissajous(writer)
}

func master(writer http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(writer, "Welcome, my dear golang enjoyers!\n")
	// buf, err := ioutil.ReadFile("photo_2023-06-20_20-43-14.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// writer.Header().Set("Content-Type", "image/jpg")
	// writer.Write(buf)

}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Количество полных колебаний x
		res     = 0.001 // Угловое разрешение
		size    = 500   // Канва изображения охватывает [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10мс)
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // Относительная частота колебаний у
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blacklndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
	// Примечание: игнорируем ошибки
}
