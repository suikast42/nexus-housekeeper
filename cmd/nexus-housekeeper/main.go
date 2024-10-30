package main

import (
	"fmt"
	"github.com/suikast42/nexus-housekeeper/app"
)

func main() {
	instance := app.AppInstance()

	instance.Logger().Info().Msg(fmt.Sprintf("Starting nexus housekeeper in version %s  with config\n%s ", instance.Version(), instance.Config()))

	err := instance.DeleteOldContainers()
	if err != nil {
		instance.Logger().Err(err).Msg("Can't delete images from nexus")
	}
	//configuration, _, err := instance.ClientNexus().ComponentsApi.GetComponents(instance.Context(), "dockerLocal", &nexus3.ComponentsApiGetComponentsOpts{})

}
