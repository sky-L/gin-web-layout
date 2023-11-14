package test

import (
	"context"
	"fmt"
	"github.com/avast/retry-go/v4"
	"github.com/briandowns/spinner"
	"github.com/skylee/gin-web-layout/utils"
	"net"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {

	conn, _ := net.Dial("udp", "8.8.8.8:53")
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	fmt.Println(localAddr.String())

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Color("yellow")

	println("🔍 find main module:\n")
	s.Suffix = " find module information..."
	s.Start()

	s.Stop()

	b, _ := utils.New(context.Background(), utils.WithConcurrencyNum(10))

	b.Go("l", func() (ret interface{}, err error) {
		err = retry.Do(
			func() error {
				return nil
			},
		)
		return
	})

	result, bErr := b.WaitAndGetResult()

	fmt.Println(result, bErr)
}
