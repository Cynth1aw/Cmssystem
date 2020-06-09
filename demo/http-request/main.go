package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main()  {
	fmt.Println("你好 V1.14.4")
	// 创建实例
	app := iris.New()

	app.Get("/getRequest", func(context context.Context) {
		//处理get请求，请求的url为：/getRequest
		path := context.Path()
		app.Logger().Info(path)
	})

	app.Get("/userpath", func(context context.Context) {
		//获取Path
		path := context.Path()
		app.Logger().Info(path)
		//写入返回数据：string类型
		context.WriteString("请求路径：" + path)
	})

	//2.处理Get请求 并接受参数
	app.Get("/userinfo", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		//获取get请求所携带的参数
		userName := context.URLParam("username") //
		app.Logger().Info(userName)

		pwd := context.URLParam("pwd")
		app.Logger().Info(pwd)
		//返回html数据格式
		context.HTML("<h1>" + userName + "," + pwd + "</h1>")
	})

	//3.处理Post请求 form表单的字段获取
	app.Post("/postLogin", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		//context.PostValue方法来获取post请求所提交的for表单数据
		name := context.PostValue("name")
		pwd := context.PostValue("pwd")
		app.Logger().Info(name, "  ", pwd)
		context.HTML(name)
	})

	//4、处理Post请求 Json格式数据
	/**
	 * Postman工具选择[{"key":"Content-Type","value":"application/json","description":""}]
	 * 请求内容：{"name": "davie","age": 28}
	 */
	 app.Post("/postJson", func(context context.Context) {

		//1.path
		path := context.Path()
		app.Logger().Info("请求URL：", path)

		//2.Json数据解析
		var person Person
		//context.ReadJSON()
		if err := context.ReadJSON(&person); err != nil {
			panic(err.Error())
		}

		//输出：Received: main.Person{Name:"davie", Age:28}
		context.Writef("Received: %#+v\n", person)
	})

	//5.处理Post请求 Xml格式数据
	/**
	 * 请求配置：Content-Type到application/xml（可选但最好设置）
	 * 请求内容：
	 *
	 *  <student>
	 *		<stu_name>davie</stu_name>
	 *		<stu_age>28</stu_age>
	 *	</student>
	 *
	 */

	 app.Post("/postXml", func(context context.Context) {

		//1.Path
		path := context.Path()
		app.Logger().Info("请求URL：", path)

		//2.XML数据解析
		var student Student
		if err := context.ReadXML(&student); err != nil {
			panic(err.Error())
		}
		//输出：
		context.Writef("Received：%#+v\n", student)
	})

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}

//自定义的struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//自定义的结构体
type Student struct {
	//XMLName xml.Name `xml:"student"`
	StuName string `xml:"stu_name"`
	StuAge  int    `xml:"stu_age"`
}