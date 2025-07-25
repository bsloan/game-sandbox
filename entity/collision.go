package entity

import (
	"github.com/jakecoffman/cp"
)

const (
	PlayerCollisionType cp.CollisionType = iota
	BlockCollisionType
	SlopeCollisionType
)

func GenericGroundedHandler(space *cp.Space, collisionType cp.CollisionType) {
	handler := space.NewCollisionHandler(collisionType, BlockCollisionType)

	handler.BeginFunc = func(arb *cp.Arbiter, space *cp.Space, data interface{}) bool {
		n := arb.Normal()
		grounded := n.Y > 0
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
