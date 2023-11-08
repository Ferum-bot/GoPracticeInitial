package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	url2 "net/url"
)

type Reply struct {
	Name                    string
	Login                   string
	PublicRepositoriesCount int32 `json:"public_repos"`
}

func GitHubInfo(login string) (string, int32, error) {
	url := fmt.Sprint("https://api.github.com/users/", url2.PathEscape(login))
	resp, err := http.Get(url)

	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("StatusCode: %s", resp.Status)
	}

	var githubInfo Reply
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&githubInfo); err != nil {
		return "", 0, err
	}

	return githubInfo.Name, githubInfo.PublicRepositoriesCount, nil
}

func main() {
	resp, err := http.Get("https://api.github.com/users/Ferum-bot")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s", resp.Status)
	}
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	//written, err := io.Copy(os.Stdout, resp.Body)
	//if err != nil {
	//	log.Fatalf("Error: %s", err)
	//}
	//fmt.Printf("Written bytes: %d\n", written)

	var reply Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&reply); err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Printf("%#v\n", reply)

	name, reposCount, err := GitHubInfo("JUSSIAR")
	if err == nil {
		printGithubInfo(name, reposCount)
	}
	name, reposCount, err = GitHubInfo("Ferum-bot")
	if err == nil {
		printGithubInfo(name, reposCount)
	}

	//fmt.Println(GitHubInfo("JUSSIAR"))
	//fmt.Println(GitHubInfo("Ferum-bot"))
}

func printGithubInfo(name string, reposCount int32) {
	fmt.Printf("Name: %s, Public repos count: %d \n", name, reposCount)
}
