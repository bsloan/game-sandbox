package entity

import (
	"github.com/bsloan/game-sandbox/settings"
	"github.com/jakecoffman/cp"
	"math"
)

const (
	PlayerCollisionType cp.CollisionType = iota
	BlockCollisionType
	SlopeCollisionType
	SwordDogCollisionType
)

func GenericGroundedHandler(space *cp.Space, collisionType cp.CollisionType) {
	handler := space.NewCollisionHandler(collisionType, BlockCollisionType)

	handler.BeginFunc = func(arb *cp.Arbiter, space *cp.Space, data interface{}) bool {
		n := arb.Normal()
		grounded := n.Y > 0 && math.Abs(n.X) < 0.5
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
			dynamicEntity.Body.SetVelocity(dynamicEntity.Body.Velocity().X, dynamicEntity.Body.Velocity().Y-50)
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

		//var enemyBody *cp.Body
		var playerBody *cp.Body

		if body1.UserData.(*Entity).Type == Player {
			playerBody = body1
			//enemyBody = body2
		} else {
			playerBody = body2
			//enemyBody = body1
		}

		// jolt the player backwards a bit
		if n.X > 0.5 {
			playerBody.ApplyForceAtLocalPoint(cp.Vector{X: settings.PlayerJumpInitialVelocity * 3, Y: playerBody.Position().Y}, cp.Vector{X: 0, Y: 0})
		} else if n.X < -0.5 {
			playerBody.ApplyForceAtLocalPoint(cp.Vector{X: -settings.PlayerJumpInitialVelocity * 3, Y: playerBody.Position().Y}, cp.Vector{X: 0, Y: 0})
		}

		// TODO: set Player state to Hurt ?

		// TODO: subtract damage from player health
		return true
	}
}

func InitializeCollisionHandlers(space *cp.Space) {
	// attach collision handlers to player
	GenericGroundedHandler(space, PlayerCollisionType)
	SlopeHandler(space, PlayerCollisionType)

	// attach collision handlers to enemies
	DamagePlayerHandler(space, SwordDogCollisionType)
	GenericGroundedHandler(space, SwordDogCollisionType)
	SlopeHandler(space, SwordDogCollisionType)
	ObstructedHandler(space, SwordDogCollisionType)
}
