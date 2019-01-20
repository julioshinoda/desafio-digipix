package service

import(
	"gopkg.in/resty.v1"
	"github.com/julioshinoda/desafio-digipix/app/model"
	"os"
	"crypto/tls"
)

func GetZipcodeService(zipcode string) (*model.ShipmentModel,error) {


    resty.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })


	resp, err := resty.R().
	SetHeader("Accept", "application/json").
	SetHeader("Content-Type","application/json").
	SetHeader("Authorization","Bearer "+os.Getenv("TOKEN")).
	SetResult(&model.ShipmentModel{}).
	Get(os.Getenv("BASE_URL")+"shipments/zipcode/"+zipcode)

    if err != nil {
		return nil, err
	}

    return resp.Result().(*model.ShipmentModel),nil
}

