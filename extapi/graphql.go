package extapi

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func makeGraphQLRequest(url, requestString string) ([]byte, error) {
	var str = []byte(requestString)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(str))
	req.Header.Set("Content-Type", "application/graphql")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
