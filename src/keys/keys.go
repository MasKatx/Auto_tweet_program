package keys

import (
"github.com/ChimeraCoder/anaconda"
)

func GetTwitterApi() *anaconda.TwitterApi {
    anaconda.SetConsumerKey("")
    anaconda.SetConsumerSecret("")
    api := anaconda.NewTwitterApi("", "")
    return api
}