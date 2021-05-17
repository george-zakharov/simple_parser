package main

import (
    "fmt"
    "flag"
    "strings"
    "time"
    "github.com/opesun/goquery"
)

var (
    WORKERS       int       = 2            // workers number
    REPORT_PERIOD int       = 10           // reports frequency (seconds)
    DUP_TO_STOP   int       = 500          // max requests till stop
    HASH_FILE     string    = "hash.bin"   // hashes file
    QUOTES_FILE   string    = "quotes.txt" // quotes file
)

func main() {
    fmt.Println("Hello")

    initFlags()

    ch := make(chan int)
    go send(ch)
    b := <- ch
    fmt.Println(b)
}

func send(c chan int) {
    c <- 15
}

// Get chanel to read from data with string type
func grab() <- chan string {
    c := make(chan string)
    // Creating gorutines
    for i := 0; i < WORKERS; i++ {
        go func () {
            // Collecting data in the loop
            for {
                x, err := goquery.ParseUrl("http://vpustotu.ru/moderation/")
                if err == nil {
                    if s := strings.TrimSpace(x.Find(".fi_text").Text()); s != "" {
                        c <- s // Send data into the chanel
                    }
                }

                time.Sleep(100 * time.Millisecond)
            }
        }()
    }

    fmt.Println("Launch threads: ", WORKERS)
    return c
}

func initFlags() {
    // Add flags (lookup with -h parameter)
    flag.IntVar(&WORKERS, "w", WORKERS, "workers number")
    flag.IntVar(&REPORT_PERIOD, "r", REPORT_PERIOD, "reports frequency (seconds)")
    flag.IntVar(&DUP_TO_STOP, "d", DUP_TO_STOP, "max requests till stop")
    flag.StringVar(&HASH_FILE, "hf", HASH_FILE, "hashes file")
    flag.StringVar(&QUOTES_FILE, "qf", QUOTES_FILE, "quotes file")
    // Parse flags and set them to vars with & symbol
    flag.Parse()
}