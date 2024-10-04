package routes

import(
	"net/http"
	"github.com/Consuelo-Fraga/loja-digiport-backend/controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	
)

func HandleRequest(){
	route := mux.NewRouter()

	route.HandlerFunc("/produtos", controllers.BuscaProdutosHandler).Methods("GET")
	route.HandlerFunc("/produtos", controllers.BuscaProdutosHandler).Methods("GET")
	route.HandlerFunc("/produtos", controllers.BuscaProdutosHandler).Methods("POST")
	route.HandlerFunc("/produtos/{id}", controllers.RemoveProdutosHandler).Methods("DELETE")
	route.HandlerFunc("/produtos", controllers.AtualizarProdutosHandler).Methods("PUT")



c := cors.New(cors.Options{
	AllowedOrigins: []string {"*"},
	AllowedMethods: []string {"GET", "POST", "DELETE", "PUT"},
	AllowedHeaders: []string {"Content-Type"},
	AllowedCredentials: true,
})

handler := c.Hamdler(route)
}


http.ListenAndServer(":8080", handler)
//http://localhost:8080

