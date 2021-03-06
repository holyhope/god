package god

import (
	"context"
)

// Unit represents single service managed by either systemd or launchd.
type Unit interface {
	// Create the unit to the system.
	Create(ctx context.Context) error
	// Enable the unit.
	Enable(ctx context.Context) error
	// Delete the unit from the system.
	Delete(ctx context.Context) error
	// Disable the unit.
	Disable(ctx context.Context) error
	// Status the status of the unit.
	Status(ctx context.Context) (UnitStatus, error)
}

type UnitStatus interface {
	// Exists returns true if the unit is on the filesystem.
	Exists(ctx context.Context) bool
	// IsEnabled returns true if the unit is enable.
	IsEnabled(ctx context.Context) bool
}
