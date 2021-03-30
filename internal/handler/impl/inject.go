package impl

import "go.uber.org/dig"

func Inject(container *dig.Container) error {
	_ = container.Provide(newIndexHandler)
	return container.Provide(newUserHandler)
}
