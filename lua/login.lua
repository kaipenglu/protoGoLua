Login = function (req)
    print(req.Name .. " is logging in with password " .. req.Pwd)
    return "Welcome, " .. req.Name
end
