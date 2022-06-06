package main

import (
"fmt"
"time"
. "fmt"
. "./keys"
)

func main() {
    chStop := make(chan int, 1)
    TimerFunc(chStop)

    time.Sleep(time.Hour * 24 * 365)
    chStop <- 0 // Tickerをstopさせるメッセージ

    close(chStop)

    time.Sleep(time.Second * 1)

    fmt.Println("Application End.")
}

func TimerFunc(stopTimer chan int) {
    go func() {
        ticker := time.NewTicker(24 * time.Hour) // 1日間隔のTicker

    LOOP:
        for {
            select {
            case <-ticker.C:
                api := GetTwitterApi()
                text := ChooseTweet()

                tweet, err := api.PostTweet(text, nil)
                if err != nil {
                    panic(err)
                }
                Print(tweet.Text)
                //}
            case <-stopTimer:
                fmt.Println("Timer stop.")
                ticker.Stop()
                break LOOP
            }
        }
        fmt.Println("timerfunc end.")
    }()
}