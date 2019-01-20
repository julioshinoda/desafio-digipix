package controller

import(
	"github.com/desafio-digipix/app/service"
	"github.com/desafio-digipix/app/model"
)

func GetZipcode(zipcode string)(*model.ShipmentModel,error){


   return service.GetZipcodeService(zipcode)



}