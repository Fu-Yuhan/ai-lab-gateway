package main 
 
import ( 
	// "net/http" 
	"net/http/httputil" 
	"net/url" 
 
	"strings"

	"github.com/gin-gonic/gin"  
) 
 
func main() { 
	// 创建一个 Gin 路由器 
	r := gin.Default() 
 
	// 定义目标 URL 
	apiURL, _ := url.Parse("URL_ADDRESS")  
	staticURL, _ := url.Parse("URL_ADDRESS")
 
	// 创建一个反向代理 
	apiProxy := httputil.NewSingleHostReverseProxy(apiURL)
	staticProxy := httputil.NewSingleHostReverseProxy(staticURL)
 
	// 定义一个中间件，用于处理所有请求 
	r.Any("/*path", func(c *gin.Context) { 
		// 修改请求的 Host 和 URL 
		path := c.Request.URL.Path
		parts := strings.Split(path, "/")
		
		targetURL := staticURL
			proxy := staticProxy
		if parts[1] == "api" {
			targetURL = apiURL
			proxy = apiProxy
		}
		c.Request.Host = targetURL.Host 
		c.Request.URL.Scheme = targetURL.Scheme 
		c.Request.URL.Host = targetURL.Host 
 
		// 转发请求 
		proxy.ServeHTTP(c.Writer, c.Request) 
	} )
 
	// 启动 Gin 服务 
	r.Run(":8080") 
} 