package calc

import (
	"encoding/json"
	"fmt"
	"github.com/gioapp/gel"
	"github.com/marcetin/wingcal/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (w *WingCal) GenerisanjeLinkova(radovi map[int]string) {
	for rad, _ := range radovi {
		w.LinkoviIzboraVrsteRadova[rad] = new(gel.Button)
	}
	return
}

func (w *WingCal) GenerisanjeDugmicaBrisanje(radovi map[int]string) {
	for rad, _ := range radovi {
		w.LinkoviIzboraVrsteRadova[rad] = new(gel.Button)
	}
	return
}

func (w *WingCal) APIpozivIzbornik(komanda string) {
	radovi := map[int]string{}
	jsonErr := json.Unmarshal(APIpoziv(komanda), &radovi)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	w.IzbornikRadova = radovi
}

func (w *WingCal) APIpozivElement(komanda string) {
	rad := &model.WingVrstaRadova{}
	jsonErr := json.Unmarshal(APIpoziv(komanda), &rad)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	w.PrikazaniElement = rad
}

func APIpoziv(komanda string) []byte {
	url := "http://192.168.192.192:9909/" + komanda
	fmt.Println("url", url)
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "wing")
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	if body != nil {
		//defer body.Close()
	}
	return body
}
