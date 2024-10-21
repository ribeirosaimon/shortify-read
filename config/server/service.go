package server

import "github.com/ribeirosaimon/shortify-read/internal/usecase"

func WithUrlUseCase(urlUseCase usecase.UrlRecord) Option {
	return func(s *services) {
		s.urlUseCase = urlUseCase
	}
}
