package leetcode

import "strings"

func wordBreak(s string, wordDict []string) bool {
    memCache := map[string]bool{}
    return helper(s, wordDict, memCache)
}

func helper(s string, wordDict []string, memCache map[string]bool) bool {
    val, ok := memCache[s]
    if ok {
        return val
    }
    if s == "" {
        return true
    }

    for _, word := range wordDict {
        if strings.HasPrefix(s, word) {
            suffix := s[len(word):]
            if helper(suffix, wordDict, memCache) {
                memCache[s] = true
                return true
            }
        }
    }
    memCache[s] = false
    return false
}
