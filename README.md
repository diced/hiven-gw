# hiven-gw

This can be useful for microservicing hiven bots.

# how 2 use

In order to accomplish 0 loss on messages, hiven-gw uses `RPUSH` and `BLPOP` in redis.
Now create a `.env` and put the contents in

```bash
TOKEN="your hiven token (found in localStorage)"
REDIS="localhost:6379"
LIST="gateway"
DISABLED_EVENTS="TYPING_START" # comma separeted list of disabled events (will not push to redis)
ZLIB=true # enable if you would like to recieve compressed zlib from hiven
DEBUG=true # enable if you want to recieve debug msgs (op, hb)
```

Then you will need to build hiven-gw by doing `go build`
Now you can run the executable it creates: `./hivengw`

# recieving

You can run `BLPOP gateway 0` in redis to get data. You can do a while loop that does that command.

# fetching gateway runtime stats

Gateway runtime stats are sent every 30 seconds to the configured list (e.g. `gateway`) (on every heartbeat). The payload will look like this:
_Note: stats event can't be disabled with DISABLED_EVENTS_

```js
{
  "e": "stats",
  "stats": {
    Alloc: 17310048,
    TotalAlloc: 36768752,
    Sys: 74793984,
    Lookups: 0,
    Mallocs: 274900,
    Frees: 266333,
    HeapAlloc: 17310048,
    HeapSys: 66715648,
    HeapIdle: 47513600,
    HeapInuse: 19202048,
    HeapReleased: 46555136,
    HeapObjects: 8567,
    StackInuse: 393216,
    StackSys: 393216,
    MSpanInuse: 95608,
    MSpanSys: 114688,
    MCacheInuse: 6944,
    MCacheSys: 16384,
    BuckHashSys: 1448755,
    GCSys: 5071632,
    OtherSys: 1033661,
    NextGC: 18121728,
    LastGC: 1605726071857147600,
    PauseTotalNs: 757207,
    PauseNs: [],
    PauseEnd: [],
    NumGC: 6,
    NumForcedGC: 0,
    GCCPUFraction: 0.00010876622264406884,
    EnableGC: true,
    DebugGC: false,
    BySize: []
  }
}
```
