package hello
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/Users/feiwang/workspace/grumpy/build/src/__python__/hello.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßget := πg.InternStr("get")
		ßget_key_value := πg.InternStr("get_key_value")
		ßget_key_value_b := πg.InternStr("get_key_value_b")
		ßglobal_meta := πg.InternStr("global_meta")
		ßhello := πg.InternStr("hello")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Dict
		_ = πTemp004
		var πTemp005 []πg.Param
		_ = πTemp005
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: from hello_b import get_key_value_b
			πF.SetLineno(1)
			if πTemp002, πE = πg.ImportModule(πF, "hello_b"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßget_key_value_b, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßget_key_value_b.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 2: global_meta = {}
			πF.SetLineno(2)
			πTemp004 = πg.NewDict()
			πTemp001 = πTemp004.ToObject()
			if πE = πF.Globals().SetItem(πF, ßglobal_meta.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 3: def hello(name, metas):
			πF.SetLineno(3)
			πTemp005 = make([]πg.Param, 2)
			πTemp005[0] = πg.Param{Name: "name", Def: nil}
			πTemp005[1] = πg.Param{Name: "metas", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("hello", "/Users/feiwang/workspace/grumpy/build/src/__python__/hello.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var µmetas *πg.Object = πArgs[1]; _ = µmetas
				var µa *πg.Object = πg.UnboundLocal; _ = µa
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 4: a = get_key_value_b(metas, name)
					πF.SetLineno(4)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µmetas, "metas"); πE != nil {
						continue
					}
					πTemp001[0] = µmetas
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[1] = µname
					if πTemp002, πE = πg.ResolveGlobal(πF, ßget_key_value_b); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µa = πTemp003
					// line 5: print "hello, world: %s %s %s" % (name, metas, a)
					πF.SetLineno(5)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmetas, "metas"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple3(µname, µmetas, µa).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("hello, world: %s %s %s").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßhello.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: def get_key_value(reqs, key):
			πF.SetLineno(8)
			πTemp005 = make([]πg.Param, 2)
			πTemp005[0] = πg.Param{Name: "reqs", Def: nil}
			πTemp005[1] = πg.Param{Name: "key", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("get_key_value", "/Users/feiwang/workspace/grumpy/build/src/__python__/hello.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µreqs *πg.Object = πArgs[0]; _ = µreqs
				var µkey *πg.Object = πArgs[1]; _ = µkey
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 9: return "result: %s %s" % (reqs.get(key, ""), global_meta)
					πF.SetLineno(9)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp003[0] = µkey
					πTemp003[1] = ß.ToObject()
					if πE = πg.CheckLocal(πF, µreqs, "reqs"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µreqs, ßget, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßglobal_meta); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp005, πTemp004).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("result: %s %s").ToObject(), πTemp002); πE != nil {
						continue
					}
					πR = πTemp001
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßget_key_value.ToObject(), πTemp003); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("hello", Code)
}
