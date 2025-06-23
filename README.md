# slicex

`slicex` is a Go package providing performance-optimized utility functions for slice manipulation not directly available in the `slices` standard library's.

## Installation
```
go get github.com/pippellia-btc/slicex
```

## Features
- **Efficient Set Operations**:   
Implements common set operations (e.g., difference, partition, symmetric difference) for slices using an efficient map-based approach.

- **Fast K-Element Selection**:   
Includes a custom algorithm for efficiently finding the Min-K and Max-K elements within a slice.

- **Generic Support**:   
All functions are built with Go generics, ensuring type safety and reusability across data types.

## Benchmarks

You can find the benchmarks [here](/bench.md).

**TLDR**:
- k-element selection: 10-100x faster then sorting the full slice

## Purpose

While Go's standard slices package provides foundational utilities, `slicex` aims to complement it by offering specialized and performant functions that address frequently encountered challenges in data processing. It's ideal for projects requiring highly efficient slice transformations.
