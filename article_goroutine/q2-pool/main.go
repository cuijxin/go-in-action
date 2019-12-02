package main

import (
	"c_program_language/patterns/pool"
)

const (
	maxGoroutines = 25 // 要使用的goroutine的数量
    pooledResources = 2 // 池中的资源的数量
)

// dbConnection模拟要共享的资源
type dbConnection struct {
	ID int32
}

