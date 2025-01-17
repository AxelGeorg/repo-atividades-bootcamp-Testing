//go:build integration

package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"testdoubles/internal/handler"
	"testdoubles/internal/hunter"
	"testdoubles/internal/positioner"
	"testdoubles/internal/prey"
	"testdoubles/internal/simulator"

	"github.com/stretchr/testify/assert"
)

// setupServer deve ser ajustado para chamar o método ConfigureHunter
func setupServer() *http.ServeMux {
	ps := positioner.NewPositionerDefault()
	sm := simulator.NewCatchSimulatorDefault(&simulator.ConfigCatchSimulatorDefault{
		Positioner: ps,
	})

	ht := hunter.NewWhiteShark(hunter.ConfigWhiteShark{
		Speed:     3.0,
		Position:  &positioner.Position{X: 0.0, Y: 0.0, Z: 0.0},
		Simulator: sm,
	})

	pr := prey.NewTuna(0.4, &positioner.Position{X: 0.0, Y: 0.0, Z: 0.0})
	h := handler.NewHunter(ht, pr)

	mux := http.NewServeMux()
	mux.HandleFunc("/hunter/configure-prey", h.ConfigurePrey)
	mux.HandleFunc("/hunter/configure-hunter", h.ConfigureHunter())
	mux.HandleFunc("/hunt", h.Hunt())

	return mux
}

func TestIntegration_ConfigurePrey(t *testing.T) {
	requestBody := handler.RequestBodyConfigPrey{
		Speed: 10.5,
		Position: &positioner.Position{
			X: 100,
			Y: 200,
		},
	}
	body, _ := json.Marshal(requestBody)

	mux := setupServer()
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := http.Post(server.URL+"/hunter/configure-prey", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Não foi possível fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Erro ao ler a resposta: %v", err)
	}

	assert.Equal(t, "A presa está configurada corretamente", string(responseBody))
}
