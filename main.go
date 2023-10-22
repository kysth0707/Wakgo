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
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var goLinks = map[string]int{
	"공지":  2244865,
	"도우미": 9600844,
}

func main() {
	fmt.Println("Starting Server...")

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", mainPage)
	router.GET("/link", getLinks)
	router.GET("/links", getLinks)
	router.GET("/:goTo", getGoLink)

	go testCodes()
	go autoSaveGoLinks()
	router.Run(":1987")
}

// ============ goLinks DB 함수 ============
func autoSaveGoLinks() {
	// ======= 새 파일 생성 =======
	now := time.Now()
	formatedNow := now.Format("2006-01-02_15-04-05")
	fileName := fmt.Sprintf("database\\%s.csv", formatedNow)
	fmt.Println("Saving file at ", formatedNow, " / file : ", fileName)
	// txt := ""
	// for key, value := range goLinks {
	// 	txt := txt + "\n" + fmt.Sprintf("%s, %d", repCommaEnter(key), value)
	// }
	// os.WriteFile(fileName, []byte(txt), 0644)
}

func loadGoLinks() {
	// ======= 파일 읽기 =======
	data, err := os.ReadFile("datas\\1.txt")
	if err != nil {
		return
	}

	fmt.Print(string(data))
}

// ============ , \n 제거 함수 ============
func repCommaEnter(txt string) string {
	nocomma := strings.ReplaceAll(txt, ",", "")
	noenter := strings.ReplaceAll(nocomma, "\n", "")
	return noenter
}

// ============ 라우터로 넘어온 함수 ============
func mainPage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": true, "explain": "메인페이지입니당"})
}

func getLinks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "링크 모음임")
}

func getGoLink(c *gin.Context) {
	goTo := c.Param("goTo")
	if goTo == "favicon.ico" {
		return
	}

	val, ok := goLinks[goTo]
	if !ok {
		// c.HTML(http.StatusOK, "nodata.html", gin.H{
		// 	"goTo": goTo,
		// })
		c.IndentedJSON(http.StatusOK, gin.H{"goTo": goTo, "status": false, "explain": "해당하는 링크가 없어요 ㅠㅠ"})
		return
	}

	c.Redirect(
		http.StatusPermanentRedirect,
		fmt.Sprintf("https://cafe.naver.com/steamindiegame/%d", val),
	)
}

// ============ 코드 테스트 ============

func testCodes() {
	time.Sleep(1 * time.Second)
	testNum := 0

	testIt("http://localhost:1987/links", "\"링크 모음임\"", &testNum)
	testIt("http://localhost:1987/link", "\"링크 모음임\"", &testNum)
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
