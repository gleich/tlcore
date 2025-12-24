package db

import (
	"fmt"

	"go.mattglei.ch/timber"
	"go.mattglei.ch/tlcore/pkg/timelog"
	"gorm.io/gorm"
)

func RunMigrations(database *gorm.DB) error {
	types := map[string]any{
		"task": &timelog.Task{},
	}
	typeCount := len(types)
	timber.Infof("running migrations (%d total)", typeCount)

	i := 1
	for name, model := range types {
		err := database.AutoMigrate(model)
		if err != nil {
			return fmt.Errorf("%w failed to run migration for %s", err, name)
		}
		timber.Donef("ran migration for %s [%d/%d]", name, i, typeCount)
		i++
	}
	return nil
}
