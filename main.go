package main

import (
	"github.com/naufalfmm/dayatani-farmer-api/app"
)

//	@title						Dayatani Farmer API
//	@version					1.0
//	@description				Prototype of farmer API
//
//	@securityDefinitions.basic	BasicAuth
//	@in							header
//	@name						Authorization
func main() {
	app.Init().Run()
}
