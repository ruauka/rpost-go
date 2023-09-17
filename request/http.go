package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	errUnexpectedHTTPStatus  = errors.New("битый json, кривой URL или сервис упал")
	errUnexpectedContentType = errors.New("no app/json")
)

func MakeRequest(url string, p []byte) (string, error) {
	//payload, _ := json.Marshal(&p)

	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s", url), bytes.NewBuffer(p))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("%s: %w", resp.Status, errUnexpectedHTTPStatus)
	}
	ct := resp.Header.Get("Content-Type")
	if ct != "application/json" {
		return "", fmt.Errorf("%w: %s", errUnexpectedContentType, ct)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	res, _ := json.MarshalIndent(result, "", "      ")
	return string(res), nil
}
