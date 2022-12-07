package EduTen

import (
	"image/color"

	m "github.com/IgneousRed/gomisc"
	eb "github.com/hajimehoshi/ebiten/v2"
)

type Image = eb.Image

var emptyImg = eb.NewImage(1, 1)

func init() {
	emptyImg.Fill(color.White)
}

func DrawTriangles(scr *Image, trigs Trigs, clr Color) {
	colorR := float32(clr.R) / 255
	colorG := float32(clr.G) / 255
	colorB := float32(clr.B) / 255
	colorA := float32(clr.A) / 255
	verts := make([]eb.Vertex, len(trigs.verts))
	for i, v := range trigs.verts {
		verts[i].DstX = float32(v[0])
		verts[i].DstY = float32(windowSize[1] - v[1])
		verts[i].ColorR = colorR
		verts[i].ColorG = colorG
		verts[i].ColorB = colorB
		verts[i].ColorA = colorA
	}
	scr.DrawTriangles(verts, trigs.inds, emptyImg, &eb.DrawTrianglesOptions{})
}
func DrawRectangle(scr *Image, pos, size Vec2, clr Color) {
	DrawTriangles(scr, Trigs{[]Vec2{
		pos,
		pos.Add(Vec2{size[0], 0}),
		pos.Add(Vec2{0, size[1]}),
		pos.Add(Vec2{size[0], size[1]}),
	}, []uint16{0, 1, 2, 1, 2, 3}}, clr)
}
func DrawLine(scr *Image, a, b Vec2, thickness float64, clr Color) {
	normal := b.Sub(a).Rot90().MagSet(thickness * .5)
	DrawTriangles(scr, Trigs{[]Vec2{
		a.Sub(normal),
		a.Add(normal),
		b.Sub(normal),
		b.Add(normal),
	}, []uint16{0, 1, 2, 1, 2, 3}}, clr)
}
func DrawCircle(scr *Image, pos Vec2, size float64, points int, clr Color) {
	verts := make([]Vec2, points)
	for i := range verts {
		verts[i] = Rad(m.Tau * float64(i) / float64(points)).
			Vec2().Mul1(size).Add(pos)
	}
	inds := make([]uint16, 0, (points-2)*3)
	for i := 2; i < points; i++ {
		inds = append(inds, 0, uint16(i-1), uint16(i))
	}
	DrawTriangles(scr, Trigs{verts, inds}, clr)
}
