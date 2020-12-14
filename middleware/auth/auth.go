package auth

import "github.com/gin-gonic/gin"

func GetUserIdAndNickname(c *gin.Context) (int64, string) {

	userID, found := c.Get("userID")
	if !found {
		return int64(0), ""
	}

	nickname, found := c.Get("nickname")
	if !found {
		return int64(userID.(float64)), ""
	}

	return int64(userID.(float64)), nickname.(string)
}
