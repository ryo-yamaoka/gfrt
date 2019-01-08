[![CircleCI](https://circleci.com/gh/ryo-yamaoka/gfrt.svg?style=svg)](https://circleci.com/gh/ryo-yamaoka/gfrt)

# gfrt

GFRT (Go Feed Redirect Tester) is a tool of RSS feed redirect test

## Usage

* `gfrt`: start server
  * `-v`: print software version
  * `-p INT`: designate listen port number (default: 80)
  * `-r URL`: disignate redirect destrination URL (default: http://www.example.com/)

### Start server

```bash
$ ./gfrt -p 8080
2019/01/08 10:38:14 listening: 8080
```


### Get RSS feed or try redirect

If you want to redirect tracking, you need set `-L` curl option.

```bash
$ curl -L -X GET http://localhost:8080/
```

### Switch mode

Switch to redirect mode

```bash
$ curl -X PUT http://localhost:8080/
```

Back to RSS feed mode

```bash
$ curl -X DELETE http://localhost:8080/
```
