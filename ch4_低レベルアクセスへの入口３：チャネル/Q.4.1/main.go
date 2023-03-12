package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// コマンドライン引数を定義
	var duration time.Duration
	flag.DurationVar(&duration, "duration", 5*time.Second, "duration to wait before sending the notification")

	// コマンドライン引数を解析
	flag.Parse()

	// 指定した時間が経過したらチャネルにメッセージが送信される
	timer := time.After(duration)

	// チャネルからメッセージを受信するまで待つ
	<-timer

	// 時間が経過したら実行される処理
	fmt.Println("The notification has been sent after", duration.Seconds(), "seconds.")
}
