local key = KEYS[1]
local uniqueId = ARGV[1]

if redis.call("get", key) == uniqueId then
    redis.call("del", key)
    return 1
end
return -1
