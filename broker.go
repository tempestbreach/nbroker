package nbroker

// MAKE "UPSTREAM" AND "DOWNSTREAM" BROKERS? Or Routers?

import(
    "log"

    "github.com/tempestbreach/augnats"
    "github.com/nats-io/nats.go"
)

type BrokerConfig struct {
    Handler                 augnats.Handler
    Conn                    *nats.Conn
    EncConn                 *nats.EncodedConn
    // Maybe omit next
    ListenSubject           string
}

type Broker struct {
    handler                 augnats.Handler
    conn                    *nats.Conn
    eConn                   *nats.EncodedConn
    SelfSubject             string
    ListenSubject           string
}

// type Broker struct {
//     downstreamHandler           augnats.Handler
//     upstreamHandler             augnats.Handler
//     conn                        *nats.Conn
//     eConn                       *nats.EncodedConn
//     ListenSubject               string
// }

func NewBroker(bc BrokerConfig) *Broker {

    b := &Broker{
        handler: bc.Handler,
        conn: bc.Conn,
        eConn: bc.EncConn,
        ListenSubject: bc.ListenSubject,
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

        // msg :=

    }

    sub.Unsubscribe()
    sub.Drain()
}
