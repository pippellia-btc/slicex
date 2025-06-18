# slicex

`slicex` is a Go package providing performance-optimized utility functions for slice manipulation not directly available in the `slices` standard library's.

## Installation
```
go get github.com/pippellia-btc/slicex
```

## Features
- **Efficient Set Operations**:   
Implements common set operations (e.g., difference, partition, symmetric difference) for slices, leveraging optimized algorithms for enhanced performance compared to a map-based approach.

- **Top Element Discovery**:   
Includes a custom algorithm for efficiently finding the top-K elements within a slice

- **Generic Support**:   
All functions are built with Go generics, ensuring type safety and reusability across various comparable and ordered data types.

## Purpose

While Go's standard slices package provides foundational utilities, slicex aims to complement it by offering specialized and performant functions that address frequently encountered challenges in data processing. It's ideal for projects requiring highly efficient slice transformations and analytical operations.
