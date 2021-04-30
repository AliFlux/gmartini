package gmartini

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"math"
	"os"
	"testing"

	gmu "github.com/engelsjk/gomathutils"
	"github.com/fogleman/gg"
)

func drawVertices(dc *gg.Context, mesh *Mesh) {
	dc.SetRGB(0, 0, 0)
	for i := 0; i < (len(mesh.Vertices) - 2); i += 2 {
		dc.DrawCircle(float64(mesh.Vertices[i]), float64(mesh.Vertices[i+1]), 0.5)
		dc.Fill()
	}
}

func drawTriangles(dc *gg.Context, mesh *Mesh) {
	dc.ClearPath()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(0.5)
	for i := 0; i < (len(mesh.Triangles) - 3); i += 3 {
		a, b, c := mesh.Triangles[i], mesh.Triangles[i+1], mesh.Triangles[i+2]
		ax, ay := float64(mesh.Vertices[2*a]), float64(mesh.Vertices[2*a+1])
		bx, by := float64(mesh.Vertices[2*b]), float64(mesh.Vertices[2*b+1])
		cx, cy := float64(mesh.Vertices[2*c]), float64(mesh.Vertices[2*c+1])
		dc.MoveTo(ax, ay)
		dc.LineTo(bx, by)
		dc.LineTo(cx, cy)
		dc.LineTo(ax, ay)
	}
	dc.Stroke()
}

func load(terrainFile, encoding string, gridSize int32, maxError float32) ([]float32, *Mesh) {

	file, err := os.Open(terrainFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	terrain, err := DecodeElevation(img, encoding, true)
	if err != nil {
		panic(err)
	}

	martini, err := New(OptionGridSize(gridSize))
	if err != nil {
		panic(err)
	}

	tile, err := martini.CreateTile(terrain)
	if err != nil {
		panic(err)
	}

	mesh := tile.GetMesh(OptionMaxError(maxError))
	return terrain, mesh
}

func TestDrawVerticesErr5(t *testing.T) {

	var terrainFile string = "data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513
	var maxError float32 = 5.0
	var imageFile string = "test/vertices-%d.png"

	_, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawVertices(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}

func TestDrawVerticesErr50(t *testing.T) {

	var terrainFile string = "data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513
	var maxError float32 = 50.0
	var imageFile string = "test/vertices-%d.png"

	_, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawVertices(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}

func TestDrawVerticesErr500(t *testing.T) {

	var terrainFile string = "data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513
	var maxError float32 = 500.0
	var imageFile string = "test/vertices-%d.png"

	_, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawVertices(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}

func TestDrawTrianglesErr5(t *testing.T) {

	var terrainFile string = "data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513
	var maxError float32 = 5.0
	var imageFile string = "test/triangles-%d.png"

	_, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawTriangles(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}

func TestDrawTrianglesErr50(t *testing.T) {

	var terrainFile string = "data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513
	var maxError float32 = 50.0
	var imageFile string = "test/triangles-%d.png"

	_, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawTriangles(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}

func TestDrawTrianglesErr500(t *testing.T) {

	var terrainFile string = "data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513
	var maxError float32 = 500.0
	var imageFile string = "test/triangles-%d.png"

	_, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawTriangles(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}

func TestDrawAll(t *testing.T) {

	var terrainFile string = "data/fuji.png"
	var encoding string = "mapbox"
	var gridSize int32 = 513
	var maxError float32 = 50.0
	var imageFile string = "test/martini-%d.png"

	terrain, mesh := load(terrainFile, encoding, gridSize, maxError)

	dc := gg.NewContext(512, 512)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawVertices(dc, mesh)
	drawTriangles(dc, mesh)
	dc.SavePNG(fmt.Sprintf(imageFile, int(maxError)))
	t.Logf("test image saved at %s", fmt.Sprintf(imageFile, int(maxError)))
}
