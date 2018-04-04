package main

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
	// "github.com/heroku/docker-registry-client/registry"
	// "github.com/docker/distribution/digest"
	// "github.com/docker/distribution/manifest"
	// "github.com/docker/libtrust"
)

const url = "http://localhost:5000"

var regMgr *RegistryManager

func main() {
	// var err error
	regMgr, _ = CreateManager(url, "", "")
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/api/repositories", getRepositoryListHndl)
	// e.GET("/api/repositories/<repoid>/details", getRepoDetails)
	assetHandler := http.FileServer(rice.MustFindBox("web").HTTPBox())
	e.GET("/", echo.WrapHandler(assetHandler))
	e.GET("/*", echo.WrapHandler(assetHandler))
	e.Logger.Fatal(e.Start(":8080"))
}

// Repodetails a
type Repodetails struct {
	Name     string
	TagCount int      `json:"tagcount"`
	TagList  []string `json:"tags"`
}

// func getRepoDetails(c echo.Context) error {
// 	repoName := c.QueryParam("name")
// 	var apiResult struct {
// 		Name string   `json:"name"`
// 		Tags []string `json:"tags"`
// 	}
// 	if err := apiGet("/v2/"+repoName+"/tags/list", &apiResult); err != nil {
// 		return err
// 	}
// 	repoDetails := Repodetails{
// 		Name:     apiResult.Name,
// 		TagList:  apiResult.Tags,
// 		TagCount: len(apiResult.Tags),
// 	}
// 	return c.JSON(http.StatusOK, repoDetails)

// }

func getRepositoryListHndl(c echo.Context) error {
	repolist, _ := regMgr.ListRepositories()
	return c.JSON(http.StatusOK, repolist)
}

var baseurl = "http://localhost:5000"

// func apiGet(path string, result interface{}) error {
// 	println(baseurl + path)
// 	resp, err := http.Get(baseurl + path)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()
// 	if 400 <= resp.StatusCode && resp.StatusCode < 500 {
// 		b, _ := ioutil.ReadAll(resp.Body)
// 		println(string(b))
// 		return errors.New("error in talking to registry")
// 	}
// 	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
// 		return err
// 	}
// 	fmt.Println(result)
// 	return nil
// }
