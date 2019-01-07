package mqttData

import (
	"time"
)

type Topic string

type Msg struct {
	payload   string
	timestamp time.Time
	topic     Topic
	from      User
}

type User struct {
	name    string
	subs    map[Topic]*Subscriber
	pubs    map[Topic]*Publisher
	ip      IP
	service *Service
}

func (u *User) GetName() string {
	return u.name
}

func NewUser(name string, service *Service) *User {
	return &User{
		name:    name,
		subs:    make(map[Topic]*Subscriber),
		pubs:    make(map[Topic]*Publisher),
		service: service,
	}
}

func (u *User) newSuber(topic Topic, callback func(data ...interface{}), args []interface{}) {
	sub := NewSubscriber(topic, callback, args)
	u.subs[topic] = sub

}
func (u *User) newPuber(topic Topic) *Publisher {
	pub := NewPublisher(topic, u.service)
	u.pubs[topic] = pub
	return pub
}

type Publisher struct {
	topic   Topic
	service *Service
}

func NewPublisher(topic Topic, service *Service) *Publisher {
	return &Publisher{topic: topic, service: service}
}

func (p *Publisher) publish(m Msg) {
	p.service.putMsg(m)
}

type Subscriber struct {
	topic    Topic
	callback func(data ...interface{})
	args     []interface{}
}

func NewSubscriber(topic Topic, callback func(data ...interface{}), args []interface{}) *Subscriber {
	return &Subscriber{topic: topic, callback: callback, args: args}
}

type IP string
type Service struct {
	msgChan chan Msg
	user    map[IP]*User
}

func NewService() *Service {
	s := new(Service)
	s.msgChan = make(chan Msg)
	s.user = make(map[IP]*User)
	return s
}
func (s *Service) getMsg() Msg {
	return <-s.msgChan
}
func (s *Service) running() {
	msg := s.getMsg()
	//fmt.Println(msg)
	s.broadcastMsg(msg)
	s.running()
}

func (s *Service) putMsg(msg Msg) {
	s.msgChan <- msg
}

func (s *Service) connect(user *User) {
	s.user[user.ip] = user
}

func (s *Service) getUsersSubTopic(topic Topic) []*User {
	users := make([]*User, 0)
	for _, user := range s.user {
		_, ok := user.subs[topic]
		if ok == true {
			users = append(users, user)
		}
	}
	return users
}

func (s *Service) broadcastMsg(msg Msg) {
	users := s.getUsersSubTopic(msg.topic)
	for _, user := range users {
		sub := user.subs[msg.topic]
		sub.callback(msg, sub.args)
	}
}
