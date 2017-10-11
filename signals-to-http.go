package main

import "fmt"
import "os"
import "os/signal"
import . "time"
import "net/http"
import "strings"
import "strconv"
const DEFAULT_URL string ="http://google.com"
const DEFAULT_SECS int =9

func queryUrl(url string) {
  fmt.Printf("getting %s\n", url)
  resp, err := http.Get(url)
	if resp != nil {
    fmt.Println(resp)
    defer resp.Body.Close()
  } else {
    fmt.Println(err)
  }
}

func getUrl() string {
  ret_url := DEFAULT_URL
  if len(os.Args) > 2 {
    url := strings.Join(os.Args[2:], "")
    if len(url) > 0 {
      ret_url=url
    }
  }
  return ret_url
}

func getNbSec() int {
  ret_secs:=DEFAULT_SECS
  if len(os.Args) > 1 && os.Args[1] != "" {
    res, err := strconv.Atoi(os.Args[1])
    if err == nil {
      ret_secs=res
    }
  }
  return ret_secs
}

func main() {
  delay := Duration(getNbSec()) * Second
  sigs := make(chan os.Signal, 1)
  done := make(chan bool, 1)
  args := getUrl()

  signal.Notify(sigs)
	go func() {
		sig := <-sigs
		fmt.Println()
    fmt.Printf("Received signal %s, sleeping for %d secs\n", sig, int(delay.Seconds()))
    Sleep(delay)
		queryUrl(args)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
