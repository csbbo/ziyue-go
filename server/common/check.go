package common

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context, formData interface{}, loginRequire bool) interface{}{
	if loginRequire == true {
		session := sessions.Default(c)
		username := session.Get("user")
		if username == nil {
			return "需要登录"
		}
	}
	if formData != nil {
		err := c.ShouldBind(formData)
		return err
	}
	return nil
}
