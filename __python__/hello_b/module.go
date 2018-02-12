package hello_b
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/Users/feiwang/workspace/grumpy/build/src/__python__/hello_b.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßget := πg.InternStr("get")
		ßget_key_value_b := πg.InternStr("get_key_value_b")
		ßglobal_meta_b := πg.InternStr("global_meta_b")
		ßhello_b := πg.InternStr("hello_b")
		var πTemp001 *πg.Dict
		_ = πTemp001
		var πTemp002 *πg.Object
		_ = πTemp002
		var πTemp003 []πg.Param
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: global_meta_b = {}
			πF.SetLineno(1)
			πTemp001 = πg.NewDict()
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßglobal_meta_b.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 2: def hello_b(name, metas):
			πF.SetLineno(2)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "name", Def: nil}
			πTemp003[1] = πg.Param{Name: "metas", Def: nil}
			πTemp002 = πg.NewFunction(πg.NewCode("hello_b", "/Users/feiwang/workspace/grumpy/build/src/__python__/hello_b.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var µmetas *πg.Object = πArgs[1]; _ = µmetas
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
					// line 3: print "hello, world: %s %s" % (name, metas)
					πF.SetLineno(3)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmetas, "metas"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µname, µmetas).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("hello, world: %s %s").ToObject(), πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßhello_b.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 6: def get_key_value_b(reqs, key):
			πF.SetLineno(6)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "reqs", Def: nil}
			πTemp003[1] = πg.Param{Name: "key", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("get_key_value_b", "/Users/feiwang/workspace/grumpy/build/src/__python__/hello_b.py", πTemp003, 0,
				func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 7: return "result: %s %s" % (reqs.get(key, ""), global_meta_b)
					πF.SetLineno(7)
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
					if πTemp004, πE = πg.ResolveGlobal(πF, ßglobal_meta_b); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßget_key_value_b.ToObject(), πTemp004); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("hello_b", Code)
}
