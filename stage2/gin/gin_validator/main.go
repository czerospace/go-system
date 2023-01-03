package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)

// 定义全局的翻译器
var trans ut.Translator

// LoginForm 定义一个结构体，并对字段进行约束
type LoginForm struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required,min=3,max=10"`
	Password string `form:"password" json:"password" xml:"password" binding:"required,min=6,max=10"`
}

// SignUpForm 定义一个结构体，并对字段进行约束
type SignUpForm struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=150"`
	Name       string `json:"name" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// removeTopStruct 去掉数据中结构体的名称，例如 去掉 LoginForm.password 中的 LoginForm
/*
		"error": {
	        "LoginForm.password": "password长度必须至少为6个字符"
	    }
*/
func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	// fmt.Println(fileds) // map[LoginForm.password:password长度必须至少为6个字符]
	for filed, err := range fileds {
		// fmt.Println(filed) // LoginForm.password
		// fmt.Println(err) // password长度必须至少为6个字符
		// fmt.Println(strings.Index(filed, ".")) // 9
		// fmt.Println( filed[strings.Index(filed, ".")+1:]) // password

		// 以 map[LoginForm.password:password长度必须至少为6个字符] 为例
		// 将 err 赋值给 key  filed[strings.Index(filed, ".")+1:]
		// 即: rsp[password] = password长度必须至少为6个字符
		rsp[filed[strings.Index(filed, ".")+1:]] = err
	}
	// fmt.Println("res是:", rsp) // map[password:password长度必须至少为6个字符]
	return rsp
}

// InitTrans 将英文翻译为中文
func InitTrans(locale string) (err error) {
	// 修改 gin 框架中的 validator 引擎属性，实现定制
	// 将 binding.Validator.Engine() 变成 validator.Validate
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取 json tag 的自定义方法,将 go 语言结构中的大写字段改为 json 中的小写字段
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			// 以逗号分割，获取name
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			// Name       string `json:"-"` 如果 tag 是 - 表示不处理
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		// 第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			en_translations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

func main() {
	if err := InitTrans("zh"); err != nil {
		fmt.Println("初始化翻译器错误")
		return
	}
	router := gin.Default()
	router.POST("loginJSON", func(c *gin.Context) {
		var loginForm LoginForm
		if err := c.ShouldBind(&loginForm); err != nil {
			// 把 err 类型转换为 validator.ValidationErrors
			errs, ok := err.(validator.ValidationErrors)
			// 如果转换 err 错误，直接返回 转换的错误，不是 validataor 错误
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
			}
			// 如果 err 能够翻译为 errs
			c.JSON(http.StatusBadRequest, gin.H{
				// errs.Translate 拿到的翻译器为 trans
				"error": removeTopStruct(errs.Translate(trans)),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
	})
	router.POST("/signup", func(c *gin.Context) {
		var signUpForm SignUpForm
		if err := c.ShouldBind(&signUpForm); err != nil {
			//fmt.Println(err.Error())
			//c.JSON(http.StatusBadRequest, gin.H{
			//	"error": err.Error(),
			//})
			// 把 err 类型转换为 validator.ValidationErrors
			errs, ok := err.(validator.ValidationErrors)
			// 如果转换 err 错误，直接返回 转换的错误，不是 validataor 错误
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
			}
			// 如果 err 能够翻译为 errs
			c.JSON(http.StatusBadRequest, gin.H{
				// errs.Translate 拿到的翻译器为 trans
				"error": removeTopStruct(errs.Translate(trans)),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})
	err := router.Run(":8083")
	if err != nil {
		panic(err)
	}
}
