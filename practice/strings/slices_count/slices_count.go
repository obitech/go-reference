/**
* See: http://www.golangbootcamp.com/book/_single-page#uid91
* - Given a list of names, you need to organize each name within a slice based on its length.
**/

package slices_count

import (
    "unicode/utf8"
)

var names = []string{"aaa", "bbb", "cc"}

func strLen(str string) int {
    return utf8.RuneCountInString(str)
}

// Iterate over array and get lengths of longest and shortest strings
func minMax(names []string) (minLeng, maxLeng int) {
    minLeng, maxLeng = strLen(names[0]), 0

    for _, v := range names {
        leng := strLen(v)
        if leng < minLeng {
            minLeng = leng
        } else if leng > maxLeng {
            maxLeng = leng
        }
    }
    return
}

func Organize(names []string) [][]string {
    minLeng, maxLeng := minMax(names)
    var out [][]string
    for i := minLeng; i <= maxLeng; i++ {
        tmp := []string{}
        for _, v := range names {
            if strLen(v) == i {
                tmp = append(tmp, v)
            }
        }
        out = append(out, tmp)
    }

    return out
}
