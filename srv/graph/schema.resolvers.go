package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/housekpr/core/srv/devices/pump"
	"github.com/housekpr/core/srv/devices/sensors/distance"
	"github.com/housekpr/core/srv/graph/generated"
	"github.com/housekpr/core/srv/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SetPumpState(ctx context.Context, input model.PumpState) (*model.Pump, error) {
	if input.State {
		log.Println("Starting the pump.")
		pump.Start()
	} else {
		log.Println("Stopping the pump.")
		pump.Stop()
	}
	r.pump.State = input.State
	log.Println("Setting pump state to", input.State, "completed.")
	return &r.pump, nil
}

func (r *queryResolver) ViewPump(ctx context.Context) (*model.Pump, error) {
	// Read the State of the pump from GPIO and convert it to bool
	var pumpState bool = pump.ReadState() != 0
	log.Println("Current pump state according to GPIO is", pumpState)
	if r.pump.State != pumpState {
		log.Println("[Warning] The state in the app and in GPIO is not in sync!", pumpState)
		log.Println("[Warning] Setting the app state to the GPIO state. ")
		r.pump.State = pumpState
	}
	return &r.pump, nil
}

func (r *queryResolver) GetTankLevel(ctx context.Context) (*model.Tank, error) {
	// Get Tank Level
	r.tank.Level = distance.GetTankLevel()
	return &r.tank, nil
}

func (r *queryResolver) FindUsers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
