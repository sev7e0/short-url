package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"short-link-go/global"
	"strings"
)

func InitTrans(locale string) (err error) {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New()
		enT := en.New()
		universalTranslator := ut.New(enT, zhT, enT)
		global.Trans, ok = universalTranslator.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "en":
			_ = entranslations.RegisterDefaultTranslations(validate, global.Trans)
		case "zh":
			_ = zhtranslations.RegisterDefaultTranslations(validate, global.Trans)
		default:
			_ = entranslations.RegisterDefaultTranslations(validate, global.Trans)
		}

		return
	}
	return
}
