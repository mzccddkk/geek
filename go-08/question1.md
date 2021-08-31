### 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

- 先查看使用到的 redis-benchmark 命令参数：

```bash
$ redis-benchmark --help
-d <size>          Data size of SET/GET value in bytes (default 3)
-t <tests>         Only run the comma separated list of tests. The test
                    names are the same as the ones produced as output.
-q                 Quiet. Just show query/sec values
-P <numreq>        Pipeline <numreq> requests. Default 1 (no pipeline).
```
> Redis 6.0 版本
>
> 如果使用 -P 参数指定通过管道传输请求，吞吐将会更高

- 10 byte：

```bash
$ redis-benchmark -h 127.0.0.1 -p 6379 -d 10 -t get,set -q
SET: 109051.26 requests per second
GET: 109890.11 requests per second
```

- 20 byte：

```bash
$ redis-benchmark -h 127.0.0.1 -p 6379 -d 20 -t get,set -q
SET: 108459.87 requests per second
GET: 107066.38 requests per second
```

- 50 byte：

```bash
$ redis-benchmark -h 127.0.0.1 -p 6379 -d 50 -t get,set -q
SET: 103734.44 requests per second
GET: 110253.59 requests per second
```

- 100 byte：

```bash
$ redis-benchmark -h 127.0.0.1 -p 6379 -d 100 -t get,set -q
SET: 107181.13 requests per second
GET: 108813.92 requests per second
```

- 200 byte：

```bash
$ redis-benchmark -h 127.0.0.1 -p 6379 -d 200 -t get,set -q
SET: 108813.92 requests per second
GET: 110497.24 requests per second
```

- 1000 byte：

```bash
$ redis-benchmark -h 127.0.0.1 -p 6379 -d 1000 -t get,set -q
SET: 105042.02 requests per second
GET: 110132.16 requests per second
```

- 5000 byte：

```bash
$ redis-benchmark -h 127.0.0.1 -p 6379 -d 5000 -t get,set -q
SET: 103950.10 requests per second
GET: 103842.16 requests per second
```

- 10000 byte：

```bash
$ redis-benchmark -h 127.0.0.1 -p 6379 -d 10000 -t get,set -q
SET: 97276.27 requests per second
GET: 97276.27 requests per second
```

> 上面测试数据没有开启 AOF 日志、多线程

可以看到，随着 value 增大，达到一定阈值之后，吞吐会随之降低，所以要避免 **bigkey**
