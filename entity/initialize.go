package entity

import (
	"github.com/bsloan/game-sandbox/asset"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

func InitializePlayer(space *cp.Space, x, y float64) *Entity {
	idleRight := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerIdleRight1,
			asset.PlayerIdleRight2,
			asset.PlayerIdleRight3,
			asset.PlayerIdleRight4,
		},
		AnimationSpeed: 0.1,
	}
	idleLeft := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerIdleLeft1,
			asset.PlayerIdleLeft2,
			asset.PlayerIdleLeft3,
			asset.PlayerIdleLeft4,
		},
		AnimationSpeed: 0.1,
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
		AnimationSpeed: 0.2,
	}
	moveLeft := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerMoveLeft1,
			asset.PlayerMoveLeft2,
			asset.PlayerMoveLeft3,
			asset.PlayerMoveLeft4,
			asset.PlayerMoveLeft5,
			asset.PlayerMoveLeft6,
		},
		AnimationSpeed: 0.2,
	}
	jumpRight := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerJumpRight1,
		},
	}
	jumpLeft := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerJumpLeft1,
		},
	}
	fallRight := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerFallRight1,
		},
	}
	fallLeft := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerFallLeft1,
		},
	}
	activeRight := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerMoveRight4,
			asset.PlayerMoveRight5,
			asset.PlayerMoveRight6,
			asset.PlayerMoveRight6,
			asset.PlayerMoveRight2,
			asset.PlayerMoveRight2,
		},
		AnimationSpeed:        0.4,
		EntityStateTransition: Idle,
	}
	activeLeft := Animation{
		Frames: []*ebiten.Image{
			asset.PlayerMoveLeft4,
			asset.PlayerMoveLeft5,
			asset.PlayerMoveLeft6,
			asset.PlayerMoveLeft6,
			asset.PlayerMoveLeft2,
			asset.PlayerMoveLeft2,
		},
		AnimationSpeed:        0.4,
		EntityStateTransition: Idle,
	}
	player := Entity{
		Type:   Player,
		State:  Idle,
		Facing: Right,
		Animations: map[EntityState]*Animation{
			Idle:         &idleRight,
			IdleRight:    &idleRight,
			IdleLeft:     &idleLeft,
			MovingRight:  &moveRight,
			MovingLeft:   &moveLeft,
			JumpingRight: &jumpRight,
			JumpingLeft:  &jumpLeft,
			FallingRight: &fallRight,
			FallingLeft:  &fallLeft,
			ActiveRight:  &activeRight,
			ActiveLeft:   &activeLeft,
		},
		Body:  cp.NewBody(1, cp.INFINITY),
		Boost: 0,
	}

	player.Body.UserData = &player
	space.AddBody(player.Body)
	player.Body.SetPosition(cp.Vector{X: x, Y: y})
	//playerShape := space.AddShape(cp.NewBox(player.Body, 8, 7, 8))
	// radius was: 10
	playerShape := space.AddShape(cp.NewCircle(player.Body, 11, cp.Vector{X: 0, Y: 0}))

	playerShape.SetElasticity(0)
	playerShape.SetFriction(0.75)
	playerShape.SetCollisionType(PlayerCollisionType)
	player.Shape = playerShape

	player.MaxHealth = 50
	player.Health = player.MaxHealth

	return &player
}

func InitializePlayerSword(space *cp.Space, x, y float64) *Entity {
	slashRight := Animation{
		Frames: []*ebiten.Image{
			asset.WhiteSlashRight1,
			asset.WhiteSlashRight2,
			asset.WhiteSlashRight3,
			asset.WhiteSlashRight4,
			asset.WhiteSlashRight5,
			asset.WhiteSlashRight6,
		},
		AnimationSpeed:        0.4,
		EntityStateTransition: Dead,
	}
	slashLeft := Animation{
		Frames: []*ebiten.Image{
			asset.WhiteSlashLeft1,
			asset.WhiteSlashLeft2,
			asset.WhiteSlashLeft3,
			asset.WhiteSlashLeft4,
			asset.WhiteSlashLeft5,
			asset.WhiteSlashLeft6,
		},
		AnimationSpeed:        0.4,
		EntityStateTransition: Dead,
	}
	sword := Entity{
		Type:   PlayerWeapon,
		State:  Idle,
		Facing: Right,
		Animations: map[EntityState]*Animation{
			Idle:        nil,
			ActiveRight: &slashRight,
			ActiveLeft:  &slashLeft,
		},
		Body: cp.NewBody(0.001, cp.INFINITY),
	} // mass is set to near-zero to ignore gravity. setting mass to true 0 causes weird physics bugs

	sword.Body.UserData = &sword
	space.AddBody(sword.Body)
	sword.Body.SetPosition(cp.Vector{X: x, Y: y})
	sword.AttackDamage = 2
	return &sword
}

func InitializeSwordDog(space *cp.Space, x, y float64) *Entity {
	idleRight := Animation{
		Frames: []*ebiten.Image{
			asset.SwordDogIdleRight1,
			asset.SwordDogIdleRight2,
			asset.SwordDogIdleRight3,
			asset.SwordDogIdleRight4,
			asset.SwordDogIdleRight5,
			asset.SwordDogIdleRight6,
		},
		AnimationSpeed: 0.1,
	}
	idleLeft := Animation{
		Frames: []*ebiten.Image{
			asset.SwordDogIdleLeft1,
			asset.SwordDogIdleLeft2,
			asset.SwordDogIdleLeft3,
			asset.SwordDogIdleLeft4,
			asset.SwordDogIdleLeft5,
			asset.SwordDogIdleLeft6,
		},
		AnimationSpeed: 0.1,
	}
	runRight := Animation{
		Frames: []*ebiten.Image{
			asset.SwordDogRunRight1,
			asset.SwordDogRunRight2,
			asset.SwordDogRunRight3,
			asset.SwordDogRunRight4,
			asset.SwordDogRunRight5,
			asset.SwordDogRunRight6,
		},
		AnimationSpeed: 0.2,
	}
	runLeft := Animation{
		Frames: []*ebiten.Image{
			asset.SwordDogRunLeft1,
			asset.SwordDogRunLeft2,
			asset.SwordDogRunLeft3,
			asset.SwordDogRunLeft4,
			asset.SwordDogRunLeft5,
			asset.SwordDogRunLeft6,
		},
		AnimationSpeed: 0.2,
	}
	bigSlashRight := Animation{
		Frames: []*ebiten.Image{
			asset.SwordDogBigSlashRight1,
			asset.SwordDogBigSlashRight2,
			asset.SwordDogBigSlashRight3,
			asset.SwordDogBigSlashRight4,
			asset.SwordDogBigSlashRight5,
			asset.SwordDogBigSlashRight6,
		},
		AnimationSpeed:        0.2,
		EntityStateTransition: MovingRight,
	}
	bigSlashLeft := Animation{
		Frames: []*ebiten.Image{
			asset.SwordDogBigSlashLeft1,
			asset.SwordDogBigSlashLeft2,
			asset.SwordDogBigSlashLeft3,
			asset.SwordDogBigSlashLeft4,
			asset.SwordDogBigSlashLeft5,
			asset.SwordDogBigSlashLeft6,
		},
		AnimationSpeed:        0.3,
		EntityStateTransition: MovingLeft,
	}
	downSlashRight := Animation{
		Frames: []*ebiten.Image{
			asset.SwordDogDownSlashRight1,
			asset.SwordDogDownSlashRight2,
			asset.SwordDogDownSlashRight3,
			asset.SwordDogDownSlashRight4,
			asset.SwordDogDownSlashRight5,
		},
		AnimationSpeed:        0.3,
		EntityStateTransition: MovingRight,
	}
	downSlashLeft := Animation{
		Frames: []*ebiten.Image{
			asset.SwordDogDownSlashLeft1,
			asset.SwordDogDownSlashLeft2,
			asset.SwordDogDownSlashLeft3,
			asset.SwordDogDownSlashLeft4,
			asset.SwordDogDownSlashLeft5,
		},
		AnimationSpeed:        0.3,
		EntityStateTransition: MovingLeft,
	}
	dying := Animation{
		Frames: []*ebiten.Image{
			asset.EnemyDeath1,
			asset.EnemyDeath2,
			asset.EnemyDeath3,
			asset.EnemyDeath4,
			asset.EnemyDeath5,
			asset.EnemyDeath6,
		},
		AnimationSpeed:        0.3,
		EntityStateTransition: Dead,
	}
	swordDog := Entity{
		Type:          SwordDog,
		State:         IdleLeft,
		RememberState: ActiveLeft,
		Facing:        Left,
		Animations: map[EntityState]*Animation{
			IdleRight:    &idleRight,
			IdleLeft:     &idleLeft,
			MovingRight:  &runRight,
			MovingLeft:   &runLeft,
			ActiveRight:  &bigSlashRight,
			ActiveLeft:   &bigSlashLeft,
			ActiveRight2: &downSlashRight,
			ActiveLeft2:  &downSlashLeft,
			Dying:        &dying,
		},
		Body: cp.NewBody(1, cp.INFINITY),
	}
	swordDog.Body.UserData = &swordDog
	space.AddBody(swordDog.Body)
	swordDog.Body.SetPosition(cp.Vector{X: x, Y: y})
	swordDogShape := space.AddShape(cp.NewCircle(swordDog.Body, 11, cp.Vector{X: 10, Y: 0}))
	swordDogShape.SetElasticity(0)
	swordDogShape.SetFriction(0.75)
	swordDogShape.SetCollisionType(GenericEnemyCollisionType)
	swordDog.Shape = swordDogShape
	swordDog.AttackDamage = 2
	swordDog.MaxHealth = 6
	swordDog.Health = swordDog.MaxHealth
	return &swordDog
}

func InitializeAlligator(space *cp.Space, x, y float64) *Entity {
	idleRight := Animation{
		Frames: []*ebiten.Image{
			asset.AlligatorIdleRight1,
			asset.AlligatorIdleRight2,
			asset.AlligatorIdleRight3,
			asset.AlligatorIdleRight4,
		},
		AnimationSpeed: 0.1,
	}
	idleLeft := Animation{
		Frames: []*ebiten.Image{
			asset.AlligatorIdleLeft1,
			asset.AlligatorIdleLeft2,
			asset.AlligatorIdleLeft3,
			asset.AlligatorIdleLeft4,
		},
		AnimationSpeed: 0.1,
	}
	runRight := Animation{
		Frames: []*ebiten.Image{
			asset.AlligatorRunRight1,
			asset.AlligatorRunRight2,
			asset.AlligatorRunRight3,
			asset.AlligatorRunRight4,
			asset.AlligatorRunRight5,
			asset.AlligatorRunRight6,
			asset.AlligatorRunRight7,
			asset.AlligatorRunRight8,
			asset.AlligatorRunRight9,
		},
		AnimationSpeed: 0.2,
	}
	runLeft := Animation{
		Frames: []*ebiten.Image{
			asset.AlligatorRunLeft1,
			asset.AlligatorRunLeft2,
			asset.AlligatorRunLeft3,
			asset.AlligatorRunLeft4,
			asset.AlligatorRunLeft5,
			asset.AlligatorRunLeft6,
			asset.AlligatorRunLeft7,
			asset.AlligatorRunLeft8,
			asset.AlligatorRunLeft9,
		},
		AnimationSpeed: 0.2,
	}
	dying := Animation{
		Frames: []*ebiten.Image{
			asset.EnemyDeath1,
			asset.EnemyDeath2,
			asset.EnemyDeath3,
			asset.EnemyDeath4,
			asset.EnemyDeath5,
			asset.EnemyDeath6,
		},
		AnimationSpeed:        0.3,
		EntityStateTransition: Dead,
	}
	slashRight := Animation{
		Frames: []*ebiten.Image{
			asset.AlligatorSlashRight1,
			asset.AlligatorSlashRight2,
			asset.AlligatorSlashRight3,
			asset.AlligatorSlashRight4,
			asset.AlligatorSlashRight5,
			asset.AlligatorSlashRight6,
			asset.AlligatorSlashRight7,
			asset.AlligatorSlashRight8,
			asset.AlligatorSlashRight9,
			asset.AlligatorSlashRight10,
			asset.AlligatorSlashRight11,
		},
		AnimationSpeed:        0.3,
		EntityStateTransition: MovingRight,
	}
	slashLeft := Animation{
		Frames: []*ebiten.Image{
			asset.AlligatorSlashLeft1,
			asset.AlligatorSlashLeft2,
			asset.AlligatorSlashLeft3,
			asset.AlligatorSlashLeft4,
			asset.AlligatorSlashLeft5,
			asset.AlligatorSlashLeft6,
			asset.AlligatorSlashLeft7,
			asset.AlligatorSlashLeft8,
			asset.AlligatorSlashLeft9,
			asset.AlligatorSlashLeft10,
			asset.AlligatorSlashLeft11,
		},
		AnimationSpeed:        0.3,
		EntityStateTransition: MovingLeft,
	}
	alligator := Entity{
		Type:          Alligator,
		State:         IdleLeft,
		RememberState: ActiveLeft,
		Facing:        Left,
		Animations: map[EntityState]*Animation{
			IdleRight:   &idleRight,
			IdleLeft:    &idleLeft,
			MovingRight: &runRight,
			MovingLeft:  &runLeft,
			ActiveRight: &slashRight,
			ActiveLeft:  &slashLeft,
			Dying:       &dying,
		},
		Body: cp.NewBody(1, cp.INFINITY),
	}
	alligator.Body.UserData = &alligator
	space.AddBody(alligator.Body)
	alligator.Body.SetPosition(cp.Vector{X: x, Y: y})
	alligatorShape := space.AddShape(cp.NewCircle(alligator.Body, 11, cp.Vector{X: 10, Y: 0}))
	alligatorShape.SetElasticity(0)
	alligatorShape.SetFriction(0.75)
	alligatorShape.SetCollisionType(GenericEnemyCollisionType)
	alligator.Shape = alligatorShape
	alligator.AttackDamage = 4
	alligator.MaxHealth = 12
	alligator.Health = alligator.MaxHealth
	return &alligator
}

func InitializeFrog(space *cp.Space, x, y float64) *Entity {
	idleRight := Animation{
		Frames: []*ebiten.Image{
			asset.FrogIdleRight1,
			asset.FrogIdleRight1,
			asset.FrogIdleRight1,
			asset.FrogIdleRight1,
			asset.FrogIdleRight1,
			asset.FrogIdleRight1,
			asset.FrogIdleRight1,
			asset.FrogIdleRight2,
			asset.FrogIdleRight3,
			asset.FrogIdleRight4,
		},
		AnimationSpeed: 0.1,
	}
	idleLeft := Animation{
		Frames: []*ebiten.Image{
			asset.FrogIdleLeft1,
			asset.FrogIdleLeft1,
			asset.FrogIdleLeft1,
			asset.FrogIdleLeft1,
			asset.FrogIdleLeft1,
			asset.FrogIdleLeft1,
			asset.FrogIdleLeft1,
			asset.FrogIdleLeft2,
			asset.FrogIdleLeft3,
			asset.FrogIdleLeft4,
		},
		AnimationSpeed: 0.1,
	}
	dying := Animation{
		Frames: []*ebiten.Image{
			asset.EnemyDeath1,
			asset.EnemyDeath2,
			asset.EnemyDeath3,
			asset.EnemyDeath4,
			asset.EnemyDeath5,
			asset.EnemyDeath6,
		},
		AnimationSpeed:        0.3,
		EntityStateTransition: Dead,
	}
	frog := Entity{
		Type:   Frog,
		State:  IdleLeft,
		Facing: Left,
		Animations: map[EntityState]*Animation{
			IdleRight: &idleRight,
			IdleLeft:  &idleLeft,
			Dying:     &dying,
		},
		Body: cp.NewBody(1, cp.INFINITY),
	}
	frog.Body.UserData = &frog
	space.AddBody(frog.Body)
	frog.Body.SetPosition(cp.Vector{X: x, Y: y})
	frogShape := space.AddShape(cp.NewCircle(frog.Body, 16, cp.Vector{X: 2, Y: -6}))
	frogShape.SetElasticity(0)
	frogShape.SetFriction(0.75)
	frogShape.SetCollisionType(GenericEnemyCollisionType)
	frog.Shape = frogShape
	frog.AttackDamage = 4
	frog.MaxHealth = 2
	frog.Health = frog.MaxHealth
	return &frog
}
