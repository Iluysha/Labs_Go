package main
// Import necessary packages
import ( 
	"fmt" 
	"net/http"
	"time" 
	"encoding/json"
	"log"
)

func handleStart(resWr http.ResponseWriter, req *http.Request) {
	resWr.Write([]byte("Hello from server!"))
	fmt.Println("Page success")
}

func handleCurrentTime(resWr http.ResponseWriter, req *http.Request) {
	jsonCur, err := json.Marshal(map[string]string{ "time": time.Now().Format(time.RFC3339) })
	if err != nil { // Error
		log.Println(err)
		return
	}
	fmt.Println("Time success")
	fmt.Println(time.Now().Format(time.RFC3339))
	resWr.Write(jsonCur) // Sending time
}

func main() {
	// Request handles
	http.HandleFunc("/", handleStart)
	http.HandleFunc("/time", handleCurrentTime)
	// Deffered call
	defer http.ListenAndServe(":8795", nil)
	// Server status message
	fmt.Println("Server started successfully and listen on 8795 port")
}
