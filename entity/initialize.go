package entity

import (
	"github.com/bsloan/game-sandbox/asset"
	"github.com/hajimehoshi/ebiten/v2"
)

func InitializePlayer(x, y float64) *Entity {
	idle := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerIdleRight1,
			asset.PlayerIdleRight2,
			asset.PlayerIdleRight3,
			asset.PlayerIdleRight4,
		},
		AnimationSpeed: 0.05,
	}
	moveRight := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerMoveRight1,
			asset.PlayerMoveRight2,
			asset.PlayerMoveRight3,
			asset.PlayerMoveRight4,
			asset.PlayerMoveRight5,
			asset.PlayerMoveRight6,
		},
		AnimationSpeed: 0.1,
	}
	player := &Entity{
		Type:  Player,
		State: Idle,
		XPos:  x,
		YPos:  y,
		Animations: map[EntityState]*Animation{
			Idle:        &idle,
			MovingRight: &moveRight,
		},
	}
	return player
}
