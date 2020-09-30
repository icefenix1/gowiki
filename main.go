package main

import (
	request "github.com/icefenix1/gowiki/request"
)

func main() {
	en := request.Request("絵文字", "en")
	jp := request.Request("絵文字", "jp")
	request.Print(request.Combine(en, jp))
}
