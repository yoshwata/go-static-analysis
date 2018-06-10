package main

import "fmt"
import "regexp"
import "github.com/xrash/smetrics"
import "github.com/juju/utils/set"

func jaccard(a, b set.Strings) float64 {
	return float64(a.Intersection(b).Size()) / float64(a.Union(b).Size())
}

func ngrams(s string, n int) set.Strings {
	// var result set.Strings
	var result = set.NewStrings(s)
	for i := 0; i < len(s)-n+1; i++ {
		result.Add(s[i : i+n])
	}
	return result
}

// nolint
func main() {
    str1 := `sfunc handleResponse(bodyBytes []byte, statusCode int) ([]byte, error) {
	switch statusCode / 100 {
	case 2:
		return bodyBytes, nil
	case 4:
		res := new(ResponseError)
		err := json.Unmarshal(bodyBytes, res)
		if err != nil {
			return nil, fmt.Errorf("Screwdriver API Response unparseable: status=%d, err=%v", statusCode, err)
		}
		return nil, res
	case 5:
		return nil, fmt.Errorf("Screwdriver API has internal server error: statusCode=%d", statusCode)
	default:
		return nil, fmt.Errorf("Unknown error happen while communicate with Screwdriver API: Statuscode=%d", statusCode)
	}
}`
    str2 := `func handleResponse(res *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Reading response Body from Store API: %v", err)
	}

	switch res.StatusCode / 100 {
	case 2:
		return body, nil
	case 4:
		res := new(ResponseError)
		err = json.Unmarshal(body, res)
		if err != nil {
			return nil, fmt.Errorf("Unparseable error response from Store API: %v", err)
		}
		return nil, res
	case 5:
		return nil, fmt.Errorf("%v: Store API has internal server error", res.StatusCode)
	default:
		return nil, fmt.Errorf("Unknown error happen while communicate with Store API")
	}
}`
    var _ = str2

    str3 := `func (c client) GetCommand() (*Command, error) {
	command := new(Command)
	command.Spec = c.spec
	uri, err := c.commandURL()
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("GET", uri, strings.NewReader(""))
	if err != nil {
		return nil, fmt.Errorf("Failed to create request about command to Store API: %v", err)
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.jwt))
	res, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to get command from Store API: %v", err)
	}
	defer res.Body.Close()
	body, err := handleResponse(res)
	if err != nil {
		return nil, err
	}
	command.Type = parseContentType(res.Header.Get("Content-type"))
	command.Body = body
	return command, nil
}`
    var _ = str3

    str4 := `func newClient(sdAPI, sdToken string) *client {
	return &client{
		baseURL: sdAPI,
		jwt:     sdToken,
		client:  &http.Client{Timeout: timeoutSec * time.Second},
	}
}`
    var _ = str4

    r := regexp.MustCompile(`(?ms)<think>(.*)</think>`)
    const s = `That is
    <think>
    FOOBAR
    </think>`
    fmt.Printf("%#v\n", r.FindStringSubmatch(s))
    fmt.Printf("%#v\n", smetrics.Jaro(str1, str2))
    fmt.Printf("%#v\n", smetrics.Jaro(str1, str3))
    fmt.Printf("%#v\n", smetrics.Jaro(str1, str4))

	a := "Flughafen Leipzig"
	b := "Flughafen zig"

	fmt.Println(ngrams(a, 3))
	fmt.Println(ngrams(b, 3))

	fmt.Println(jaccard(ngrams(str1, 3), ngrams(str1, 3)))
	fmt.Println(jaccard(ngrams(str1, 3), ngrams(str2, 3)))
	fmt.Println(jaccard(ngrams(str1, 3), ngrams(str3, 3)))
	fmt.Println(jaccard(ngrams(str1, 3), ngrams(str4, 3)))
}
