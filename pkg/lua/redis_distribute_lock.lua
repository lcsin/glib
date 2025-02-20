local key = KEYS[1]
local uniqueId = ARGV[1]
local ttl = tonumber(ARGV[2])

if redis.call("setnx", key, uniqueId) == 1 then
    redis.call("expire", key, ttl)
    return 1
elseif redis.call("get", key) == uniqueId then
    redis.call("expire", key, ttl)
    return 1
end
return -1

