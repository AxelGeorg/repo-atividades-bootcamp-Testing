package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testdoubles/internal/hunter"
	"testdoubles/internal/positioner"
	"testdoubles/internal/prey"
	"testdoubles/internal/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Testa a configuração da presa
func TestHunter_ConfigurePrey_Success(t *testing.T) {
	requestBody := RequestBodyConfigPrey{
		Speed: 10.5,
		Position: &positioner.Position{
			X: 100,
			Y: 200,
		},
	}
	body, _ := json.Marshal(requestBody)

	req, err := http.NewRequest(http.MethodPost, "/hunter/configure-prey", bytes.NewReader(body))
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	ps := positioner.NewPositionerDefault()
	sm := simulator.NewCatchSimulatorDefault(&simulator.ConfigCatchSimulatorDefault{Positioner: ps})

	ht := hunter.NewWhiteShark(hunter.ConfigWhiteShark{
		Speed:     3.0,
		Position:  &positioner.Position{X: 0.0, Y: 0.0, Z: 0.0},
		Simulator: sm,
	})

	pr := prey.NewTuna(0.4, &positioner.Position{X: 0.0, Y: 0.0, Z: 0.0})
	h := NewHunter(ht, pr)

	h.ConfigurePrey(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "A presa está configurada corretamente", recorder.Body.String())
}

// Testa a configuração do caçador
func TestHunter_ConfigureHunter(t *testing.T) {
	requestBody := RequestBodyConfigHunter{
		Speed: 15.0,
		Position: &positioner.Position{
			X: 50,
			Y: 100,
		},
	}
	body, _ := json.Marshal(requestBody)

	req, err := http.NewRequest(http.MethodPost, "/hunter/configure-hunter", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Não foi possível criar a requisição: %v", err)
	}

	recorder := httptest.NewRecorder()

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

	h := NewHunter(ht, pr)

	h.ConfigureHunter()(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "O caçador está configurado corretamente", recorder.Body.String())
}

// Testa quando o tubarão consegue caçar a presa
func TestHandler_Hunt_Success(t *testing.T) {
	shark := hunter.NewHunterMock()
	shark.HuntFunc = func(pr prey.Prey) (duration float64, err error) {
		return 15.0, nil
	}

	hd := NewHunter(shark, nil)
	funcHunt := hd.Hunt()
	expectedBody := "Caçador executado com sucesso - 15.000000"

	req := httptest.NewRequest(http.MethodPost, "/hunt", nil)
	res := httptest.NewRecorder()

	funcHunt(res, req)

	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.Equal(t, expectedBody, res.Body.String())
}

// Testa quando o tubarão não consegue caçar a presa
func TestHandler_Hunt_Failure(t *testing.T) {
	shark := hunter.NewHunterMock()
	shark.HuntFunc = func(pr prey.Prey) (duration float64, err error) {
		return 0, nil
	}

	hd := NewHunter(shark, nil)
	funcHunt := hd.Hunt()
	expectedBody := "Caçador executado com sucesso - 0.000000"

	req := httptest.NewRequest(http.MethodPost, "/hunt", nil)
	res := httptest.NewRecorder()

	funcHunt(res, req)

	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.Equal(t, expectedBody, res.Body.String())
}

// Testa manuseio de erro inesperado
func TestHandler_Hunt_StatusBadRequest(t *testing.T) {
	shark := hunter.NewHunterMock()
	shark.HuntFunc = func(pr prey.Prey) (duration float64, err error) {
		return 0, fmt.Errorf("um erro inesperado")
	}

	hd := NewHunter(shark, nil)
	funcHunt := hd.Hunt()

	req := httptest.NewRequest(http.MethodPost, "/hunt", nil)
	res := httptest.NewRecorder()

	funcHunt(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	assert.Contains(t, res.Body.String(), "Erro: um erro inesperado")
}
