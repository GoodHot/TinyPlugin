package plugin

import "_github.com/labstack/echo"

type HTTPContext struct {
	Context echo.Context // echo上下文
	UserID  string
}
