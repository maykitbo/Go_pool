package logorand

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

type Logo struct {
	text, file_name string
	x, y, count     int
	img             *image.RGBA
}

func Init(t, fname string, xx, yy int) Logo {
	rand.Seed(time.Now().UnixNano())
	return Logo{
		text:      t,
		file_name: fname,
		x:         xx,
		y:         yy,
		count:     1,
	}
}

func InitClient(fname string, xx, yy int) {
	rand.Seed(time.Now().UnixNano())
	logo := Logo{
		text:      RandomString(true),
		file_name: fname,
		x:         xx,
		y:         yy,
		count:     1,
	}
	for {
		logo.Create()
		logo.Save()
		time.Sleep(time.Second)
	}
}

func randomColor() color.RGBA {
	return color.RGBA{uint8(rand.Intn(254) + 1), uint8(rand.Intn(254) + 1), uint8(rand.Intn(254) + 1), 255}
}

type layerRandomizer struct {
	rand_float     float64
	palette        []color.RGBA
	x, y, k, trans int
	img            *image.RGBA
}

func (lr layerRandomizer) colorMix(x, y, g int) {
	lr.img.SetRGBA(x, y, color.RGBA{
		uint8((int(lr.palette[g].R)*lr.trans + int(lr.img.RGBAAt(x, y).R)*(255-lr.trans)) / 255),
		uint8((int(lr.palette[g].G)*lr.trans + int(lr.img.RGBAAt(x, y).G)*(255-lr.trans)) / 255),
		uint8((int(lr.palette[g].B)*lr.trans + int(lr.img.RGBAAt(x, y).B)*(255-lr.trans)) / 255),
		255,
	})
}

func (lr *layerRandomizer) fullLoop(f func(int, int, int) bool) {
	lr.paletteR()
	for y := 0; y < lr.y; y++ {
		for x := 0; x < lr.x; x++ {
			for g := 0; g < lr.k; g++ {
				if f(x, y, g) {
					lr.colorMix(x, y, g)
				}
			}
		}
	}
}

func (lr layerRandomizer) layerType1() { // noise
	lr.fullLoop(func(x, y, g int) bool {
		return (x*(y+1))%(g+1) == 1
	})
}

func (lr layerRandomizer) layerType2() { // horizontal waves
	lr.fullLoop(func(x, y, g int) bool {
		waveWidth := float64(lr.x) / float64(lr.k)
		waveHeight := (lr.rand_float + 1.0) * 7.0
		wave := math.Sin(float64(x)*2*math.Pi/waveWidth+float64(g)*2*math.Pi/float64(lr.k)) * waveHeight
		return wave+waveWidth*float64(g+1) > float64(y) && wave+waveWidth*float64(g) < float64(y)
	})
}

func (lr layerRandomizer) layerType3() { // vertical waves
	lr.fullLoop(func(x, y, g int) bool {
		waveWidth := float64(lr.x) / float64(lr.k)
		waveHeight := (rand.Float64() + 0.1) * 15.0
		wave := math.Sin(float64(y)*2*math.Pi/waveWidth+float64(g)*2*math.Pi/float64(lr.k)) * waveHeight
		return wave+waveWidth*float64(g+1) > float64(x) && wave+waveWidth*float64(g) < float64(x)
	})
}

func (lr layerRandomizer) layerType4() { // circles from the center
	lr.fullLoop(func(x, y, g int) bool {
		lr.rand_float = float64(lr.x/2) * math.Sqrt2 / float64(lr.k)
		distanceFromCenter := math.Sqrt(math.Pow(float64(x-lr.y/2), 2) + math.Pow(float64(y-lr.x/2), 2))
		return distanceFromCenter < float64(g+1)*lr.rand_float && distanceFromCenter > float64(g)*lr.rand_float
	})
}

func (lr layerRandomizer) layerType5() { // beam segments
	lr.fullLoop(func(x, y, g int) bool {
		angle := math.Atan2(float64(y-lr.y/2), float64(x-lr.x/2)) + math.Pi
		return int(angle*float64(lr.k)/(2*math.Pi)) == g
	})
}

func (lr layerRandomizer) layerType6() { // circles with random center, R, count
	lr.paletteR()
	for g := 0; g < lr.k; g++ {
		center_x, center_y, R := rand.Intn(lr.x), rand.Intn(lr.x), rand.Intn(lr.x/2)+lr.x/11
		for y := center_y - R; y < center_y+R; y++ {
			chord := math.Sqrt(math.Pow(float64(R), 2.0)-math.Pow(float64(center_y-y), 2.0)) * (lr.rand_float + 1)
			for x := center_x - int(chord); x < center_x+int(chord); x++ {
				lr.colorMix(x, y, g)
			}
		}
	}
}

func (lr layerRandomizer) layerType7() { // squares
	lr.paletteR()
	for g := 0; g < lr.k; g++ {
		center_x, center_y, R := rand.Intn(lr.x), rand.Intn(lr.x), rand.Intn(lr.x/2)+lr.x/11
		for y := center_y - R; y < center_y+R; y++ {
			for x := center_x - R; x < center_x+R; x++ {
				lr.colorMix(x, y, g)
			}
		}
	}
}

func (lr layerRandomizer) colorgGlare(x, y, k int) {
	lr.img.SetRGBA(x, y, color.RGBA{
		uint8((int(255-lr.img.RGBAAt(x, y).R)*(255-k) + int(lr.img.RGBAAt(x, y).R)*k) / 255),
		uint8((int(255-lr.img.RGBAAt(x, y).G)*(255-k) + int(lr.img.RGBAAt(x, y).G)*k) / 255),
		uint8((int(255-lr.img.RGBAAt(x, y).B)*(255-k) + int(lr.img.RGBAAt(x, y).B)*k) / 255),
		255,
	})
}

func (lr layerRandomizer) layerType8() { // glare
	count := rand.Intn(7) + 2
	for i := 0; i < count; i++ {
		center_x := rand.Intn(lr.x-6) + 3
		center_y := rand.Intn(lr.y-6) + 3
		l := rand.Intn(lr.x/25) + 7
		for j := 1; j <= l; j++ {
			lr.colorgGlare(center_x+j, center_y+j, j*255/l)
			lr.colorgGlare(center_x-j, center_y+j, j*255/l)
			lr.colorgGlare(center_x+j, center_y-j, j*255/l)
			lr.colorgGlare(center_x-j, center_y-j, j*255/l)
		}
	}
}

func (lr layerRandomizer) layerType9() { // circles from random place
	a, b := rand.Intn(300), rand.Intn(300)
	lr.fullLoop(func(x, y, g int) bool {
		lr.rand_float = float64(lr.x/2) * math.Sqrt2 / float64(lr.k)
		distanceFromCenter := math.Sqrt(math.Pow(float64(x-a), 2) + math.Pow(float64(y-b), 2))
		return distanceFromCenter < float64(g+1)*lr.rand_float && distanceFromCenter > float64(g)*lr.rand_float
	})
}

func (lr *layerRandomizer) easyLoop(f func(int, int)) {
	lr.paletteR()
	for y := 0; y < lr.y; y++ {
		for x := 0; x < lr.x; x++ {
			f(x, y)
		}
	}
}

func (lr layerRandomizer) filter1() {
	lr.easyLoop(func(x, y int) {
		lr.img.SetRGBA(x, y, color.RGBA{
			uint8(float64(lr.img.RGBAAt(x, y).R) * math.Pow(math.Sin(float64(x)/100), 2)),
			uint8(float64(lr.img.RGBAAt(x, y).G) * math.Pow(math.Sin(float64(y)/100), 2)),
			uint8(float64(lr.img.RGBAAt(x, y).B) * math.Pow(math.Sin(float64(x+y)/100), 2)),
			255,
		})
	})
}

func (lr layerRandomizer) filter2() {
	lr.easyLoop(func(x, y int) {
		lr.img.SetRGBA(x, y, color.RGBA{
			uint8(float64(lr.img.RGBAAt(x, y).R) * math.Pow(math.Cos(float64(y)/10), 2)),
			uint8(float64(lr.img.RGBAAt(x, y).G) * math.Pow(math.Cos(float64(x)/10), 2)),
			uint8(float64(lr.img.RGBAAt(x, y).B) * math.Pow(math.Cos(float64(x+y)/10), 2)),
			255,
		})
	})
}

func (lr layerRandomizer) filter3() {
	lr.easyLoop(func(x, y int) {
		lr.img.SetRGBA(x, y, color.RGBA{
			uint8(math.Sin(float64(lr.img.RGBAAt(x, y).R))*128 + 128),
			uint8(math.Sin(float64(lr.img.RGBAAt(x, y).G))*128 + 128),
			uint8(math.Sin(float64(lr.img.RGBAAt(x, y).B))*128 + 128),
			255,
		})
	})
}

func (lr layerRandomizer) filter4() {
	lr.easyLoop(func(x, y int) {
		lr.img.SetRGBA(x, y, color.RGBA{
			uint8(math.Pow(math.Cos(float64(lr.img.RGBAAt(x, y).R)), 2) * 128),
			uint8(math.Pow(math.Cos(float64(lr.img.RGBAAt(x, y).G)), 2) * 128),
			uint8(math.Pow(math.Cos(float64(lr.img.RGBAAt(x, y).B)), 2) * 128),
			255,
		})
	})
}

func (lr layerRandomizer) colorVignette(x, y int, k float64) {
	lr.img.SetRGBA(x, y, color.RGBA{
		uint8(float64(lr.img.RGBAAt(x, y).R) * k),
		uint8(float64(lr.img.RGBAAt(x, y).G) * k),
		uint8(float64(lr.img.RGBAAt(x, y).B) * k),
		255,
	})
}

func (lr layerRandomizer) chamfer1() { // vignette
	center_x, center_y := lr.x/2, lr.y/2
	max_radius := (rand.Float64() + 1.0) * math.Min(float64(center_x), float64(center_y))
	for y := 0; y < lr.y; y++ {
		for x := 0; x < lr.x; x++ {
			radius := math.Sqrt(math.Pow(float64(center_y-y), 2) + math.Pow(float64(center_x-x), 2))
			lr.colorVignette(x, y, 1-math.Min(1, radius/max_radius))
		}
	}
}

func (lr *layerRandomizer) paletteR() { // random set of colors
	lr.trans = rand.Intn(150) + 50
	lr.rand_float = rand.Float64()
	lr.k = rand.Intn(7) + 2
	lr.palette = make([]color.RGBA, lr.k)
	for g := 0; g < lr.k; g++ {
		lr.palette[g] = randomColor()
	}
}

func (lr layerRandomizer) oneLayer() {
	if rand.Float64() > 0.65 {
		lr.layerType1()
	}
	if rand.Float64() > 0.65 {
		lr.layerType2()
	}
	if rand.Float64() > 0.65 {
		lr.layerType3()
	}
	if rand.Float64() > 0.65 {
		lr.layerType4()
	}
	if rand.Float64() > 0.65 {
		lr.layerType5()
	}
	if rand.Float64() > 0.65 {
		lr.layerType6()
	}
	if rand.Float64() > 0.65 {
		lr.layerType7()
	}
	if rand.Float64() > 0.65 {
		lr.layerType8()
	}
	if rand.Float64() > 0.75 {
		lr.layerType9()
	}
	if rand.Float64() > 0.75 {
		lr.layerType9()
	}
	if rand.Float64() > 0.85 {
		lr.filter1()
	}
	if rand.Float64() > 0.85 {
		lr.filter2()
	}
	if rand.Float64() > 0.85 {
		lr.filter3()
	}
	if rand.Float64() > 0.85 {
		lr.filter4()
	}
	if rand.Float64() > 0.55 {
		lr.chamfer1()
	}
}

func initLayerRandomizer(x, y int, i *image.RGBA) (lr layerRandomizer) {
	lr.x, lr.y = x, y
	lr.img = i
	return
}

func (l *Logo) Create() {
	if l.count%(rand.Intn(6)+1) == 0 {
		l.text = RandomString(true)
	}
	l.img = image.NewRGBA(image.Rect(0, 0, l.x, l.y))
	draw.Draw(l.img, image.Rect(0, 0, l.x, l.y), &image.Uniform{randomColor()}, image.ZP, draw.Src)
	lr := initLayerRandomizer(l.x, l.y, l.img)
	lr.oneLayer()
	l.addText()
	l.count++
}

func (l *Logo) addText() {
	fontBytes, err := ioutil.ReadFile("docs/ttfs/" + ttfName())
	if err != nil {
		fmt.Println(err)
		// return
	}
	font, err := truetype.Parse(fontBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	extra := image.NewRGBA(image.Rect(0, 0, l.x, l.y))
	draw.Draw(extra, image.Rect(0, 0, l.x, l.y), &image.Uniform{color.Black}, image.ZP, draw.Src)
	scale := (rand.Float64() + 1.5) * 20
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(scale)
	c.SetClip(extra.Bounds())
	c.SetDst(extra)
	c.SetSrc(image.White)
	dx := (float64(l.x) - float64(len(l.text))*scale*0.6) / 2
	pt := freetype.Pt(int(dx), int(float64(l.y)-scale*1.5)/2)
	_, err = c.DrawString(l.text, pt)
	if err != nil {
		fmt.Println(err)
		return
	}
	for x := 0; x < l.x; x++ {
		for y := 0; y < l.y; y++ {
			if extra.RGBAAt(x, y).B == 255 && extra.RGBAAt(x, y).R == 255 && extra.RGBAAt(x, y).G == 255 {
				l.img.SetRGBA(x, y, color.RGBA{
					uint8(math.Pow(math.Cos(float64(l.img.RGBAAt(x, y).R)), 2) * float64(l.img.RGBAAt(x, y).G)),
					uint8(math.Pow(math.Cos(float64(l.img.RGBAAt(x, y).G)), 2) * float64(l.img.RGBAAt(x, y).B)),
					uint8(math.Pow(math.Cos(float64(l.img.RGBAAt(x, y).B)), 2) * float64(l.img.RGBAAt(x, y).R)),
					255,
				})
			}
		}
	}
}

func ttfName() string {
	switch rand.Intn(11) + 2 {
	case 2:
		return "Go-Bold.ttf"
	case 3:
		return "Go-Italic.ttf"
	case 4:
		return "Go-Medium-Italic.ttf"
	case 5:
		return "Go-Medium.ttf"
	case 6:
		return "Go-Mono-Bold-Italic.ttf"
	case 7:
		return "Go-Mono-Bold.ttf"
	case 8:
		return "Go-Mono-Italic.ttf"
	case 9:
		return "Go-Mono.ttf"
	case 10:
		return "Go-Regular.ttf"
	case 11:
		return "Go-Smallcaps-Italic.ttf"
	case 12:
		return "Go-Smallcaps.ttf"
	}
	return "Go-Bold-Italic.ttf"
}

func (l Logo) Save() {
	f, er := os.Create("buff.png")
	if er != nil {
		fmt.Println(er)
		return
	}
	defer f.Close()
	er = png.Encode(f, l.img)
	if er != nil {
		fmt.Println(er)
		return
	}
	er = os.Rename("buff.png", l.file_name)
	if er != nil {
		fmt.Println(er)
	}
}

func RandomString(le bool) string {
	const consonants = "bcdfghjklmnpqrstvwxz"
	const vowels = "aeiouy"
	l := 0
	if le == true {
		l = 2 + rand.Intn(10)
	} else {
		l = 200 + rand.Intn(300)
	}

	f := rand.Intn(2)
	str := make([]byte, l)
	for k := 0; k < l; k++ {
		if k%(rand.Intn(7)+2) == 0 && k != l-2 {
			str[k] = ' '
		} else if (f+k)%2 == 1 {
			str[k] = vowels[rand.Intn(len(vowels))] - 'a' + 'A'
		} else {
			str[k] = consonants[rand.Intn(len(consonants))] - 'a' + 'A'
		}
	}
	return string(str)
}
