package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Ma-Leal/cep-race/internal/dto"
	"github.com/Ma-Leal/cep-race/internal/entity"
)

func GetViaCep(ch chan entity.Address, w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("cep")
	if param == "" {
		json.NewEncoder(w).Encode("Wrong URL")
		return
	}
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", param)
	res, err := http.Get(url)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("Failed to get URL %s, error %s", url, err))
		return
	}

	defer res.Body.Close()

	var viaCepInput dto.ViaCepInput

	body, err := io.ReadAll(res.Body)
	if err != nil {
		json.NewEncoder(w).Encode("failed to read body")
		return
	}

	err = json.Unmarshal(body, &viaCepInput)
	if err != nil {
		json.NewEncoder(w).Encode("failed to unmarshal")
		return
	}

	address := entity.NewAddress(
		viaCepInput.Cep,
		viaCepInput.City,
		viaCepInput.State,
		viaCepInput.Neighborhood,
		viaCepInput.Street,
		url)

	ch <- *address
}
