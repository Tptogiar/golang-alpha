package x_time_rate

import (
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

func TestAlpha(t *testing.T) {

	// 初始化一个限速器，每秒产生10个令牌，桶的大小为100个
	// 初始化状态桶是满的
	var limiter = rate.NewLimiter(10, 100)

	for i := 0; i < 20; i++ {
		if limiter.AllowN(time.Now(), 25) {
			fmt.Printf("%03d Ok  %sn", i, time.Now().Format("2006-01-02 15:04:05.000"))
		} else {
			fmt.Printf("%03d Err %sn", i, time.Now().Format("2006-01-02 15:04:05.000"))
		}
		time.Sleep(500 * time.Millisecond)
	}

}
