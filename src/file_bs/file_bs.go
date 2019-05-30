package main

import (
    "fmt"
    "flag"
    "math/rand"
    "os"
    "strings"
    "strconv"
    "sync"
)

// cl flag values
var output_dir string
var isTest bool
var file_count int
var file_length_range string

var length_min int
var length_max int
var length_delta int

func init(){
  // flag value defaults
  const (
    output_dir_d = "./output"
    isTest_d = false
    file_count_d = 4
    file_length_range_d = "256"
  )
  // flag usage descriptions
  const (
    output_dir_u = "the output directory"
    isTest_u = "collect & print test statistics"
    file_count_u = "the number of files to generate"
    file_length_range_u = "the length of each file. Ranges (n-m) are also accepted"
  )
  // command line flags definitions
  flag.StringVar(&output_dir, "out", output_dir_d, output_dir_u)
  flag.BoolVar(&isTest, "test", isTest_d, isTest_u)
  flag.IntVar(&file_count, "count", file_count_d, file_count_u)
  flag.StringVar(&file_length_range, "length", file_length_range_d, file_length_range_u)
}

func main(){
    flag.Parse()
    if (isTest) {
      fmt.Println("Test statistics activated (WIP)")
    }
    parse_file_length()
    os.MkdirAll(output_dir, os.ModePerm)
    var wg sync.WaitGroup
    wg.Add(file_count)
    for i := 0; i<file_count; i++ {
        var filename = output_dir+"/file"+strconv.Itoa(i)+".txt"
        go generate_file(filename, &wg)
    }
    wg.Wait()
}

func parse_file_length(){
    if (strings.Contains(file_length_range, "-")){
      var err, err2 error
      var length_range = strings.Split(file_length_range, "-")
      length_min, err = strconv.Atoi(length_range[0])
      length_max, err2 = strconv.Atoi(length_range[1])
      length_delta = length_max - length_min
      if (err != nil){
        panic(err)
      }
      if (err2 != nil){
        panic(err2)
      }
    } else {
      var err error
      length_min, err = strconv.Atoi(file_length_range)
      length_max = length_min
      length_delta = 0
      if (err != nil){
        panic(err)
      }
    }
}

func generate_file(filename string, wg *sync.WaitGroup){
  defer wg.Done()
  f, err := os.Create(filename)
  defer f.Close()
  if (err != nil){
    fmt.Println("Failed to create "+filename)
    return
  }
  var file_str = ""
  var str_len int
  if (length_delta == 0){
    str_len = length_min
  } else {
    str_len = rand.Intn(length_delta) + length_min
  }
  for i:=0; i<str_len; i++ {
    file_str += string(rand.Intn(255))
  }
  f.WriteString(file_str)
  f.Sync()
}
