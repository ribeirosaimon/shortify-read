package usecase

import (
	"context"
	"errors"

	"github.com/ribeirosaimon/shortify-read/internal/cache"
	"github.com/ribeirosaimon/shortify-read/internal/dto"
	"github.com/ribeirosaimon/tooltip/tlog"
)

type UrlRecord interface {
	FindUrlRecord(ctx context.Context, urlId string) (*dto.UrlRecord, error)
}

type urlRecordUseCase struct {
	cache cache.UrlRecord
}

func NewUrlRecord() *urlRecordUseCase {
	return &urlRecordUseCase{
		cache: cache.NewUrlRecord(),
	}
}

func (u *urlRecordUseCase) FindUrlRecord(ctx context.Context, urlId string) (*dto.UrlRecord, error) {
	urlRecord, err := u.cache.FindById(ctx, urlId)
	if err != nil {
		tlog.Warn("NewUrlRecord.FindUrlRecord", err.Error())
		return nil, err
	}

	if urlRecord == nil {
		tlog.Warn("NewUrlRecord.FindUrlRecord", "url not found")
		return nil, errors.New("urlRecord not found")
	}
	return urlRecord, nil
}
