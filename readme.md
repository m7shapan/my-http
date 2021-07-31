
# My HTTP
is a tool that makes http requests and prints the address of the request along with the MD5 hash of the response

My HTTP can: 
-  perform the requests in parallel
- limit the number of parallel requests, to prevent exhausting local resources. through accepting `parallel` flag to indicate this limit, and the default is 10 if the flag is not provided.
  
## Installation

Install my-http with git

```bash
  git clone https://github.com/m7shapan/my-http
  cd my-http
```
then build it
```bash
    go build -o my-http
```

Now it's ready to be used
    
## Usage Examples

basic usage

```bash
    ./my-http adjust.com
    # http://adjust.com bd8eb7f53fd65c6e1424d0b45e0d261a
```
multiple urls
```bash
    adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny
    http://google.com 0233b7f5cea5d7c50d32a60853cf9f56
    # http://adjust.com bd8eb7f53fd65c6e1424d0b45e0d261a
    # http://yandex.com e2b8d7219cc24b766a69ea337f24e805
    # http://twitter.com b7eaed42453945f8cd50b218cb481d23
    # http://facebook.com 3619f3c4c3cd8aac67bae09a2d3b1f5d
    # http://yahoo.com ee9a3d19e6bae369a51078e911e0b7b3
    # http://reddit.com/r/notfunny fac093d2bd8efcda9cf624141d9f9039
    # http://reddit.com/r/funny a37865a60c96e1d4b3d26499eaa2ca6f
```

```bash
    ./my-http -parallel=5 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny

    # http://google.com ace426ce35c142b2b1664f1bb377b9f3
    # http://adjust.com bd8eb7f53fd65c6e1424d0b45e0d261a
    # http://yandex.com 1f571d2d3ef4c92fb63c97701caa20fe
    # http://twitter.com 0e936742e85183c867ce0ebb6307d82b
    # http://facebook.com a634e3a6e57e8859609facb5f25d3543
    # http://yahoo.com 66cd87df2eebf7094120c3d3a2e04786
    # http://reddit.com/r/funny fb9833201352709d781eab8b067feca5
    # http://reddit.com/r/notfunny b800f70cbea079857156b41a348d570a

```



  