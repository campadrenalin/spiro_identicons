package art

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image/color"

	"math"
)

const quality = 500
const radiusOuter = 0.5 // Like most measurements, this is normalized to [-1,1] screen coords

type Tracer struct {
	width   float64
	height  float64
	scale   float64
	offsetX Animation
	offsetY Animation

	radiusInner  float64
	radiusMarker float64
}

func NewTracer(r *Request) Tracer {
	w := float64(r.width)
	h := float64(r.height)
	animate := func(min, max float64) Animation {
		return NewAnimation(r.Consume(), min, max)
	}
	return Tracer{
		width:   w,
		height:  h,
		scale:   (w + h) / 2,
		offsetX: animate(-0.05, 0.05),
		offsetY: animate(-0.05, 0.05),

		radiusInner:  radiusOuter / float64(r.numLeaves),
		radiusMarker: 0.08,
	}
}

func (t Tracer) PointFor(phi float64) Point {
	// Frame of reference: inner
	gearRatio := radiusOuter / t.radiusInner
	marker := Point{t.radiusMarker, 0}
	marker.Rotate(-gearRatio * phi)

	// Frame of reference: outer
	marker.Translate(radiusOuter-t.radiusInner, 0)
	marker.Rotate(phi)

	// Finishing touches
	marker.Translate(t.offsetX.Float64(), t.offsetY.Float64())
	marker.Scale(t.scale)
	marker.Translate(t.width/2, t.height/2)

	return marker
}
func (t Tracer) Points() (points [quality]Point) {
	for i := range points {
		points[i] = t.PointFor(float64(i) / quality * 2 * math.Pi)
	}
	return
}

func (t Tracer) Draw(gc *draw2dimg.GraphicContext, col color.Color) {
	points := t.Points()
	first := points[0]

	gc.SetStrokeColor(col)
	gc.SetLineWidth(t.scale * 0.02)

	gc.MoveTo(first.x, first.y)
	for _, p := range points[1:] {
		gc.LineTo(p.x, p.y)
	}
	gc.Close()
	gc.Stroke()
}
