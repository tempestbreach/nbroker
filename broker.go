package nbroker

import(
    "log"

    "github.com/nats-io/nats.go"
)

type Handler interface {
    HandleMsg(msg *nats.Msg)
}

type Broker struct {
    handler                 Handler
    conn                    *nats.Conn
    SelfSubject             string
    ListenSubject           string
}

func NewBroker(h Handler, ls string) *Broker {
    b := &Broker{
        handler: h,
        ListenSubject: ls,
    }
    return b
}

func(b *Broker) ListenAndPublish() {

    chanRecv := make(chan *nats.Msg, 64)
    sub, err := b.conn.ChanSubscribe(b.ListenSubject, chanRecv)
    if err != nil {
        log.Fatal(err)
    }

    for {
        msg := <-chanRecv
        b.handler.HandleMsg(msg)
    }

    sub.Unsubscribe()
    sub.Drain()
}
