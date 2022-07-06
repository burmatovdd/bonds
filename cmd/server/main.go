package main

import "modules/internal/server/api"

func main() {
	service := api.ApiStartService{}
	service.HandleRequest()
}
