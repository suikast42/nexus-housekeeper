package main

import (
	"fmt"
	"github.com/suikast42/nexus-housekeeper/config"
)

func main() {
	instance := config.AppInstance()

	instance.Logger().Info().Msg(fmt.Sprintf("Starting nexus housekeeper in version %s  with config\n%s ", instance.Version(), instance.Config()))
}
