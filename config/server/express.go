package server

import "github.com/ribeirosaimon/shortify-read/internal/usecase"

// Services application
// is nice only apply interfaces
type services struct {
	urlUseCase usecase.UrlRecord
}

// appServerService is my app server
// This variable cannot be null, it is where all the application services are found
// It makes testing easier, where I just need to map out what is needed
var appServerService = &services{}

// Option is the function with handler my servicees
type Option func(*services)

// NewServices
// is the only place where I can initialize the variable appServerService
// because this my services was private
func NewServices(opts ...Option) *services {
	if appServerService == nil {
		appServerService = &services{}
	}

	for _, opt := range opts {
		opt(appServerService)
	}
	return appServerService
}

// functions to get a mapped service

func GetUrlRecordUsecase() usecase.UrlRecord {
	return appServerService.urlUseCase
}
