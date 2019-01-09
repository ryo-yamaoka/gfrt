[![CircleCI](https://circleci.com/gh/ryo-yamaoka/gfrt.svg?style=svg)](https://circleci.com/gh/ryo-yamaoka/gfrt)

# gfrt

GFRT (Go Feed Redirect Tester) is a tool of RSS feed redirect test

## Usage

* `gfrt`: Start server
  * `-v`: Print software version
  * `-p INT`: Designate listen port number (default: 80)
  * `-r URL`: Disignate redirect destination URL (default: http://www.example.com/)

* Environment variables
  * GFRT_EXTERNAL_HOSTNAME
    * It sets to external link for feed and article link
    * If you need to change listen port, this is must include port number (ex: `192.0.2.1:8080`)
    * default(not set): `127.0.0.1`

### Start server

```bash
$ ./gfrt -p 8080
2019/01/08 10:38:14 listening: 8080
```

### Get RSS feed or try redirect

If you want to redirect tracking, you need to set `-L` curl option.

```bash
$ curl -L -X GET http://localhost:8080/feed
```

### Switch mode

Switch to redirect mode

```bash
$ curl -X PUT http://localhost:8080/feed
```

Back to RSS feed mode

```bash
$ curl -X DELETE http://localhost:8080/feed
```
