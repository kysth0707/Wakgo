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
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting Server...")

	router := gin.Default()
	router.GET("/link", getLinks)
	router.GET("/links", getLinks)
	router.GET("/:goTo", getGoLink)

	router.Run("localhost:1987")
}

func getLinks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "링크 모음임")
}

func getGoLink(c *gin.Context) {
	goTo := c.Param("goTo")
	c.IndentedJSON(http.StatusOK, gin.H{"니가 날린 거": goTo})
}
