--
-- Created by IntelliJ IDEA.
-- User: guoxiang.li
-- Date: 2020/07/03
-- Time: 20:03
-- To change this template use File | Settings | File Templates.
--

local function max(...)

    local args = { ... }
    local val, idx
    for i = 1, #args do
        if val == nil or args[i] > val then
            val,
            idx = args[i], i
        end
    end
    return val, idx
end

local function assert(v)
    if not v then fail() end
end

local v = max(3, 9, 7, 128, 35)
assert(v == 128)