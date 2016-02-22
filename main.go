package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

var (
	host      = flag.String("h", "", "Set mirage Host URI")
	subDomain = flag.String("s", "", "subdomain")
)

type MirageInfo struct {
	Id        string `json:"id"`
	ShortId   string `json:"short_id"`
	SubDomain string `json:"subdomain"`
	Branch    string `json:"branch"`
	Image     string `json:"image"`
	IpAddress string `json:"ipaddress"`
}

type List struct {
	Result []MirageInfo `json:"result"`
}

func main() {
	flag.Parse()

	if *host == "" {
		log.Fatalln("please your mirage host")
	}

	if *subDomain == "" {
		log.Fatalln("please set subdomain")
	}

	re := regexp.MustCompile(*subDomain)

	command := flag.Args()[0]

	list, err := getList()
	if err != nil {
		log.Fatalln(err.Error())
	}

	var info MirageInfo
	for _, v := range list.Result {
		if re.MatchString(v.SubDomain) {
			info = v
		}
	}

	fmt.Printf("docker exec -it %s %s \n", info.ShortId, command)
}

func getList() (list List, err error) {
	url := fmt.Sprintf("%s/api/list", *host)

	res, err := http.Get(url)
	if err != nil {
		return List{}, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&list)
	if err != nil {
		return List{}, err
	}

	return list, nil
}
