package main

import (
	"github.com/LuizFreitas225/user-manager-go/src/controller/middleware"
)

func main() {
	managerRouter := middleware.CreateRouterManager()
	managerRouter.Start()
}
