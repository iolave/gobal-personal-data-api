package main
import (
	"fmt"
	// "github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	// "strings"
	"net/url"
	"regexp"
	// "encoding/json"
	// "sync"
)

// type TestData struct {
//     Data string `json:"data"`
// }

// var testData = TestData{Data: "some string"}

// postAlbums adds an album from JSON received in the request body.
// func testFunction(c *gin.Context) {
// 	data := url.Values{
//         "action":       {"search_by_rut"},
//         "occupation": {"19408138-3"},
//     }

// 	resp, err := http.PostForm("https://rutificador.org/backend.php",data)

// 	var res map[string]interface{}

//     json.NewDecoder(resp.Body).Decode(&res)

//     fmt.Println(res["form"])
//     fmt.Println(resp.Body)
// 	fmt.Println(err)

//     c.IndentedJSON(http.StatusOK, testData)
// }

func main() {

	// router := gin.Default()

    // router.GET("/api", testFunction)

    // router.Run("localhost:8080")

	getRutificadorData()
}

//checks and prints a message if a website is up or down
func getRutificadorData() {
	const myUrl = "https://rutificador.org/backend.php"

	// form data
	data := url.Values{}
	data.Add("action", "search_by_name")
	data.Add("name", "ignacio olave")

	resp, err := http.PostForm(myUrl, data)
	if err != nil {panic(err)}


	content, _ := ioutil.ReadAll(resp.Body)

	var str = string(content)
	// fmt.Println(str)


	re := regexp.MustCompile(`<tr>(.+?)</tr>`)
	submatchall := re.FindAllStringSubmatch(str,-1)

	fmt.Println(submatchall)
	// var header = ""
	// for _, element := range submatchall {
	// 	header = element[1]
	// }
	// fmt.Println(header)

	// re2 := regexp.MustCompile(`<th.*?>(.+?)</th>`)
	// submatchall2 := re2.FindAllStringSubmatch(header,-1)
	// fmt.Println(submatchall2)
	for _, element := range submatchall {
		fmt.Println(element[1])
	}

	defer resp.Body.Close()
}