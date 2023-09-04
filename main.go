package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	detail "github.com/shnartho/go-devops-project/details"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking the application's health")

	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")
	hostname, err := detail.GetHostname()
	if err != nil {
		panic(err)
	}
	ip, _ := detail.GetIPAddress()
	fmt.Println(hostname, ip)

	response := map[string]string{
		"hostname": hostname,
		"ip":       ip.String(),
	}
	json.NewEncoder(w).Encode(response)

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)

	log.Println("Server is starting on port 8089...")
	log.Fatal(http.ListenAndServe(":8089", r))
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, you've requested: %s with token: %s\n", r.URL.Path, r.URL.Query().Get("token"))
// }

// func main() {
// 	http.HandleFunc("/", rootHandler)

// 	fs := http.FileServer(http.Dir("static/"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	log.Println("Web server has started")
// 	http.ListenAndServe(":80", nil)
// }

// package main

// import (
// 	"fmt"
// 	"unsafe"

// 	geo "github.com/shnartho/go-devops-project/geometry"

// 	"rsc.io/quote"
// )

// func rectProps(length, width float64) (float64, float64) {
// 	area := length * width
// 	perimeter := 2 * (length * width)
// 	return area, perimeter
// }

// func main() {
// 	var x int = 10
// 	var y, z = 2, 3
// 	name := "Nayem"
// 	fmt.Println("Hello world! a simple devops project")
// 	fmt.Println(quote.Go())
// 	fmt.Println(x, y, z)
// 	fmt.Printf("Type of name is %T and size is %d", name, unsafe.Sizeof(name))

// 	a, p := rectProps(2, 5)
// 	fmt.Printf("Area is %f and perimeter is %f", a, p)

// 	var daysOfTheMonth = map[string]any{
// 		"Jan": 31,
// 		"Feb": map[string]int{"leapyear": 29, "non-leapyear": 28},
// 		"Mar": 31,
// 	}

// 	fmt.Println("\n February month consist of", daysOfTheMonth["Feb"].(map[string]int)["leapyear"], "days")

// 	area := geo.Area(3, 4)
// 	diagonal := geo.Diagonal(5, 6)
// 	fmt.Println(area, diagonal)

// }
