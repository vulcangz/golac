# golac

`golac` is a Golang wrapper for [PaddleHub LAC](https://www.paddlepaddle.org.cn/hubdetail?name=lac&en_category=LexicalAnalysis). 

## Install

Download and install it:

```shell
go get github.com/vulcangz/golac
```

Make sure that you can run PaddleHub LAC on [command line](https://github.com/PaddlePaddle/PaddleHub#%E5%BF%AB%E9%80%9F%E4%BD%93%E9%AA%8C):

```shell
hub run lac --input_text "今天是个好日子"
```

## Usage

A simple code for using `golac` is:

```go
package main

import (
  "fmt"
	
  "github.com/vulcangz/golac" // exposes "golac"
)

func main() {
  text := `天气预报说今天要下雨`

  // LocalExec connector is responsible to run PaddleHub LAC process.
  c := golac.NewLocalExec(nil)
  c.Option = "--input_text"   // default option

  // LAC text
  doc, err := c.Run(text)
  if err != nil {
    fmt.Println(err.Error())
  }
  
  d, _ := golac.Decode(doc)
  for _, v := range d {
    for i := 0; i < len(v.Word); i++ {
      fmt.Printf("%s(%s) ", v.Word[i], v.Tag[i])
    }
  }
}

```

Output:

```text
天气预报(n) 说(v) 今天(TIME) 要(v) 下雨(v)
```

## LICENSE

MIT
