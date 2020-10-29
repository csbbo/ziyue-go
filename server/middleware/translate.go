package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

func Translation() gin.HandlerFunc {
	return func(c *gin.Context) {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			v.RegisterTagNameFunc(func(fld reflect.StructField) string {
				name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
				if name == "-" {
					return ""
				}
				return name
			})

			zhT := zh.New()
			enT := en.New()

			locale := c.DefaultQuery("locale", "zh")
			uni := ut.New(enT, zhT, enT)
			var ok bool
			trans, ok := uni.GetTranslator(locale)
			if !ok {
				fmt.Errorf("uni.GetTranslator(%s) failed", locale)
			}

			switch locale {
			case "en":
				enTranslations.RegisterDefaultTranslations(v, trans)
			case "zh":
				zhTranslations.RegisterDefaultTranslations(v, trans)
			default:
				enTranslations.RegisterDefaultTranslations(v, trans)
			}
			c.Set("trans", trans)
		}
		c.Next()
	}
}
