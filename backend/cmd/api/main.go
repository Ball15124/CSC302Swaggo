//go:generate swag init --dir ../.. --generalInfo cmd/api/main.go --output ./docs
package main

import "backend/cmd/api/di"

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @host localhost:8000
// @BasePath /
// @schemes http
// @produce application/json
// @consumes application/json
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email
// @license.name Apache 2.0
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
