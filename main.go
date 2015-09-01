package main
// could have been done with an httppie bash script, but this was not going to be fun.

import (
  "os"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
  "time"
  "math/rand"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  url := buildUrl()
  for i := 0; i < 5; i++ {
    if i > 0 {
      waitTime := rand.Intn(10) + 1;
      fmt.Printf("Pausing for %d seconds\n", waitTime)
      time.Sleep( time.Duration(waitTime) * time.Second)
    }
    sendMessage(url, "Hello World")
  }
}

func buildUrl() string {
  arguments := os.Args[1:]
  if len(arguments) != 2 {
    fmt.Println("Fatal: Missing host & port.")
    fmt.Printf("Usage: %s server port\n", os.Args[0])
    os.Exit(-1)
  }
  server := arguments[0]
  port := arguments[1]
  return fmt.Sprintf("http://%s:%s",server, port)
}

func sendMessage(url string, message string) {
  response, err := http.Post(url, "text/plain", toBuffer(message))
  if err != nil {
    panic(err)
  }
  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    panic(err)
  }
  fmt.Println("<-- ", string(body))
}

func toBuffer(message string) *bytes.Buffer {
  fmt.Println("--> ", message)
  return bytes.NewBuffer([]byte(message))
}
