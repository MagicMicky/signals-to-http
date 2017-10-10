package main

import "fmt"
import "os"
import "os/signal"
import . "time"
import "net/http"
import "strings"

func getGoogle(url string) {
  fmt.Printf("getting %s\n", url)
  resp, err := http.Get(url)
	if resp != nil {
    fmt.Println(resp)
    defer resp.Body.Close()
  } else {
    fmt.Println(err)
  }
}

func getUrl(args []string) string {
  url := strings.Join(args, "")
  if len(url) == 0 {
    return "http://google.com"
  }
  return url
}

func main() {
  delay := 9 * Second
  sigs := make(chan os.Signal, 1)
  done := make(chan bool, 1)
  args := getUrl(os.Args[1:])

  signal.Notify(sigs)
	go func() {
		sig := <-sigs
		fmt.Println()
    fmt.Printf("Received signal %s, sleeping for %d secs\n", sig, int(delay.Seconds()))
    Sleep(delay)
		getGoogle(args)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
