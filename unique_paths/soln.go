package leetcode

type key struct {
    m int
    n int
}

var memCache = make(map[key]int)

func uniquePaths(m int, n int) int {
    k := key{
        m: m,
        n: n,
    }
    if m == 1 && n == 1 {
        return 1
    }
    if m == 0 || n == 0 {
        return 0
    }
    val, ok := memCache[k]
    if ok {
        return val
    }
    val, ok = memCache[key{m:n, n:m}]
    if ok {
        return val
    }
   memCache[k] = uniquePaths(m-1, n) + uniquePaths(m, n-1)
   return memCache[k]
}