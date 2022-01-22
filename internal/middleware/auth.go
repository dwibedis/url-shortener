package middleware

import (
	"dwibedis/url-shortener/internal/repository"
	"github.com/gin-gonic/gin"
	"log"
)

const AllowedIpRateLimit = 4

func Authenticate(c *gin.Context) {
	log.Println("Checking the Auth Eligibility!!")
	/**
	disabling auth segment!!
	 */
	c.Next()
	//ip:= c.ClientIP()
	//log.Printf("cllientIp: %v\n", ip)
	//if !ipRateLimit(c, ip) {
	//	c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"error": "rate limit exceeded, please mailto dwibedisatyaprakash@gmail.com to continue!!"})
	//}
	//err := repository.IncrIpRateLimit(c, ip)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"error": "something went wrong, please mailto dwibedisatyaprakash@gmail.com to continue!!"})
	//}
	//c.Next()
}

func ipRateLimit(c * gin.Context, ip string) bool {
	ipHitsInLastMin := repository.GetIpRate(c, ip)
	log.Printf("ip %v, hits till now: %v\n", ip, ipHitsInLastMin)
	return ipHitsInLastMin <= AllowedIpRateLimit
}

/**
method checks if the number of entries in redis > max, then auto fail over!!
 */
func storageFailOver(c *gin.Context) bool {
	return true
}