package game

import (
	"fmt"
	"image"
	"log"
	"math"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/kyeett/gomponents/direction"

	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"

	ase "github.com/kyeett/GoAseprite"
	tiled "github.com/lafriks/go-tiled"

	"github.com/kyeett/dungeon/assets"
	"github.com/kyeett/dungeon/resolvutil"
	"github.com/kyeett/gomponents/components"
)

func (g *Game) initMap() {

	for _, l := range g.m.Layers {
		for i, t := range l.Tiles {
			if t.IsNil() {
				continue
			}

			sx, sy := i%g.m.Width, i/g.m.Width

			x, y := TilesheetCoords(g.m.Tilesets[0], t.ID)
			srcRect := image.Rect(0, 0, g.m.TileWidth, g.m.TileHeight).Add(image.Pt(x, y))

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(-float64(g.m.TileWidth/2), -float64(g.m.TileHeight/2))
			if t.VerticalFlip {
				op.GeoM.Scale(1, -1)
			}
			if t.HorizontalFlip {
				op.GeoM.Scale(-1, 1)
			}
			if t.DiagonalFlip {
				op.GeoM.Rotate(3 * math.Pi / 2)
				op.GeoM.Scale(1, -1)
			}
			op.GeoM.Translate(float64(g.m.TileWidth/2), float64(g.m.TileHeight/2))
			op.GeoM.Translate(float64(sx*g.m.TileWidth), float64(sy*g.m.TileHeight))

			g.backgroundImg.DrawImage(g.spriteImg.SubImage(srcRect).(*ebiten.Image), op)
		}
	}

	g.staticSpace = resolv.Space{}
	for _, og := range g.m.ObjectGroups {
		for _, o := range og.Objects {
			g.parseObject(o)
		}
	}

}

func TilesheetCoords(t *tiled.Tileset, ID uint32) (int, int) {
	y := (ID) / uint32(t.Columns)
	x := (ID) % uint32(t.Columns)
	return int(x) * t.TileWidth, int(y) * t.TileHeight
}

func LayerTileID(l *tiled.Layer, width, x, y int) uint32 {
	return l.Tiles[width*y+x].ID
}

func New(options ...Option) (*Game, error) {

	for _, o := range options {
		switch o {
		case OptionFromDisk:
			assets.ReadFromDisk = true
		}
	}

	filename := "assets/tilesets/world-1.tmx"
	dir := filepath.Dir(filename)

	// m, err := tiled.LoadFromFile(filename)
	m, err := tiled.LoadFromReader(dir, assets.FileReaderFatal(filename))
	if err != nil {
		return &Game{}, err
	}

	// Load sprite image
	tileset := m.Tilesets[0]
	tilesetPath := tileset.Source
	path := dir + "/" + tileset.Image.Source
	// gfx.DecodePNG()
	// tilesetImg, err := gfx.OpenPNG(path)
	tilesetImg, err := gfx.DecodePNG(assets.FileReaderFatal(path))
	if err != nil {
		log.Fatal(tilesetPath, err)
	}
	spriteImg, _ := ebiten.NewImageFromImage(tilesetImg, ebiten.FilterDefault)

	bgImg, _ := ebiten.NewImage(m.Width*m.TileWidth, m.Height*m.TileHeight, ebiten.FilterDefault)
	g := Game{
		baseDir:       dir,
		spriteImg:     spriteImg,
		m:             m,
		backgroundImg: bgImg,
		entities:      components.NewMap(),
		entityList:    []string{},
	}

	g.newPlayer()
	g.initMap()

	// Add slash
	slashFile = ase.Load("assets/animation/slash.json")
	slashFile.Play("Slash")
	img, err := gfx.DecodePNG(assets.FileReaderFatal(slashFile.ImagePath))
	if err != nil {
		log.Fatal(err)
	}
	slashImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	currentTime = time.Now()
	return &g, nil
}

var slashFile ase.File
var slashImg *ebiten.Image

func (g *Game) newPlayer() {

	hitbox := gfx.R(10, 10, 22, 26)
	g.entityList = append(g.entityList, playerID)
	g.entities.Add(playerID, components.NewHitbox(hitbox))
	g.entities.Add(playerID, components.Pos{Vec: gfx.V(0, 0)})
	g.entities.Add(playerID, components.Velocity{Vec: gfx.V(0, 0)})
	g.entities.Add(playerID, components.Directed{D: direction.Left})

	playerFile := ase.Load("assets/animation/hero.json")
	playerFile.Play("Slash")
	img, err := gfx.DecodePNG(assets.FileReaderFatal(playerFile.ImagePath))
	if err != nil {
		log.Fatal(err)
	}
	playerImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	g.entities.Add(playerID, components.Drawable{Image: playerImg})
	g.entities.Add(playerID, components.Animated{Ase: playerFile})

	// Add hook
	g.entityList = append(g.entityList, hookID)
	g.entities.Add(hookID, components.Pos{Vec: gfx.V(0, 0)})
	g.entities.Add(hookID, components.Following{ID: playerID, Offset: gfx.V(16, 16)})
}

func (g *Game) parseObject(o *tiled.Object) {
	ID := g.UUID()
	switch o.Type {
	case "player":
		pos := g.Pos(playerID)
		pos.X, pos.Y = float64(o.X), float64(o.Y)
	default:

		r := gfx.R(o.X, o.Y, o.X+o.Width, o.Y+o.Height)

		rs := []gfx.Rect{}
		for _, y := range []float64{r.Min.Y, r.Max.Y} {
			for _, x := range []float64{r.Min.X, r.Max.X} {
				g.pts = append(g.pts, gfx.V(x, y))
			}
		}

		rs = append(rs, r)
		g.staticSpace.AddShape(resolvutil.ScaledRect(r, collisionScaling))

		b := components.NewHitbox(r)
		g.entities.Add(ID, b)
		g.entityList = append(g.entityList, ID)
	}
}

func (g *Game) UUID() string {
	return fmt.Sprintf("%d", rand.Intn(10000))
}
