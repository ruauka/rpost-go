package parse

import (
	"encoding/json"
	"fmt"
	"os"

	"r-post/request"
)

// ReadFile - Чтение файла.
func ReadFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("file read err: %w", err)
	}

	return file, nil
}

// JsonParse - Парсинг входящего json в структуру.
func JsonParse(payload []byte) (map[string]interface{}, error) {
	p := map[string]interface{}{}

	if err := json.Unmarshal(payload, &p); err != nil {
		return nil, fmt.Errorf("payload unmarshal err: %w", err)
	}

	return p, nil
}

func GetResp(url, path string) (string, error) {
	// Чтение файла.
	file, err := ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("не нашел такой файл")
	}

	//// Парсинг входящего json в структуру.
	//c, err := JsonParse(file)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	////fmt.Printf("%+v\n", c)

	return request.MakeRequest(url, file)
}
