This is a small software that does the following:
- Wait for any signals
- On a signal, sleep 10 secs, then query a url and log the output before exiting.

The goal was to monitor if docker shutdown's were handled gracefully when a machine stops.

The url to query and the waiting period are configurable. By default it would query Google, but it can be used to query a custom endpoint so that you are notified on a distant server of the shutdown happening (ie: if the machine hosting the container is supposed to be destroyed)

```
docker run -d MagicMicky/signals-to-http  25 http://mycustomurl/
```

