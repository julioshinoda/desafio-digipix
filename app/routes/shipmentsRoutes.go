package routes
import(
	"github.com/gin-gonic/gin"
	"github.com/desafio-digipix/app/controller"
	"strings"
	"net/http"
)

func Shipments(route *gin.Engine ){


	route.GET("/morcego", func(c *gin.Context) {


		c.JSON(200, gin.H{
			"message": "ta la",
		})
	})


	route.GET("/shipments/zipcode/:zipcode", func(c *gin.Context) {

		zipcode := strings.Trim(strings.Replace(c.Param("zipcode"), "-", "",-1)," ")


		if(len(zipcode) != 8){
			c.JSON(400, gin.H{
				"message": "CEP inválido",
			})
			return 
		}	
	
	    shipment,err := controller.GetZipcode(zipcode)

		if err != nil{
			c.JSON(401, gin.H{
				"message": err.Error(),
			})
			return
		}	

		c.JSON(http.StatusOK, shipment)
		return
	})
	 


}