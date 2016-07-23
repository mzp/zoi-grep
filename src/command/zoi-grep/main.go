package main

import (
  "fmt"
  "zoi"
  "os"
)

func main() {
  entries, err := zoi.DefaultPanels()

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  for _, v := range zoi.Filter(os.Args[1:], entries) {
    image, err := zoi.ReadImage(v.Path)

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    fmt.Println(zoi.InlineImage(image))
  }
}
