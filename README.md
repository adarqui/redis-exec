# redis-exec

Execute commands based on redis pubsub or popping queues.

## usage (for now)

```
./redis-exec 127.0.0.1:6379 resque:queue:images /tools/test.sh
```

test.sh (command) takes two arguments: queue and payload

## WIP

Just whipped this up real quick for something I need. I can clean it up later.
