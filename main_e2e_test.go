package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruckersE2E(t *testing.T) {
	t.Run("get truckers with banking data", func(t *testing.T) {

		expectedResult := "[{\"uuid\":\"7e08e496-1231-4ff4-87d4-a84d2ef53d20\",\"name\":\"João Teles\",\"banking\":{\"bank\":\"Banco do Brasil\",\"agency\":\"52320\",\"account_number\":\"10524-6\"}},{\"uuid\":\"1606aabd-07fc-4484-adf4-5cfcada8e3b8\",\"name\":\"Frederico Santos\",\"banking\":{\"bank\":\"Cooperativa Sicredi\",\"agency\":\"10\",\"account_number\":\"635-10\"}},{\"uuid\":\"082b918b-1fe3-4bb4-a928-50a2cd4234b2\",\"name\":\"Renata Silveira \",\"banking\":{\"bank\":\"Banco do Estado do Rio Grande do Sul\",\"agency\":\"5236\",\"account_number\":\"68552-10\"}},{\"uuid\":\"1b274d80-96ff-4c23-9bd9-e73ef511ebdd\",\"name\":\"Osvaldo Noronha\",\"banking\":{\"bank\":\"Banco do Brasil\",\"agency\":\"96735\",\"account_number\":\"41573-0\"}},{\"uuid\":\"e25dc6e5-dc83-480c-bf35-8d34fc717465\",\"name\":\"Cristiane Fraga\",\"banking\":{\"bank\":\"Banco Itaú\",\"agency\":\"7412\",\"account_number\":\"964\"}}]"

		router := setupRouter()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/truckers", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, expectedResult, w.Body.String())
	})
}
