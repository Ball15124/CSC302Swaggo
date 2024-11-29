//go:generate swag init --dir ../.. --generalInfo cmd/api/main.go --output docs
package main

import "backend/cmd/api/di"

// gin-swagger middleware
// swagger embed files

// @title Swinging Example API
// @version 2.5
// @description Kuy Hee Kuy Hee
func main() {
	app, err := di.InitializedApp()
	if err != nil {
		return
	}
	err = app.RunApp()
	if err != nil {
		return
	}
}
