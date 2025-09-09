package boards

import (
	_ "embed"
	"encoding/json"
	"log"
	"slices"

	"github.com/bsloan/game-sandbox/asset"
	"github.com/bsloan/game-sandbox/entity"
	"github.com/jakecoffman/cp"
)

var NoOpTiles = []int{
	asset.EMPTY,
	asset.GRASS_SLOPE_R_START,
	asset.GRASS_SLOPE_R_MIDDLE,
	asset.GRASS_SLOPE_L_START,
	asset.GRASS_SLOPE_L_MIDDLE,
	asset.GRASS_SLOPE_STEEP_R_START,
	asset.GRASS_SLOPE_STEEP_L_START,
	asset.GRASS_FOREGROUND_1,
	asset.GRASS_FOREGROUND_2,
	asset.BRANCH_FOREGROUND_1,
	asset.BRANCH_FOREGROUND_2,
	asset.CAVE_BACKGROUND_L_1,
	asset.CAVE_BACKGROUND_L_2,
	asset.CAVE_BACKGROUND_L_3,
	asset.CAVE_BACKGROUND_L_4,
	asset.CAVE_BACKGROUND_L_5,
	asset.CAVE_BACKGROUND_L_6,
	asset.CAVE_BACKGROUND_L_7,
	asset.CAVE_BACKGROUND_R_1,
	asset.CAVE_BACKGROUND_R_2,
	asset.CAVE_BACKGROUND_R_3,
	asset.CAVE_BACKGROUND_R_4,
	asset.CAVE_BACKGROUND_R_5,
	asset.CAVE_BACKGROUND_R_6,
	asset.CAVE_BACKGROUND_R_7,
	asset.CAVE_BACKGROUND_ROCKS,
}

var ForegroundTiles = []int{
	asset.GRASS_FOREGROUND_1,
	asset.GRASS_FOREGROUND_2,
	asset.BRANCH_FOREGROUND_1,
	asset.BRANCH_FOREGROUND_2,
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

func (gb *Gameboard) LoadGameboard(mapData []byte, space *cp.Space, registry *entity.Registry) {
	err := json.Unmarshal(mapData, &gb)
	if err != nil {
		log.Fatal("Error unmarshaling JSON for map:", err)
	}
	gb.TileWidth = len(gb.Map[0])
	gb.TileHeight = len(gb.Map)
	gb.PixelWidth = gb.TileWidth * gb.TileSize
	gb.PixelHeight = gb.TileHeight * gb.TileSize
	gb.initializeTiles(space)
	gb.initializeEntities(space, registry)
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
				if tile == asset.GRASS_SLOPE_R_MAGIC_ROOT {
					// hard code a line segment of a pre-defined length and slope, starting at this tile
					// this is a hack but it works. this particular slope root creates a 26 degree positive
					// slope that is 6 tiles wide and 6 tiles high.
					vert1 := cp.Vector{X: 4, Y: 0}
					vert2 := cp.Vector{X: 80, Y: -34} // was -36
					tileShape = cp.NewSegment(tileBody, vert2, vert1, 3)
					tileShape.SetFriction(0.1)
					tileShape.SetCollisionType(entity.SlopeCollisionType)
				} else if tile == asset.GRASS_SLOPE_L_MAGIC_ROOT {
					vert1 := cp.Vector{X: -4, Y: 0}
					vert2 := cp.Vector{X: -80, Y: -34} // was: -36
					tileShape = cp.NewSegment(tileBody, vert2, vert1, 3)
					tileShape.SetFriction(0.1)
					tileShape.SetCollisionType(entity.SlopeCollisionType)
				} else if tile == asset.GRASS_SLOPE_STEEP_R_MAGIC_ROOT {
					vert1 := cp.Vector{X: 4, Y: 0}    // 4
					vert2 := cp.Vector{X: 29, Y: -24} // 30, -24 better)
					tileShape = cp.NewSegment(tileBody, vert2, vert1, 3)
					tileShape.SetFriction(0.1)
					tileShape.SetCollisionType(entity.SlopeCollisionType)
				} else if tile == asset.GRASS_SLOPE_STEEP_L_MAGIC_ROOT {
					vert1 := cp.Vector{X: -4, Y: 0}    // -4
					vert2 := cp.Vector{X: -29, Y: -24} // -30, -24
					tileShape = cp.NewSegment(tileBody, vert2, vert1, 3)
					tileShape.SetFriction(0.1)
					tileShape.SetCollisionType(entity.SlopeCollisionType)
				} else {
					// by default, just make it a regular block
					tileShape = cp.NewBox(tileBody, 16, 16, 0)
					tileShape.SetFriction(1)
					tileShape.SetCollisionType(entity.BlockCollisionType)
				}
				tileShape.SetElasticity(0)
				space.AddBody(tileBody)
				space.AddShape(tileShape)
			}
		}
	}
}

func (gb *Gameboard) initializeEntities(space *cp.Space, registry *entity.Registry) {
	// FIXME: initialize from the map instead of hard-coding
	swordDog1 := entity.InitializeSwordDog(space, 555, 414)
	swordDog2 := entity.InitializeSwordDog(space, 493, 414)
	alligator1 := entity.InitializeAlligator(space, 820, 414)

	frog1 := entity.InitializeFrog(space, 1100, 400)
	frog2 := entity.InitializeFrog(space, 300, 285)

	eagle1 := entity.InitializeEagle(space, 300, 250)

	registry.AddEntity(swordDog1)
	registry.AddEntity(swordDog2)
	registry.AddEntity(alligator1)
	registry.AddEntity(frog1)
	registry.AddEntity(frog2)
	registry.AddEntity(eagle1)
}
