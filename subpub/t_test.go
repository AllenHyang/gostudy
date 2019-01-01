package subpub

import (
	"fmt"
	"testing"
	"time"
)

func Test_pub(t *testing.T) {
	s := NewService()
	go func() {
		s.running()
	}()

	go func() {
		userp := NewUser("name_pub", s)
		p := userp.newPuber("/test")
		for i := 0; i < 10; i++ {

			msg := Msg{payload: fmt.Sprintf("%d here", i), timestamp: time.Now(), topic: "/test", from: *userp}
			p.publish(msg)
			time.Sleep(time.Second * 2)
		}
	}()

	user := NewUser("test1", s)
	user.newSuber("/test", callback, nil)
	s.connect(user)
	<-make(chan interface{})

}
func callback(data ...interface{}) {
	fmt.Print("I'm callback")
	fmt.Println(data...)
}
