package main

import (
	"azt.com/api-sample/infra/router"
)

func init() {

}
func main() {
	r := router.CreateRoute()
	r.Run(":8080")
}
