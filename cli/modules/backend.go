package modules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/logrusorgru/aurora"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	BackendURL = "http://127.0.0.1:8080"
	QueryURL   = BackendURL + "/query"
	LoginURL   = BackendURL + "/login"
)

type Backend struct {
	Name  string
	token string
}

var fctPtr = map[string]func(*Backend, []string) error{
	"login": func(b *Backend, cmd []string) error { return b.Login(cmd) },
	"query": func(b *Backend, cmd []string) error { return b.QueryData(cmd) },
}

func NewBackendModule() Executor {
	return &Backend{Name: "Backend"}
}

func (b *Backend) String() string {
	return b.Name
}

func (b *Backend) Login(cmd []string) error {
	if len(cmd) != 2 {
		return fmt.Errorf("wrong number of arguments : {%s}", strings.Join(cmd, " "))
	}
	u, _ := url.Parse(LoginURL)
	values, _ := url.ParseQuery(u.RawQuery)
	values.Set("login", cmd[0])
	values.Set("pass", cmd[1])
	u.RawQuery = values.Encode()
	resp, err := http.Post(fmt.Sprintf("%v", u), "text/plain", nil)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	b.token = string(bytes)
	fmt.Println(aurora.Green(cmd[0] + " logged in !"))
	return nil
}

func (b *Backend) QueryData(cmd []string) error {
	if len(cmd) == 0 {
		return b.JsonRequest("", "")
	}
	if cmd[0] == "exporter" {
		if len(cmd) == 3 {
			return b.JsonRequest(QueryURL+"/"+cmd[1], cmd[2])
		} else {
			return b.JsonRequest(QueryURL+"/"+cmd[1], "")
		}
	}
	if len(cmd) > 2 {
		return fmt.Errorf("wrong number of arguments : {%s}", strings.Join(cmd, " "))
	}
	return b.JsonRequest(QueryURL, cmd[0])
}

func (b *Backend) Execute(cmd []string) error {
	fmt.Printf("Executing module : %v\n", b)
	if fct, ok := fctPtr[cmd[0]]; !ok {
		return fmt.Errorf("instruction '%s' for module '%s' unknown", cmd[0], b)
	} else {
		return fct(b, cmd[1:])
	}
}

func (b *Backend) LoggedRequest(url string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+b.token)
	return http.DefaultClient.Do(req)
}

func (b *Backend) JsonRequest(route, dataValue string) error {
	u, _ := url.Parse(route)
	values, _ := url.ParseQuery(u.RawQuery)
	values.Set("datas", dataValue)
	u.RawQuery = values.Encode()
	resp, err := b.LoggedRequest(fmt.Sprintf("%v", u))
	if err != nil {
		return err
	}
	datas, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var out bytes.Buffer
	err = json.Indent(&out, datas, "", " ")
	if err != nil {
		return err
	}
	_, err = out.WriteTo(os.Stdout)
	fmt.Println()
	return err
}
