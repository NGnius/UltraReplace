package main

import (
        "fmt"
        "flag"
        )
// command line flag values
var replacements string
var input_dir string
var output_dir string

func init(){
  // flag value defaults
  const (
      output_dir_d = "./output"
  )
  // flag usage
  const (
      output_dir_u = "the output directory"
  )
  // command line flags definitions
  flag.StringVar(&output_dir, "out", output_dir_d, output_dir_u)
}

func main(){
    fmt.Println("Not Implemented")
}
