# elephant
Cache like elephant

## Objective
I needed a flexible way to cache requests to a paid API because my SOAR solution used to consume the quota for requests with the same values. I thought I could create a cache system that retrieves responses from a local version if the same value was recently requested. Nginx and Caddy server caches were difficult to manage because I had no control over cache invalidation. Therefore, this tool uses Badger KV DB to store cached responses with customized time-to-live configurations.

## setup

```
git clone https://github.com/alwashali/elephant.git
cd elephant
go run . 
```

```
NAME:
   Elephant cache - Run the server with TTL to cache everything passes through

USAGE:
   Elephant cache [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h   show help (default: false)
   --ttl value  Time to live for the cache
```

## Time to live option 
Run the tool with --ttl option and specify the duration for the cache live time.

```
go run . run --ttl 1d 
```

