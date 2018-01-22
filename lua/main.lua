pcall = unsafe_pcall
xpcall = unsafe_xpcall

require "login"

print("Lua state is initializing!")

local handleError = function(error)
    error = debug.traceback(error)
    print(error)
end


LoginEnterPoint = function(req)
    return xpcall(Login, handleError, req)
end
