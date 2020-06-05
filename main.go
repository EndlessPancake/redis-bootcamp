package main

import (
  "github.com/garyburd/redigo/redis"
  "fmt"
  "os"
  "bufio"
  // "time"
)

func redis_set(key string, value string, c redis.Conn){
  // c.Do("SET", key, value)
  c.Do("PFADD", key, value)
}

// func redis_get(key string, c redis.Conn) string{
func redis_get(key string, c redis.Conn) int{
  // s, err := redis.String(c.Do("GET", key))
  s, err := redis.Int(c.Do("PFCOUNT", key))
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return s
}

func redis_connection() redis.Conn {
  const IP_PORT = "0.0.0.0:6379"
  c, err := redis.Dial("tcp", IP_PORT)
  if err != nil {
    panic(err)
  }
  return c
}

func main() {
  c := redis_connection()
  defer c.Close()

  var fp *os.File
  var err error
  // var key = "kyufu_paper"
  var key = os.Args[2]

  if len(os.Args) < 3 {
     fp  = os.Stdin
  } else {
     fp, err = os.Open(os.Args[1])
     if err != nil {
        panic(err)
     }
     defer fp.Close()
  }

  scanner := bufio.NewScanner(fp)
  for scanner.Scan() {
      // fmt.Println(scanner.Text())
      redis_set(key, scanner.Text(), c)
      // time.Sleep(time.Millisecond * 200)
  }
  if err := scanner.Err(); err != nil {
     panic(err)
  }

  // PFCount
  s := redis_get(key, c)
  fmt.Println(s)
}
