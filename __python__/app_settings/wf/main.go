package main
import (
	"grumpy"
	_ "__python__/traceback"
	_ "__python__/hello"
	mod "__python__/app_settings"

)
func main() {
	grumpy.ImportModule(grumpy.NewRootFrame(), "traceback")
	// 如何直接调用代码呢？

	// f := grumpy.NewRootFrame()

	// f.code = code
	// f.globals = m.Dict()

	// πTemp002, πE = grumpy.ImportModule(f, "hello")

	grumpy.RunMain(mod.Code)
}
