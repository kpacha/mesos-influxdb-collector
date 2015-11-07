package reader

import (
	"io/ioutil"
	"log"
	"net/http"
)

func ReadUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error connecting to", url)
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading from", resp.Body)
		return body, err
	}
	return body, nil
}
