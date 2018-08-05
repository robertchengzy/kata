##### redis 批量删除keys
`
src/redis-cli -h 127.0.0.1 -p 62224 -a abc#12345 keys "point_recharge_*" | xargs src/redis-cli -h 127.0.0.1 -p 62224 -a abc#12345 del
`