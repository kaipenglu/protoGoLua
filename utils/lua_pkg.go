package utils

import "github.com/aarzilli/golua/lua"

func AddPkgPath(L *lua.State, addPath string) {
    L.GetGlobal("package")
    L.GetField(-1, "path")
    luaPkgPath := L.ToString(-1) // grab path string from top of stack
    luaPkgPath = addPath + "/?.lua;" + luaPkgPath
    L.Pop(1); // get rid of the string on the stack we just pushed on line 5
    L.PushString(luaPkgPath) // push the new one
    L.SetField(-2, "path") // set the field "path" in table at -2 with value at top of stack
    L.Pop(1) // get rid of package table from top of stack
}
