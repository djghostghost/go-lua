--
-- Created by IntelliJ IDEA.
-- User: guoxiang.li
-- Date: 2020/07/02
-- Time: 10:54
-- To change this template use File | Settings | File Templates.
--

local t = { "a", "b", "c" }

t[2] = "B"
t["foo"] = "Bar"

local s = t[3] .. t[2] .. t[1] .. t["foo"] .. #t


