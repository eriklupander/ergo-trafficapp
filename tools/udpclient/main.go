package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/vmihailenco/msgpack/v5"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type GeoPos struct {
	ID        string
	Lon       float64
	Lat       float64
	Emergency bool
}

func main() {
	if len(os.Args) < 2 {
		panic("must pass address argument")
	}
	// Resolve the string address to a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	workers := 2
	if len(os.Args) == 3 {
		workers, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
	total := atomic.Int64{}
	go func() {
		fmt.Printf("UDP packets sent: ")
		for {
			fmt.Printf("%d ... ", total.Load())
			time.Sleep(time.Second * 5)
		}
	}()
	wg := sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		// Dial to the address with UDP
		conn, err := net.DialUDP("udp", nil, udpAddr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		id := uuid.NewString()

		wg.Add(1)
		go func(i int) {

			ticker := time.NewTicker(time.Second)
			timeout := time.After(time.Hour)
			for {
				select {
				case <-ticker.C:
					msg, err := msgpack.Marshal(GeoPos{
						ID:        id,
						Lon:       rndPos(),
						Lat:       rndPos(),
						Emergency: rand.Intn(1000)%1000 == 0, // 1 per 1000 events is emergency.
					})
					if err != nil {
						panic(err.Error())
					}

					time.Sleep(time.Millisecond * time.Duration(rand.Intn(500))) // sleep between 0 to 500 ms so we don't burst all events at the same time.
					// Send a message to the server
					_, err = conn.Write(msg)

					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					total.Add(1)
				case <-timeout:
					wg.Done()
					return
				}
			}
		}(i)
	}
	wg.Wait()
}

func rndPos() float64 {
	return -90 + 180*rand.Float64()
}
