-- 限流对象：例如ip
local key = KEYS[1]
-- 窗口大小：请求时间的范围，例如1秒内100个请求，这里的窗口大小就是1秒
local window = tonumber(ARGV[1])
-- 阈值：请求数量的阈值，例如1秒内100个请求，这里的阈值就是100
local threshold = tonumber(ARGV[2])
-- 当前请求的时间
local now = tonumber(ARGV[3])
-- 窗口的起始时间
local min = now - window

redis.call("ZREMRANGEBYSCORE", key, "-inf", min)
if redis.call("ZCOUNT", key, "-inf", '+inf') >= threshold then
    return true
else
    redis.call("ZADD", key, now, now)
    redis.call("PEXPIRE", key, window)
    return false
end
