# hiven-gw

This can be useful for microservicing hiven bots.

# how 2 use

In order to accomplish 0 loss on messages, hiven-gw uses `RPUSH` and `BLPOP` in redis.
Now create a `.env` and put the contents in

```bash
TOKEN="your hiven token (found in localStorage)"
REDIS="localhost:6379"
LIST="gateway"
ZLIB=true # enable if you would like to recieve compressed zlib from hiven
DEBUG=true # enable if you want to recieve debug msgs (op, hb)
```

Then you will need to build hiven-gw by doing `go build`
Now you can run the executable it creates: `./hivengw`

# recieving

You can run `BLPOP gateway 0` in redis to get data. You can do a while loop that does that command.
