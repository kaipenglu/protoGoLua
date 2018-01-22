package hdl

import (
    "fmt"
    "../pb/pbcpl"
    "github.com/stevedonovan/luar"
    "github.com/aarzilli/golua/lua"
)

func LoginReq(L *lua.State, msg interface{}) interface{} {
    req := msg.(*pbcpl.LoginReq)
    L.GetGlobal("LoginEnterPoint")
    luar.GoToLua(L, req)
    L.Call(1, 2)

    defer L.Pop(2)

    success := L.ToBoolean(1)
    if !success {
        return nil
    }

    res := &pbcpl.LoginRes{}
    cmd := pbcpl.PACKET_ID_PACKET_LOGIN_RES
    res.Cmd = &cmd
    s := L.ToString(2)
    res.Ans = &s
    return res
}

func LoginRes(L *lua.State, msg interface{}) interface{} {
    req := msg.(*pbcpl.LoginRes)
    fmt.Println(*req.Ans)
    return nil
}
