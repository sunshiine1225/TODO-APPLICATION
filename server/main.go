package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/sunshiine1225/TODO-APPLICATION/router"	
)

func main(){
	r := router.Router()
	fmt.Println("starting the server no port 9000")
	log.Fatal(http.ListenAndServe(":9000",r))
}