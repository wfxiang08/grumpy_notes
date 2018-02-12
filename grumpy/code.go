// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grumpy

import (
	"reflect"
)

// CodeType is the object representing the Python 'code' type.
var CodeType = newBasisType("code", reflect.TypeOf(Code{}), toCodeUnsafe, ObjectType)

// CodeFlag is a switch controlling the behavior of a Code object.
type CodeFlag int

const (
	// CodeFlagVarArg means a Code object accepts *arg parameters.
	CodeFlagVarArg CodeFlag = 4
	// CodeFlagKWArg means a Code object accepts **kwarg parameters.
	CodeFlagKWArg CodeFlag = 8
)

// Code represents Python 'code' objects.
type Code struct {
	Object // Code也是一个对象，有自己基本的属性：package，文件名，等
	name     string `attr:"co_name"`
	filename string `attr:"co_filename"`
	// argc is the number of positional arguments.
	argc      int      `attr:"co_argcount"`
	flags     CodeFlag `attr:"co_flags"`
	paramSpec *ParamSpec
	fn        func(*Frame, []*Object) (*Object, *BaseException) // Code最终需要有一个可执行的片段，例如：package的初始化
}

// NewCode creates a new Code object that executes the given fn.
func NewCode(name, filename string, params []Param, flags CodeFlag,
	fn func(*Frame, []*Object) (*Object, *BaseException)) *Code {
	// 构建一块Code
	s := NewParamSpec(name, params, flags&CodeFlagVarArg != 0, flags&CodeFlagKWArg != 0)

	return &Code{Object{typ: CodeType}, name, filename, len(params), flags, s, fn}
}

func toCodeUnsafe(o *Object) *Code {
	return (*Code)(o.toPointer())
}

// Eval runs the code object c in the context of the given globals.
func (c *Code) Eval(f *Frame, globals *Dict, args Args, kwargs KWArgs) (*Object, *BaseException) {
	// 如何执行代码呢？
	validated := f.MakeArgs(c.paramSpec.Count)

	// 处理好参数的逻辑
	if raised := c.paramSpec.Validate(f, validated, args, kwargs); raised != nil {
		return nil, raised
	}
	oldExc, oldTraceback := f.ExcInfo()

	// 创建新的Frame, 使用globals来替换
	next := newChildFrame(f)
	next.code = c
	next.globals = globals

	// 执行Codedefn
	ret, raised := c.fn(next, validated)
	next.release()

	// 回收内存
	f.FreeArgs(validated)

	if raised == nil {
		// Restore exc_info to what it was when we left the previous
		// frame.
		f.RestoreExc(oldExc, oldTraceback)
		if ret == nil {
			ret = None
		}
	} else {
		_, tb := f.ExcInfo()
		if f.code != nil {
			// The root frame has no code object so don't include it
			// in the traceback.
			tb = newTraceback(f, tb)
		}
		f.RestoreExc(raised, tb)
	}
	return ret, raised
}
