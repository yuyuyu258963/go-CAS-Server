package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 设置请求的合法域名
func pathVerify(c *gin.Context) {
	// fmt.Println(c.Request.RequestURI)
}

// 加载CAS登录页面
func casLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// 检测登录状态
func casCheckLoadingState(c *gin.Context) {
	cookie, err := c.Cookie(CasLoinTicketName)
	fmt.Println(cookie)
	// service := c.DefaultQuery("service", "")
	loginState := CheckLoadingState(cookie)
	if err == nil && loginState == nil {
		// 这里的ticket应该重新生成
		// c.Redirect(http.StatusMovedPermanently, service+"?ticket="+cookie)
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功了",
			"ticket":  cookie,
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请重新登陆",
		})
	}
}

// cas 登入
func casLogin(c *gin.Context) {
	params := PraseLoginInfoFromRequest(c)
	isAllowedDomain := checkAllowedDomains(params.Service)
	isAllowedUser := params.Check()

	if isAllowedDomain && isAllowedUser {
		// 这里应该有个生成 ticket 的操作 目前以 uuid 替代
		uidTicket := getUUid()
		InsertMapData(params.Service, uidTicket, params.Username)
		// 登陆成功时向登陆页面写入 cookie
		c.SetCookie(CasLoinTicketName, uidTicket, 100000, "/cas", ".ywhabc.com", false, false)
		c.Redirect(http.StatusMovedPermanently, params.Service+"?ticket="+uidTicket)
	} else if isAllowedUser && !isAllowedDomain {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "域名不合法",
		})
	} else {
		c.String(http.StatusBadRequest, "登录出错")
	}
}

// cas 登出
func casLogout(c *gin.Context) {
	cookie, err := c.Cookie(CasLoinTicketName)
	message := "成功登出"
	state := RemoveLoginHistory(cookie)
	if state != nil || err != nil {
		message = "登出失败"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

// 非法页面跳转到登录页面
func redirect(c *gin.Context) {
	fmt.Println("非法页面")
	c.Redirect(http.StatusMovedPermanently, "/cas/loginPage")
}

// 登出接口
func router() {
	r := gin.Default()
	r.LoadHTMLFiles("./html/index.html")
	g := r.Group("/cas")
	g.Use(pathVerify)

	g.GET("/loginPage", casLoginPage)
	g.GET("/checkLoginState", casCheckLoadingState)
	g.GET("/login", casLogin)
	g.GET("/logout", casLogout)
	r.GET("/", redirect)
	g.GET("/", redirect)

	r.Run(":889")
}
