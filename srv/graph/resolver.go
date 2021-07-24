package graph

import "github.com/housekpr/core/srv/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	pump model.Pump
	tank model.Tank
}
