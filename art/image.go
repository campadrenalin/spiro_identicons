package art

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"image/draw"
	"image/png"

	"crypto/md5"
	"io"
)

var encoder = png.Encoder{CompressionLevel: png.BestSpeed}

type Request struct {
	bytes         []byte
	width, height int
	numLeaves     int
}

func NewRequest(seed string) Request {
	sum := md5.Sum([]byte(seed))
	r := Request{
		bytes:  sum[:],
		width:  400,
		height: 400,
	}
	r.numLeaves = r.Consume()%5 + 3 // [3-7]
	return r
}
func (r *Request) Consume() int {
	first := r.bytes[0]
	r.bytes = r.bytes[1:]
	return int(first)
}

func (r *Request) RenderPNG(w io.Writer) {
	bg_color := color.RGBA{0, 0, 0, 0xff}
	img := image.NewRGBA(image.Rect(0, 0, r.width, r.height))
	draw.Draw(img, img.Bounds(), &image.Uniform{bg_color}, image.ZP, draw.Src)

	// Draw actual spiro
	gc := draw2dimg.NewGraphicContext(img)
	t := NewTracer(r)
	t.Draw(gc, color.White)

	encoder.Encode(w, img)
}
