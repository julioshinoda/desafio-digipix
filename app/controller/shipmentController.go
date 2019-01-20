package controller

import(
	"github.com/julioshinoda/desafio-digipix/app/service"
	"github.com/julioshinoda/desafio-digipix/app/model"
)

func GetZipcode(zipcode string)(*model.ShipmentModel,error){


   return service.GetZipcodeService(zipcode)



}