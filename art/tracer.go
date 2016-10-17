package art

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image/color"

	"math"
)

const quality = 500
const radiusOuter = float64(2 * 3 * 5 * 7) // Divides evenly for all the primes/leaf counts we care about

type Tracer struct {
	centerX float64
	centerY float64

	radiusInner  float64
	radiusMarker float64
}

func NewTracer(r *Request) Tracer {
	return Tracer{
		centerX: float64(r.width) / 2,
		centerY: float64(r.height) / 2,

		radiusInner:  radiusOuter / float64(r.numLeaves),
		radiusMarker: 100,
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
	marker.Translate(t.centerX, t.centerY)

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
	gc.SetLineWidth(5)

	gc.MoveTo(first.x, first.y)
	for _, p := range points[1:] {
		gc.LineTo(p.x, p.y)
	}
	gc.Close()
	gc.Stroke()
}
