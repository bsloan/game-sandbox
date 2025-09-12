package entity

import (
	"math"

	"github.com/bsloan/game-sandbox/settings"
	"github.com/jakecoffman/cp"
)

const (
	PlayerCollisionType cp.CollisionType = iota
	PlayerSwordCollisionType
	BlockCollisionType
	SlopeCollisionType
	GenericEnemyCollisionType
	FrogCollisionType
	EagleCollisionType
	GemCollisionType
)

func GenericGroundedHandler(space *cp.Space, collisionType cp.CollisionType) {
	handler := space.NewCollisionHandler(collisionType, BlockCollisionType)

	handler.BeginFunc = func(arb *cp.Arbiter, space *cp.Space, data interface{}) bool {
		n := arb.Normal()
		grounded := n.Y > 0 && math.Abs(n.X) < 0.5 // check for small x overlap to help not get stuck on walls
		if grounded {
			body1, body2 := arb.Bodies()
			if body1.UserData != nil {
				body1.UserData.(*Entity).Grounded = true
				body1.UserData.(*Entity).Shape.SetFriction(0.75)
			} else {
				body2.UserData.(*Entity).Grounded = true
				body1.UserData.(*Entity).Shape.SetFriction(0.75)
			}
		}
		return true
	}
}

func SlopeHandler(space *cp.Space, collisionType cp.CollisionType) {
	handler := space.NewCollisionHandler(collisionType, SlopeCollisionType)

	handler.BeginFunc = func(arb *cp.Arbiter, space *cp.Space, data interface{}) bool {
		n := arb.Normal()
		body1, body2 := arb.Bodies()
		// determine which of the colliding entities is the non-static one
		// e.g., which is the moving entity such as the player, not the tile
		var dynamicEntity *Entity
		if body1.UserData != nil {
			dynamicEntity = body1.UserData.(*Entity)
		} else {
			dynamicEntity = body2.UserData.(*Entity)
		}

		grounded := n.Y > 0
		sloping := n.X > 0 || n.X < 0
		if sloping {
			// entities need a little uphill boost to get up the hill - apply boost by entity type
			// TODO: uphillBoost could be an attribute of the Entity
			uphillBoost := 0.0
			if dynamicEntity.Type == Player {
				uphillBoost = 50
			} else if dynamicEntity.Type == SwordDog {
				uphillBoost = 75
			}
			dynamicEntity.Body.SetVelocity(dynamicEntity.Body.Velocity().X, dynamicEntity.Body.Velocity().Y-uphillBoost)
			dynamicEntity.OnSlope = true
		}
		if grounded {
			dynamicEntity.Grounded = true
		}
		return true
	}

	handler.SeparateFunc = func(arb *cp.Arbiter, space *cp.Space, data interface{}) {
		body1, body2 := arb.Bodies()
		if body1.UserData != nil {
			body1.UserData.(*Entity).OnSlope = false
		} else {
			body2.UserData.(*Entity).OnSlope = false
		}
	}
}

func ObstructedHandler(space *cp.Space, collisionType cp.CollisionType) {
	handler := space.NewCollisionHandler(collisionType, BlockCollisionType)
	handler.BeginFunc = func(arb *cp.Arbiter, space *cp.Space, data interface{}) bool {
		n := arb.Normal()
		body1, body2 := arb.Bodies()
		var dynamicEntity *Entity
		if body1.UserData != nil {
			dynamicEntity = body1.UserData.(*Entity)
		} else {
			dynamicEntity = body2.UserData.(*Entity)
		}
		if n.X > 0.5 {
			dynamicEntity.Facing = Left
			dynamicEntity.State = MovingLeft
		} else if n.X < -0.5 {
			dynamicEntity.Facing = Right
			dynamicEntity.State = MovingRight
		}
		return true
	}
}

func DamagePlayerHandler(space *cp.Space, collisionType cp.CollisionType) {
	handler := space.NewCollisionHandler(collisionType, PlayerCollisionType)

	handler.BeginFunc = func(arb *cp.Arbiter, space *cp.Space, data interface{}) bool {
		n := arb.Normal()
		body1, body2 := arb.Bodies()

		var enemyBody *cp.Body
		var playerBody *cp.Body

		if body1.UserData.(*Entity).Type == Player {
			playerBody = body1
			enemyBody = body2
		} else {
			playerBody = body2
			enemyBody = body1
		}

		// jolt the player backwards a bit
		if n.X > 0.75 && playerBody.UserData.(*Entity).Damaged <= 0 {
			playerBody.ApplyForceAtLocalPoint(cp.Vector{X: settings.PlayerJumpInitialVelocity * 3, Y: 0}, cp.Vector{X: 0, Y: 0})
			playerBody.UserData.(*Entity).Damaged = 5
		} else if n.X < -0.75 {
			playerBody.ApplyForceAtLocalPoint(cp.Vector{X: -settings.PlayerJumpInitialVelocity * 3, Y: 0}, cp.Vector{X: 0, Y: 0})
			playerBody.UserData.(*Entity).Damaged = 5
		}
		if n.Y > 0.75 && playerBody.UserData.(*Entity).Damaged <= 0 {
			playerBody.ApplyForceAtLocalPoint(cp.Vector{X: 0, Y: settings.PlayerJumpInitialVelocity * 3}, cp.Vector{X: 0, Y: 0})
			playerBody.UserData.(*Entity).Damaged = 5
		} else if n.Y < -0.75 {
			playerBody.ApplyForceAtLocalPoint(cp.Vector{X: 0, Y: -settings.PlayerJumpInitialVelocity * 3}, cp.Vector{X: 0, Y: 0})
			playerBody.UserData.(*Entity).Damaged = 5
		}

		// subtract enemy's attack damage from player's health
		playerBody.UserData.(*Entity).Health -= enemyBody.UserData.(*Entity).AttackDamage

		return true
	}
}

func PlayerSwordHandler(space *cp.Space, collisionType cp.CollisionType) {
	handler := space.NewCollisionHandler(collisionType, PlayerSwordCollisionType)

	handler.BeginFunc = func(arb *cp.Arbiter, space *cp.Space, data interface{}) bool {
		body1, body2 := arb.Bodies()

		var enemyBody *cp.Body
		var swordBody *cp.Body

		if body1.UserData.(*Entity).Type == PlayerWeapon {
			swordBody = body1
			enemyBody = body2
		} else {
			swordBody = body2
			enemyBody = body1
		}

		enemyBody.UserData.(*Entity).Damaged = 3

		// subtract attack damage from other entity's health
		enemyBody.UserData.(*Entity).Health -= swordBody.UserData.(*Entity).AttackDamage

		return true
	}
}

func GemHandler(space *cp.Space, collisionType cp.CollisionType) {
	handler := space.NewCollisionHandler(collisionType, GemCollisionType)

	handler.BeginFunc = func(arb *cp.Arbiter, space *cp.Space, data interface{}) bool {
		body1, body2 := arb.Bodies()

		var gemBody *cp.Body
		var otherBody *cp.Body

		if body1.UserData.(*Entity).Type == Gem {
			gemBody = body1
			otherBody = body2
		} else {
			gemBody = body2
			otherBody = body1
		}

		if otherBody.UserData.(*Entity).Type == Player {
			gemBody.UserData.(*Entity).State = Dying
			gemBody.EachShape(func(shape *cp.Shape) {
				gemBody.RemoveShape(shape)
				space.RemoveShape(shape)
			})
		}

		// ignore entity/gem collisions
		return false
	}
}

func InitializeCollisionHandlers(space *cp.Space) {
	// attach collision handlers to player
	GenericGroundedHandler(space, PlayerCollisionType)
	SlopeHandler(space, PlayerCollisionType)
	GemHandler(space, PlayerCollisionType)

	// attach collision handlers to generic enemies
	DamagePlayerHandler(space, GenericEnemyCollisionType)
	GenericGroundedHandler(space, GenericEnemyCollisionType)
	SlopeHandler(space, GenericEnemyCollisionType)
	ObstructedHandler(space, GenericEnemyCollisionType)
	PlayerSwordHandler(space, GenericEnemyCollisionType)
	GemHandler(space, GenericEnemyCollisionType)

	// custom enemy collision handlers
	// frog
	DamagePlayerHandler(space, FrogCollisionType)
	GenericGroundedHandler(space, FrogCollisionType)
	SlopeHandler(space, FrogCollisionType)
	PlayerSwordHandler(space, FrogCollisionType)
	GemHandler(space, FrogCollisionType)
	// eagle
	DamagePlayerHandler(space, EagleCollisionType)
	PlayerSwordHandler(space, EagleCollisionType)
	ObstructedHandler(space, EagleCollisionType)
	GemHandler(space, EagleCollisionType)
}
