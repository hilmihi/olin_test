package main

import (
	"olin_test/driver/router"
	"olin_test/usecase"
)

func main() {
	useCase := usecase.NewAppUseCase()

	router.Init(useCase)
}
