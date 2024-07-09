package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"github.com/canyouhearthemusic/proxy-server/internal/models"
	"github.com/canyouhearthemusic/proxy-server/pkg/utils"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

var (
	requests sync.Map
)

func HandleProxyRequest(w http.ResponseWriter, r *http.Request) {
	var proxyReq models.ProxyRequest
	if err := json.NewDecoder(r.Body).Decode(&proxyReq); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	proxyReq.MethodToUpper()

	client := &http.Client{}
	req, err := http.NewRequest(proxyReq.Method, proxyReq.URL, nil)
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer(err))
		return
	}

	for key, value := range proxyReq.Headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer(err))
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		render.Render(w, r, utils.ErrInternalServer(err))
		return
	}

	proxyRes := models.ProxyResponse{
		ID:     uuid.New().String(),
		Status: resp.StatusCode,
		Headers: func() map[string]string {
			headers := make(map[string]string)
			for k, v := range resp.Header {
				headers[k] = v[0]
			}
			return headers
		}(),
		Length: len(body),
	}

	requests.Store(proxyRes.ID, proxyReq)
	render.JSON(w, r, proxyRes)
}
