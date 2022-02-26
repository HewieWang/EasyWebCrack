package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "bufio"
    "io"
    "strconv"
)

var author string = `
*****************************************************
*                                                   *
****************    Code By HewieWang   *************
*                                                   *
*****************************************************
`

func get_form() {
  
}

func crack(url string,count int)  {
    countString := strconv.Itoa(count)
    fmt.Println("task: "+countString)
    fmt.Println(url)
    fmt.Println("")
}

func main() {
    fmt.Println(author)
    if strings.HasPrefix(os.Args[1], "http") {
        crack(os.Args[1],1)
    } else {
      fi, err := os.Open(os.Args[1])
      if err != nil {
          log.Printf("Error: %s\n", err)
          return
      }
      defer fi.Close()

      br := bufio.NewReader(fi)
      count:=0
      for {
          a, _, c := br.ReadLine()
          if c == io.EOF {
              break
          }
          count++
          crack(string(a),count)
      }
    }
}
