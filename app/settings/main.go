package main

import (
	"grumpy"
	_ "__python__/traceback"
	_ "__python__/hello"
	"log"
)

func main() {
	// 1. 调用基础的库
	grumpy.ImportModule(grumpy.NewRootFrame(), "traceback")

	rootFrame := grumpy.NewRootFrameWithMeta("__main__", "settings/main.go")

	var packages []*grumpy.Object
	var methodHello *grumpy.Object
	var methodGetKV *grumpy.Object
	var basicEx *grumpy.BaseException

	strGetKeyValue := grumpy.InternStr("get_key_value")
	strHello := grumpy.InternStr("hello")
	strB := grumpy.InternStr("b")

	// 2. 引用Python编译之后的代码
	if packages, basicEx = grumpy.ImportModule(rootFrame, "hello"); basicEx != nil {
		log.Printf("Exception: %v", basicEx)
		return
	}

	// Python中的Frame该如何处理？每次调用都使用不同的Frame？函数退出之后如何Reset状态呢？
	// 3. 准备封装代码
	packageHello := packages[0]
	if methodHello, basicEx = grumpy.GetAttr(rootFrame, packageHello, strHello, nil); basicEx != nil {
		log.Printf("Exception: %v", basicEx)
		return
	}

	// 然后函数：methodHello， methodGetKV在rootFrame可见
	if basicEx = rootFrame.Globals().SetItem(rootFrame, strHello.ToObject(), methodHello); basicEx != nil {
		log.Printf("Exception: %v", basicEx)
		return
	}

	if methodGetKV, basicEx = grumpy.GetAttr(rootFrame, packageHello, strGetKeyValue, nil); basicEx != nil {
		log.Printf("Exception: %v", basicEx)
		return
	}
	if basicEx = rootFrame.Globals().SetItem(rootFrame, strGetKeyValue.ToObject(), methodGetKV); basicEx != nil {
		log.Printf("Exception: %v", basicEx)
		return
	}

	// 准备工作做好了，然后准备调用函数

	// NewDict()不需要主动回收
	dictInst := grumpy.NewDict()
	if basicEx = dictInst.SetItem(rootFrame,
		grumpy.InternStr("a").ToObject(),
		grumpy.NewInt(1).ToObject()); basicEx != nil {
		log.Printf("Exception: %v", basicEx)
		return
	}
	if basicEx = dictInst.SetItem(rootFrame,
		grumpy.InternStr("b").ToObject(),
		grumpy.InternStr("hh").ToObject()); basicEx != nil {
		log.Printf("Exception: %v", basicEx)
		return
	}

	// 修改Globals是否会存在污染呢？
	if basicEx = rootFrame.Globals().SetItem(rootFrame,
		grumpy.InternStr("d").ToObject(),
		dictInst.ToObject()); basicEx != nil {
		log.Printf("Exception: %v", basicEx)
		return
	}

	// line 7: b = get_key_value(d, "b")
	argsTuple := rootFrame.MakeArgs(2)
	var instanceD *grumpy.Object
	if instanceD, basicEx = grumpy.ResolveGlobal(rootFrame, grumpy.InternStr("d")); basicEx != nil {
		log.Printf("Exception: %v", basicEx)
		return
	}
	argsTuple[0] = instanceD
	argsTuple[1] = strB.ToObject()
	if methodGetKV, basicEx = grumpy.ResolveGlobal(rootFrame, strGetKeyValue); basicEx != nil {
		log.Printf("Exception: %v", basicEx)
		return
	}

	var result *grumpy.Object
	if result, basicEx = methodGetKV.Call(rootFrame, argsTuple, nil); basicEx != nil {
		log.Printf("Exception: %v", basicEx)
		return
	}
	rootFrame.FreeArgs(argsTuple)

	log.Printf("Result: %s", result.String())

}
