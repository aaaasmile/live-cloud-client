package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"live-cloud-client/conf"
	"log"
	"net/http"
)

type ResType int

const (
	RTFile = iota
	RTDir
)

type Resource struct {
	Name string
	Type ResType
}

type ListReq struct {
	Path string
}

type ListResp struct {
	Resources []Resource
}

func List(remotePath string) error {
	log.Println("List for ", remotePath)
	tk, err := getReqToken()
	if err != nil {
		return err
	}
	var bearer = "Bearer " + tk
	url := conf.Current.ServiceURL + "Api/List"
	lr := ListReq{
		Path: remotePath,
	}
	b, err := json.Marshal(lr)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(b)
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	rawbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("error on List request: %s", string(rawbody))
	}
	listResp := ListResp{}
	if err := json.Unmarshal(rawbody, &listResp); err != nil {
		return err
	}

	fmt.Println("*** list: ", listResp)
	return nil
}
