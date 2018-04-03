package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"

	// "github.com/heroku/docker-registry-client/registry"
	"github.com/rusenask/docker-registry-client/registry"
	// "github.com/docker/distribution/digest"
	// "github.com/docker/distribution/manifest"
	// "github.com/docker/libtrust"
)

func main() {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/api/repository", getRepositoryListHndl)
	e.GET("/api/repository/details", getRepoDetails)
	e.Static("/", ".")

	url := "http://localhost:5000"
	username := "" // anonymous
	password := "" // anonymous
	hub, _ := registry.New(url, username, password)

	fmt.Println(hub.Repositories())
	e.Logger.Fatal(e.Start(":8080"))
}

// Repodetails a
type Repodetails struct {
	Name     string
	TagCount int      `json:"tagcount"`
	TagList  []string `json:"tags"`
}

func getRepoDetails(c echo.Context) error {
	repoName := c.QueryParam("name")
	var apiResult struct {
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	}
	if err := apiGet("/v2/"+repoName+"/tags/list", &apiResult); err != nil {
		return err
	}
	repoDetails := Repodetails{
		Name:     apiResult.Name,
		TagList:  apiResult.Tags,
		TagCount: len(apiResult.Tags),
	}
	return c.JSON(http.StatusOK, repoDetails)

}

func getRepositoryListHndl(c echo.Context) error {
	var result map[string][]string
	if err := apiGet("/v2/_catalog", &result); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result["repositories"])
}

var baseurl = "http://localhost:5000"

func apiGet(path string, result interface{}) error {
	println(baseurl + path)
	resp, err := http.Get(baseurl + path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if 400 <= resp.StatusCode && resp.StatusCode < 500 {
		b, _ := ioutil.ReadAll(resp.Body)
		println(string(b))
		return errors.New("error in talking to registry")
	}
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
