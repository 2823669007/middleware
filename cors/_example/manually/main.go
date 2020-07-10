package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	crs := func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		if ctx.Method() == iris.MethodOptions {
			ctx.Header("Access-Control-Methods", "POST, PUT, PATCH, DELETE")
			ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Content-Type")
			ctx.Header("Access-Control-Max-Age", "86400")
			ctx.StatusCode(iris.StatusNoContent)
			return
		}

		ctx.Next()
	} // or	"github.com/iris-contrib/middleware/cors"

	v1 := app.Party("/api/v1", crs).AllowMethods(iris.MethodOptions) // <- important for the preflight.
	{
		v1.Post("/mailer", func(ctx iris.Context) {
			var any iris.Map
			err := ctx.ReadJSON(&any)
			if err != nil {
				ctx.StopWithError(iris.StatusBadRequest, err)
				return
			}
			ctx.Application().Logger().Infof("received %#+v", any)

			ctx.JSON(iris.Map{"message": "ok"})
		})

		v1.Get("/home", func(ctx iris.Context) {
			ctx.WriteString("Hello from /home")
		})
		v1.Get("/about", func(ctx iris.Context) {
			ctx.WriteString("Hello from /about")
		})
		v1.Post("/send", func(ctx iris.Context) {
			ctx.WriteString("sent")
		})
		v1.Put("/send", func(ctx iris.Context) {
			ctx.WriteString("updated")
		})
		v1.Delete("/send", func(ctx iris.Context) {
			ctx.WriteString("deleted")
		})
	}

	app.Listen(":80", iris.WithTunneling)
}
