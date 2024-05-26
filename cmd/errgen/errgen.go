package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"hash/fnv"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	generatedPath = "pkg/errors/ecode.go"
	configPath    = "configs/errors.json"
)

type errorJson struct {
	Err        string            `json:"error"`
	HttpStatus int               `json:"http_status"`
	Msg        map[string]string `json:"msg"`
}

var tplHelperFunc = `
// GetHTTPErrData returning http status, error message and error code
func GetHTTPErrData(err error) (int, uint32, string) {
	code, _ := GetCode(err)

	return getHttpStatus(code), code, getErrText(code)
}

func getErrText(code uint32) string {
	if _, ok := HttpResponse[code]; !ok {
		return unknown
	}

	if _, ok := HttpResponse[code]["text"]; !ok {
		return unknown
	}

	if status, ok := HttpResponse[code]["text"].(string); ok {
		return status
	}

	return unknown
}

func getHttpStatus(code uint32) int {
	if _, ok := HttpResponse[code]; !ok {
		return http.StatusInternalServerError
	}

	if _, ok := HttpResponse[code]["status"]; !ok {
		return http.StatusInternalServerError
	}

	if status, ok := HttpResponse[code]["status"].(int); ok {
		return status
	}

	return http.StatusInternalServerError
}`

func main() {
	ecodePkg := "// Code generated, DO NOT EDIT.\npackage errors"

	errJson, err := parseErrJson()
	if err != nil {
		log.Fatalf("failed %v", err)
	}

	if len(errJson) == 0 {
		log.Fatal("file must contain at least one error")
	}

	formatedErrWebJson, err := format.Source([]byte(fmt.Sprintf("%s\n\n%s", ecodePkg, genErrData(errJson))))
	if err != nil {
		log.Fatalf("failed format: %v", err)
	}

	if err := saveFile(
		generatedPath,      // path
		formatedErrWebJson, // data
	); err != nil {
		log.Fatalf("failed save %s: %v", generatedPath, err)
	}
}

func saveFile(filepath string, data []byte) error {
	if err := os.WriteFile(filepath, data, 0644); err != nil {
		return fmt.Errorf("write in file=%s: %v", filepath, err)
	}

	return nil
}

func parseErrJson() ([]errorJson, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("read errors.json: %v", err)
	}

	var errJson []errorJson

	if err := json.Unmarshal(data, &errJson); err != nil {
		return nil, fmt.Errorf("parse errors.json: %v", err)
	}

	return errJson, nil
}

func genErrData(errJson []errorJson) string {
	var ecode, errs string

	ecodeHttpStatus := make([]string, 0, len(errJson))
	eWebJson := make(map[uint32]map[string]string, len(errJson))

	for i := range errJson {
		errJson[i].Err = strings.TrimSpace(errJson[i].Err)

		if !validateErrText(errJson[i].Err) {
			log.Fatalf("invalid error: %s", errJson[i].Err)
		}

		code := getUniqKey(errJson[i].Err)

		if _, ok := eWebJson[code]; ok {
			log.Fatalf("duplicate error: %s", errJson[i].Err)
		}

		if errJson[i].Msg == nil {
			errJson[i].Msg = make(map[string]string, 1)
		}

		errText := upperFirst(errJson[i].Err)

		if eText, ok := errJson[i].Msg["en"]; !ok {
			errJson[i].Msg["en"] = errText
		} else if eText != "" {
			errText = eText
		}

		eWebJson[code] = errJson[i].Msg

		ecodeHttpStatus = append(
			ecodeHttpStatus,

			fmt.Sprintf(
				`%d: {"status": %d, "text": "%s"}`,

				code,
				getHttpStatus(errJson[i].HttpStatus),
				errText,
			),
		)

		eName := generateErrVar(errJson[i].Err)

		ecode = fmt.Sprintf(
			"%sCode%s uint32 = %d\n",

			ecode,
			eName,
			code,
		)

		errs = fmt.Sprintf(
			"%s%s error = New(\"%s\", Code%s)\n",

			errs,
			eName,
			errJson[i].Err,
			eName,
		)
	}

	errWebJson, err := json.MarshalIndent(eWebJson, "", "    ")
	if err != nil {
		log.Fatalf("failed encode web errors: %v", err)
	}

	ecode = fmt.Sprintf(`
		import "net/http"

		const unknown = "unknown"

		// Example usage for output errors.response.json
		// echo.GET("/errors", func(ctx echo.Context) error {
		//   return ctx.JSONBlob(http.StatusOK, errors.ResponseList)
		// })
		var ResponseList []byte = []byte(`+"`%s`"+`) // %s

		// Error codes
		const (%s)

		// Error Variables
		var (%s)

		// Hash map data by error codes
		var HttpResponse = map[uint32]map[string]any{`+"\n%s,\n}\n%s\n",
		string(errWebJson),
		"`", // cheating vscode
		ecode,
		errs,
		strings.Join(ecodeHttpStatus, ",\n"),
		tplHelperFunc,
	)

	return ecode
}

func generateErrVar(str string) string {
	return strings.Join(strings.Split(strings.Title(strings.ToLower(str)), " "), "")
}

func getHttpStatus(status int) int {
	if status < 100 || status > 999 {
		return 500
	}

	return status
}

func validateErrText(errStr string) bool {
	if errStr == "" {
		return false
	}

	word := []rune(errStr)

	if _, err := strconv.Atoi(string(word[0])); err == nil {
		return false
	}

	return true
}

func upperFirst(errStr string) string {
	words := strings.Split(errStr, " ")
	word := []rune(words[0])

	firstWord := fmt.Sprintf("%s%s", strings.ToUpper(string(word[0])), string(word[1:]))

	if len(words) == 1 {
		return firstWord
	}

	return fmt.Sprintf("%s %s", firstWord, strings.Join(words[1:], " "))
}

func getUniqKey(errStr string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(errStr))

	return hash.Sum32()
}
