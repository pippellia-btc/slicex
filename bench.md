# Benchmarks
```
goos: linux
goarch: amd64
pkg: github.com/pippellia-btc/slicex
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
```

## Set Operations

### Unique
```
BenchmarkUnique/size=1000-4                34848             34632 ns/op           29980 B/op          3 allocs/op
BenchmarkUnique/size=10000-4                3511            347192 ns/op          262256 B/op          6 allocs/op
BenchmarkUnique/size=100000-4                358           3239547 ns/op         2195487 B/op          3 allocs/op
BenchmarkUnique/size=1000000-4                21          48853057 ns/op        30285861 B/op          3 allocs/op
```

### Intersection
```
BenchmarkIntersection/size=1000-4                  10479            116972 ns/op           66113 B/op         10 allocs/op
BenchmarkIntersection/size=10000-4                   837           1322413 ns/op          555259 B/op         30 allocs/op
BenchmarkIntersection/size=100000-4                   92          12754598 ns/op         4265785 B/op         35 allocs/op
BenchmarkIntersection/size=1000000-4                   5         214926981 ns/op        67499571 B/op         29 allocs/op
```

### Union
```
BenchmarkUnion/size=1000-4                         14812             83503 ns/op          122904 B/op          3 allocs/op
BenchmarkUnion/size=10000-4                         1011           1015307 ns/op         1024024 B/op          3 allocs/op
BenchmarkUnion/size=100000-4                          90          12214997 ns/op         8773658 B/op          3 allocs/op
BenchmarkUnion/size=1000000-4                          5         207448430 ns/op        121135166 B/op         3 allocs/op
```

### Difference
```
BenchmarkDifference/size=1000-4                    25191             47789 ns/op           29984 B/op          3 allocs/op
BenchmarkDifference/size=10000-4                    1921            538609 ns/op          263147 B/op         10 allocs/op
BenchmarkDifference/size=100000-4                    162           7206230 ns/op         2204841 B/op         14 allocs/op
BenchmarkDifference/size=1000000-4                     9         117624942 ns/op        30287816 B/op         12 allocs/op
```

### SymmetricDifference
```
BenchmarkSymmetricDifference/size=1000-4                   14655             81519 ns/op           53304 B/op          5 allocs/op
BenchmarkSymmetricDifference/size=10000-4                   1015           1054238 ns/op          663815 B/op         11 allocs/op
BenchmarkSymmetricDifference/size=100000-4                   105          11024835 ns/op         5460620 B/op         20 allocs/op
BenchmarkSymmetricDifference/size=1000000-4                    5         204104461 ns/op        56560636 B/op         22 allocs/op
```


### Partition
```
BenchmarkPartition/size=1000-4                             14172             82753 ns/op           57392 B/op         15 allocs/op
BenchmarkPartition/size=10000-4                             1077           1111989 ns/op          800251 B/op         28 allocs/op
BenchmarkPartition/size=100000-4                             105          11291308 ns/op         5957518 B/op         41 allocs/op
BenchmarkPartition/size=1000000-4                              5         213663805 ns/op        61784083 B/op         51 allocs/op
```

## K-Element Selection

### MinK
```
BenchmarkMinK/min_10/1000-4              4115961               288 ns/op               0 B/op          0 allocs/op
BenchmarkMinK/min_10/10000-4              416628              2746 ns/op               0 B/op          0 allocs/op
BenchmarkMinK/min_10/100000-4              42150             27978 ns/op               0 B/op          0 allocs/op
BenchmarkMinK/min_10/1000000-4              2562            412954 ns/op               0 B/op          0 allocs/op

BenchmarkMinKNaive/min_10/1000-4         1484443               806 ns/op               0 B/op          0 allocs/op
BenchmarkMinKNaive/min_10/10000-4         151060              7806 ns/op               0 B/op          0 allocs/op
BenchmarkMinKNaive/min_10/100000-4         15339             77716 ns/op               0 B/op          0 allocs/op
BenchmarkMinKNaive/min_10/1000000-4         1396            803882 ns/op               0 B/op          0 allocs/op
```

### MaxK

```
BenchmarkTopK/top_10/1000-4              2496135               474.8 ns/op            56 B/op          2 allocs/op
BenchmarkTopK/top_10/10000-4              329167              3546 ns/op              56 B/op          2 allocs/op
BenchmarkTopK/top_10/100000-4              31326             34302 ns/op              56 B/op          2 allocs/op
BenchmarkTopK/top_10/1000000-4              2490            416383 ns/op              56 B/op          2 allocs/op

BenchmarkTopKNaive/top_10/1000-4          240920              4662 ns/op               0 B/op          0 allocs/op
BenchmarkTopKNaive/top_10/10000-4          34471             33364 ns/op               0 B/op          0 allocs/op
BenchmarkTopKNaive/top_10/100000-4          2053            533211 ns/op               0 B/op          0 allocs/op
BenchmarkTopKNaive/top_10/1000000-4          360           3267533 ns/op               0 B/op          0 allocs/op
```

**Note**: `TopK` is considerably slower then `MinK`, since it uses the less optimized `sort.Slice` instead of `slices.Sort` to return elements in descending order.

Normally, the recommended way to sort in descending order is by calling `slices.SortFunc`, which is however extremely inefficient as you can see from the `TopKNaive` results.