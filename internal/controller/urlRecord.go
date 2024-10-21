package controller

import (
	"context"
	"net/http"

	"github.com/ribeirosaimon/shortify-read/config/server"
	"github.com/ribeirosaimon/shortify-read/internal/usecase"
	"github.com/ribeirosaimon/tooltip/response"
)

type UrlRecord struct {
	urlRecord usecase.UrlRecord
}

func NewUrlRecord() *UrlRecord {
	return &UrlRecord{
		urlRecord: server.GetUrlRecordUsecase(),
	}
}

// NewUrlRecord
// @Summary Get URL
// @Description Busca um registro de URL no banco de dados
// @Produce  json
// @Param id path string true "ID da URL"
// @Router /{id} [get]
// @Success 307 {object} dto.UrlRecord "Success"
// @Failure 404 {string} string "URL não encontrada"
func (u *UrlRecord) NewUrlRecord(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	url, err := getIDFromURL(r.URL.Path)
	if err != nil {
		response.BadRequest(w, err)
	}

	record, err := u.urlRecord.FindUrlRecord(ctx, url)

	if err != nil {
		response.BadRequest(w, err)
		return
	}

	http.Redirect(w, r, record.Url, http.StatusTemporaryRedirect)
}
