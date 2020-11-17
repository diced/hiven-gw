# hiven-gw

This can be useful for microservicing hiven bots.

# how 2 use
In order to accomplish 0 loss on messages, hiven-gw uses `RPUSH` and `BLPOP` in redis.
Now create a `.env` and put the contents in
```bash
TOKEN="your hiven token (found in localStorage)"
REDIS="localhost:6379"
LIST="gateway"
```
Then you will need to build hiven-gw by doing `go build`
Now you can run the executable it creates: `./higw`

# recieving
You can run `BLPOP gateway 0` in redis to get data. You can do a while loop that does that command.