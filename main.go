// 아이 코드 부끄러워

/*
	          .---.                  ,-.
	         /. ./|              ,--/ /|
	     .--'.  ' ;            ,--. :/ |               ,---.
	    /__./ \ : |            :  : ' /    ,----._,.  '   ,'\
	.--'.  '   \' .  ,--.--.   |  '  /    /   /  ' / /   /   |

/___/ \ |    ' ' /       \  '  |  :   |   :     |.   ; ,. :
;   \  \;      :.--.  .-. | |  |   \  |   | .\  .'   | |: :

	\   ;  `      | \__\/: . . '  : |. \ .   ; ';  |'   | .; :
	 .   \    .\  ; ," .--.; | |  | ' \ \'   .   . ||   :    |
	  \   \   ' \ |/  /  ,.  | '  : |--'  `---`-'| | \   \  /
	   :   '  |--";  :   .'   \;  |,'     .'__/\_: |  `----'
	    \   \ ;   |  ,     .-./'--'       |   :    :
	     '---"     `--`---'                \   \  /
	                                        `--`-'

go로 만든 왁물원 go

http://wakgo.kro.kr/공지
*/

package main

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var goLinks = make(map[string]int)

func main() {
	fmt.Println("Starting Server...")
	goLinks["공지"] = 2244865

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/link", getLinks)
	router.GET("/links", getLinks)
	router.GET("/:goTo", getGoLink)

	go testCodes()
	router.Run("localhost:1987")
}

// ============ 라우터로 넘어온 함수 ============

func getLinks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "링크 모음임")
}

func getGoLink(c *gin.Context) {
	goTo := c.Param("goTo")
	goTo = html.UnescapeString(goTo)
	val, ok := goLinks[goTo]
	if !ok {
		c.HTML(http.StatusOK, "nodata.html", gin.H{
			"goTo": goTo,
		})
		return
	}

	fmt.Println(val)
	c.Redirect(
		http.StatusMovedPermanently,
		"https://cafe.naver.com/steamindiegame/"+string(val),
	)
	// c.IndentedJSON(http.StatusOK, gin.H{"니가 날린 거": goTo, "안녕": val})
}

// ============ 코드 테스트 ============

func testCodes() {
	time.Sleep(1 * time.Second)
	testNum := 0

	testIt("http://localhost:1987/links", "\"링크 모음임\"", &testNum)
}

func testIt(url string, targetResponse string, testNumAddr *int) {
	*testNumAddr += 1
	response, _ := http.Get(url)
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK && targetResponse == string(body) {
		fmt.Print(*testNumAddr)
		fmt.Print(". Pass / ")
		fmt.Println(url)
	} else {
		fmt.Print(*testNumAddr)
		fmt.Print(". Fail / ")
		fmt.Println(url)
	}
	// fmt.Println(response.StatusCode, http.StatusOK)
	// fmt.Println(targetResponse, string(body))
}
