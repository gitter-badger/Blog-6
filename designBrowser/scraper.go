package designBrowser

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetMore(w http.ResponseWriter, r *http.Request, q string) {
	url := fmt.Sprintf("http://design-seeds.com/index.php/P%v", q)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	w.Write(body)
}
