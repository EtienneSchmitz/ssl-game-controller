package engine

import (
	"github.com/RoboCup-SSL/ssl-game-controller/internal/app/geom"
	"github.com/RoboCup-SSL/ssl-game-controller/internal/app/state"
	"github.com/RoboCup-SSL/ssl-game-controller/internal/app/statemachine"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const ballSteadyThreshold = 0.2
const robotRadius = 0.09
const distanceThreshold = 0.05

func createGameEventChange(eventType state.GameEvent_Type, event *state.GameEvent) *statemachine.Change {
	event.Type = &eventType
	event.Origin = []string{changeOriginEngine}
	return &statemachine.Change{
		Origin: &changeOriginEngine,
		Change: &statemachine.Change_AddGameEvent{
			AddGameEvent: &statemachine.AddGameEvent{
				GameEvent: event,
			},
		},
	}
}

func (e *Engine) robotsInsideRadius(robots []*Robot, pos *geom.Vector2, radius float64) bool {
	for _, robot := range robots {
		distance := robot.Pos.DistanceTo(pos)
		if distance < radius {
			return true
		}
	}

	return false
}

func goDur(duration *durationpb.Duration) time.Duration {
	return duration.AsDuration()
}

func goTime(timestamp *timestamppb.Timestamp) time.Time {
	return timestamp.AsTime()
}

func addDur(duration *durationpb.Duration, delta time.Duration) {
	*duration = *durationpb.New(duration.AsDuration() + delta)
}

func (e *Engine) ballSteady() bool {
	if e.gcState.TrackerStateGc.Ball == nil || e.gcState.TrackerStateGc.Ball.Vel == nil {
		return true
	}
	return e.gcState.TrackerStateGc.Ball.Vel.ToVector2().Length() < ballSteadyThreshold
}
