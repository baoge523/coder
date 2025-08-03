package background_controller

import "context"

type BackgroundController interface {
	Start(ctx context.Context) error

	Stop() error
}
