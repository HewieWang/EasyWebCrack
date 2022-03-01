package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "bufio"
    "io"
    "strconv"
    "net/http"
    "io/ioutil"
    "time"
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
    method := "GET"

    client := &http.Client {
      Timeout: 10 * time.Second,
    }
    req, err := http.NewRequest(method, url, nil)

    if err != nil {
      fmt.Println(err)
      return
    }
    req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36")
    req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
    req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
    res, err := client.Do(req)
    if err != nil {
      fmt.Println(err)
      return
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
      fmt.Println(err)
      return
    }
    html:=string(body)

    captchas := []string{"验证码", "验 证 码","点击更换", "点击刷新","看不清","认证码","安全问题"}
    for _ , s := range captchas {
        if strings.Contains(html,s) {
            log.Panic("Find captcha in "+url)
        }
    }

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
