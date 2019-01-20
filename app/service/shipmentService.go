package service

import(
	"gopkg.in/resty.v1"
	"github.com/desafio-digipix/app/model"
	"errors"
	"os"
)

func GetZipcodeService(zipcode string) (*model.ShipmentModel,error) {

	resp, err := resty.R().
	SetHeader("Accept", "application/json").
	SetHeader("Content-Type","application/json").
	SetHeader("Authorization","Bearer "+os.Getenv("TOKEN")).
	SetResult(&model.ShipmentModel{}).
	Get(os.Getenv("BASE_URL")+"shipments/zipcode/"+zipcode)

    if err != nil {
		return nil, errors.New("Erro no retorno da API")
	}

    return resp.Result().(*model.ShipmentModel),nil
}