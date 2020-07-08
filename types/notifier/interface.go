package notifier

import (
	"github.com/torusresearch/statping/types/failures"
	"github.com/torusresearch/statping/types/services"
)

// Notifier interface is required to create a new Notifier
type Notifier interface {
	OnSuccess(services.Service) (string, error)                   // OnSuccess is triggered when a service is successful
	OnFailure(services.Service, failures.Failure) (string, error) // OnFailure is triggered when a service is failing
	OnTest() (string, error)                                      // OnTest is triggered for testing
	OnSave() (string, error)                                      // OnSave is triggered for when saved
}
