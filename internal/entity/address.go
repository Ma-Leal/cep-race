package entity

type Address struct {
	Cep          string `json:"cep"`
	City         string `json:"cidade"`
	State        string `json:"estado"`
	Neighborhood string `json:"bairro"`
	Street       string `json:"logradouro"`
	Source       string `json:"url"`
}

func NewAddress(cep, city, state, neighborhood, street, source string) *Address {
	return &Address{
		Cep:          cep,
		City:         city,
		State:        state,
		Neighborhood: neighborhood,
		Street:       street,
		Source:       source,
	}

}
