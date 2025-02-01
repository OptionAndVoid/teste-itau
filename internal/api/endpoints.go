package api

import (
	"net/http"
)

type Transacao interface {
	GetTransacao(w http.ResponseWriter, r *http.Request)
	PostTransacao(w http.ResponseWriter, r *http.Request)
	DeleteTransacao(w http.ResponseWriter, r *http.Request)
}

type TransacaoController struct{}

func (t *TransacaoController) GetTransacao(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"route": "GetTransacao"})
}

func (t *TransacaoController) PostTransacao(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"route": "PostTransacao"})
}

func (t *TransacaoController) DeleteTransacao(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"route": "DeleteTransacao"})
}
