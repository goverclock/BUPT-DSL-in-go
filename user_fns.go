package DSL

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// implement here and add to executor.go/initUserFuncs()

func catfact(s *Script, args []string) {
	type Response struct {
		Fact   string `json:"fact"`
		Length int    `json:"length"`
	}

	response, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	fmt.Println("客服:" + responseObject.Fact)

	s.finish(position{})
}

func dogfact(s *Script, args []string) {
	type Response struct {
		Facts   []string `json:"facts"`
		Success int    `json:"success"`
	}

	response, err := http.Get("https://dog-api.kinduff.com/api/facts")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	fmt.Println("客服:" + responseObject.Facts[0])

	s.finish(position{})
}
