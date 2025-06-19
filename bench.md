# Benchmarks
```
goos: linux
goarch: amd64
pkg: github.com/pippellia-btc/slicex
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
```

## Set operations

### Unique
```
BenchmarkUnique/size=1000-4               238820              4812 ns/op               0 B/op          0 allocs/op
BenchmarkUnique/size=10000-4               29473             40503 ns/op               0 B/op          0 allocs/op
BenchmarkUnique/size=100000-4               1118           1045025 ns/op               0 B/op          0 allocs/op
BenchmarkUnique/size=1000000-4               228           4876423 ns/op               0 B/op          0 allocs/op

BenchmarkUniqueMap/size=1000-4             42171             27692 ns/op           21808 B/op          3 allocs/op
BenchmarkUniqueMap/size=10000-4             4363            269876 ns/op          180321 B/op          5 allocs/op
BenchmarkUniqueMap/size=100000-4             691           1725758 ns/op         1392696 B/op          3 allocs/op
BenchmarkUniqueMap/size=1000000-4             25          46673943 ns/op        22284364 B/op         10 allocs/op
```

### Intersection
```
BenchmarkIntersection/size=1000-4                 106520             11005 ns/op               0 B/op          0 allocs/op
BenchmarkIntersection/size=10000-4                 14288             83212 ns/op               0 B/op          0 allocs/op
BenchmarkIntersection/size=100000-4                  612           1941758 ns/op               0 B/op          0 allocs/op
BenchmarkIntersection/size=1000000-4                  91          11199676 ns/op               0 B/op          0 allocs/op

BenchmarkIntersectionMap/size=1000-4               31039             39581 ns/op           21785 B/op          2 allocs/op
BenchmarkIntersectionMap/size=10000-4               3058            336565 ns/op          180249 B/op          2 allocs/op
BenchmarkIntersectionMap/size=100000-4               422           2834571 ns/op         1392664 B/op          2 allocs/op
BenchmarkIntersectionMap/size=1000000-4               12          87378222 ns/op        22283784 B/op         10 allocs/op
```

### Union
```
BenchmarkUnion/size=1000-4                         43617             27460 ns/op               0 B/op          0 allocs/op
BenchmarkUnion/size=10000-4                         3157            358336 ns/op               0 B/op          0 allocs/op
BenchmarkUnion/size=100000-4                         306           3946495 ns/op               0 B/op          0 allocs/op
BenchmarkUnion/size=1000000-4                         43          26823770 ns/op               0 B/op          0 allocs/op

BenchmarkUnionMap/size=1000-4                      13630             86043 ns/op          144688 B/op          5 allocs/op
BenchmarkUnionMap/size=10000-4                      1393            880048 ns/op         1458224 B/op          6 allocs/op
BenchmarkUnionMap/size=100000-4                      163           7121089 ns/op        13860915 B/op          6 allocs/op
BenchmarkUnionMap/size=1000000-4                       7         153296023 ns/op        171302973 B/op         6 allocs/op
```

### Difference
```
BenchmarkDifference/size=1000-4                    74428             15730 ns/op               0 B/op          0 allocs/op
BenchmarkDifference/size=10000-4                    7848            143296 ns/op               0 B/op          0 allocs/op
BenchmarkDifference/size=100000-4                    628           1867687 ns/op               0 B/op          0 allocs/op
BenchmarkDifference/size=1000000-4                    96          14950663 ns/op               0 B/op          0 allocs/op

BenchmarkDifferenceMap/size=1000-4                 37393             29675 ns/op           21789 B/op          2 allocs/op
BenchmarkDifferenceMap/size=10000-4                 3832            318160 ns/op          180393 B/op          6 allocs/op
BenchmarkDifferenceMap/size=100000-4                 540           2165198 ns/op         1392666 B/op          2 allocs/op
BenchmarkDifferenceMap/size=1000000-4                 21          48517385 ns/op        22284460 B/op         11 allocs/op
```

### Differences
```
BenchmarkDifferences/size=1000-4                   73292             16011 ns/op               0 B/op          0 allocs/op
BenchmarkDifferences/size=10000-4                   7792            143103 ns/op               0 B/op          0 allocs/op
BenchmarkDifferences/size=100000-4                   633           1852744 ns/op               0 B/op          0 allocs/op
BenchmarkDifferences/size=1000000-4                   67          15435585 ns/op               0 B/op          0 allocs/op

BenchmarkDifferencesMap/size=1000-4                18253             63708 ns/op           43578 B/op          4 allocs/op
BenchmarkDifferencesMap/size=10000-4                1832            678309 ns/op          360787 B/op         13 allocs/op
BenchmarkDifferencesMap/size=100000-4                260           4584786 ns/op         2785332 B/op          4 allocs/op
BenchmarkDifferencesMap/size=1000000-4                10         107374478 ns/op        44568921 B/op         22 allocs/op
```

### SymmetricDifference
```
BenchmarkSymmetricDifference/size=1000-4                   78373             14934 ns/op               0 B/op          0 allocs/op
BenchmarkSymmetricDifference/size=10000-4                   4524            251034 ns/op               0 B/op          0 allocs/op
BenchmarkSymmetricDifference/size=100000-4                   670           1747275 ns/op               0 B/op          0 allocs/op
BenchmarkSymmetricDifference/size=1000000-4                   80          15744703 ns/op               0 B/op          0 allocs/op

BenchmarkSymmetricDifferenceMap/size=1000-4                22468             52929 ns/op           43568 B/op          4 allocs/op
BenchmarkSymmetricDifferenceMap/size=10000-4                1594            708030 ns/op          360785 B/op         13 allocs/op
BenchmarkSymmetricDifferenceMap/size=100000-4                253           4724050 ns/op         2785333 B/op          4 allocs/op
BenchmarkSymmetricDifferenceMap/size=1000000-4                22          51974736 ns/op        44564549 B/op          4 allocs/op
```


### Partition
```
BenchmarkPartition/size=1000-4                             57817             20484 ns/op               8 B/op          1 allocs/op
BenchmarkPartition/size=10000-4                             7545            144560 ns/op               8 B/op          1 allocs/op
BenchmarkPartition/size=100000-4                             636           1854845 ns/op               8 B/op          1 allocs/op
BenchmarkPartition/size=1000000-4                             79          18950444 ns/op               8 B/op          1 allocs/op

BenchmarkPartitionMap/size=1000-4                          29502             41093 ns/op           43672 B/op          9 allocs/op
BenchmarkPartitionMap/size=10000-4                          2121            501720 ns/op          360601 B/op          9 allocs/op
BenchmarkPartitionMap/size=100000-4                          362           3328196 ns/op         2785435 B/op          9 allocs/op
BenchmarkPartitionMap/size=1000000-4                          26          42737413 ns/op        44564646 B/op          9 allocs/op
```

