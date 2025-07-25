package boards

import (
	_ "embed"
	"encoding/json"
	"github.com/bsloan/game-sandbox/entity"
	"github.com/jakecoffman/cp"
	"log"
	"slices"
)

const (
	EMPTY = iota
	GRASS_LEFT_EDGE
	GRASS_MIDDLE
	GRASS_RIGHT_EDGE
	DIRT_CENTER
	GRASS_SLOPE_R_START
	GRASS_SLOPE_R_BASE_1
	GRASS_SLOPE_R_MIDDLE
	GRASS_SLOPE_R_BASE_2
	GRASS_SLOPE_R_MAGIC_ROOT
	GRASS_SLOPE_L_START
	GRASS_SLOPE_L_BASE_1
	GRASS_SLOPE_L_MIDDLE
	GRASS_SLOPE_L_BASE_2
	GRASS_SLOPE_L_MAGIC_ROOT
)

var NoOpTiles = []int{
	EMPTY,
	GRASS_SLOPE_R_START,
	GRASS_SLOPE_R_MIDDLE,
	GRASS_SLOPE_L_START,
	GRASS_SLOPE_L_MIDDLE,
}

var (
	//go:embed map.json
	Level1Map []byte
)

type Gameboard struct {
	Map         [][]int `json:"map"`
	TileSize    int     `json:"tilesize"`
	TileWidth   int
	TileHeight  int
	PixelWidth  int
	PixelHeight int
}

func (gb *Gameboard) LoadGameboard(mapData []byte, space *cp.Space) {
	err := json.Unmarshal(mapData, &gb)
	if err != nil {
		log.Fatal("Error unmarshaling JSON for map:", err)
	}
	gb.TileWidth = len(gb.Map[0])
	gb.TileHeight = len(gb.Map)
	gb.PixelWidth = gb.TileWidth * gb.TileSize
	gb.PixelHeight = gb.TileHeight * gb.TileSize
	gb.initializeTiles(space)
}

func (gb *Gameboard) initializeTiles(space *cp.Space) {
	for ty := 0; ty < gb.TileHeight; ty++ {
		for tx := 0; tx < gb.TileWidth; tx++ {
			tile := gb.Map[ty][tx]
			if !slices.Contains(NoOpTiles, tile) {
				x, y := float64(tx*16), float64(ty*16)
				tileBody := cp.NewStaticBody()
				tileBody.SetPosition(cp.Vector{X: x, Y: y})
				var tileShape *cp.Shape
				if tile == GRASS_MIDDLE || tile == DIRT_CENTER || tile == GRASS_SLOPE_R_BASE_1 || tile == GRASS_SLOPE_R_BASE_2 || tile == GRASS_SLOPE_L_BASE_1 || tile == GRASS_SLOPE_L_BASE_2 {
					// these are center blocks - make the collision shape a normal 16x16 square
					tileShape = cp.NewBox(tileBody, 16, 16, 0)
					tileShape.SetFriction(1)
					tileShape.SetCollisionType(entity.BlockCollisionType)
				} else if tile == GRASS_LEFT_EDGE || tile == GRASS_RIGHT_EDGE {
					// these are edge blocks - collision shape is narrower
					tileShape = cp.NewBox(tileBody, 8, 16, 0)
					tileShape.SetFriction(1)
					tileShape.SetCollisionType(entity.BlockCollisionType)
				} else if tile == GRASS_SLOPE_R_MAGIC_ROOT {
					// hard code a line segment of a pre-defined length and slope, starting at this tile
					// this is a hack but it works. this particular slope root creates a 26 degree positive
					// slope that is 6 tiles wide and 6 tiles high.
					vert1 := cp.Vector{X: 4, Y: 0}
					vert2 := cp.Vector{X: 80, Y: -36}
					tileShape = cp.NewSegment(tileBody, vert2, vert1, 3)
					tileShape.SetFriction(0.1)
					tileShape.SetCollisionType(entity.SlopeCollisionType)
				} else if tile == GRASS_SLOPE_L_MAGIC_ROOT {
					vert1 := cp.Vector{X: -4, Y: 0}
					vert2 := cp.Vector{X: -80, Y: -36}
					tileShape = cp.NewSegment(tileBody, vert2, vert1, 3)
					tileShape.SetFriction(0.1)
					tileShape.SetCollisionType(entity.SlopeCollisionType)
				} else {
					// unrecognized tile - panic
				}
				tileShape.SetElasticity(0)
				space.AddBody(tileBody)
				space.AddShape(tileShape)
			}
		}
	}
}
