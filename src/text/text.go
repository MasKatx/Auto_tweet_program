package text

import (
    "time"
    "fmt"
)

func ChooseTweet()string{

    TweetList :=  [10]string{
        "1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"0"
    }

    d := time.Now().Day()
    m := time.Now().Month()

    TweetContent := TweetList[(d-1)%10]
    TweetOfToday := fmt.Sprintf("【%d月%d日】\n %s",m,d, TweetContent)
    return TweetOfToday
}