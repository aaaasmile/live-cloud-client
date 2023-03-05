package cmd

import (
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
	url := conf.Current.ServiceURL + "Api/List"
	lr := ListReq{
		Path: remotePath,
	}

	req, err := getRequestWithAuthHeader(url, lr)
	if err != nil {
		return err
	}
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
