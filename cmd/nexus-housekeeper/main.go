package main

import (
	"fmt"
	"github.com/suikast42/nexus-housekeeper/app"
)

func main() {
	instance := app.AppInstance()

	instance.Logger().Info().Msg(fmt.Sprintf("Starting nexus housekeeper in version %s  with config\n%s ", instance.Version(), instance.Config()))

	instance.DeleteOldContainers()
	//configuration, _, err := instance.ClientNexus().ComponentsApi.GetComponents(instance.Context(), "dockerLocal", &nexus3.ComponentsApiGetComponentsOpts{})

}
