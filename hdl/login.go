package hdl

import (
    "../pb/pbcpl"
    "github.com/stevedonovan/luar"
    "github.com/aarzilli/golua/lua"
)

func CalcCard(L *lua.State, msg interface{}) interface{} {
    req := msg.(*pbcpl.CalcCardReq)

    L.GetGlobal("Login")
    luar.GoToLua(L, req)
    L.Call(1, 2)

    noError := L.ToBoolean(1)
    codeSuc := pbcpl.RET_CODE_RET_SUCCESS
    codeErr := pbcpl.RET_CODE_RET_SERVER_ERROR
    if noError {
        res.Msgno = &codeSuc
    } else {
        res.Msgno = &codeErr
    }

    L.Pop(1)
    return res
}
