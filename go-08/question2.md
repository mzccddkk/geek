### 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

- value 10 byte：

```bash
# 清空数据库 & 输出内存占用
127.0.0.1:6379> FLUSHALL
127.0.0.1:6379> INFO memory
used_memory:1586000
used_memory_human:1.51M

# 插入 50w 测试数据：key 10 byte，value 10 byte
127.0.0.1:6379> EVAL "for i=1,500000 do redis.call('SET', KEYS[1]..string.format('%06d', i), string.format('%010d', i)) end" 1 "key:"
127.0.0.1:6379> INFO memory
used_memory:45780912
used_memory_human:43.66M
```

> 50w 测试数据占用总内存：45780912 - 1586000 = 44194912(byte) ≈ 42(mb)
>
> 50w 测试数据平均每个 key 占用内存：44194912 / 500000 ≈ 88(byte)


- value 20 byte：

```bash
# 清空数据库 & 输出内存占用
127.0.0.1:6379> FLUSHALL
127.0.0.1:6379> INFO memory
used_memory:1658928
used_memory_human:1.58M

# 插入 50w 测试数据：key 10 byte，value 20 byte
127.0.0.1:6379> EVAL "for i=1,500000 do redis.call('SET', KEYS[1]..string.format('%06d', i), string.format('%020d', i)) end" 1 "key:"
127.0.0.1:6379> INFO memory
used_memory:53853904
used_memory_human:51.36M
```

> 50w 测试数据占用总内存：53853904 - 1658928 = 52194976(byte) ≈ 50(mb)
>
> 50w 测试数据平均每个 key 占用内存：52194976 / 500000 ≈ 104(byte)


- value 50 byte：

```bash
# 清空数据库 & 输出内存占用
127.0.0.1:6379> FLUSHALL
127.0.0.1:6379> INFO memory
used_memory:1659792
used_memory_human:1.58M

# 插入 50w 测试数据：key 10 byte，value 50 byte
127.0.0.1:6379> EVAL "for i=1,500000 do redis.call('SET', KEYS[1]..string.format('%06d', i), string.format('%050d', i)) end" 1 "key:"
127.0.0.1:6379> INFO memory
used_memory:69920224
used_memory_human:66.68M
```

> 50w 测试数据占用总内存：69920224 - 1659792 = 68260432(byte) ≈ 65(mb)
>
> 50w 测试数据平均每个 key 占用内存：52194976 / 500000 ≈ 137(byte)



可以看到，随着 value 增大，平均每个 key 占用的内存也越来越大

> Redis 容量预估：http://www.redis.cn/redis_memory/