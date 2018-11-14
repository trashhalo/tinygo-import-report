
# Tinygo Import Report
This project imports each package in the stdlib and reports if it imports cleanly in tinygo. 
A package with a check may work a package with a x will definately not currently work.
If a package has a x in imported? but check in importedWithInit? it means you need to build your project
with -initinterp.

| Package | Imported? | ImportedWithInit? |
| --- | --- | --- |
| archive/tar |  [:x:](#archivetar)  |  [:x:](#archivetar)  | 
| archive/zip |  [:x:](#archivezip)  |  [:x:](#archivezip)  | 
| bufio |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| bytes |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| compress/bzip2 |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| compress/flate |  [:x:](#compressflate)  |  :heavy_check_mark:  | 
| compress/gzip |  [:x:](#compressgzip)  |  :heavy_check_mark:  | 
| compress/lzw |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| compress/zlib |  [:x:](#compresszlib)  |  :heavy_check_mark:  | 
| container |  [:x:](#container)  |  [:x:](#container)  | 
| container/heap |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| container/list |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| container/ring |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| context |  [:x:](#context)  |  [:x:](#context)  | 
| crypto |  [:x:](#crypto)  |  :heavy_check_mark:  | 
| crypto/aes |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| crypto/cipher |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| crypto/des |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| crypto/dsa |  [:x:](#cryptodsa)  |  [:x:](#cryptodsa)  | 
| crypto/ecdsa |  [:x:](#cryptoecdsa)  |  [:x:](#cryptoecdsa)  | 
| crypto/elliptic |  [:x:](#cryptoelliptic)  |  [:x:](#cryptoelliptic)  | 
| crypto/hmac |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| crypto/md5 |  [:x:](#cryptomd5)  |  [:x:](#cryptomd5)  | 
| crypto/rand |  [:x:](#cryptorand)  |  [:x:](#cryptorand)  | 
| crypto/rc4 |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| crypto/rsa |  [:x:](#cryptorsa)  |  [:x:](#cryptorsa)  | 
| crypto/sha1 |  [:x:](#cryptosha1)  |  [:x:](#cryptosha1)  | 
| crypto/sha256 |  [:x:](#cryptosha256)  |  [:x:](#cryptosha256)  | 
| crypto/sha512 |  [:x:](#cryptosha512)  |  [:x:](#cryptosha512)  | 
| crypto/subtle |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| crypto/tls |  [:x:](#cryptotls)  |  [:x:](#cryptotls)  | 
| crypto/x509/pkix |  [:x:](#cryptox509pkix)  |  [:x:](#cryptox509pkix)  | 
| database/sql |  [:x:](#databasesql)  |  [:x:](#databasesql)  | 
| database/sql/driver |  [:x:](#databasesqldriver)  |  [:x:](#databasesqldriver)  | 
| encoding |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| encoding/ascii85 |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| encoding/asn1 |  [:x:](#encodingasn1)  |  [:x:](#encodingasn1)  | 
| encoding/base32 |  [:x:](#encodingbase32)  |  [:x:](#encodingbase32)  | 
| encoding/base64 |  [:x:](#encodingbase64)  |  [:x:](#encodingbase64)  | 
| encoding/binary |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| encoding/csv |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| encoding/gob |  [:x:](#encodinggob)  |  [:x:](#encodinggob)  | 
| encoding/hex |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| encoding/json |  [:x:](#encodingjson)  |  [:x:](#encodingjson)  | 
| encoding/pem |  [:x:](#encodingpem)  |  [:x:](#encodingpem)  | 
| encoding/xml |  [:x:](#encodingxml)  |  [:x:](#encodingxml)  | 
| errors |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| expvar |  [:x:](#expvar)  |  [:x:](#expvar)  | 
| flag |  [:x:](#flag)  |  [:x:](#flag)  | 
| fmt |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| hash |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| hash/adler32 |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| hash/crc32 |  [:x:](#hashcrc32)  |  :heavy_check_mark:  | 
| hash/crc64 |  [:x:](#hashcrc64)  |  :heavy_check_mark:  | 
| hash/fnv |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| html |  [:x:](#html)  |  [:x:](#html)  | 
| html/template |  [:x:](#htmltemplate)  |  [:x:](#htmltemplate)  | 
| image |  [:x:](#image)  |  [:x:](#image)  | 
| image/color |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| image/color/palette |  [:x:](#imagecolorpalette)  |  [:x:](#imagecolorpalette)  | 
| image/draw |  [:x:](#imagedraw)  |  [:x:](#imagedraw)  | 
| image/gif |  [:x:](#imagegif)  |  [:x:](#imagegif)  | 
| image/jpeg |  [:x:](#imagejpeg)  |  [:x:](#imagejpeg)  | 
| image/png |  [:x:](#imagepng)  |  [:x:](#imagepng)  | 
| index/suffixarray |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| io |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| io/ioutil |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| log |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| log/syslog |  [:x:](#logsyslog)  |  [:x:](#logsyslog)  | 
| math |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| math/big |  [:x:](#mathbig)  |  [:x:](#mathbig)  | 
| math/bits |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| math/cmplx |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| math/rand |  [:x:](#mathrand)  |  [:x:](#mathrand)  | 
| mime |  [:x:](#mime)  |  [:x:](#mime)  | 
| mime/multipart |  [:x:](#mimemultipart)  |  [:x:](#mimemultipart)  | 
| mime/quotedprintable |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| net |  [:x:](#net)  |  [:x:](#net)  | 
| net/http |  [:x:](#nethttp)  |  [:x:](#nethttp)  | 
| net/http/cgi |  [:x:](#nethttpcgi)  |  [:x:](#nethttpcgi)  | 
| net/http/cookiejar |  [:x:](#nethttpcookiejar)  |  [:x:](#nethttpcookiejar)  | 
| net/http/fcgi |  [:x:](#nethttpfcgi)  |  [:x:](#nethttpfcgi)  | 
| net/http/httptest |  [:x:](#nethttphttptest)  |  [:x:](#nethttphttptest)  | 
| net/http/httptrace |  [:x:](#nethttphttptrace)  |  [:x:](#nethttphttptrace)  | 
| net/http/httputil |  [:x:](#nethttphttputil)  |  [:x:](#nethttphttputil)  | 
| net/http/pprof |  [:x:](#nethttppprof)  |  [:x:](#nethttppprof)  | 
| net/mail |  [:x:](#netmail)  |  [:x:](#netmail)  | 
| net/rpc |  [:x:](#netrpc)  |  [:x:](#netrpc)  | 
| net/rpc/jsonrpc |  [:x:](#netrpcjsonrpc)  |  [:x:](#netrpcjsonrpc)  | 
| net/smtp |  [:x:](#netsmtp)  |  [:x:](#netsmtp)  | 
| net/textproto |  [:x:](#nettextproto)  |  [:x:](#nettextproto)  | 
| net/url |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| os |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| os/exec |  [:x:](#osexec)  |  [:x:](#osexec)  | 
| os/signal |  [:x:](#ossignal)  |  [:x:](#ossignal)  | 
| os/user |  [:x:](#osuser)  |  [:x:](#osuser)  | 
| path |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| path/filepath |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| plugin |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| reflect |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| regexp |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| regexp/syntax |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| runtime |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| runtime/cgo |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| runtime/debug |  [:x:](#runtimedebug)  |  [:x:](#runtimedebug)  | 
| runtime/msan |  [:x:](#runtimemsan)  |  [:x:](#runtimemsan)  | 
| runtime/pprof |  [:x:](#runtimepprof)  |  [:x:](#runtimepprof)  | 
| runtime/race |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| runtime/trace |  [:x:](#runtimetrace)  |  [:x:](#runtimetrace)  | 
| sort |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| strconv |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| strings |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| sync |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| sync/atomic |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| syscall |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| syscall/js |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| testing |  [:x:](#testing)  |  [:x:](#testing)  | 
| testing/iotest |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| testing/quick |  [:x:](#testingquick)  |  [:x:](#testingquick)  | 
| text/scanner |  [:x:](#textscanner)  |  :heavy_check_mark:  | 
| text/tabwriter |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| text/template |  [:x:](#texttemplate)  |  [:x:](#texttemplate)  | 
| text/template/parse |  [:x:](#texttemplateparse)  |  [:x:](#texttemplateparse)  | 
| time |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| unicode |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| unicode/utf16 |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| unicode/utf8 |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| unsafe |  :heavy_check_mark:  |  :heavy_check_mark:  | 



## archive/tar
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/archive_tar:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/archive/tar/common.go:445:15: DeepEqual not declared by package reflect
error: couldn't load packages due to errors: archive/tar
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/archive_tar:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/archive/tar/common.go:445:15: DeepEqual not declared by package reflect
error: couldn't load packages due to errors: archive/tar
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## archive/zip
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/archive_zip:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/archive/zip/register.go:106:21: Map not declared by package sync
/usr/local/go/src/archive/zip/register.go:107:21: Map not declared by package sync
error: couldn't load packages due to errors: archive/zip
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/archive_zip:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/archive/zip/register.go:106:21: Map not declared by package sync
/usr/local/go/src/archive/zip/register.go:107:21: Map not declared by package sync
error: couldn't load packages due to errors: archive/zip
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## bufio
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/bufio:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/bufio:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## bytes
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/bytes:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/bytes:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## compress/bzip2
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/compress_bzip2:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/compress_bzip2:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## compress/flate
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/compress_flate:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: init call: compress/flate.generateFixedOffsetEncoding
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/compress_flate:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## compress/gzip
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/compress_gzip:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: init call: compress/flate.generateFixedOffsetEncoding
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/compress_gzip:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## compress/lzw
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/compress_lzw:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/compress_lzw:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## compress/zlib
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/compress_zlib:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: init call: compress/flate.generateFixedOffsetEncoding
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/compress_zlib:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## container
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/container:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/go/src/github.com/trashhalo/tinygo-import-report/main.go:5:4: could not import container (invalid package name: "")
error: couldn't load packages due to errors: github.com/trashhalo/tinygo-import-report
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/container:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/go/src/github.com/trashhalo/tinygo-import-report/main.go:5:4: could not import container (invalid package name: "")
error: couldn't load packages due to errors: github.com/trashhalo/tinygo-import-report
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## container/heap
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/container_heap:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/container_heap:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## container/list
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/container_list:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/container_list:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## container/ring
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/container_ring:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/container_ring:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## context
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/context:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
error: couldn't load packages due to errors: context
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/context:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
error: couldn't load packages due to errors: context
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## crypto
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: init: slice expression with bounds
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## crypto/aes
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_aes:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_aes:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## crypto/cipher
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_cipher:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_cipher:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## crypto/des
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_des:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_des:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## crypto/dsa
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_dsa:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: init call: math/rand.NewSource
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_dsa:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !35947
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !35972
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36022
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36111
error: Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !35947
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !35972
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36022
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36111

Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## crypto/ecdsa
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_ecdsa:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
error: couldn't load packages due to errors: encoding/asn1
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_ecdsa:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
error: couldn't load packages due to errors: encoding/asn1
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## crypto/elliptic
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_elliptic:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: init call: math/rand.NewSource
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_elliptic:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36561
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36586
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36636
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36725
error: Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36561
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36586
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36636
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36725

Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## crypto/hmac
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_hmac:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_hmac:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## crypto/md5
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_md5:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: init: slice expression with bounds
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_md5:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: binop type: [4]byte
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## crypto/rand
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_rand:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: init call: math/rand.NewSource
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_rand:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36186
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36211
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36261
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36350
error: Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36186
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36211
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36261
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36350

Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## crypto/rc4
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_rc4:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_rc4:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## crypto/rsa
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_rsa:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: init: slice expression with bounds
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_rsa:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36380
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36405
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36455
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36544
error: Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36380
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36405
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36455
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36544

Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## crypto/sha1
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_sha1:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: init: slice expression with bounds
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_sha1:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Both operands to a binary operator are not of the same type!
  %slice.len = sub i64 %8, i32 0, !dbg !5677
Both operands to a binary operator are not of the same type!
  %slice.len11 = sub i64 %88, i32 0, !dbg !5712
error: Both operands to a binary operator are not of the same type!
  %slice.len = sub i64 %8, i32 0, !dbg !5677
Both operands to a binary operator are not of the same type!
  %slice.len11 = sub i64 %88, i32 0, !dbg !5712

Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## crypto/sha256
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_sha256:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: init: slice expression with bounds
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_sha256:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Both operands to a binary operator are not of the same type!
  %slice.len = sub i64 %8, i32 0, !dbg !5730
Both operands to a binary operator are not of the same type!
  %slice.len16 = sub i64 %117, i32 0, !dbg !5778
error: Both operands to a binary operator are not of the same type!
  %slice.len = sub i64 %8, i32 0, !dbg !5730
Both operands to a binary operator are not of the same type!
  %slice.len16 = sub i64 %117, i32 0, !dbg !5778

Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## crypto/sha512
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_sha512:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: init: slice expression with bounds
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_sha512:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Both operands to a binary operator are not of the same type!
  %slice.len = sub i64 %8, i32 0, !dbg !5792
Both operands to a binary operator are not of the same type!
  %slice.len17 = sub i64 %124, i32 0, !dbg !5843
error: Both operands to a binary operator are not of the same type!
  %slice.len = sub i64 %8, i32 0, !dbg !5792
Both operands to a binary operator are not of the same type!
  %slice.len17 = sub i64 %124, i32 0, !dbg !5843

Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## crypto/subtle
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_subtle:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_subtle:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## crypto/tls
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_tls:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
error: couldn't load packages due to errors: internal/singleflight, context, vendor/golang_org/x/crypto/cryptobyte and 2 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_tls:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
error: couldn't load packages due to errors: internal/singleflight, net, encoding/asn1 and 2 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## crypto/x509/pkix
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_x509_pkix:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
error: couldn't load packages due to errors: encoding/asn1
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_x509_pkix:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
error: couldn't load packages due to errors: encoding/asn1
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## database/sql
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/database_sql:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/database/sql/driver/types.go:227:20: invalid operation: rv.Type().Elem() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/database/sql/sql.go:537:7: Locker not declared by package sync
/usr/local/go/src/database/sql/sql.go:3051:7: Locker not declared by package sync
/usr/local/go/src/database/sql/sql.go:3073:23: Locker not declared by package sync
/usr/local/go/src/database/sql/sql.go:3069:29: Stack not declared by package runtime
/usr/local/go/src/database/sql/sql.go:3062:5: invalid operation: dr (variable of type driverResult) has no field or method Lock
/usr/local/go/src/database/sql/sql.go:3063:11: invalid operation: dr (variable of type driverResult) has no field or method Unlock
/usr/local/go/src/database/sql/sql.go:3056:5: invalid operation: dr (variable of type driverResult) has no field or method Lock
/usr/local/go/src/database/sql/sql.go:3057:11: invalid operation: dr (variable of type driverResult) has no field or method Unlock
/usr/local/go/src/database/sql/sql.go:2505:5: invalid operation: ds (variable of type *driverStmt) has no field or method Lock
/usr/local/go/src/database/sql/sql.go:2506:11: invalid operation: ds (variable of type *driverStmt) has no field or method Unlock
/usr/local/go/src/database/sql/sql.go:2334:5: invalid operation: ds (variable of type *driverStmt) has no field or method Lock
/usr/local/go/src/database/sql/sql.go:2335:11: invalid operation: ds (variable of type *driverStmt) has no field or method Unlock
/usr/local/go/src/database/sql/sql.go:2642:22: invalid operation: rs.closemu (variable of type sync.RWMutex) has no field or method RLocker
/usr/local/go/src/database/sql/convert.go:500:20: invalid operation: rv.Type().Elem() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/database/sql/convert.go:357:31: invalid operation: sv.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/database/sql/convert.go:367:41: invalid operation: sv.Type() (value of type reflect.Type) has no field or method ConvertibleTo
/usr/local/go/src/database/sql/convert.go:368:13: invalid operation: sv (variable of type reflect.Value) has no field or method Convert
/usr/local/go/src/database/sql/convert.go:380:19: Zero not declared by package reflect
/usr/local/go/src/database/sql/convert.go:383:18: New not declared by package reflect
/usr/local/go/src/database/sql/sql.go:546:5: invalid operation: ds (variable of type *driverStmt) has no field or method Lock
/usr/local/go/src/database/sql/sql.go:547:11: invalid operation: ds (variable of type *driverStmt) has no field or method Unlock
error: couldn't load packages due to errors: database/sql, database/sql/driver, context
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/database_sql:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/database/sql/driver/types.go:227:20: invalid operation: rv.Type().Elem() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/database/sql/sql.go:537:7: Locker not declared by package sync
/usr/local/go/src/database/sql/sql.go:3051:7: Locker not declared by package sync
/usr/local/go/src/database/sql/sql.go:3073:23: Locker not declared by package sync
/usr/local/go/src/database/sql/sql.go:3069:29: Stack not declared by package runtime
/usr/local/go/src/database/sql/sql.go:3062:5: invalid operation: dr (variable of type driverResult) has no field or method Lock
/usr/local/go/src/database/sql/sql.go:3063:11: invalid operation: dr (variable of type driverResult) has no field or method Unlock
/usr/local/go/src/database/sql/sql.go:3056:5: invalid operation: dr (variable of type driverResult) has no field or method Lock
/usr/local/go/src/database/sql/sql.go:3057:11: invalid operation: dr (variable of type driverResult) has no field or method Unlock
/usr/local/go/src/database/sql/sql.go:2505:5: invalid operation: ds (variable of type *driverStmt) has no field or method Lock
/usr/local/go/src/database/sql/sql.go:2506:11: invalid operation: ds (variable of type *driverStmt) has no field or method Unlock
/usr/local/go/src/database/sql/sql.go:2334:5: invalid operation: ds (variable of type *driverStmt) has no field or method Lock
/usr/local/go/src/database/sql/sql.go:2335:11: invalid operation: ds (variable of type *driverStmt) has no field or method Unlock
/usr/local/go/src/database/sql/sql.go:2642:22: invalid operation: rs.closemu (variable of type sync.RWMutex) has no field or method RLocker
/usr/local/go/src/database/sql/convert.go:500:20: invalid operation: rv.Type().Elem() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/database/sql/convert.go:357:31: invalid operation: sv.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/database/sql/convert.go:367:41: invalid operation: sv.Type() (value of type reflect.Type) has no field or method ConvertibleTo
/usr/local/go/src/database/sql/convert.go:368:13: invalid operation: sv (variable of type reflect.Value) has no field or method Convert
/usr/local/go/src/database/sql/convert.go:380:19: Zero not declared by package reflect
/usr/local/go/src/database/sql/convert.go:383:18: New not declared by package reflect
/usr/local/go/src/database/sql/sql.go:546:5: invalid operation: ds (variable of type *driverStmt) has no field or method Lock
/usr/local/go/src/database/sql/sql.go:547:11: invalid operation: ds (variable of type *driverStmt) has no field or method Unlock
error: couldn't load packages due to errors: context, database/sql, database/sql/driver
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## database/sql/driver
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/database_sql_driver:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/database/sql/driver/types.go:227:20: invalid operation: rv.Type().Elem() (value of type reflect.Type) has no field or method Implements
error: couldn't load packages due to errors: context, database/sql/driver
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/database_sql_driver:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/database/sql/driver/types.go:227:20: invalid operation: rv.Type().Elem() (value of type reflect.Type) has no field or method Implements
error: couldn't load packages due to errors: database/sql/driver, context
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## encoding
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## encoding/ascii85
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_ascii85:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_ascii85:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## encoding/asn1
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_asn1:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
error: couldn't load packages due to errors: encoding/asn1
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_asn1:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
error: couldn't load packages due to errors: encoding/asn1
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## encoding/base32
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_base32:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: init call: encoding/base32.NewEncoding
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_base32:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
unsupported instruction during init evaluation:
  %3 = ptrtoint i8* %0 to i32, !dbg !14
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## encoding/base64
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_base64:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: init call: encoding/base64.NewEncoding
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_base64:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
unsupported instruction during init evaluation:
  %3 = ptrtoint i8* %0 to i32, !dbg !14
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## encoding/binary
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_binary:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_binary:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## encoding/csv
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_csv:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_csv:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## encoding/gob
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_gob:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/encoding/gob/type.go:39:24: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:801:26: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:802:26: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:844:8: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:866:8: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:867:9: invalid operation: rt (variable of type reflect.Type) has no field or method PkgPath
/usr/local/go/src/encoding/gob/type.go:868:21: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:870:21: invalid operation: rt (variable of type reflect.Type) has no field or method PkgPath
/usr/local/go/src/encoding/gob/type.go:870:42: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:715:28: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:722:31: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:496:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/type.go:514:37: invalid operation: t.Elem() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:531:17: invalid operation: typ (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:119:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/encoding/gob/type.go:126:9: invalid operation: rt (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/gob/type.go:142:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:643:70: invalid operation: f (variable of type reflect.StructField) has no field or method Index
/usr/local/go/src/encoding/gob/encode.go:603:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:562:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/encoder.go:127:29: invalid operation: st (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/encode.go:320:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/encoding/gob/decode.go:1258:17: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:376:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/gob/decode.go:233:18: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:1202:17: invalid operation: base (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/decode.go:1118:30: invalid operation: srt (variable of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/encoding/gob/decode.go:1019:31: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:948:22: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:825:35: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:656:21: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:625:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:564:19: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:568:18: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:568:27: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:569:18: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:569:28: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:570:19: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:571:19: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:575:9: invalid operation: value (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/encoding/gob/decode.go:466:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
error: couldn't load packages due to errors: encoding/gob
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_gob:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/encoding/gob/type.go:39:24: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:801:26: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:802:26: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:844:8: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:866:8: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:867:9: invalid operation: rt (variable of type reflect.Type) has no field or method PkgPath
/usr/local/go/src/encoding/gob/type.go:868:21: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:870:21: invalid operation: rt (variable of type reflect.Type) has no field or method PkgPath
/usr/local/go/src/encoding/gob/type.go:870:42: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:715:28: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:722:31: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:496:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/type.go:514:37: invalid operation: t.Elem() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:531:17: invalid operation: typ (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:119:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/encoding/gob/type.go:126:9: invalid operation: rt (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/gob/type.go:142:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:643:70: invalid operation: f (variable of type reflect.StructField) has no field or method Index
/usr/local/go/src/encoding/gob/encode.go:603:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:562:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/encoder.go:127:29: invalid operation: st (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/encode.go:320:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/encoding/gob/decode.go:1258:17: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:376:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/gob/decode.go:233:18: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:1202:17: invalid operation: base (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/decode.go:1118:30: invalid operation: srt (variable of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/encoding/gob/decode.go:1019:31: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:948:22: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:825:35: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:656:21: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:625:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:564:19: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:568:18: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:568:27: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:569:18: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:569:28: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:570:19: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:571:19: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:575:9: invalid operation: value (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/encoding/gob/decode.go:466:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
error: couldn't load packages due to errors: encoding/gob
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## encoding/hex
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_hex:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_hex:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## encoding/json
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_json:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1176:51: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:749:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:698:11: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:703:9: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:391:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:395:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:400:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:404:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:364:11: WaitGroup not declared by package sync
/usr/local/go/src/encoding/json/decode.go:459:41: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:483:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:485:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:913:18: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:934:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:963:6: invalid operation: v (variable of type reflect.Value) has no field or method SetBytes
/usr/local/go/src/encoding/json/decode.go:967:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1001:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1009:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/encoding/json/decode.go:1017:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/encoding/json/decode.go:1025:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowFloat
/usr/local/go/src/encoding/json/decode.go:634:40: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:652:12: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:657:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:657:24: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:664:18: MakeMap not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:704:23: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:706:25: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:741:25: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:748:38: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:789:19: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:793:31: invalid operation: reflect.ValueOf(key) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:794:17: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:31: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:805:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:809:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:813:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:817:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:822:6: invalid operation: v (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/encoding/json/decode.go:526:8: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:557:14: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:17: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:27: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:563:13: Copy not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:567:7: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:599:17: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:604:6: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
error: couldn't load packages due to errors: encoding/json
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_json:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1176:51: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:749:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:698:11: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:703:9: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:391:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:395:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:400:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:404:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:364:11: WaitGroup not declared by package sync
/usr/local/go/src/encoding/json/decode.go:459:41: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:483:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:485:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:913:18: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:934:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:963:6: invalid operation: v (variable of type reflect.Value) has no field or method SetBytes
/usr/local/go/src/encoding/json/decode.go:967:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1001:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1009:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/encoding/json/decode.go:1017:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/encoding/json/decode.go:1025:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowFloat
/usr/local/go/src/encoding/json/decode.go:634:40: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:652:12: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:657:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:657:24: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:664:18: MakeMap not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:704:23: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:706:25: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:741:25: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:748:38: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:789:19: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:793:31: invalid operation: reflect.ValueOf(key) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:794:17: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:31: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:805:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:809:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:813:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:817:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:822:6: invalid operation: v (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/encoding/json/decode.go:526:8: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:557:14: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:17: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:27: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:563:13: Copy not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:567:7: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:599:17: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:604:6: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
error: couldn't load packages due to errors: encoding/json
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## encoding/pem
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_pem:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: init call: encoding/base64.NewEncoding
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_pem:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
unsupported instruction during init evaluation:
  %3 = ptrtoint i8* %0 to i32, !dbg !14
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## encoding/xml
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_xml:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/encoding/xml/typeinfo.go:47:19: Map not declared by package sync
/usr/local/go/src/encoding/xml/typeinfo.go:318:14: invalid operation: typ (variable of type reflect.Type) has no field or method FieldByIndex
/usr/local/go/src/encoding/xml/typeinfo.go:319:14: invalid operation: typ (variable of type reflect.Type) has no field or method FieldByIndex
/usr/local/go/src/encoding/xml/typeinfo.go:114:29: invalid operation: f (variable of type *reflect.StructField) has no field or method Index
/usr/local/go/src/encoding/xml/typeinfo.go:117:11: invalid operation: f (variable of type *reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/xml/typeinfo.go:168:20: invalid operation: f (variable of type *reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/xml/typeinfo.go:175:19: invalid operation: f (variable of type *reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/xml/typeinfo.go:63:10: invalid operation: f (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/xml/typeinfo.go:63:30: invalid operation: f (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/xml/typeinfo.go:63:46: invalid operation: f (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/xml/typeinfo.go:68:9: invalid operation: f (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/xml/read.go:602:20: New not declared by package reflect
/usr/local/go/src/encoding/xml/read.go:660:7: invalid operation: dst (variable of type reflect.Value) has no field or method SetBytes
/usr/local/go/src/encoding/xml/read.go:188:7: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/xml/read.go:329:20: New not declared by package reflect
/usr/local/go/src/encoding/xml/read.go:334:38: invalid operation: val.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:342:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:347:38: invalid operation: val.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:353:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:393:17: Append not declared by package reflect
/usr/local/go/src/encoding/xml/read.go:393:37: Zero not declared by package reflect
/usr/local/go/src/encoding/xml/read.go:397:6: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/xml/read.go:557:70: invalid operation: saveData.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:566:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:244:20: New not declared by package reflect
/usr/local/go/src/encoding/xml/read.go:248:38: invalid operation: val.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:255:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:261:38: invalid operation: val.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:268:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:277:19: Append not declared by package reflect
/usr/local/go/src/encoding/xml/read.go:277:39: Zero not declared by package reflect
/usr/local/go/src/encoding/xml/read.go:281:8: invalid operation: val (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/xml/marshal.go:638:16: invalid operation: typ (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/xml/marshal.go:639:26: invalid operation: typ (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/xml/marshal.go:643:33: invalid operation: typ.Elem() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/xml/marshal.go:813:38: invalid operation: vf.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:825:39: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:766:12: Copy not declared by package reflect
/usr/local/go/src/encoding/xml/marshal.go:544:38: invalid operation: val.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:557:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:569:38: invalid operation: val.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:580:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:427:31: invalid operation: typ (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:432:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:438:31: invalid operation: typ (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:443:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:486:15: invalid operation: typ (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/xml/typeinfo.go:356:20: New not declared by package reflect
error: couldn't load packages due to errors: encoding/xml
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_xml:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/encoding/xml/typeinfo.go:47:19: Map not declared by package sync
/usr/local/go/src/encoding/xml/typeinfo.go:318:14: invalid operation: typ (variable of type reflect.Type) has no field or method FieldByIndex
/usr/local/go/src/encoding/xml/typeinfo.go:319:14: invalid operation: typ (variable of type reflect.Type) has no field or method FieldByIndex
/usr/local/go/src/encoding/xml/typeinfo.go:114:29: invalid operation: f (variable of type *reflect.StructField) has no field or method Index
/usr/local/go/src/encoding/xml/typeinfo.go:117:11: invalid operation: f (variable of type *reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/xml/typeinfo.go:168:20: invalid operation: f (variable of type *reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/xml/typeinfo.go:175:19: invalid operation: f (variable of type *reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/xml/typeinfo.go:63:10: invalid operation: f (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/xml/typeinfo.go:63:30: invalid operation: f (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/xml/typeinfo.go:63:46: invalid operation: f (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/xml/typeinfo.go:68:9: invalid operation: f (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/xml/read.go:602:20: New not declared by package reflect
/usr/local/go/src/encoding/xml/read.go:660:7: invalid operation: dst (variable of type reflect.Value) has no field or method SetBytes
/usr/local/go/src/encoding/xml/read.go:188:7: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/xml/read.go:329:20: New not declared by package reflect
/usr/local/go/src/encoding/xml/read.go:334:38: invalid operation: val.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:342:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:347:38: invalid operation: val.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:353:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:393:17: Append not declared by package reflect
/usr/local/go/src/encoding/xml/read.go:393:37: Zero not declared by package reflect
/usr/local/go/src/encoding/xml/read.go:397:6: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/xml/read.go:557:70: invalid operation: saveData.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:566:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:244:20: New not declared by package reflect
/usr/local/go/src/encoding/xml/read.go:248:38: invalid operation: val.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:255:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:261:38: invalid operation: val.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:268:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/read.go:277:19: Append not declared by package reflect
/usr/local/go/src/encoding/xml/read.go:277:39: Zero not declared by package reflect
/usr/local/go/src/encoding/xml/read.go:281:8: invalid operation: val (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/xml/marshal.go:638:16: invalid operation: typ (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/xml/marshal.go:639:26: invalid operation: typ (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/xml/marshal.go:643:33: invalid operation: typ.Elem() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/xml/marshal.go:813:38: invalid operation: vf.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:825:39: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:766:12: Copy not declared by package reflect
/usr/local/go/src/encoding/xml/marshal.go:544:38: invalid operation: val.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:557:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:569:38: invalid operation: val.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:580:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:427:31: invalid operation: typ (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:432:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:438:31: invalid operation: typ (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:443:37: invalid operation: pv.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/xml/marshal.go:486:15: invalid operation: typ (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/xml/typeinfo.go:356:20: New not declared by package reflect
error: couldn't load packages due to errors: encoding/xml
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## errors
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/errors:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/errors:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## expvar
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/expvar:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1176:51: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:749:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:698:11: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:703:9: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:391:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:395:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:400:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:404:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:364:11: WaitGroup not declared by package sync
/usr/local/go/src/encoding/json/decode.go:459:41: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:483:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:485:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:913:18: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:934:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:963:6: invalid operation: v (variable of type reflect.Value) has no field or method SetBytes
/usr/local/go/src/encoding/json/decode.go:967:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1001:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1009:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/encoding/json/decode.go:1017:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/encoding/json/decode.go:1025:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowFloat
/usr/local/go/src/encoding/json/decode.go:634:40: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:652:12: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:657:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:657:24: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:664:18: MakeMap not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:704:23: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:706:25: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:741:25: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:748:38: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:789:19: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:793:31: invalid operation: reflect.ValueOf(key) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:794:17: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:31: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:805:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:809:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:813:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:817:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:822:6: invalid operation: v (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/encoding/json/decode.go:526:8: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:557:14: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:17: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:27: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:563:13: Copy not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:567:7: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:599:17: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:604:6: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
/usr/local/go/src/expvar/expvar.go:102:14: Map not declared by package sync
/usr/local/go/src/expvar/expvar.go:251:17: Map not declared by package sync
/usr/local/go/src/expvar/expvar.go:341:23: MemStats not declared by package runtime
/usr/local/go/src/expvar/expvar.go:342:10: ReadMemStats not declared by package runtime
error: couldn't load packages due to errors: encoding/json, encoding/asn1, internal/singleflight and 7 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/expvar:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1176:51: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:749:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:698:11: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:703:9: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:391:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:395:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:400:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:404:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:364:11: WaitGroup not declared by package sync
/usr/local/go/src/encoding/json/decode.go:459:41: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:483:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:485:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:913:18: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:934:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:963:6: invalid operation: v (variable of type reflect.Value) has no field or method SetBytes
/usr/local/go/src/encoding/json/decode.go:967:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1001:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1009:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/encoding/json/decode.go:1017:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/encoding/json/decode.go:1025:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowFloat
/usr/local/go/src/encoding/json/decode.go:634:40: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:652:12: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:657:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:657:24: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:664:18: MakeMap not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:704:23: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:706:25: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:741:25: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:748:38: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:789:19: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:793:31: invalid operation: reflect.ValueOf(key) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:794:17: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:31: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:805:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:809:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:813:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:817:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:822:6: invalid operation: v (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/encoding/json/decode.go:526:8: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:557:14: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:17: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:27: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:563:13: Copy not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:567:7: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:599:17: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:604:6: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
/usr/local/go/src/expvar/expvar.go:102:14: Map not declared by package sync
/usr/local/go/src/expvar/expvar.go:251:17: Map not declared by package sync
/usr/local/go/src/expvar/expvar.go:341:23: MemStats not declared by package runtime
/usr/local/go/src/expvar/expvar.go:342:10: ReadMemStats not declared by package runtime
error: couldn't load packages due to errors: net, encoding/asn1, vendor/golang_org/x/crypto/cryptobyte and 7 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## flag
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/flag:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/flag/flag.go:410:15: New not declared by package reflect
/usr/local/go/src/flag/flag.go:412:15: Zero not declared by package reflect
error: couldn't load packages due to errors: flag
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/flag:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/flag/flag.go:410:15: New not declared by package reflect
/usr/local/go/src/flag/flag.go:412:15: Zero not declared by package reflect
error: couldn't load packages due to errors: flag
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## fmt
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/fmt:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/fmt:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## hash
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/hash:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/hash:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## hash/adler32
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/hash_adler32:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/hash_adler32:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## hash/crc32
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/hash_crc32:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: init call: hash/crc32.simpleMakeTable
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/hash_crc32:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## hash/crc64
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/hash_crc64:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: init call: hash/crc64.makeTable
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/hash_crc64:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## hash/fnv
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/hash_fnv:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/hash_fnv:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## html
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/html:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: init call: strings.NewReplacer
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/html:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
unsupported instruction during init evaluation:
  %5 = ptrtoint i8* %0 to i32, !dbg !17
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## html/template
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/html_template:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1176:51: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:749:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:698:11: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:703:9: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:391:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:395:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:400:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:404:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:364:11: WaitGroup not declared by package sync
/usr/local/go/src/encoding/json/decode.go:459:41: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:483:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:485:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/funcs.go:140:19: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:142:18: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/funcs.go:145:70: invalid operation: value.Type() (value of type reflect.Type) has no field or method ConvertibleTo
/usr/local/go/src/text/template/funcs.go:146:17: invalid operation: value (variable of type reflect.Value) has no field or method Convert
/usr/local/go/src/text/template/funcs.go:93:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:32: invalid operation: typ (variable of type reflect.Type) has no field or method Out
/usr/local/go/src/text/template/funcs.go:75:88: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumOut
/usr/local/go/src/encoding/json/decode.go:913:18: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:934:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:963:6: invalid operation: v (variable of type reflect.Value) has no field or method SetBytes
/usr/local/go/src/encoding/json/decode.go:967:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1001:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1009:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/encoding/json/decode.go:1017:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/encoding/json/decode.go:1025:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowFloat
/usr/local/go/src/encoding/json/decode.go:634:40: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:652:12: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:657:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:657:24: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/funcs.go:196:45: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/funcs.go:203:17: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:248:92: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:250:15: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/funcs.go:252:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:256:17: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:267:11: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:268:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:278:14: invalid operation: v (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:949:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:949:50: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:950:30: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:664:18: MakeMap not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:704:23: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:706:25: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:741:25: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:748:38: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:789:19: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:793:31: invalid operation: reflect.ValueOf(key) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:794:17: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:31: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:805:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:809:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:813:31: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:950:79: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:159:16: Error not declared by package runtime
/usr/local/go/src/encoding/json/decode.go:817:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:822:6: invalid operation: v (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/encoding/json/decode.go:526:8: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:557:14: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:17: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:27: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:563:13: Copy not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:567:7: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:599:17: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:604:6: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:866:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:856:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:845:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:834:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:823:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:812:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:769:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:793:10: invalid operation: typ (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:722:13: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:728:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:735:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:735:33: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:738:20: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:748:59: invalid operation: value.Type().Elem() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:753:16: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:653:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:654:18: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:656:79: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:658:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:659:69: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:663:71: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/exec.go:670:32: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:673:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:674:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:674:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:681:12: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:681:19: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:682:10: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:686:13: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:695:16: invalid operation: fun (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:584:19: invalid operation: ptr (variable of type reflect.Value) has no field or method MethodByName
/usr/local/go/src/text/template/exec.go:591:33: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/text/template/exec.go:596:22: invalid operation: receiver (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/text/template/exec.go:612:21: invalid operation: nameVal.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:612:50: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/exec.go:622:23: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:529:25: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:431:56: invalid operation: value.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:375:20: invalid operation: val (variable of type reflect.Value) has no field or method Recv
/usr/local/go/src/html/template/js.go:135:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:57: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
error: couldn't load packages due to errors: html/template, encoding/json, text/template/parse and 1 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/html_template:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1176:51: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:749:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:698:11: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:703:9: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:391:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:395:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:400:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:404:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:364:11: WaitGroup not declared by package sync
/usr/local/go/src/encoding/json/decode.go:459:41: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:483:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:485:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:913:18: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:934:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:963:6: invalid operation: v (variable of type reflect.Value) has no field or method SetBytes
/usr/local/go/src/encoding/json/decode.go:967:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1001:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1009:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/encoding/json/decode.go:1017:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/encoding/json/decode.go:1025:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowFloat
/usr/local/go/src/encoding/json/decode.go:634:40: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:652:12: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:657:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:657:24: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:664:18: MakeMap not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:704:23: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:706:25: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:741:25: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:748:38: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:789:19: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:793:31: invalid operation: reflect.ValueOf(key) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:794:17: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:31: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:805:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:809:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:813:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:817:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:822:6: invalid operation: v (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/text/template/funcs.go:140:19: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:142:18: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/funcs.go:145:70: invalid operation: value.Type() (value of type reflect.Type) has no field or method ConvertibleTo
/usr/local/go/src/text/template/funcs.go:146:17: invalid operation: value (variable of type reflect.Value) has no field or method Convert
/usr/local/go/src/text/template/funcs.go:93:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:32: invalid operation: typ (variable of type reflect.Type) has no field or method Out
/usr/local/go/src/text/template/funcs.go:75:88: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:196:45: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/funcs.go:203:17: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:248:92: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:250:15: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/funcs.go:252:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:256:17: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:267:11: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:268:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:278:14: invalid operation: v (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:949:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:949:50: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:950:30: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:950:79: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:159:16: Error not declared by package runtime
/usr/local/go/src/encoding/json/decode.go:526:8: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:557:14: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:17: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:27: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:563:13: Copy not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:567:7: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:599:17: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:604:6: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:866:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:856:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:845:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:834:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:823:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:812:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:769:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:793:10: invalid operation: typ (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:722:13: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:728:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:735:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:735:33: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:738:20: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:748:59: invalid operation: value.Type().Elem() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:753:16: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:653:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:654:18: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:656:79: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:658:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:659:69: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:663:71: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/exec.go:670:32: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:673:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:674:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:674:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:681:12: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:681:19: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:682:10: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:686:13: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:695:16: invalid operation: fun (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:584:19: invalid operation: ptr (variable of type reflect.Value) has no field or method MethodByName
/usr/local/go/src/text/template/exec.go:591:33: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/text/template/exec.go:596:22: invalid operation: receiver (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/text/template/exec.go:612:21: invalid operation: nameVal.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:612:50: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/exec.go:622:23: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:529:25: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:431:56: invalid operation: value.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:375:20: invalid operation: val (variable of type reflect.Value) has no field or method Recv
/usr/local/go/src/html/template/js.go:135:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:57: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
error: couldn't load packages due to errors: text/template/parse, html/template, text/template and 1 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## image
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Basic Block in function 'github.com/trashhalo/tinygo-import-report.init' does not have terminator!
label %entry
error: Basic Block in function 'github.com/trashhalo/tinygo-import-report.init' does not have terminator!
label %entry

Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: interp: store on a pointer bitcast

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*PointerCastValue).Store(0xc003b7ce80, 0x7fa6ac024a70)
	/go/src/github.com/aykevl/tinygo/interp/values.go:374 +0x39
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000143670, 0x2a6fe40, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc000143650, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:108 +0x16c8
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0039d9940, 0x7fa6ac04da48, 0x0, 0x0, 0x0, 0xc00391b974, 0x5, 0x0, 0x0, 0xc000143710, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc0039d9940, 0x7fa6ac04da48, 0x0, 0x0, 0x0, 0xc00391b974, 0x5, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x2939c20, 0x292fb40, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7fffbb7d9f21, 0x29, 0x7fffbb7d9ed8, 0x3b, 0xc000122340, 0xc000143f58, 0xc000143dc0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:65 +0x1dd3
main.Build(0x7fffbb7d9f21, 0x29, 0x7fffbb7d9ed8, 0x3b, 0x7fffbb7d9f1c, 0x4, 0xc0000a1f58, 0xc000128000, 0x120)
	/go/src/github.com/aykevl/tinygo/main.go:219 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:476 +0xa16
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 2

```

## image/color
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_color:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_color:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## image/color/palette
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_color_palette:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Basic Block in function 'github.com/trashhalo/tinygo-import-report.init' does not have terminator!
label %entry
error: Basic Block in function 'github.com/trashhalo/tinygo-import-report.init' does not have terminator!
label %entry

Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_color_palette:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: interp: store on a pointer bitcast

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*PointerCastValue).Store(0xc000851e00, 0x7f18540277e0)
	/go/src/github.com/aykevl/tinygo/interp/values.go:374 +0x39
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000151670, 0x7f185402b930, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc000151650, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:108 +0x16c8
github.com/aykevl/tinygo/interp.(*Eval).function(0xc000865980, 0x7f1854028658, 0x0, 0x0, 0x0, 0xc0008c8760, 0x13, 0x0, 0x0, 0xc000151710, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc000865980, 0x7f1854028658, 0x0, 0x0, 0x0, 0xc0008c8760, 0x13, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x258dd70, 0x2583c90, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7fff3a651f21, 0x29, 0x7fff3a651ed8, 0x3b, 0xc000086750, 0xc000151f58, 0xc000151dc0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:65 +0x1dd3
main.Build(0x7fff3a651f21, 0x29, 0x7fff3a651ed8, 0x3b, 0x7fff3a651f1c, 0x4, 0xc0000a1f58, 0xc0000be480, 0x120)
	/go/src/github.com/aykevl/tinygo/main.go:219 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:476 +0xa16
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 2

```

## image/draw
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_draw:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Basic Block in function 'github.com/trashhalo/tinygo-import-report.init' does not have terminator!
label %entry
error: Basic Block in function 'github.com/trashhalo/tinygo-import-report.init' does not have terminator!
label %entry

Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_draw:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: interp: store on a pointer bitcast

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*PointerCastValue).Store(0xc00261a740, 0x7ff8dc024800)
	/go/src/github.com/aykevl/tinygo/interp/values.go:374 +0x39
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000119670, 0x7ff8dc9505b0, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc000119650, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:108 +0x16c8
github.com/aykevl/tinygo/interp.(*Eval).function(0xc00386e100, 0x7ff8dc01b178, 0x0, 0x0, 0x0, 0xc00415ae44, 0x5, 0x0, 0x0, 0xc000119710, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc00386e100, 0x7ff8dc01b178, 0x0, 0x0, 0x0, 0xc00415ae44, 0x5, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x29b0c20, 0x29a6b70, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffc9fe6ef21, 0x29, 0x7ffc9fe6eed8, 0x3b, 0xc0000e6340, 0xc000119f58, 0xc000119dc0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:65 +0x1dd3
main.Build(0x7ffc9fe6ef21, 0x29, 0x7ffc9fe6eed8, 0x3b, 0x7ffc9fe6ef1c, 0x4, 0xc00008ff58, 0xc0000f0000, 0x120)
	/go/src/github.com/aykevl/tinygo/main.go:219 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:476 +0xa16
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 2

```

## image/gif
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_gif:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Basic Block in function 'github.com/trashhalo/tinygo-import-report.init' does not have terminator!
label %entry
Both operands to a binary operator are not of the same type!
  %slice.cap = sub i32 1024, i8 %37, !dbg !8447
Call parameter type does not match function signature!
  %37 = load i8, i8* %36, !dbg !8446
 i32  call void @runtime.sliceBoundsCheck(i32 1024, i8 %37, i8 %35), !dbg !8447
Invalid InsertValueInst operands!
  %39 = insertvalue { i8*, i32, i32 } %38, i8 %slice.len, 1, !dbg !8447
Both operands to a binary operator are not of the same type!
  %slice.len = sub i8 %45, i32 0, !dbg !8541
Call parameter type does not match function signature!
  %45 = load i8, i8* %44, !dbg !8540
 i32  call void @runtime.sliceBoundsCheck(i32 1024, i32 0, i8 %45), !dbg !8541
Invalid InsertValueInst operands!
  %47 = insertvalue { i8*, i32, i32 } %46, i8 %slice.len, 1, !dbg !8541
Both operands to a binary operator are not of the same type!
  %slice.len = sub i8 %6, i32 0, !dbg !8822
Call parameter type does not match function signature!
  %6 = extractvalue { i8, %runtime._interface } %5, 0, !dbg !8816
 i32  call void @runtime.sliceBoundsCheck(i32 1024, i32 0, i8 %6), !dbg !8822
Invalid InsertValueInst operands!
  %18 = insertvalue { i8*, i32, i32 } %17, i8 %slice.len, 1, !dbg !8822
error: Basic Block in function 'github.com/trashhalo/tinygo-import-report.init' does not have terminator!
label %entry
Both operands to a binary operator are not of the same type!
  %slice.cap = sub i32 1024, i8 %37, !dbg !8447
Call parameter type does not match function signature!
  %37 = load i8, i8* %36, !dbg !8446
 i32  call void @runtime.sliceBoundsCheck(i32 1024, i8 %37, i8 %35), !dbg !8447
Invalid InsertValueInst operands!
  %39 = insertvalue { i8*, i32, i32 } %38, i8 %slice.len, 1, !dbg !8447
Both operands to a binary operator are not of the same type!
  %slice.len = sub i8 %45, i32 0, !dbg !8541
Call parameter type does not match function signature!
  %45 = load i8, i8* %44, !dbg !8540
 i32  call void @runtime.sliceBoundsCheck(i32 1024, i32 0, i8 %45), !dbg !8541
Invalid InsertValueInst operands!
  %47 = insertvalue { i8*, i32, i32 } %46, i8 %slice.len, 1, !dbg !8541
Both operands to a binary operator are not of the same type!
  %slice.len = sub i8 %6, i32 0, !dbg !8822
Call parameter type does not match function signature!
  %6 = extractvalue { i8, %runtime._interface } %5, 0, !dbg !8816
 i32  call void @runtime.sliceBoundsCheck(i32 1024, i32 0, i8 %6), !dbg !8822
Invalid InsertValueInst operands!
  %18 = insertvalue { i8*, i32, i32 } %17, i8 %slice.len, 1, !dbg !8822

Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_gif:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Both operands to a binary operator are not of the same type!
  %slice.cap = sub i32 1024, i8 %37, !dbg !34428
Call parameter type does not match function signature!
  %37 = load i8, i8* %36, !dbg !34427
 i32  call void @runtime.sliceBoundsCheck(i32 1024, i8 %37, i8 %35), !dbg !34428
Invalid InsertValueInst operands!
  %39 = insertvalue { i8*, i32, i32 } %38, i8 %slice.len, 1, !dbg !34428
Both operands to a binary operator are not of the same type!
  %slice.len = sub i8 %45, i32 0, !dbg !34522
Call parameter type does not match function signature!
  %45 = load i8, i8* %44, !dbg !34521
 i32  call void @runtime.sliceBoundsCheck(i32 1024, i32 0, i8 %45), !dbg !34522
Invalid InsertValueInst operands!
  %47 = insertvalue { i8*, i32, i32 } %46, i8 %slice.len, 1, !dbg !34522
Both operands to a binary operator are not of the same type!
  %slice.len = sub i8 %6, i32 0, !dbg !34803
Call parameter type does not match function signature!
  %6 = extractvalue { i8, %runtime._interface } %5, 0, !dbg !34797
 i32  call void @runtime.sliceBoundsCheck(i32 1024, i32 0, i8 %6), !dbg !34803
Invalid InsertValueInst operands!
  %18 = insertvalue { i8*, i32, i32 } %17, i8 %slice.len, 1, !dbg !34803
error: Both operands to a binary operator are not of the same type!
  %slice.cap = sub i32 1024, i8 %37, !dbg !34428
Call parameter type does not match function signature!
  %37 = load i8, i8* %36, !dbg !34427
 i32  call void @runtime.sliceBoundsCheck(i32 1024, i8 %37, i8 %35), !dbg !34428
Invalid InsertValueInst operands!
  %39 = insertvalue { i8*, i32, i32 } %38, i8 %slice.len, 1, !dbg !34428
Both operands to a binary operator are not of the same type!
  %slice.len = sub i8 %45, i32 0, !dbg !34522
Call parameter type does not match function signature!
  %45 = load i8, i8* %44, !dbg !34521
 i32  call void @runtime.sliceBoundsCheck(i32 1024, i32 0, i8 %45), !dbg !34522
Invalid InsertValueInst operands!
  %47 = insertvalue { i8*, i32, i32 } %46, i8 %slice.len, 1, !dbg !34522
Both operands to a binary operator are not of the same type!
  %slice.len = sub i8 %6, i32 0, !dbg !34803
Call parameter type does not match function signature!
  %6 = extractvalue { i8, %runtime._interface } %5, 0, !dbg !34797
 i32  call void @runtime.sliceBoundsCheck(i32 1024, i32 0, i8 %6), !dbg !34803
Invalid InsertValueInst operands!
  %18 = insertvalue { i8*, i32, i32 } %17, i8 %slice.len, 1, !dbg !34803

Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## image/jpeg
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_jpeg:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: full slice expressions (with max): []byte
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_jpeg:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: full slice expressions (with max): []byte
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## image/png
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_png:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: init call: compress/flate.generateFixedOffsetEncoding
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_png:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: interp: store on a pointer bitcast

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*PointerCastValue).Store(0xc0068ecfa0, 0x7fb3cc035930)
	/go/src/github.com/aykevl/tinygo/interp/values.go:374 +0x39
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000123670, 0x7fb3cca95b00, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc000123650, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:108 +0x16c8
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0018ab540, 0x7fb3cc0a5118, 0x0, 0x0, 0x0, 0xc007090f90, 0x5, 0x0, 0x0, 0xc000123710, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc0018ab540, 0x7fb3cc0a5118, 0x0, 0x0, 0x0, 0xc007090f90, 0x5, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x1ca0c20, 0x1c96b70, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7fff41175f21, 0x29, 0x7fff41175ed8, 0x3b, 0xc0000869c0, 0xc000123f58, 0xc000123dc0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:65 +0x1dd3
main.Build(0x7fff41175f21, 0x29, 0x7fff41175ed8, 0x3b, 0x7fff41175f1c, 0x4, 0xc0000a1f58, 0xc0000be5a0, 0x120)
	/go/src/github.com/aykevl/tinygo/main.go:219 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:476 +0xa16
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 2

```

## index/suffixarray
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/index_suffixarray:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/index_suffixarray:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## io
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/io:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/io:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## io/ioutil
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/io_ioutil:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/io_ioutil:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## log
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/log:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/log:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## log/syslog
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/log_syslog:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
error: couldn't load packages due to errors: context, internal/singleflight, net
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/log_syslog:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
error: couldn't load packages due to errors: net, context, internal/singleflight
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## math
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/math:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/math:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## math/big
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/math_big:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: init call: math/rand.NewSource
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/math_big:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !35941
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !35966
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36016
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36105
error: Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !35941
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !35966
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36016
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36105

Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## math/bits
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/math_bits:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/math_bits:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## math/cmplx
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/math_cmplx:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/math_cmplx:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## math/rand
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/math_rand:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: init call: math/rand.NewSource
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/math_rand:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: interp: GEP on a constant

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*LocalValue).GetElementPtr(0xc000a984d0, 0xc0006cf6b8, 0x2, 0x2, 0x2, 0x2)
	/go/src/github.com/aykevl/tinygo/interp/values.go:87 +0x167
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00015f098, 0x7f43f4156130, 0x0, 0x7ed582, 0x4, 0x7f43f4156180, 0x7f43f4158b28, 0x7f43f41307f0, 0x0, 0xfb744a3300000000, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:127 +0x193b
github.com/aykevl/tinygo/interp.(*Eval).function(0xc000f995c0, 0x7f43f4013a58, 0xc000ecea80, 0x2, 0x2, 0xc00066c590, 0x9, 0x7ed582, 0x4, 0x853240, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00015f670, 0x7f43f4051f50, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc00015f650, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:355 +0x3582
github.com/aykevl/tinygo/interp.(*Eval).function(0xc000f995c0, 0x7f43f400e288, 0x0, 0x0, 0x0, 0xc00066c590, 0x9, 0x0, 0x0, 0xc00015f710, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc000f995c0, 0x7f43f400e288, 0x0, 0x0, 0x0, 0xc00066c590, 0x9, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x1f8ed80, 0x1f84cd0, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffcf6aa0f21, 0x29, 0x7ffcf6aa0ed8, 0x3b, 0xc00013c340, 0xc00015ff58, 0xc00015fdc0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:65 +0x1dd3
main.Build(0x7ffcf6aa0f21, 0x29, 0x7ffcf6aa0ed8, 0x3b, 0x7ffcf6aa0f1c, 0x4, 0xc000081f58, 0xc000142000, 0x120)
	/go/src/github.com/aykevl/tinygo/main.go:219 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:476 +0xa16
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 2

```

## mime
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/mime:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
error: couldn't load packages due to errors: mime
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/mime:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
error: couldn't load packages due to errors: mime
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## mime/multipart
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/mime_multipart:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
error: couldn't load packages due to errors: mime, net, internal/singleflight and 1 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/mime_multipart:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
error: couldn't load packages due to errors: net, context, mime and 1 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## mime/quotedprintable
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/mime_quotedprintable:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/mime_quotedprintable:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## net
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
error: couldn't load packages due to errors: internal/singleflight, context, net
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
error: couldn't load packages due to errors: net, internal/singleflight, context
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## net/http
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
error: couldn't load packages due to errors: net, encoding/asn1, internal/singleflight and 5 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
error: couldn't load packages due to errors: net, context, encoding/asn1 and 5 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## net/http/cgi
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_cgi:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
error: couldn't load packages due to errors: internal/singleflight, encoding/asn1, net and 5 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_cgi:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
error: couldn't load packages due to errors: mime, net/http, net and 5 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## net/http/cookiejar
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_cookiejar:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
error: couldn't load packages due to errors: mime, context, net/http and 5 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_cookiejar:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
error: couldn't load packages due to errors: net/http, vendor/golang_org/x/crypto/cryptobyte, net/http/httptrace and 5 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## net/http/fcgi
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_fcgi:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
error: couldn't load packages due to errors: net/http/httptrace, context, encoding/asn1 and 5 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_fcgi:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
error: couldn't load packages due to errors: internal/singleflight, mime, context and 5 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## net/http/httptest
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_httptest:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/flag/flag.go:410:15: New not declared by package reflect
/usr/local/go/src/flag/flag.go:412:15: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
/usr/local/go/src/net/http/httptest/server.go:44:10: WaitGroup not declared by package sync
error: couldn't load packages due to errors: net/http/httptrace, net/http/httptest, net/http and 7 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_httptest:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/flag/flag.go:410:15: New not declared by package reflect
/usr/local/go/src/flag/flag.go:412:15: Zero not declared by package reflect
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
/usr/local/go/src/net/http/httptest/server.go:44:10: WaitGroup not declared by package sync
error: couldn't load packages due to errors: context, mime, net/http/httptest and 7 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## net/http/httptrace
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_httptrace:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
error: couldn't load packages due to errors: net/http/httptrace, net, vendor/golang_org/x/crypto/cryptobyte and 3 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_httptrace:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
error: couldn't load packages due to errors: net, context, internal/singleflight and 3 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## net/http/httputil
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_httputil:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
error: couldn't load packages due to errors: internal/singleflight, mime, net/http/httptrace and 5 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_httputil:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
error: couldn't load packages due to errors: net/http, net/http/httptrace, context and 5 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## net/http/pprof
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_pprof:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/runtime/trace/trace.go:147:10: StopTrace not declared by package runtime
/usr/local/go/src/runtime/trace/trace.go:124:20: StartTrace not declared by package runtime
/usr/local/go/src/runtime/trace/trace.go:129:20: ReadTrace not declared by package runtime
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1176:51: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:749:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:698:11: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:703:9: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:391:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:395:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:400:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:404:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:364:11: WaitGroup not declared by package sync
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/encoding/json/decode.go:459:41: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:483:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:485:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:913:18: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:934:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:963:6: invalid operation: v (variable of type reflect.Value) has no field or method SetBytes
/usr/local/go/src/encoding/json/decode.go:967:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1001:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1009:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/encoding/json/decode.go:1017:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/encoding/json/decode.go:1025:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowFloat
/usr/local/go/src/encoding/json/decode.go:634:40: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:652:12: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:657:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:657:24: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:664:18: MakeMap not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:704:23: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:706:25: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:741:25: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:748:38: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:789:19: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:793:31: invalid operation: reflect.ValueOf(key) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:794:17: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:31: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:805:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:809:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/text/template/funcs.go:140:19: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:142:18: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/funcs.go:145:70: invalid operation: value.Type() (value of type reflect.Type) has no field or method ConvertibleTo
/usr/local/go/src/text/template/funcs.go:146:17: invalid operation: value (variable of type reflect.Value) has no field or method Convert
/usr/local/go/src/text/template/funcs.go:93:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:32: invalid operation: typ (variable of type reflect.Type) has no field or method Out
/usr/local/go/src/text/template/funcs.go:75:88: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumOut
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/text/template/funcs.go:196:45: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/funcs.go:203:17: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:248:92: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:250:15: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/funcs.go:252:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:256:17: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/encoding/json/decode.go:813:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:817:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:822:6: invalid operation: v (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/encoding/json/decode.go:526:8: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:557:14: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:17: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:27: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:563:13: Copy not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:567:7: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:599:17: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:604:6: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/text/template/funcs.go:267:11: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:268:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:278:14: invalid operation: v (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:949:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:949:50: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:950:30: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:950:79: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:159:16: Error not declared by package runtime
/usr/local/go/src/text/template/exec.go:866:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:856:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:845:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:834:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:823:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:812:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:769:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:793:10: invalid operation: typ (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:722:13: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:728:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:735:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:735:33: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:738:20: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:748:59: invalid operation: value.Type().Elem() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:753:16: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:653:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:654:18: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:656:79: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:658:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:659:69: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:663:71: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/exec.go:670:32: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:673:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:674:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:674:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/runtime/pprof/pprof.go:372:135: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:693:84: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:718:31: StackRecord not declared by package runtime
/usr/local/go/src/text/template/exec.go:681:12: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:681:19: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:682:10: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:686:13: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:695:16: invalid operation: fun (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:584:19: invalid operation: ptr (variable of type reflect.Value) has no field or method MethodByName
/usr/local/go/src/text/template/exec.go:591:33: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/text/template/exec.go:596:22: invalid operation: receiver (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/text/template/exec.go:612:21: invalid operation: nameVal.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:612:50: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/exec.go:622:23: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:529:25: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:431:56: invalid operation: value.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:375:20: invalid operation: val (variable of type reflect.Value) has no field or method Recv
/usr/local/go/src/runtime/pprof/protomem.go:15:46: MemProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/proto.go:223:20: CallersFrames not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:920:20: SetMutexProfileFraction not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:807:10: SetCPUProfileRate not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:762:10: SetCPUProfileRate not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:700:18: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:706:22: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:678:16: Stack not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:542:24: MemStats not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:546:26: MemStats not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:547:11: ReadMemStats not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:556:18: MemProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:557:19: MemProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:562:22: MemProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:563:19: MemProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:572:45: MemProfileRate not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:581:20: MemProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:596:13: MemProfileRate not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:490:20: CallersFrames not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:877:18: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:878:19: MutexProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:880:22: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:881:19: MutexProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:900:49: SetMutexProfileFraction not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:819:18: MutexProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:825:18: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:826:19: BlockProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:828:22: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:829:19: BlockProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:813:18: BlockProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:526:18: MemProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:656:63: ThreadCreateProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:650:18: ThreadCreateProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:669:60: GoroutineProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:661:17: NumGoroutine not declared by package runtime
/usr/local/go/src/html/template/js.go:135:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/runtime/pprof/pprof.go:287:15: Callers not declared by package runtime
/usr/local/go/src/html/template/content.go:143:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:57: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
error: couldn't load packages due to errors: runtime/pprof, runtime/trace, encoding/json and 11 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_pprof:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/runtime/trace/trace.go:147:10: StopTrace not declared by package runtime
/usr/local/go/src/runtime/trace/trace.go:124:20: StartTrace not declared by package runtime
/usr/local/go/src/runtime/trace/trace.go:129:20: ReadTrace not declared by package runtime
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1176:51: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
/usr/local/go/src/encoding/json/encode.go:749:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:698:11: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:703:9: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:391:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:395:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:400:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:404:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:364:11: WaitGroup not declared by package sync
/usr/local/go/src/encoding/json/decode.go:459:41: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:483:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:485:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:913:18: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:934:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:963:6: invalid operation: v (variable of type reflect.Value) has no field or method SetBytes
/usr/local/go/src/encoding/json/decode.go:967:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1001:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1009:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/encoding/json/decode.go:1017:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/encoding/json/decode.go:1025:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowFloat
/usr/local/go/src/encoding/json/decode.go:634:40: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:652:12: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:657:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:657:24: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:664:18: MakeMap not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:704:23: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:706:25: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:741:25: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:748:38: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:789:19: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:793:31: invalid operation: reflect.ValueOf(key) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:794:17: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:31: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:805:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:809:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:813:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:817:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:822:6: invalid operation: v (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/encoding/json/decode.go:526:8: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:557:14: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:17: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:27: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:563:13: Copy not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:567:7: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:599:17: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:604:6: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/runtime/pprof/pprof.go:372:135: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:693:84: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:718:31: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/protomem.go:15:46: MemProfileRecord not declared by package runtime
/usr/local/go/src/text/template/funcs.go:140:19: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:142:18: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/funcs.go:145:70: invalid operation: value.Type() (value of type reflect.Type) has no field or method ConvertibleTo
/usr/local/go/src/text/template/funcs.go:146:17: invalid operation: value (variable of type reflect.Value) has no field or method Convert
/usr/local/go/src/text/template/funcs.go:93:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:32: invalid operation: typ (variable of type reflect.Type) has no field or method Out
/usr/local/go/src/text/template/funcs.go:75:88: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:196:45: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/funcs.go:203:17: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:248:92: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:250:15: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/funcs.go:252:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:256:17: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:267:11: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:268:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:278:14: invalid operation: v (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:949:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:949:50: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:950:30: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:950:79: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:159:16: Error not declared by package runtime
/usr/local/go/src/text/template/exec.go:866:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:856:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:845:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:834:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:823:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:812:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:769:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:793:10: invalid operation: typ (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:722:13: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:728:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:735:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:735:33: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:738:20: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:748:59: invalid operation: value.Type().Elem() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:753:16: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:653:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:654:18: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:656:79: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:658:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:659:69: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:663:71: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/exec.go:670:32: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/runtime/pprof/proto.go:223:20: CallersFrames not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:920:20: SetMutexProfileFraction not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:807:10: SetCPUProfileRate not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:762:10: SetCPUProfileRate not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:700:18: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:706:22: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:678:16: Stack not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:542:24: MemStats not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:546:26: MemStats not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:547:11: ReadMemStats not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:556:18: MemProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:557:19: MemProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:562:22: MemProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:563:19: MemProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:572:45: MemProfileRate not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:581:20: MemProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:596:13: MemProfileRate not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:490:20: CallersFrames not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:877:18: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:878:19: MutexProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:880:22: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:881:19: MutexProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:900:49: SetMutexProfileFraction not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:819:18: MutexProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:825:18: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:826:19: BlockProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:828:22: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:829:19: BlockProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:813:18: BlockProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:526:18: MemProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:656:63: ThreadCreateProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:650:18: ThreadCreateProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:669:60: GoroutineProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:661:17: NumGoroutine not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:287:15: Callers not declared by package runtime
/usr/local/go/src/text/template/exec.go:673:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:674:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:674:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:681:12: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:681:19: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:682:10: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:686:13: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:695:16: invalid operation: fun (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:584:19: invalid operation: ptr (variable of type reflect.Value) has no field or method MethodByName
/usr/local/go/src/text/template/exec.go:591:33: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/text/template/exec.go:596:22: invalid operation: receiver (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/text/template/exec.go:612:21: invalid operation: nameVal.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:612:50: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/exec.go:622:23: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:529:25: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:431:56: invalid operation: value.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:375:20: invalid operation: val (variable of type reflect.Value) has no field or method Recv
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/html/template/js.go:135:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/html/template/content.go:143:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:57: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
error: couldn't load packages due to errors: vendor/golang_org/x/crypto/cryptobyte, internal/singleflight, net and 11 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## net/mail
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_mail:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
error: couldn't load packages due to errors: mime, context, internal/singleflight and 1 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_mail:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
error: couldn't load packages due to errors: net, internal/singleflight, mime and 1 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## net/rpc
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_rpc:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1176:51: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:39:24: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:801:26: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:802:26: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:844:8: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:866:8: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:867:9: invalid operation: rt (variable of type reflect.Type) has no field or method PkgPath
/usr/local/go/src/encoding/gob/type.go:868:21: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:870:21: invalid operation: rt (variable of type reflect.Type) has no field or method PkgPath
/usr/local/go/src/encoding/gob/type.go:870:42: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:715:28: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:722:31: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:496:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/type.go:514:37: invalid operation: t.Elem() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:531:17: invalid operation: typ (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:119:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/encoding/gob/type.go:126:9: invalid operation: rt (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/gob/type.go:142:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:643:70: invalid operation: f (variable of type reflect.StructField) has no field or method Index
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/encoding/json/encode.go:749:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:698:11: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:703:9: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:391:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:395:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:400:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:404:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:364:11: WaitGroup not declared by package sync
/usr/local/go/src/encoding/gob/encode.go:603:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:562:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/text/template/funcs.go:140:19: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:142:18: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/funcs.go:145:70: invalid operation: value.Type() (value of type reflect.Type) has no field or method ConvertibleTo
/usr/local/go/src/text/template/funcs.go:146:17: invalid operation: value (variable of type reflect.Value) has no field or method Convert
/usr/local/go/src/text/template/funcs.go:93:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:32: invalid operation: typ (variable of type reflect.Type) has no field or method Out
/usr/local/go/src/text/template/funcs.go:75:88: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumOut
/usr/local/go/src/encoding/json/decode.go:459:41: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:483:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:485:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:913:18: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:934:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:963:6: invalid operation: v (variable of type reflect.Value) has no field or method SetBytes
/usr/local/go/src/encoding/json/decode.go:967:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1001:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1009:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/encoding/gob/encoder.go:127:29: invalid operation: st (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:1017:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/encoding/json/decode.go:1025:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowFloat
/usr/local/go/src/encoding/json/decode.go:634:40: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:652:12: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:657:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:657:24: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:664:18: MakeMap not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:704:23: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:706:25: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:741:25: New not declared by package reflect
/usr/local/go/src/text/template/funcs.go:196:45: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/funcs.go:203:17: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:748:38: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/text/template/funcs.go:248:92: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:250:15: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/funcs.go:252:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:256:17: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:267:11: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:268:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:278:14: invalid operation: v (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:949:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:949:50: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/gob/encode.go:320:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/text/template/exec.go:950:30: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:950:79: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:1258:17: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:376:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/gob/decode.go:233:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:789:19: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:793:31: invalid operation: reflect.ValueOf(key) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:794:17: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:31: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:805:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:809:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:813:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:817:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:822:6: invalid operation: v (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/encoding/json/decode.go:526:8: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:557:14: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:17: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:27: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:563:13: Copy not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:567:7: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:599:17: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:604:6: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:159:16: Error not declared by package runtime
/usr/local/go/src/text/template/exec.go:866:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:856:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:845:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:834:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:823:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:812:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:769:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:793:10: invalid operation: typ (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:722:13: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:728:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:735:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:735:33: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:738:20: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:748:59: invalid operation: value.Type().Elem() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:753:16: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:653:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:654:18: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:656:79: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:658:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:659:69: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:663:71: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/exec.go:670:32: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:673:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:674:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:674:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:681:12: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:681:19: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:682:10: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:686:13: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:695:16: invalid operation: fun (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:584:19: invalid operation: ptr (variable of type reflect.Value) has no field or method MethodByName
/usr/local/go/src/text/template/exec.go:591:33: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/text/template/exec.go:596:22: invalid operation: receiver (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/text/template/exec.go:612:21: invalid operation: nameVal.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:612:50: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/exec.go:622:23: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:529:25: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/encoding/gob/decode.go:1202:17: invalid operation: base (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/decode.go:1118:30: invalid operation: srt (variable of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/encoding/gob/decode.go:1019:31: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:948:22: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:825:35: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/exec.go:431:56: invalid operation: value.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/gob/decode.go:656:21: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:625:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:564:19: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:568:18: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:568:27: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:569:18: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:569:28: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:570:19: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:571:19: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:575:9: invalid operation: value (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/text/template/exec.go:375:20: invalid operation: val (variable of type reflect.Value) has no field or method Recv
/usr/local/go/src/encoding/gob/decode.go:466:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/html/template/js.go:135:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:57: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
/usr/local/go/src/net/rpc/server.go:156:21: Method not declared by package reflect
/usr/local/go/src/net/rpc/server.go:190:18: Map not declared by package sync
/usr/local/go/src/net/rpc/server.go:375:70: WaitGroup not declared by package sync
/usr/local/go/src/net/rpc/server.go:288:22: invalid operation: typ (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/net/rpc/server.go:289:17: invalid operation: typ (variable of type reflect.Type) has no field or method Method
/usr/local/go/src/net/rpc/server.go:218:22: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/net/rpc/server.go:218:35: invalid operation: t (variable of type reflect.Type) has no field or method PkgPath
/usr/local/go/src/net/rpc/server.go:563:18: New not declared by package reflect
/usr/local/go/src/net/rpc/server.go:565:18: New not declared by package reflect
/usr/local/go/src/net/rpc/server.go:576:19: New not declared by package reflect
/usr/local/go/src/net/rpc/server.go:580:29: MakeMap not declared by package reflect
/usr/local/go/src/net/rpc/server.go:463:17: WaitGroup not declared by package sync
/usr/local/go/src/net/rpc/server.go:245:43: invalid operation: reflect.Indirect(s.rcvr).Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/net/rpc/server.go:268:37: PtrTo not declared by package reflect
error: couldn't load packages due to errors: context, mime, html/template and 11 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_rpc:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1176:51: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
/usr/local/go/src/encoding/gob/type.go:39:24: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:801:26: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:802:26: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:844:8: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:866:8: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:867:9: invalid operation: rt (variable of type reflect.Type) has no field or method PkgPath
/usr/local/go/src/encoding/gob/type.go:868:21: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:870:21: invalid operation: rt (variable of type reflect.Type) has no field or method PkgPath
/usr/local/go/src/encoding/gob/type.go:870:42: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:715:28: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:722:31: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:749:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:698:11: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:703:9: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:391:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:395:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:400:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:404:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:364:11: WaitGroup not declared by package sync
/usr/local/go/src/encoding/gob/type.go:496:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:459:41: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:483:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:485:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:913:18: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:934:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:963:6: invalid operation: v (variable of type reflect.Value) has no field or method SetBytes
/usr/local/go/src/encoding/json/decode.go:967:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1001:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1009:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/encoding/json/decode.go:1017:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/encoding/json/decode.go:1025:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowFloat
/usr/local/go/src/encoding/json/decode.go:634:40: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/gob/type.go:514:37: invalid operation: t.Elem() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:531:17: invalid operation: typ (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:119:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/encoding/gob/type.go:126:9: invalid operation: rt (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/gob/type.go:142:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:643:70: invalid operation: f (variable of type reflect.StructField) has no field or method Index
/usr/local/go/src/encoding/gob/encode.go:603:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:562:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:652:12: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:657:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:657:24: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:664:18: MakeMap not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:704:23: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:706:25: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:741:25: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:748:38: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:789:19: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:793:31: invalid operation: reflect.ValueOf(key) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:794:17: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:31: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:805:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:809:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:813:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:817:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:822:6: invalid operation: v (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/text/template/funcs.go:140:19: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:142:18: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/funcs.go:145:70: invalid operation: value.Type() (value of type reflect.Type) has no field or method ConvertibleTo
/usr/local/go/src/text/template/funcs.go:146:17: invalid operation: value (variable of type reflect.Value) has no field or method Convert
/usr/local/go/src/text/template/funcs.go:93:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/encoding/json/decode.go:526:8: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:557:14: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:17: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:27: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:563:13: Copy not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:567:7: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:599:17: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:604:6: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/funcs.go:95:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:32: invalid operation: typ (variable of type reflect.Type) has no field or method Out
/usr/local/go/src/text/template/funcs.go:75:88: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:196:45: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/funcs.go:203:17: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:248:92: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/encoding/gob/encoder.go:127:29: invalid operation: st (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/funcs.go:250:15: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/funcs.go:252:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:256:17: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:267:11: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:268:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:278:14: invalid operation: v (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:949:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:949:50: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:950:30: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:950:79: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:159:16: Error not declared by package runtime
/usr/local/go/src/text/template/exec.go:866:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:856:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:845:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:834:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:823:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:812:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:769:19: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:320:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/encoding/gob/decode.go:1258:17: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:376:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/gob/decode.go:233:18: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:793:10: invalid operation: typ (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:722:13: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:728:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:735:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:735:33: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:738:20: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:748:59: invalid operation: value.Type().Elem() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:753:16: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:653:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:654:18: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:656:79: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:658:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:659:69: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:663:71: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/exec.go:670:32: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:673:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:674:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:674:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:681:12: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:681:19: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:682:10: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:686:13: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:695:16: invalid operation: fun (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:584:19: invalid operation: ptr (variable of type reflect.Value) has no field or method MethodByName
/usr/local/go/src/text/template/exec.go:591:33: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/text/template/exec.go:596:22: invalid operation: receiver (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/text/template/exec.go:612:21: invalid operation: nameVal.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:612:50: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/exec.go:622:23: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:1202:17: invalid operation: base (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/decode.go:1118:30: invalid operation: srt (variable of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/encoding/gob/decode.go:1019:31: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:948:22: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:825:35: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:656:21: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:625:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/text/template/exec.go:529:25: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:431:56: invalid operation: value.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:375:20: invalid operation: val (variable of type reflect.Value) has no field or method Recv
/usr/local/go/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:564:19: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:568:18: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:568:27: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:569:18: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:569:28: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:570:19: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:571:19: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:575:9: invalid operation: value (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/encoding/gob/decode.go:466:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/html/template/js.go:135:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:57: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
/usr/local/go/src/net/rpc/server.go:156:21: Method not declared by package reflect
/usr/local/go/src/net/rpc/server.go:190:18: Map not declared by package sync
/usr/local/go/src/net/rpc/server.go:375:70: WaitGroup not declared by package sync
/usr/local/go/src/net/rpc/server.go:288:22: invalid operation: typ (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/net/rpc/server.go:289:17: invalid operation: typ (variable of type reflect.Type) has no field or method Method
/usr/local/go/src/net/rpc/server.go:218:22: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/net/rpc/server.go:218:35: invalid operation: t (variable of type reflect.Type) has no field or method PkgPath
/usr/local/go/src/net/rpc/server.go:563:18: New not declared by package reflect
/usr/local/go/src/net/rpc/server.go:565:18: New not declared by package reflect
/usr/local/go/src/net/rpc/server.go:576:19: New not declared by package reflect
/usr/local/go/src/net/rpc/server.go:580:29: MakeMap not declared by package reflect
/usr/local/go/src/net/rpc/server.go:463:17: WaitGroup not declared by package sync
/usr/local/go/src/net/rpc/server.go:245:43: invalid operation: reflect.Indirect(s.rcvr).Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/net/rpc/server.go:268:37: PtrTo not declared by package reflect
error: couldn't load packages due to errors: text/template, text/template/parse, internal/singleflight and 11 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## net/rpc/jsonrpc
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_rpc_jsonrpc:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
/usr/local/go/src/encoding/gob/type.go:39:24: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:801:26: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:802:26: Map not declared by package sync
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/gob/type.go:844:8: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:866:8: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:867:9: invalid operation: rt (variable of type reflect.Type) has no field or method PkgPath
/usr/local/go/src/encoding/gob/type.go:868:21: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:870:21: invalid operation: rt (variable of type reflect.Type) has no field or method PkgPath
/usr/local/go/src/encoding/gob/type.go:870:42: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:715:28: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:722:31: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:496:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/type.go:514:37: invalid operation: t.Elem() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:531:17: invalid operation: typ (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:119:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/encoding/gob/type.go:126:9: invalid operation: rt (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/gob/type.go:142:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:643:70: invalid operation: f (variable of type reflect.StructField) has no field or method Index
/usr/local/go/src/encoding/gob/encode.go:603:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:562:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/encoder.go:127:29: invalid operation: st (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/encode.go:320:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/encoding/gob/decode.go:1258:17: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:376:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/gob/decode.go:233:18: New not declared by package reflect
/usr/local/go/src/text/template/funcs.go:140:19: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:142:18: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/funcs.go:145:70: invalid operation: value.Type() (value of type reflect.Type) has no field or method ConvertibleTo
/usr/local/go/src/text/template/funcs.go:146:17: invalid operation: value (variable of type reflect.Value) has no field or method Convert
/usr/local/go/src/text/template/funcs.go:93:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:32: invalid operation: typ (variable of type reflect.Type) has no field or method Out
/usr/local/go/src/text/template/funcs.go:75:88: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:196:45: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/funcs.go:203:17: Zero not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/text/template/funcs.go:248:92: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:250:15: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/funcs.go:252:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:256:17: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:267:11: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:268:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:278:14: invalid operation: v (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:949:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:949:50: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:950:30: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:950:79: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:159:16: Error not declared by package runtime
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1176:51: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/text/template/exec.go:866:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:856:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:845:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:834:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:823:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:812:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:769:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:793:10: invalid operation: typ (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:722:13: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:728:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:735:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:735:33: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:738:20: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:748:59: invalid operation: value.Type().Elem() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:753:16: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:653:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:654:18: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:656:79: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:658:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:659:69: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:663:71: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/exec.go:670:32: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:673:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:674:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:674:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:681:12: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:681:19: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:682:10: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:686:13: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:695:16: invalid operation: fun (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:584:19: invalid operation: ptr (variable of type reflect.Value) has no field or method MethodByName
/usr/local/go/src/text/template/exec.go:591:33: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/text/template/exec.go:596:22: invalid operation: receiver (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/text/template/exec.go:612:21: invalid operation: nameVal.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:612:50: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/exec.go:622:23: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:529:25: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:431:56: invalid operation: value.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:375:20: invalid operation: val (variable of type reflect.Value) has no field or method Recv
/usr/local/go/src/encoding/json/encode.go:749:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:698:11: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:703:9: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:391:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:395:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:400:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:404:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:364:11: WaitGroup not declared by package sync
/usr/local/go/src/encoding/gob/decode.go:1202:17: invalid operation: base (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/decode.go:1118:30: invalid operation: srt (variable of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/encoding/gob/decode.go:1019:31: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:948:22: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:825:35: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:459:41: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:483:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:485:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:913:18: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:934:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:963:6: invalid operation: v (variable of type reflect.Value) has no field or method SetBytes
/usr/local/go/src/encoding/json/decode.go:967:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/gob/decode.go:656:21: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:625:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:564:19: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:568:18: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:568:27: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:569:18: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:569:28: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:570:19: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:571:19: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:575:9: invalid operation: value (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/encoding/gob/decode.go:466:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/encoding/json/decode.go:1001:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1009:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/encoding/json/decode.go:1017:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/encoding/json/decode.go:1025:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowFloat
/usr/local/go/src/encoding/json/decode.go:634:40: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:652:12: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:657:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:657:24: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:664:18: MakeMap not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:704:23: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:706:25: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:741:25: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:748:38: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:789:19: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:793:31: invalid operation: reflect.ValueOf(key) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:794:17: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:31: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:805:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:809:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:813:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:817:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:822:6: invalid operation: v (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/encoding/json/decode.go:526:8: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:557:14: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:17: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:27: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:563:13: Copy not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:567:7: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:599:17: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:604:6: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/html/template/js.go:135:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:57: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
/usr/local/go/src/net/rpc/server.go:156:21: Method not declared by package reflect
/usr/local/go/src/net/rpc/server.go:190:18: Map not declared by package sync
/usr/local/go/src/net/rpc/server.go:375:70: WaitGroup not declared by package sync
/usr/local/go/src/net/rpc/server.go:288:22: invalid operation: typ (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/net/rpc/server.go:289:17: invalid operation: typ (variable of type reflect.Type) has no field or method Method
/usr/local/go/src/net/rpc/server.go:218:22: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/net/rpc/server.go:218:35: invalid operation: t (variable of type reflect.Type) has no field or method PkgPath
/usr/local/go/src/net/rpc/server.go:563:18: New not declared by package reflect
/usr/local/go/src/net/rpc/server.go:565:18: New not declared by package reflect
/usr/local/go/src/net/rpc/server.go:576:19: New not declared by package reflect
/usr/local/go/src/net/rpc/server.go:580:29: MakeMap not declared by package reflect
/usr/local/go/src/net/rpc/server.go:463:17: WaitGroup not declared by package sync
/usr/local/go/src/net/rpc/server.go:245:43: invalid operation: reflect.Indirect(s.rcvr).Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/net/rpc/server.go:268:37: PtrTo not declared by package reflect
error: couldn't load packages due to errors: net/http/httptrace, net, encoding/asn1 and 11 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_rpc_jsonrpc:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:39:24: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1176:51: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:801:26: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:802:26: Map not declared by package sync
/usr/local/go/src/encoding/gob/type.go:844:8: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:866:8: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:867:9: invalid operation: rt (variable of type reflect.Type) has no field or method PkgPath
/usr/local/go/src/encoding/gob/type.go:868:21: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:870:21: invalid operation: rt (variable of type reflect.Type) has no field or method PkgPath
/usr/local/go/src/encoding/gob/type.go:870:42: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:715:28: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:722:31: invalid operation: rt (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:496:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/type.go:514:37: invalid operation: t.Elem() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:531:17: invalid operation: typ (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:119:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/encoding/gob/type.go:126:9: invalid operation: rt (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/gob/type.go:142:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:643:70: invalid operation: f (variable of type reflect.StructField) has no field or method Index
/usr/local/go/src/encoding/gob/encode.go:603:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:562:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:749:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:698:11: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:703:9: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
/usr/local/go/src/encoding/json/encode.go:391:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:395:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:400:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:404:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:364:11: WaitGroup not declared by package sync
/usr/local/go/src/encoding/json/decode.go:459:41: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:483:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:485:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:913:18: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:934:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:963:6: invalid operation: v (variable of type reflect.Value) has no field or method SetBytes
/usr/local/go/src/encoding/json/decode.go:967:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1001:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1009:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/encoding/json/decode.go:1017:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/encoding/json/decode.go:1025:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowFloat
/usr/local/go/src/encoding/json/decode.go:634:40: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/gob/encoder.go:127:29: invalid operation: st (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/encode.go:320:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/encoding/json/decode.go:652:12: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:657:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:657:24: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:664:18: MakeMap not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:1258:17: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:376:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/gob/decode.go:233:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:704:23: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:706:25: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:140:19: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:142:18: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/funcs.go:145:70: invalid operation: value.Type() (value of type reflect.Type) has no field or method ConvertibleTo
/usr/local/go/src/encoding/json/decode.go:741:25: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:748:38: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:789:19: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:793:31: invalid operation: reflect.ValueOf(key) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:794:17: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:795:31: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:805:31: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:146:17: invalid operation: value (variable of type reflect.Value) has no field or method Convert
/usr/local/go/src/text/template/funcs.go:93:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:32: invalid operation: typ (variable of type reflect.Type) has no field or method Out
/usr/local/go/src/text/template/funcs.go:75:88: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumOut
/usr/local/go/src/encoding/json/decode.go:809:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:813:31: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:817:30: invalid operation: reflect.ValueOf(n) (value of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:822:6: invalid operation: v (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/encoding/json/decode.go:526:8: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:557:14: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:17: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:27: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:563:13: Copy not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:567:7: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/text/template/funcs.go:196:45: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/funcs.go:203:17: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:248:92: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:250:15: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/funcs.go:252:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:256:17: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:267:11: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:268:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:278:14: invalid operation: v (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:949:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:949:50: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:950:30: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:950:79: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:599:17: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:604:6: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/text/template/exec.go:159:16: Error not declared by package runtime
/usr/local/go/src/text/template/exec.go:866:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:856:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:845:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:834:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:823:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:812:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:769:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:793:10: invalid operation: typ (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:722:13: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:728:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:735:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:735:33: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:738:20: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:748:59: invalid operation: value.Type().Elem() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:753:16: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:653:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:654:18: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:656:79: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:658:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:659:69: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:663:71: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/exec.go:670:32: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:673:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:674:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:674:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:681:12: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:681:19: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:682:10: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:686:13: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:695:16: invalid operation: fun (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:584:19: invalid operation: ptr (variable of type reflect.Value) has no field or method MethodByName
/usr/local/go/src/text/template/exec.go:591:33: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/text/template/exec.go:596:22: invalid operation: receiver (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/text/template/exec.go:612:21: invalid operation: nameVal.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:612:50: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:1202:17: invalid operation: base (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/text/template/exec.go:622:23: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:529:25: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/encoding/gob/decode.go:1118:30: invalid operation: srt (variable of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/text/template/exec.go:431:56: invalid operation: value.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:375:20: invalid operation: val (variable of type reflect.Value) has no field or method Recv
/usr/local/go/src/encoding/gob/decode.go:1019:31: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:948:22: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:825:35: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:656:21: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:625:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:564:19: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:568:18: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:568:27: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:569:18: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:569:28: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:570:19: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:571:19: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:575:9: invalid operation: value (variable of type reflect.Value) has no field or method SetMapIndex
/usr/local/go/src/encoding/gob/decode.go:466:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/html/template/js.go:135:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:57: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/net/http/httptrace/trace.go:202:22: MakeFunc not declared by package reflect
/usr/local/go/src/net/http/httptrace/trace.go:203:11: invalid operation: tfCopy (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/httptrace/trace.go:204:14: invalid operation: of (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/net/http/h2_bundle.go:3690:16: Cond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:6804:24: Cond not declared by package sync
/usr/local/go/src/net/http/server.go:646:16: Cond not declared by package sync
/usr/local/go/src/net/http/transport.go:2016:38: Locker not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:3137:17: Stack not declared by package runtime
/usr/local/go/src/net/http/h2_bundle.go:5870:24: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:1746:23: Stack not declared by package runtime
/usr/local/go/src/net/http/server.go:655:18: NewCond not declared by package sync
/usr/local/go/src/net/http/h2_bundle.go:7223:17: NewCond not declared by package sync
/usr/local/go/src/net/rpc/server.go:156:21: Method not declared by package reflect
/usr/local/go/src/net/rpc/server.go:190:18: Map not declared by package sync
/usr/local/go/src/net/rpc/server.go:375:70: WaitGroup not declared by package sync
/usr/local/go/src/net/rpc/server.go:288:22: invalid operation: typ (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/net/rpc/server.go:289:17: invalid operation: typ (variable of type reflect.Type) has no field or method Method
/usr/local/go/src/net/rpc/server.go:218:22: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/net/rpc/server.go:218:35: invalid operation: t (variable of type reflect.Type) has no field or method PkgPath
/usr/local/go/src/net/rpc/server.go:563:18: New not declared by package reflect
/usr/local/go/src/net/rpc/server.go:565:18: New not declared by package reflect
/usr/local/go/src/net/rpc/server.go:576:19: New not declared by package reflect
/usr/local/go/src/net/rpc/server.go:580:29: MakeMap not declared by package reflect
/usr/local/go/src/net/rpc/server.go:463:17: WaitGroup not declared by package sync
/usr/local/go/src/net/rpc/server.go:245:43: invalid operation: reflect.Indirect(s.rcvr).Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/net/rpc/server.go:268:37: PtrTo not declared by package reflect
error: couldn't load packages due to errors: mime, html/template, encoding/gob and 11 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## net/smtp
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_smtp:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
error: couldn't load packages due to errors: internal/singleflight, net, vendor/golang_org/x/crypto/cryptobyte and 2 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_smtp:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
error: couldn't load packages due to errors: vendor/golang_org/x/crypto/cryptobyte, encoding/asn1, net and 2 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## net/textproto
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_textproto:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
error: couldn't load packages due to errors: internal/singleflight, context, net
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_textproto:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
error: couldn't load packages due to errors: context, net, internal/singleflight
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## net/url
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_url:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_url:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## os
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/os:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/os:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## os/exec
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/os_exec:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
error: couldn't load packages due to errors: context
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/os_exec:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
error: couldn't load packages due to errors: context
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## os/signal
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/os_signal:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: unknown expression: select nonblocking [t10<-sig]
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/os_signal:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: unknown expression: select nonblocking [t10<-sig]
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## os/user
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/os_user:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/os/user/lookup.go:53:9: undeclared name: lookupGroupId
/usr/local/go/src/os/user/lookup.go:47:9: undeclared name: lookupGroup
/usr/local/go/src/os/user/lookup.go:41:9: undeclared name: lookupUserId
/usr/local/go/src/os/user/lookup.go:32:9: undeclared name: lookupUser
/usr/local/go/src/os/user/lookup.go:11:41: undeclared name: current
/usr/local/go/src/os/user/lookup.go:58:9: undeclared name: listGroups
error: couldn't load packages due to errors: os/user
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/os_user:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/os/user/lookup.go:53:9: undeclared name: lookupGroupId
/usr/local/go/src/os/user/lookup.go:47:9: undeclared name: lookupGroup
/usr/local/go/src/os/user/lookup.go:41:9: undeclared name: lookupUserId
/usr/local/go/src/os/user/lookup.go:32:9: undeclared name: lookupUser
/usr/local/go/src/os/user/lookup.go:11:41: undeclared name: current
/usr/local/go/src/os/user/lookup.go:58:9: undeclared name: listGroups
error: couldn't load packages due to errors: os/user
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## path
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/path:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/path:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## path/filepath
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/path_filepath:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/path_filepath:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## plugin
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/plugin:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/plugin:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## reflect
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/reflect:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/reflect:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## regexp
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/regexp:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/regexp:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## regexp/syntax
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/regexp_syntax:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/regexp_syntax:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## runtime
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/runtime:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/runtime:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## runtime/cgo
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/runtime_cgo:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/runtime_cgo:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## runtime/debug
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/runtime_debug:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/runtime/debug/stack.go:24:16: Stack not declared by package runtime
/usr/local/go/src/runtime/debug/garbage.go:37:34: MemStats not declared by package runtime
error: couldn't load packages due to errors: runtime/debug
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/runtime_debug:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/runtime/debug/stack.go:24:16: Stack not declared by package runtime
/usr/local/go/src/runtime/debug/garbage.go:37:34: MemStats not declared by package runtime
error: couldn't load packages due to errors: runtime/debug
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## runtime/msan
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/runtime_msan:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/go/src/github.com/trashhalo/tinygo-import-report/main.go:5:4: could not import runtime/msan (invalid package name: "")
error: couldn't load packages due to errors: github.com/trashhalo/tinygo-import-report
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/runtime_msan:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/go/src/github.com/trashhalo/tinygo-import-report/main.go:5:4: could not import runtime/msan (invalid package name: "")
error: couldn't load packages due to errors: github.com/trashhalo/tinygo-import-report
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## runtime/pprof
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/runtime_pprof:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/runtime/pprof/pprof.go:372:135: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:693:84: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:718:31: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/protomem.go:15:46: MemProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/proto.go:223:20: CallersFrames not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:920:20: SetMutexProfileFraction not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:807:10: SetCPUProfileRate not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:762:10: SetCPUProfileRate not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:700:18: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:706:22: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:678:16: Stack not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:542:24: MemStats not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:546:26: MemStats not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:547:11: ReadMemStats not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:556:18: MemProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:557:19: MemProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:562:22: MemProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:563:19: MemProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:572:45: MemProfileRate not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:581:20: MemProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:596:13: MemProfileRate not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:490:20: CallersFrames not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:877:18: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:878:19: MutexProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:880:22: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:881:19: MutexProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:900:49: SetMutexProfileFraction not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:819:18: MutexProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:825:18: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:826:19: BlockProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:828:22: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:829:19: BlockProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:813:18: BlockProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:526:18: MemProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:656:63: ThreadCreateProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:650:18: ThreadCreateProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:669:60: GoroutineProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:661:17: NumGoroutine not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:287:15: Callers not declared by package runtime
error: couldn't load packages due to errors: context, runtime/pprof
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/runtime_pprof:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/runtime/pprof/pprof.go:372:135: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:693:84: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:718:31: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/protomem.go:15:46: MemProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/proto.go:223:20: CallersFrames not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:920:20: SetMutexProfileFraction not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:807:10: SetCPUProfileRate not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:762:10: SetCPUProfileRate not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:700:18: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:706:22: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:678:16: Stack not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:542:24: MemStats not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:546:26: MemStats not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:547:11: ReadMemStats not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:556:18: MemProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:557:19: MemProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:562:22: MemProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:563:19: MemProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:572:45: MemProfileRate not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:581:20: MemProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:596:13: MemProfileRate not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:490:20: CallersFrames not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:877:18: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:878:19: MutexProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:880:22: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:881:19: MutexProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:900:49: SetMutexProfileFraction not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:819:18: MutexProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:825:18: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:826:19: BlockProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:828:22: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:829:19: BlockProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:813:18: BlockProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:526:18: MemProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:656:63: ThreadCreateProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:650:18: ThreadCreateProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:669:60: GoroutineProfile not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:661:17: NumGoroutine not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:287:15: Callers not declared by package runtime
error: couldn't load packages due to errors: context, runtime/pprof
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## runtime/race
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/runtime_race:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/runtime_race:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## runtime/trace
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/runtime_trace:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/runtime/trace/trace.go:147:10: StopTrace not declared by package runtime
/usr/local/go/src/runtime/trace/trace.go:124:20: StartTrace not declared by package runtime
/usr/local/go/src/runtime/trace/trace.go:129:20: ReadTrace not declared by package runtime
error: couldn't load packages due to errors: runtime/trace, context
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/runtime_trace:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/runtime/trace/trace.go:147:10: StopTrace not declared by package runtime
/usr/local/go/src/runtime/trace/trace.go:124:20: StartTrace not declared by package runtime
/usr/local/go/src/runtime/trace/trace.go:129:20: ReadTrace not declared by package runtime
error: couldn't load packages due to errors: context, runtime/trace
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## sort
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/sort:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/sort:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## strconv
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/strconv:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/strconv:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## strings
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/strings:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/strings:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## sync
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/sync:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/sync:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## sync/atomic
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/sync_atomic:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/sync_atomic:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## syscall
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/syscall:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/syscall:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## syscall/js
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/syscall_js:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/syscall_js:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## testing
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/testing:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/runtime/debug/stack.go:24:16: Stack not declared by package runtime
/usr/local/go/src/runtime/debug/garbage.go:37:34: MemStats not declared by package runtime
/usr/local/go/src/flag/flag.go:410:15: New not declared by package reflect
/usr/local/go/src/flag/flag.go:412:15: Zero not declared by package reflect
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/runtime/trace/trace.go:147:10: StopTrace not declared by package runtime
/usr/local/go/src/runtime/trace/trace.go:124:20: StartTrace not declared by package runtime
/usr/local/go/src/runtime/trace/trace.go:129:20: ReadTrace not declared by package runtime
/usr/local/go/src/testing/benchmark.go:26:22: MemStats not declared by package runtime
/usr/local/go/src/testing/testing.go:339:46: Frame not declared by package runtime
/usr/local/go/src/testing/testing.go:1136:11: MemProfileRate not declared by package runtime
/usr/local/go/src/testing/testing.go:1165:11: SetBlockProfileRate not declared by package runtime
/usr/local/go/src/testing/testing.go:1168:11: SetMutexProfileFraction not declared by package runtime
/usr/local/go/src/testing/testing.go:693:15: Callers not declared by package runtime
/usr/local/go/src/testing/testing.go:697:20: CallersFrames not declared by package runtime
/usr/local/go/src/testing/benchmark.go:616:14: WaitGroup not declared by package sync
/usr/local/go/src/testing/benchmark.go:493:15: Callers not declared by package runtime
/usr/local/go/src/testing/benchmark.go:101:11: ReadMemStats not declared by package runtime
/usr/local/go/src/testing/benchmark.go:90:11: ReadMemStats not declared by package runtime
/usr/local/go/src/testing/benchmark.go:76:11: ReadMemStats not declared by package runtime
/usr/local/go/src/testing/testing.go:660:10: Goexit not declared by package runtime
/usr/local/go/src/testing/testing.go:590:10: Goexit not declared by package runtime
/usr/local/go/src/testing/testing.go:351:15: Callers not declared by package runtime
/usr/local/go/src/testing/testing.go:355:20: CallersFrames not declared by package runtime
/usr/local/go/src/testing/testing.go:356:43: Frame not declared by package runtime
/usr/local/go/src/testing/testing.go:356:43: Frame not declared by package runtime
/usr/local/go/src/testing/testing.go:356:43: Frame not declared by package runtime
/usr/local/go/src/testing/testing.go:370:22: CallersFrames not declared by package runtime
/usr/local/go/src/testing/testing.go:849:15: Callers not declared by package runtime
/usr/local/go/src/testing/testing.go:882:11: Goexit not declared by package runtime
/usr/local/go/src/testing/allocs.go:27:23: MemStats not declared by package runtime
/usr/local/go/src/testing/allocs.go:28:10: ReadMemStats not declared by package runtime
/usr/local/go/src/testing/allocs.go:37:10: ReadMemStats not declared by package runtime
error: couldn't load packages due to errors: flag, context, runtime/debug and 2 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/testing:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/runtime/debug/stack.go:24:16: Stack not declared by package runtime
/usr/local/go/src/runtime/debug/garbage.go:37:34: MemStats not declared by package runtime
/usr/local/go/src/flag/flag.go:410:15: New not declared by package reflect
/usr/local/go/src/flag/flag.go:412:15: Zero not declared by package reflect
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/runtime/trace/trace.go:147:10: StopTrace not declared by package runtime
/usr/local/go/src/runtime/trace/trace.go:124:20: StartTrace not declared by package runtime
/usr/local/go/src/runtime/trace/trace.go:129:20: ReadTrace not declared by package runtime
/usr/local/go/src/testing/benchmark.go:26:22: MemStats not declared by package runtime
/usr/local/go/src/testing/testing.go:339:46: Frame not declared by package runtime
/usr/local/go/src/testing/testing.go:1136:11: MemProfileRate not declared by package runtime
/usr/local/go/src/testing/testing.go:1165:11: SetBlockProfileRate not declared by package runtime
/usr/local/go/src/testing/testing.go:1168:11: SetMutexProfileFraction not declared by package runtime
/usr/local/go/src/testing/testing.go:693:15: Callers not declared by package runtime
/usr/local/go/src/testing/testing.go:697:20: CallersFrames not declared by package runtime
/usr/local/go/src/testing/benchmark.go:616:14: WaitGroup not declared by package sync
/usr/local/go/src/testing/benchmark.go:493:15: Callers not declared by package runtime
/usr/local/go/src/testing/benchmark.go:101:11: ReadMemStats not declared by package runtime
/usr/local/go/src/testing/benchmark.go:90:11: ReadMemStats not declared by package runtime
/usr/local/go/src/testing/benchmark.go:76:11: ReadMemStats not declared by package runtime
/usr/local/go/src/testing/testing.go:660:10: Goexit not declared by package runtime
/usr/local/go/src/testing/testing.go:590:10: Goexit not declared by package runtime
/usr/local/go/src/testing/testing.go:351:15: Callers not declared by package runtime
/usr/local/go/src/testing/testing.go:355:20: CallersFrames not declared by package runtime
/usr/local/go/src/testing/testing.go:356:43: Frame not declared by package runtime
/usr/local/go/src/testing/testing.go:356:43: Frame not declared by package runtime
/usr/local/go/src/testing/testing.go:356:43: Frame not declared by package runtime
/usr/local/go/src/testing/testing.go:370:22: CallersFrames not declared by package runtime
/usr/local/go/src/testing/testing.go:849:15: Callers not declared by package runtime
/usr/local/go/src/testing/testing.go:882:11: Goexit not declared by package runtime
/usr/local/go/src/testing/allocs.go:27:23: MemStats not declared by package runtime
/usr/local/go/src/testing/allocs.go:28:10: ReadMemStats not declared by package runtime
/usr/local/go/src/testing/allocs.go:37:10: ReadMemStats not declared by package runtime
error: couldn't load packages due to errors: testing, context, runtime/trace and 2 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## testing/iotest
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/testing_iotest:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/testing_iotest:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## testing/quick
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/testing_quick:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/flag/flag.go:410:15: New not declared by package reflect
/usr/local/go/src/flag/flag.go:412:15: Zero not declared by package reflect
/usr/local/go/src/testing/quick/quick.go:350:25: invalid operation: f (variable of type reflect.Type) has no field or method In
/usr/local/go/src/testing/quick/quick.go:352:95: invalid operation: f (variable of type reflect.Type) has no field or method In
/usr/local/go/src/testing/quick/quick.go:319:43: invalid operation: xType (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/testing/quick/quick.go:329:26: invalid operation: x (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/testing/quick/quick.go:330:26: invalid operation: y (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/testing/quick/quick.go:332:15: DeepEqual not declared by package reflect
/usr/local/go/src/testing/quick/quick.go:272:11: invalid operation: fType (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/testing/quick/quick.go:275:11: invalid operation: fType (variable of type reflect.Type) has no field or method Out
/usr/local/go/src/testing/quick/quick.go:279:43: invalid operation: fType (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/testing/quick/quick.go:289:12: invalid operation: fVal (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/testing/quick/quick.go:67:22: Zero not declared by package reflect
/usr/local/go/src/testing/quick/quick.go:71:15: New not declared by package reflect
/usr/local/go/src/testing/quick/quick.go:107:17: MakeMap not declared by package reflect
/usr/local/go/src/testing/quick/quick.go:109:36: invalid operation: concrete (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/testing/quick/quick.go:118:18: Zero not declared by package reflect
/usr/local/go/src/testing/quick/quick.go:124:18: New not declared by package reflect
error: couldn't load packages due to errors: testing/quick, flag
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/testing_quick:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/flag/flag.go:410:15: New not declared by package reflect
/usr/local/go/src/flag/flag.go:412:15: Zero not declared by package reflect
/usr/local/go/src/testing/quick/quick.go:350:25: invalid operation: f (variable of type reflect.Type) has no field or method In
/usr/local/go/src/testing/quick/quick.go:352:95: invalid operation: f (variable of type reflect.Type) has no field or method In
/usr/local/go/src/testing/quick/quick.go:319:43: invalid operation: xType (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/testing/quick/quick.go:329:26: invalid operation: x (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/testing/quick/quick.go:330:26: invalid operation: y (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/testing/quick/quick.go:332:15: DeepEqual not declared by package reflect
/usr/local/go/src/testing/quick/quick.go:272:11: invalid operation: fType (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/testing/quick/quick.go:275:11: invalid operation: fType (variable of type reflect.Type) has no field or method Out
/usr/local/go/src/testing/quick/quick.go:279:43: invalid operation: fType (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/testing/quick/quick.go:289:12: invalid operation: fVal (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/testing/quick/quick.go:67:22: Zero not declared by package reflect
/usr/local/go/src/testing/quick/quick.go:71:15: New not declared by package reflect
/usr/local/go/src/testing/quick/quick.go:107:17: MakeMap not declared by package reflect
/usr/local/go/src/testing/quick/quick.go:109:36: invalid operation: concrete (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/testing/quick/quick.go:118:18: Zero not declared by package reflect
/usr/local/go/src/testing/quick/quick.go:124:18: New not declared by package reflect
error: couldn't load packages due to errors: testing/quick, flag
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## text/scanner
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/text_scanner:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
error: todo: init: map key not implemented: rune
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/text_scanner:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## text/tabwriter
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/text_tabwriter:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/text_tabwriter:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## text/template
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/text_template:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
/usr/local/go/src/text/template/funcs.go:140:19: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:142:18: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/funcs.go:145:70: invalid operation: value.Type() (value of type reflect.Type) has no field or method ConvertibleTo
/usr/local/go/src/text/template/funcs.go:146:17: invalid operation: value (variable of type reflect.Value) has no field or method Convert
/usr/local/go/src/text/template/funcs.go:93:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:32: invalid operation: typ (variable of type reflect.Type) has no field or method Out
/usr/local/go/src/text/template/funcs.go:75:88: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:196:45: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/funcs.go:203:17: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:248:92: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:250:15: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/funcs.go:252:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:256:17: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:267:11: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:268:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:278:14: invalid operation: v (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:949:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:949:50: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:950:30: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:950:79: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:159:16: Error not declared by package runtime
/usr/local/go/src/text/template/exec.go:866:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:856:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:845:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:834:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:823:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:812:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:769:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:793:10: invalid operation: typ (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:722:13: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:728:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:735:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:735:33: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:738:20: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:748:59: invalid operation: value.Type().Elem() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:753:16: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:653:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:654:18: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:656:79: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:658:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:659:69: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:663:71: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/exec.go:670:32: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:673:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:674:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:674:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:681:12: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:681:19: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:682:10: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:686:13: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:695:16: invalid operation: fun (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:584:19: invalid operation: ptr (variable of type reflect.Value) has no field or method MethodByName
/usr/local/go/src/text/template/exec.go:591:33: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/text/template/exec.go:596:22: invalid operation: receiver (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/text/template/exec.go:612:21: invalid operation: nameVal.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:612:50: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/exec.go:622:23: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:529:25: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:431:56: invalid operation: value.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:375:20: invalid operation: val (variable of type reflect.Value) has no field or method Recv
error: couldn't load packages due to errors: text/template, text/template/parse
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/text_template:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
/usr/local/go/src/text/template/funcs.go:140:19: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:142:18: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/funcs.go:145:70: invalid operation: value.Type() (value of type reflect.Type) has no field or method ConvertibleTo
/usr/local/go/src/text/template/funcs.go:146:17: invalid operation: value (variable of type reflect.Value) has no field or method Convert
/usr/local/go/src/text/template/funcs.go:93:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:32: invalid operation: typ (variable of type reflect.Type) has no field or method Out
/usr/local/go/src/text/template/funcs.go:75:88: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:196:45: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/funcs.go:203:17: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:248:92: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:250:15: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/funcs.go:252:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:256:17: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:267:11: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/funcs.go:268:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/funcs.go:278:14: invalid operation: v (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:949:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:949:50: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:950:30: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:950:79: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:159:16: Error not declared by package runtime
/usr/local/go/src/text/template/exec.go:866:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:856:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:845:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:834:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:823:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:812:20: New not declared by package reflect
/usr/local/go/src/text/template/exec.go:769:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:793:10: invalid operation: typ (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:722:13: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:728:19: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:735:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:735:33: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:738:20: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:748:59: invalid operation: value.Type().Elem() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:753:16: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:653:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:654:18: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:656:79: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:658:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:659:69: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:663:71: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/exec.go:670:32: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:673:9: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:674:18: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:674:25: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:681:12: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:681:19: invalid operation: typ (variable of type reflect.Type) has no field or method NumIn
/usr/local/go/src/text/template/exec.go:682:10: invalid operation: typ (variable of type reflect.Type) has no field or method IsVariadic
/usr/local/go/src/text/template/exec.go:686:13: invalid operation: typ (variable of type reflect.Type) has no field or method In
/usr/local/go/src/text/template/exec.go:695:16: invalid operation: fun (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/text/template/exec.go:584:19: invalid operation: ptr (variable of type reflect.Value) has no field or method MethodByName
/usr/local/go/src/text/template/exec.go:591:33: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/text/template/exec.go:596:22: invalid operation: receiver (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/text/template/exec.go:612:21: invalid operation: nameVal.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:612:50: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/exec.go:622:23: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:529:25: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:431:56: invalid operation: value.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:375:20: invalid operation: val (variable of type reflect.Value) has no field or method Recv
error: couldn't load packages due to errors: text/template, text/template/parse
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## text/template/parse
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/text_template_parse:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
error: couldn't load packages due to errors: text/template/parse
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/text_template_parse:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
error: couldn't load packages due to errors: text/template/parse
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## time
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/time:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/time:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## unicode
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/unicode:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/unicode:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## unicode/utf16
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/unicode_utf16:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/unicode_utf16:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## unicode/utf8
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/unicode_utf8:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/unicode_utf8:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## unsafe
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/unsafe:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/unsafe:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

