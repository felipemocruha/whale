package main

type IFn interface {
	invoke() (interface{}, error)
	invoke(arg1 interface{}) (interface{}, error)
	invoke(arg1, arg2 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5, arg6 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5, arg6,
		arg7 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5, arg6, arg7,
		arg8 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5, arg6, arg7,
		arg8, arg9 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5, arg6, arg7,
		arg8, arg9, arg10 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5, arg6, arg7,
		arg8, arg9, arg10, arg11 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5, arg6, arg7,
		arg8, arg9, arg10, arg11, arg12 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5, arg6, arg7,
		arg8, arg9, arg10, arg11, arg12, arg13 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5, arg6, arg7,
		arg8, arg9, arg10, arg11, arg12, arg13, arg14 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5, arg6, arg7,
		arg8, arg9, arg10, arg11, arg12, arg13, arg14,
		arg15 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5, arg6, arg7,
		arg8, arg9, arg10, arg11, arg12, arg13, arg14,
		arg15, arg16 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5, arg6, arg7,
		arg8, arg9, arg10, arg11, arg12, arg13, arg14,
		arg15, arg16, arg17 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5, arg6, arg7,
		arg8, arg9, arg10, arg11, arg12, arg13, arg14,
		arg15, arg16, arg17, arg18 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5, arg6, arg7,
		arg8, arg9, arg10, arg11, arg12, arg13, arg14,
		arg15, arg16, arg17, arg18, arg19 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5, arg6, arg7,
		arg8, arg9, arg10, arg11, arg12, arg13, arg14,
		arg15, arg16, arg17, arg18, arg19, arg20 interface{}) (interface{}, error)
	invoke(arg1, arg2, arg3, arg4, arg5, arg6, arg7,
		arg8, arg9, arg10, arg11, arg12, arg13, arg14,
		arg15, arg16, arg17, arg18, arg19, arg20 interface{},
		args ...interface{}) (interface{}, error)
	applyTo(arglist ISeq) (interface{}, error)
}
