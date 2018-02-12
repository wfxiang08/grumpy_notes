package __main__
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "__python__/app_settings/wf/t01/module.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßa := πg.InternStr("a")
		ßb := πg.InternStr("b")
		ßd := πg.InternStr("d")
		ßdd := πg.InternStr("dd")
		ßget_key_value := πg.InternStr("get_key_value")
		ßhello := πg.InternStr("hello")
		ßhh := πg.InternStr("hh")
		ßwangfei := πg.InternStr("wangfei")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Dict
		_ = πTemp004
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: from app_settings.wf.hello import hello, get_key_value
			πF.SetLineno(1)
			if πTemp002, πE = πg.ImportModule(πF, "app_settings.wf.hello"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßhello, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßhello.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßget_key_value, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßget_key_value.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 3: d = {"a": 1, "b": "hh"}
			πF.SetLineno(3)
			πTemp004 = πg.NewDict()
			if πE = πTemp004.SetItem(πF, ßa.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßb.ToObject(), ßhh.ToObject()); πE != nil {
				continue
			}
			πTemp001 = πTemp004.ToObject()
			if πE = πF.Globals().SetItem(πF, ßd.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: hello("wangfei", d)
			πF.SetLineno(5)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = ßwangfei.ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			πTemp002[1] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßhello); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 7: b = get_key_value(d, "b")
			πF.SetLineno(7)
			πTemp002 = πF.MakeArgs(2)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			πTemp002[0] = πTemp001
			πTemp002[1] = ßb.ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßget_key_value); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßb.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 8: print b
			πF.SetLineno(8)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßb); πE != nil {
				continue
			}
			πTemp002[0] = πTemp001
			if πE = πg.Print(πF, πTemp002, true); πE != nil {
				continue
			}
			// line 10: dd = get_key_value(d, "d")
			πF.SetLineno(10)
			πTemp002 = πF.MakeArgs(2)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßd); πE != nil {
				continue
			}
			πTemp002[0] = πTemp001
			πTemp002[1] = ßd.ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßget_key_value); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßdd.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 11: print dd
			πF.SetLineno(11)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßdd); πE != nil {
				continue
			}
			πTemp002[0] = πTemp001
			if πE = πg.Print(πF, πTemp002, true); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("__main__", Code)
}
