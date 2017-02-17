package main

import (
       "bufio"
       "fmt"
       "sort"
       "os"
)

func main() {
     // use bufio for buffered io, faster than fmt.Scan
     in := bufio.NewReader(os.Stdin)

     var len1 int
     fmt.Fscan(in, &len1)
     l1 := make([]int, 100)
     for i := 0; i < len1; i++ {
         var value int
         fmt.Fscan(in, &value)
         l1[value % 100]++
     }

     items := make([]int, 0)

     var len2 int
     fmt.Fscan(in, &len2)
     for i := 0; i < len2; i++ {
         var value int
         fmt.Fscan(in, &value)
         l1[value % 100]--

         if l1[value % 100] == -1 {
             items = append(items, value)
         }
     }

     sort.Ints(items)

     for _, item := range items {
         fmt.Printf("%d ", item)
     }

}
