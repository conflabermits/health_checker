package hcfunc

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"
)

type ShortOutput struct {
	Name       string `json:"name"`
	StatusCode string `json:"statusCode"`
}

func Health_checker_http_req(url string, hostHeader string) string {
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err.Error()
	} else {
		// override Host header, if specified
		if len(hostHeader) > 0 {
			req.Host = hostHeader
		}
		req.Header.Set("user-agent", "health_checker-go")
		//req.Header.Add("X-Forwarded-Proto", "https")
		resp, err := client.Do(req)
		if err != nil {
			return err.Error()
		} else {
			body, err := io.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				return err.Error()
			} else {
				return string(body)
			}
		}
	}
}

func Parse_health_checker_json(jsonString string, depth string) string {

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonMap)
	var response string

	if depth == "dynamic" {
		comp := jsonMap["components"]
		broken_components := make([]map[string]interface{}, 0)
		for _, value := range comp.([]interface{}) {
			value_map := value.(map[string]interface{})
			if value_map["statusCode"] != "OK" {
				broken_components = append(broken_components, value_map)
			}
		}
		if len(broken_components) > 0 {
			jsonMap["broken_components"] = broken_components
		}
		delete(jsonMap, "components")
		dynamic_json, err := json.MarshalIndent(jsonMap, "", "    ")
		if err != nil {
			fmt.Println(err)
		}
		response = string(dynamic_json)
	} else if depth == "short" {
		short_output := ShortOutput{Name: jsonMap["name"].(string), StatusCode: jsonMap["statusCode"].(string)}
		short_json, err := json.MarshalIndent(short_output, "", "    ")
		if err != nil {
			fmt.Println(err)
		}
		response = string(short_json)
	} else if depth == "full" {
		full_json, err := json.MarshalIndent(jsonMap, "", "    ")
		if err != nil {
			fmt.Println(err)
		}
		response = string(full_json)
	}
	return response
}

type ResultDetails struct {
	Success  bool
	URL      string
	Depth    string
	Response string
}

func Web(port string) {

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		reqURL := r.FormValue("url")
		fmt.Println("Request URL: " + reqURL)
		reqDepth := r.FormValue("depth")
		fmt.Println("Request Depth: " + reqDepth)
		httpResponse := Health_checker_http_req(reqURL, "health_checker_web")
		response := Parse_health_checker_json(httpResponse, reqDepth)
		fmt.Println("Response: " + response)

		result := ResultDetails{
			Success:  true,
			URL:      reqURL,
			Depth:    reqDepth,
			Response: response,
		}
		tmpl.Execute(w, result)
	})

	http.ListenAndServe(":"+port, nil)
}
