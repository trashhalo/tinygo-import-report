
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
| compress/flate |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| compress/gzip |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| compress/lzw |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| compress/zlib |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| container |  [:x:](#container)  |  [:x:](#container)  | 
| container/heap |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| container/list |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| container/ring |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| context |  [:x:](#context)  |  [:x:](#context)  | 
| crypto |  :heavy_check_mark:  |  :heavy_check_mark:  | 
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
| crypto/sha1 |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| crypto/sha256 |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| crypto/sha512 |  :heavy_check_mark:  |  :heavy_check_mark:  | 
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
| hash/crc32 |  :heavy_check_mark:  |  :heavy_check_mark:  | 
| hash/crc64 |  :heavy_check_mark:  |  :heavy_check_mark:  | 
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
| text/scanner |  :heavy_check_mark:  |  :heavy_check_mark:  | 
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
Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !35958
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !35983
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36033
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36122
error: Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !35958
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !35983
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36033
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36122

Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_dsa:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !35958
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !35983
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36033
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36122
error: Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !35958
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !35983
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36033
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36122

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
Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36572
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36597
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36647
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36736
error: Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36572
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36597
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36647
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36736

Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_elliptic:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36572
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36597
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36647
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36736
error: Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36572
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36597
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36647
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36736

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
panic: interp: load from a pointer bitcast: &PointerCastValue{Value: &GlobalValue{crypto/md5$alloc}, CastType: PointerType(ArrayType(IntegerType(8 bits)[4]))}

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*PointerCastValue).Load(0xc001560520, 0x7f492023fa00)
	/go/src/github.com/aykevl/tinygo/interp/values.go:369 +0x1ff
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000111058, 0x7f492023efe0, 0x0, 0x7eb6c2, 0x4, 0x7f492023efe0, 0x7f4920240410, 0x7f4920008488, 0x0, 0xe4cf1ff300000000, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:96 +0x695f
github.com/aykevl/tinygo/interp.(*Eval).function(0xc00155b940, 0x7f4920019b08, 0x0, 0x0, 0x0, 0xc000b4a5a0, 0xa, 0x7eb6c2, 0x4, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000111648, 0x7f4920185cc0, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc000111628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc00155b940, 0x7f4920019958, 0x0, 0x0, 0x0, 0xc000b4a5a0, 0xa, 0x0, 0x0, 0xc0001116e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc00155b940, 0x7f4920019958, 0x0, 0x0, 0x0, 0xc000b4a5a0, 0xa, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x258bc20, 0x2581b40, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffc1db5af21, 0x29, 0x7ffc1db5aed8, 0x3b, 0xc000086d00, 0xc000111f48, 0xc000111da8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffc1db5af21, 0x29, 0x7ffc1db5aed8, 0x3b, 0x7ffc1db5af1c, 0x4, 0xc0000a1f48, 0x90, 0xc0000a1eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 2

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_md5:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: interp: load from a pointer bitcast: &PointerCastValue{Value: &GlobalValue{crypto/md5$alloc}, CastType: PointerType(ArrayType(IntegerType(8 bits)[4]))}

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*PointerCastValue).Load(0xc00130a080, 0x7f0f9823f700)
	/go/src/github.com/aykevl/tinygo/interp/values.go:369 +0x1ff
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00015b058, 0x7f0f9823ed40, 0x0, 0x7eb6c2, 0x4, 0x7f0f9823ed40, 0x7f0f98240170, 0x7f0f980081e8, 0x0, 0xe97376ab00000000, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:96 +0x695f
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0015e9c40, 0x7f0f98019868, 0x0, 0x0, 0x0, 0xc000fe4760, 0xa, 0x7eb6c2, 0x4, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00015b648, 0x7f0f98185a30, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc00015b628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0015e9c40, 0x7f0f980196b8, 0x0, 0x0, 0x0, 0xc000fe4760, 0xa, 0x0, 0x0, 0xc00015b6e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc0015e9c40, 0x7f0f980196b8, 0x0, 0x0, 0x0, 0xc000fe4760, 0xa, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x11bdc20, 0x11b3b70, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7fff3e9acf21, 0x29, 0x7fff3e9aced8, 0x3b, 0xc0001321a0, 0xc00015bf48, 0xc00015bda8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7fff3e9acf21, 0x29, 0x7fff3e9aced8, 0x3b, 0x7fff3e9acf1c, 0x4, 0xc0000a1f48, 0x425124, 0xc0000a1eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 2

```

## crypto/rand
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_rand:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36197
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36222
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36272
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36361
error: Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36197
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36222
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36272
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36361

Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_rand:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36197
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36222
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36272
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36361
error: Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36197
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36222
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36272
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36361

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
Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36391
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36416
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36466
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36555
error: Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36391
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36416
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36466
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36555

Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_rsa:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36391
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36416
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36466
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36555
error: Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !36391
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !36416
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36466
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36555

Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## crypto/sha1
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_sha1:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_sha1:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## crypto/sha256
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_sha256:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_sha256:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

## crypto/sha512
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_sha512:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/crypto_sha512:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report

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
error: couldn't load packages due to errors: encoding/asn1, net, internal/singleflight and 2 more
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
error: couldn't load packages due to errors: encoding/asn1, context, internal/singleflight and 2 more
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
error: couldn't load packages due to errors: context, database/sql, database/sql/driver
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
error: couldn't load packages due to errors: context, database/sql/driver
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
panic: branch was not true or false

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00014a478, 0x7f375c96c880, 0x0, 0xc001068240, 0xc, 0xc, 0x4, 0xc, 0x4, 0x4, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:422 +0x69da
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0031cc140, 0x7f375c02ca68, 0xc00305c340, 0x3, 0x4, 0xc000e598c0, 0xf, 0xc001068240, 0xc, 0xc000000000, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00014aa68, 0x7f375c99ef10, 0x7f375c99ed70, 0xc0010681b8, 0x8, 0x0, 0x0, 0xc00288c358, 0x1, 0x1, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0031cc140, 0x7f375c030ff8, 0xc0031d2900, 0x5, 0x8, 0xc000e598c0, 0xf, 0xc0010681b8, 0x8, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00014b058, 0x7f375c9afc00, 0x7f375c9afa40, 0x7eb6c2, 0x4, 0x0, 0x0, 0xc00288c330, 0x1, 0x1, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0031cc140, 0x7f375c032c98, 0xc0031b2820, 0x2, 0x2, 0xc000e598c0, 0xf, 0x7eb6c2, 0x4, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00014b648, 0x7f375c054360, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc00014b628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0031cc140, 0x7f375c032d98, 0x0, 0x0, 0x0, 0xc000e598c0, 0xf, 0x0, 0x0, 0xc00014b6e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc0031cc140, 0x7f375c032d98, 0x0, 0x0, 0x0, 0xc000e598c0, 0xf, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x10b5c20, 0x10abb40, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffddb93cf21, 0x29, 0x7ffddb93ced8, 0x3b, 0xc000128340, 0xc00014bf48, 0xc00014bda8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffddb93cf21, 0x29, 0x7ffddb93ced8, 0x3b, 0x7ffddb93cf1c, 0x4, 0xc0000a1f48, 0x24, 0xc0000a1eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 2

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_base32:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: branch was not true or false

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000162478, 0x7f015806f320, 0x0, 0xc002a75030, 0xc, 0xc, 0x4, 0xc, 0x4, 0x4, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:422 +0x69da
github.com/aykevl/tinygo/interp.(*Eval).function(0xc003b043c0, 0x7f015402ca68, 0xc0021ec900, 0x3, 0x4, 0xc003652b80, 0xf, 0xc002a75030, 0xc, 0xc000000000, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000162a68, 0x7f01580a1f60, 0x7f01580a1dc0, 0xc002a74fa8, 0x8, 0x0, 0x0, 0xc0025180b0, 0x1, 0x1, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc003b043c0, 0x7f0154030ff8, 0xc003ad5480, 0x5, 0x8, 0xc003652b80, 0xf, 0xc002a74fa8, 0x8, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000163058, 0x7f01580b2c60, 0x7f01580b2aa0, 0x7eb6c2, 0x4, 0x0, 0x0, 0xc002518088, 0x1, 0x1, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc003b043c0, 0x7f0154032c98, 0xc0039fa9c0, 0x2, 0x2, 0xc003652b80, 0xf, 0x7eb6c2, 0x4, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000163648, 0x7f01580309d0, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc000163628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc003b043c0, 0x7f0154032d98, 0x0, 0x0, 0x0, 0xc003652b80, 0xf, 0x0, 0x0, 0xc0001636e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc003b043c0, 0x7f0154032d98, 0x0, 0x0, 0x0, 0xc003652b80, 0xf, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x2488c20, 0x247eb70, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffcf3448f21, 0x29, 0x7ffcf3448ed8, 0x3b, 0xc000144410, 0xc000163f48, 0xc000163da8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffcf3448f21, 0x29, 0x7ffcf3448ed8, 0x3b, 0x7ffcf3448f1c, 0x4, 0xc0000a1f48, 0x24, 0xc0000a1eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 2

```

## encoding/base64
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_base64:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: branch was not true or false

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00012a478, 0x290a9e0, 0x0, 0xc0018b79a0, 0xc, 0xc, 0x4, 0xc, 0x4, 0x4, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:422 +0x69da
github.com/aykevl/tinygo/interp.(*Eval).function(0xc00175db80, 0x2759888, 0xc0018c2f00, 0x3, 0x4, 0xc000aeb0a0, 0xf, 0xc0018b79a0, 0xc, 0xc000000000, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00012aa68, 0x2942680, 0x29424e0, 0xc0018b7910, 0x8, 0x0, 0x0, 0xc001370d20, 0x1, 0x1, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc00175db80, 0x27548b8, 0xc001896080, 0x5, 0x8, 0xc000aeb0a0, 0xf, 0xc0018b7910, 0x8, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00012b058, 0x2955b50, 0x2955900, 0x7eb6c2, 0x4, 0x0, 0x0, 0xc001370cf0, 0x1, 0x1, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc00175db80, 0x2761cb8, 0xc00189a800, 0x2, 0x2, 0xc000aeb0a0, 0xf, 0x7eb6c2, 0x4, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00012b648, 0x28ca4e0, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc00012b628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc00175db80, 0x2761db8, 0x0, 0x0, 0x0, 0xc000aeb0a0, 0xf, 0x0, 0x0, 0xc00012b6e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc00175db80, 0x2761db8, 0x0, 0x0, 0x0, 0xc000aeb0a0, 0xf, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x2746d70, 0x273cc90, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffec4c83f21, 0x29, 0x7ffec4c83ed8, 0x3b, 0xc0000e64e0, 0xc00012bf48, 0xc00012bda8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffec4c83f21, 0x29, 0x7ffec4c83ed8, 0x3b, 0x7ffec4c83f1c, 0x4, 0xc000081f48, 0x24, 0xc000081eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 2

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_base64:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: branch was not true or false

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000146478, 0x7fa40c1c2660, 0x0, 0xc001de6650, 0xc, 0xc, 0x4, 0xc, 0x4, 0x4, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:422 +0x69da
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0006f09c0, 0x7fa40c0112a8, 0xc0006f1500, 0x3, 0x4, 0xc001d869e0, 0xf, 0xc001de6650, 0xc, 0xc000000000, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000146a68, 0x7fa40c1fa310, 0x7fa40c1fa170, 0xc001de65c0, 0x8, 0x0, 0x0, 0xc001ddb540, 0x1, 0x1, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0006f09c0, 0x7fa40c00bec8, 0xc000465780, 0x5, 0x8, 0xc001d869e0, 0xf, 0xc001de65c0, 0x8, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000147058, 0x7fa40c20d800, 0x7fa40c20d5b0, 0x7eb6c2, 0x4, 0x0, 0x0, 0xc001ddb518, 0x1, 0x1, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0006f09c0, 0x7fa40c0198d8, 0xc00127aec0, 0x2, 0x2, 0xc001d869e0, 0xf, 0x7eb6c2, 0x4, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000147648, 0x7fa40c182100, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc000147628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0006f09c0, 0x7fa40c0199d8, 0x0, 0x0, 0x0, 0xc001d869e0, 0xf, 0x0, 0x0, 0xc0001476e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc0006f09c0, 0x7fa40c0199d8, 0x0, 0x0, 0x0, 0xc001d869e0, 0xf, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x21d7c20, 0x21cdb40, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffe83083f21, 0x29, 0x7ffe83083ed8, 0x3b, 0xc0000868f0, 0xc000147f48, 0xc000147da8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffe83083f21, 0x29, 0x7ffe83083ed8, 0x3b, 0x7ffe83083f1c, 0x4, 0xc0000a1f48, 0x90, 0xc0000a1eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 2

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
panic: branch was not true or false

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000162478, 0x7fc024970590, 0x0, 0xc0049ac910, 0xc, 0xc, 0x4, 0xc, 0x4, 0x4, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:422 +0x69da
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0034b2900, 0x7fc02402d938, 0xc0033e6f40, 0x3, 0x4, 0xc0044a1ce0, 0xf, 0xc0049ac910, 0xc, 0xc000000000, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000162a68, 0x7fc0249a2600, 0x7fc0249a2460, 0xc0049ac880, 0x8, 0x0, 0x0, 0xc0044a3cf8, 0x1, 0x1, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0034b2900, 0x7fc024031ec8, 0xc00363a200, 0x5, 0x8, 0xc0044a1ce0, 0xf, 0xc0049ac880, 0x8, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000163058, 0x7fc0249b5e80, 0x7fc0249b5c30, 0x7eb6c2, 0x4, 0x0, 0x0, 0xc0044a3cd0, 0x1, 0x1, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0034b2900, 0x7fc024033d78, 0xc002fad060, 0x2, 0x2, 0xc0044a1ce0, 0xf, 0x7eb6c2, 0x4, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000163648, 0x7fc02492fb10, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc000163628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0034b2900, 0x7fc024033e78, 0x0, 0x0, 0x0, 0xc0044a1ce0, 0xf, 0x0, 0x0, 0xc0001636e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc0034b2900, 0x7fc024033e78, 0x0, 0x0, 0x0, 0xc0044a1ce0, 0xf, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x2c3bc20, 0x2c31b70, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffc78d37f21, 0x29, 0x7ffc78d37ed8, 0x3b, 0xc000140340, 0xc000163f48, 0xc000163da8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffc78d37f21, 0x29, 0x7ffc78d37ed8, 0x3b, 0x7ffc78d37f1c, 0x4, 0xc00009ff48, 0x24, 0xc00009feb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 2

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/encoding_pem:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: branch was not true or false

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00014a478, 0x7efef4970460, 0x0, 0xc0042ca020, 0xc, 0xc, 0x4, 0xc, 0x4, 0x4, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:422 +0x69da
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0038f50c0, 0x7efef402d7e8, 0xc00389db00, 0x3, 0x4, 0xc0042ae2a0, 0xf, 0xc0042ca020, 0xc, 0xc000000000, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00014aa68, 0x7efef49a24e0, 0x7efef49a2340, 0xc0042a9f90, 0x8, 0x0, 0x0, 0xc00459ae78, 0x1, 0x1, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0038f50c0, 0x7efef4031d78, 0xc0038eee00, 0x5, 0x8, 0xc0042ae2a0, 0xf, 0xc0042a9f90, 0x8, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00014b058, 0x7efef49b5d50, 0x7efef49b5b00, 0x7eb6c2, 0x4, 0x0, 0x0, 0xc00459ae50, 0x1, 0x1, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0038f50c0, 0x7efef4033c28, 0xc0034c3aa0, 0x2, 0x2, 0xc0042ae2a0, 0xf, 0x7eb6c2, 0x4, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00014b648, 0x7efef492f9c0, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc00014b628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0038f50c0, 0x7efef4033d28, 0x0, 0x0, 0x0, 0xc0042ae2a0, 0xf, 0x0, 0x0, 0xc00014b6e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc0038f50c0, 0x7efef4033d28, 0x0, 0x0, 0x0, 0xc0042ae2a0, 0xf, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x24a3c20, 0x2499b40, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffee19d1f21, 0x29, 0x7ffee19d1ed8, 0x3b, 0xc000128340, 0xc00014bf48, 0xc00014bda8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffee19d1f21, 0x29, 0x7ffee19d1ed8, 0x3b, 0x7ffee19d1f1c, 0x4, 0xc0000a1f48, 0x24, 0xc0000a1eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 2

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
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1176:51: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:749:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:698:11: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:703:9: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
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
error: couldn't load packages due to errors: expvar, vendor/golang_org/x/crypto/cryptobyte, internal/singleflight and 7 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/expvar:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
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
error: couldn't load packages due to errors: internal/singleflight, encoding/json, net and 7 more
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
unsupported instruction during init evaluation:
  %7 = inttoptr i32 %6 to i8*, !dbg !19
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/html:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
unsupported instruction during init evaluation:
  %7 = inttoptr i32 %6 to i8*, !dbg !19
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
/usr/local/go/src/text/template/exec.go:866:20: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:526:8: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
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
/usr/local/go/src/encoding/json/decode.go:557:14: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:17: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:558:27: invalid operation: v (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:563:13: Copy not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:567:7: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:599:17: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:604:6: invalid operation: v (variable of type reflect.Value) has no field or method SetLen
/usr/local/go/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:375:20: invalid operation: val (variable of type reflect.Value) has no field or method Recv
/usr/local/go/src/html/template/js.go:135:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:57: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
error: couldn't load packages due to errors: html/template, encoding/json, text/template and 1 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/html_template:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/text/template/funcs.go:140:19: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:142:18: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/funcs.go:145:70: invalid operation: value.Type() (value of type reflect.Type) has no field or method ConvertibleTo
/usr/local/go/src/text/template/funcs.go:146:17: invalid operation: value (variable of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1176:51: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/text/template/funcs.go:93:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:11: invalid operation: typ (variable of type reflect.Type) has no field or method NumOut
/usr/local/go/src/text/template/funcs.go:95:32: invalid operation: typ (variable of type reflect.Type) has no field or method Out
/usr/local/go/src/text/template/funcs.go:75:88: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumOut
/usr/local/go/src/encoding/json/encode.go:749:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:698:11: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:703:9: invalid operation: t (variable of type reflect.Type) has no field or method Key
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
/usr/local/go/src/encoding/json/encode.go:391:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:395:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:400:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:404:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:364:11: WaitGroup not declared by package sync
/usr/local/go/src/encoding/json/decode.go:459:41: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/text/template/exec.go:529:25: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/encoding/json/decode.go:483:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:485:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:431:56: invalid operation: value.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:375:20: invalid operation: val (variable of type reflect.Value) has no field or method Recv
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
/usr/local/go/src/html/template/js.go:135:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:57: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
error: couldn't load packages due to errors: html/template, text/template/parse, encoding/json and 1 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## image
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: interp: store on a pointer bitcast

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*PointerCastValue).Store(0xc0032508c0, 0x7f18ec024a80)
	/go/src/github.com/aykevl/tinygo/interp/values.go:374 +0x39
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000139648, 0x7f18ec936620, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc000139628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:108 +0x18c0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc003c62b00, 0x7f18ec0370d8, 0x0, 0x0, 0x0, 0xc003700e84, 0x5, 0x0, 0x0, 0xc0001396e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc003c62b00, 0x7f18ec0370d8, 0x0, 0x0, 0x0, 0xc003700e84, 0x5, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x116bd70, 0x1161c90, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffe4265cf21, 0x29, 0x7ffe4265ced8, 0x3b, 0xc000118340, 0xc000139f48, 0xc000139da8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffe4265cf21, 0x29, 0x7ffe4265ced8, 0x3b, 0x7ffe4265cf1c, 0x4, 0xc000081f48, 0x24, 0xc000081eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 2

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: interp: store on a pointer bitcast

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*PointerCastValue).Store(0xc002e50880, 0x7f0648024930)
	/go/src/github.com/aykevl/tinygo/interp/values.go:374 +0x39
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00014b648, 0x7f06489364e0, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc00014b628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:108 +0x18c0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc000996280, 0x7f0648036f88, 0x0, 0x0, 0x0, 0xc0017c7014, 0x5, 0x0, 0x0, 0xc00014b6e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc000996280, 0x7f0648036f88, 0x0, 0x0, 0x0, 0xc0017c7014, 0x5, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x1b46c20, 0x1b3cb40, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7fff52bedf21, 0x29, 0x7fff52beded8, 0x3b, 0xc000128340, 0xc00014bf48, 0xc00014bda8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7fff52bedf21, 0x29, 0x7fff52beded8, 0x3b, 0x7fff52bedf1c, 0x4, 0xc0000a1f48, 0x24, 0xc0000a1eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
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
panic: interp: store on a pointer bitcast

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*PointerCastValue).Store(0xc00075dc00, 0x7f89ac1d4d90)
	/go/src/github.com/aykevl/tinygo/interp/values.go:374 +0x39
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00014b648, 0x7f89ac0178e0, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc00014b628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:108 +0x18c0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc000822040, 0x7f89ac013188, 0x0, 0x0, 0x0, 0xc000872580, 0x13, 0x0, 0x0, 0xc00014b6e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc000822040, 0x7f89ac013188, 0x0, 0x0, 0x0, 0xc000872580, 0x13, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x1ed1c20, 0x1ec7b40, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffc09cecf21, 0x29, 0x7ffc09ceced8, 0x3b, 0xc000128340, 0xc00014bf48, 0xc00014bda8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffc09cecf21, 0x29, 0x7ffc09ceced8, 0x3b, 0x7ffc09cecf1c, 0x4, 0xc0000a1f48, 0x24, 0xc0000a1eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 2

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_color_palette:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: interp: store on a pointer bitcast

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*PointerCastValue).Store(0xc0006e3280, 0x7fa6741cc0e0)
	/go/src/github.com/aykevl/tinygo/interp/values.go:374 +0x39
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00014b648, 0x7fa66c0178e0, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc00014b628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:108 +0x18c0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc0007dfd40, 0x7fa66c013188, 0x0, 0x0, 0x0, 0xc0005d2100, 0x13, 0x0, 0x0, 0xc00014b6e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc0007dfd40, 0x7fa66c013188, 0x0, 0x0, 0x0, 0xc0005d2100, 0x13, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x2536c20, 0x252cb40, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffc7fa72f21, 0x29, 0x7ffc7fa72ed8, 0x3b, 0xc000128340, 0xc00014bf48, 0xc00014bda8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffc7fa72f21, 0x29, 0x7ffc7fa72ed8, 0x3b, 0x7ffc7fa72f1c, 0x4, 0xc0000a1f48, 0x24, 0xc0000a1eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 2

```

## image/draw
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_draw:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: interp: store on a pointer bitcast

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*PointerCastValue).Store(0xc0038eabc0, 0x7f4b40024ac0)
	/go/src/github.com/aykevl/tinygo/interp/values.go:374 +0x39
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000153648, 0x7f4b40938110, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc000153628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:108 +0x18c0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc003b9e600, 0x7f4b4001b108, 0x0, 0x0, 0x0, 0xc00337e104, 0x5, 0x0, 0x0, 0xc0001536e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc003b9e600, 0x7f4b4001b108, 0x0, 0x0, 0x0, 0xc00337e104, 0x5, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x1c87c20, 0x1c7db40, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffd23611f21, 0x29, 0x7ffd23611ed8, 0x3b, 0xc000134340, 0xc000153f48, 0xc000153da8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffd23611f21, 0x29, 0x7ffd23611ed8, 0x3b, 0x7ffd23611f1c, 0x4, 0xc0000a1f48, 0x425124, 0xc0000a1eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 2

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_draw:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: interp: store on a pointer bitcast

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*PointerCastValue).Store(0xc00320c520, 0x7f3088024970)
	/go/src/github.com/aykevl/tinygo/interp/values.go:374 +0x39
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00014f648, 0x7f3088937fc0, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc00014f628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:108 +0x18c0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc00382ba40, 0x7f308801afb8, 0x0, 0x0, 0x0, 0xc0039de914, 0x5, 0x0, 0x0, 0xc00014f6e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc00382ba40, 0x7f308801afb8, 0x0, 0x0, 0x0, 0xc0039de914, 0x5, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x171cc20, 0x1712b40, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffdc6f9ef21, 0x29, 0x7ffdc6f9eed8, 0x3b, 0xc0000fc680, 0xc00014ff48, 0xc00014fda8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffdc6f9ef21, 0x29, 0x7ffdc6f9eed8, 0x3b, 0x7ffdc6f9ef1c, 0x4, 0xc0000a1f48, 0x24, 0xc0000a1eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 2

```

## image/gif
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_gif:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: interp: store on a pointer bitcast

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*PointerCastValue).Store(0xc005ab76c0, 0x2c61da0)
	/go/src/github.com/aykevl/tinygo/interp/values.go:374 +0x39
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000101648, 0x3612990, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc000101628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:108 +0x18c0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc004cac4c0, 0x2c8acd8, 0x0, 0x0, 0x0, 0xc002e89844, 0x5, 0x0, 0x0, 0xc0001016e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc004cac4c0, 0x2c8acd8, 0x0, 0x0, 0x0, 0xc002e89844, 0x5, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x2c2dc20, 0x2c23b70, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffd7f72bf21, 0x29, 0x7ffd7f72bed8, 0x3b, 0xc000086b60, 0xc000101f48, 0xc000101da8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffd7f72bf21, 0x29, 0x7ffd7f72bed8, 0x3b, 0x7ffd7f72bf1c, 0x4, 0xc0000a1f48, 0x90, 0xc0000a1eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 2

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_gif:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: interp: store on a pointer bitcast

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*PointerCastValue).Store(0xc004d27cc0, 0x7fd0600327d0)
	/go/src/github.com/aykevl/tinygo/interp/values.go:374 +0x39
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc00012b648, 0x7fd0609e3610, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc00012b628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:108 +0x18c0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc002dd96c0, 0x7fd06005b918, 0x0, 0x0, 0x0, 0xc0066b4104, 0x5, 0x0, 0x0, 0xc00012b6e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc002dd96c0, 0x7fd06005b918, 0x0, 0x0, 0x0, 0xc0066b4104, 0x5, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x16e7c10, 0x16ddb60, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffe97f71f21, 0x29, 0x7ffe97f71ed8, 0x3b, 0xc0001021a0, 0xc00012bf48, 0xc00012bda8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffe97f71f21, 0x29, 0x7ffe97f71ed8, 0x3b, 0x7ffe97f71f1c, 0x4, 0xc0000a1f48, 0x425124, 0xc0000a1eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 2

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
panic: interp: store on a pointer bitcast

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*PointerCastValue).Store(0xc005c3e760, 0x7fc108015770)
	/go/src/github.com/aykevl/tinygo/interp/values.go:374 +0x39
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000145648, 0x1895910, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc000145628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:108 +0x18c0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc00208d500, 0x7fc100033f78, 0x0, 0x0, 0x0, 0xc0012cf450, 0x5, 0x0, 0x0, 0xc0001456e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc00208d500, 0x7fc100033f78, 0x0, 0x0, 0x0, 0xc0012cf450, 0x5, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0xecbd70, 0xec1c90, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7fffb6208f21, 0x29, 0x7fffb6208ed8, 0x3b, 0xc000124340, 0xc000145f48, 0xc000145da8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7fffb6208f21, 0x29, 0x7fffb6208ed8, 0x3b, 0x7fffb6208f1c, 0x4, 0xc000081f48, 0x24, 0xc000081eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 2

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/image_png:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: interp: store on a pointer bitcast

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*PointerCastValue).Store(0xc005b09a00, 0x7fefe4035a90)
	/go/src/github.com/aykevl/tinygo/interp/values.go:374 +0x39
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000149648, 0x7fefe4a33120, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc000149628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:108 +0x18c0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc000d59d40, 0x7fefe4064358, 0x0, 0x0, 0x0, 0xc006f67270, 0x5, 0x0, 0x0, 0xc0001496e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc000d59d40, 0x7fefe4064358, 0x0, 0x0, 0x0, 0xc006f67270, 0x5, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x2d62d70, 0x2d58c90, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffcb2f76f21, 0x29, 0x7ffcb2f76ed8, 0x3b, 0xc000128340, 0xc000149f48, 0xc000149da8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffcb2f76f21, 0x29, 0x7ffcb2f76ed8, 0x3b, 0x7ffcb2f76f1c, 0x4, 0xc000081f48, 0x24, 0xc000081eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
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
error: couldn't load packages due to errors: internal/singleflight, net, context
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
error: couldn't load packages due to errors: context, internal/singleflight, net
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
Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !35952
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !35977
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36027
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36116
error: Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !35952
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !35977
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36027
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36116

Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/math_big:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !35952
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !35977
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36027
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36116
error: Both operands to a binary operator are not of the same type!
  %24 = sub i32 %21, i64 %19, !dbg !35952
Both operands to a binary operator are not of the same type!
  %23 = sub i32 %20, i64 %4, !dbg !35977
Both operands to a binary operator are not of the same type!
  %37 = sub i32 %34, i64 %32, !dbg !36027
Both operands to a binary operator are not of the same type!
  %20 = sub i32 %17, i64 %15, !dbg !36116

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
panic: interp: GEP on a constant

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*LocalValue).GetElementPtr(0xc000f2b1f0, 0xc000f12e48, 0x2, 0x2, 0x2, 0x2)
	/go/src/github.com/aykevl/tinygo/interp/values.go:87 +0x167
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000161058, 0x7f57ec0f9cf0, 0x0, 0x7eb6c2, 0x4, 0x7f57ec0f9d40, 0x7f57ec0fc6d8, 0x7f57ec0ce320, 0x0, 0x66dcec8400000000, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:127 +0x1b33
github.com/aykevl/tinygo/interp.(*Eval).function(0xc000efa6c0, 0x7f57e800d958, 0xc000972160, 0x2, 0x2, 0xc000ad8600, 0x9, 0x7eb6c2, 0x4, 0x851600, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000161648, 0x7f57e803f300, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc000161628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc000efa6c0, 0x7f57e800ae58, 0x0, 0x0, 0x0, 0xc000ad8600, 0x9, 0x0, 0x0, 0xc0001616e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc000efa6c0, 0x7f57e800ae58, 0x0, 0x0, 0x0, 0xc000ad8600, 0x9, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x2498c20, 0x248eb70, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffd90f86f21, 0x29, 0x7ffd90f86ed8, 0x3b, 0xc000144340, 0xc000161f48, 0xc000161da8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffd90f86f21, 0x29, 0x7ffd90f86ed8, 0x3b, 0x7ffd90f86f1c, 0x4, 0xc0000a1f48, 0x24, 0xc0000a1eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 2

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/math_rand:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
panic: interp: GEP on a constant

goroutine 1 [running]:
github.com/aykevl/tinygo/interp.(*LocalValue).GetElementPtr(0xc0009513c0, 0xc0007b49e8, 0x2, 0x2, 0x2, 0x2)
	/go/src/github.com/aykevl/tinygo/interp/values.go:87 +0x167
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000133058, 0x7f0aac14a650, 0x0, 0x7eb6c2, 0x4, 0x7f0aac14a6a0, 0x7f0aac14d038, 0x7f0aac120b40, 0x0, 0x2a2e299900000000, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:127 +0x1b33
github.com/aykevl/tinygo/interp.(*Eval).function(0xc000f26880, 0x7f0aac00d958, 0xc000e911c0, 0x2, 0x2, 0xc000018930, 0x9, 0x7eb6c2, 0x4, 0x851600, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*frame).evalBasicBlock(0xc000133648, 0x7f0aac03f300, 0x0, 0x0, 0x0, 0x44ca7c, 0xc000000300, 0x300000002, 0xc000000300, 0xc000133628, ...)
	/go/src/github.com/aykevl/tinygo/interp/frame.go:363 +0x3af0
github.com/aykevl/tinygo/interp.(*Eval).function(0xc000f26880, 0x7f0aac00ae58, 0x0, 0x0, 0x0, 0xc000018930, 0x9, 0x0, 0x0, 0xc0001336e8, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:93 +0x208
github.com/aykevl/tinygo/interp.(*Eval).Function(0xc000f26880, 0x7f0aac00ae58, 0x0, 0x0, 0x0, 0xc000018930, 0x9, 0x0, 0x0, 0x0, ...)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:76 +0x85
github.com/aykevl/tinygo/interp.Run(0x216ad70, 0x2160c90, 0x0, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/interp/interp.go:65 +0x4d8
main.Compile(0x7ffe51e9bf21, 0x29, 0x7ffe51e9bed8, 0x3b, 0xc000070750, 0xc000133f48, 0xc000133da8, 0x0, 0x0)
	/go/src/github.com/aykevl/tinygo/main.go:67 +0x1df0
main.Build(0x7ffe51e9bf21, 0x29, 0x7ffe51e9bed8, 0x3b, 0x7ffe51e9bf1c, 0x4, 0xc000081f48, 0x90, 0xc000081eb0)
	/go/src/github.com/aykevl/tinygo/main.go:224 +0xe2
main.main()
	/go/src/github.com/aykevl/tinygo/main.go:484 +0xa8d
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
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
error: couldn't load packages due to errors: internal/singleflight, net, mime and 1 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/mime_multipart:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
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
error: couldn't load packages due to errors: mime, net, context and 1 more
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
error: couldn't load packages due to errors: context, net, internal/singleflight
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
error: couldn't load packages due to errors: net/http, net/http/httptrace, internal/singleflight and 5 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
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
error: couldn't load packages due to errors: mime, net, net/http/httptrace and 5 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## net/http/cgi
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_cgi:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
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
error: couldn't load packages due to errors: internal/singleflight, net/http, net and 5 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_cgi:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
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
error: couldn't load packages due to errors: net/http/httptrace, net, mime and 5 more
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
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
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
error: couldn't load packages due to errors: net, vendor/golang_org/x/crypto/cryptobyte, net/http/httptrace and 5 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_cookiejar:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
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
error: couldn't load packages due to errors: context, net, encoding/asn1 and 5 more
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
error: couldn't load packages due to errors: vendor/golang_org/x/crypto/cryptobyte, encoding/asn1, net/http and 5 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_fcgi:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
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
error: couldn't load packages due to errors: net/http/httptrace, internal/singleflight, encoding/asn1 and 5 more
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
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
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
error: couldn't load packages due to errors: net, flag, context and 7 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_httptest:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/flag/flag.go:410:15: New not declared by package reflect
/usr/local/go/src/flag/flag.go:412:15: Zero not declared by package reflect
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
/usr/local/go/src/net/http/httptest/server.go:44:10: WaitGroup not declared by package sync
error: couldn't load packages due to errors: encoding/asn1, flag, net/http/httptrace and 7 more
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
error: couldn't load packages due to errors: net, internal/singleflight, net/http/httptrace and 3 more
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
error: couldn't load packages due to errors: context, net, internal/singleflight and 3 more
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
error: couldn't load packages due to errors: context, net/http, encoding/asn1 and 5 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_httputil:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
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
error: couldn't load packages due to errors: context, encoding/asn1, mime and 5 more
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
/usr/local/go/src/runtime/trace/trace.go:147:10: StopTrace not declared by package runtime
/usr/local/go/src/runtime/trace/trace.go:124:20: StartTrace not declared by package runtime
/usr/local/go/src/runtime/trace/trace.go:129:20: ReadTrace not declared by package runtime
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1176:51: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:749:16: PtrTo not declared by package reflect
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
/usr/local/go/src/encoding/json/encode.go:698:11: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:703:9: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:391:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:395:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:400:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:404:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:364:11: WaitGroup not declared by package sync
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
/usr/local/go/src/encoding/json/decode.go:459:41: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:483:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:485:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
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
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:549:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:14: DeepEqual not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:558:47: Zero not declared by package reflect
/usr/local/go/src/encoding/asn1/marshal.go:447:18: invalid operation: t.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/marshal.go:479:89: invalid operation: t.Field(startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/marshal.go:483:99: invalid operation: t.Field(i + startingField) (value of type reflect.StructField) has no field or method Tag
/usr/local/go/src/html/template/js.go:135:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/asn1/common.go:174:26: invalid operation: t (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:658:80: invalid operation: ifaceType (variable of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/asn1.go:810:115: invalid operation: fieldType (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/asn1/asn1.go:831:12: Copy not declared by package reflect
/usr/local/go/src/encoding/asn1/asn1.go:901:27: invalid operation: structType.Field(i) (value of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/asn1/asn1.go:919:100: invalid operation: field (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/asn1/asn1.go:932:12: Copy not declared by package reflect
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/html/template/content.go:143:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:57: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
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
error: couldn't load packages due to errors: runtime/trace, mime, text/template and 11 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_http_pprof:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/runtime/trace/trace.go:147:10: StopTrace not declared by package runtime
/usr/local/go/src/runtime/trace/trace.go:124:20: StartTrace not declared by package runtime
/usr/local/go/src/runtime/trace/trace.go:129:20: ReadTrace not declared by package runtime
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
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
/usr/local/go/src/text/template/funcs.go:140:19: Zero not declared by package reflect
/usr/local/go/src/text/template/funcs.go:142:18: invalid operation: value.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/funcs.go:145:70: invalid operation: value.Type() (value of type reflect.Type) has no field or method ConvertibleTo
/usr/local/go/src/text/template/funcs.go:146:17: invalid operation: value (variable of type reflect.Value) has no field or method Convert
/usr/local/go/src/encoding/json/decode.go:459:41: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:483:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:485:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:913:18: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:934:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
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
/usr/local/go/src/runtime/pprof/pprof.go:372:135: BlockProfileRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:693:84: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:718:31: StackRecord not declared by package runtime
/usr/local/go/src/runtime/pprof/protomem.go:15:46: MemProfileRecord not declared by package runtime
/usr/local/go/src/text/template/funcs.go:278:14: invalid operation: v (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/runtime/pprof/proto.go:223:20: CallersFrames not declared by package runtime
/usr/local/go/src/encoding/json/decode.go:160:15: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:949:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:949:50: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:950:30: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:950:79: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:159:16: Error not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:920:20: SetMutexProfileFraction not declared by package runtime
/usr/local/go/src/runtime/pprof/pprof.go:807:10: SetCPUProfileRate not declared by package runtime
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
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:584:19: invalid operation: ptr (variable of type reflect.Value) has no field or method MethodByName
/usr/local/go/src/text/template/exec.go:591:33: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/text/template/exec.go:596:22: invalid operation: receiver (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/text/template/exec.go:612:21: invalid operation: nameVal.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:612:50: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/exec.go:622:23: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:529:25: cannot convert nil (untyped nil value) to reflect.Type
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
/usr/local/go/src/text/template/exec.go:431:56: invalid operation: value.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:375:20: invalid operation: val (variable of type reflect.Value) has no field or method Recv
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
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:266:58: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/vendor/golang_org/x/crypto/cryptobyte/asn1.go:273:59: invalid operation: reflect.ValueOf(out).Elem() (value of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/html/template/js.go:135:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:57: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
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
error: couldn't load packages due to errors: runtime/pprof, vendor/golang_org/x/crypto/cryptobyte, runtime/trace and 11 more
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
error: couldn't load packages due to errors: context, net, internal/singleflight and 1 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_mail:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
error: couldn't load packages due to errors: context, net, internal/singleflight and 1 more
Makefile:9: recipe for target 'build-init' failed
make: *** [build-init] Error 1

```

## net/rpc
Without Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_rpc:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
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
/usr/local/go/src/encoding/gob/type.go:496:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/type.go:514:37: invalid operation: t.Elem() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:531:17: invalid operation: typ (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:119:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/encoding/gob/type.go:126:9: invalid operation: rt (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/gob/type.go:142:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:643:70: invalid operation: f (variable of type reflect.StructField) has no field or method Index
/usr/local/go/src/encoding/gob/encode.go:603:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:562:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/encoding/gob/encoder.go:127:29: invalid operation: st (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/gob/encode.go:320:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/encoding/gob/decode.go:1258:17: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:376:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/gob/decode.go:233:18: New not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
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
/usr/local/go/src/encoding/gob/decode.go:1202:17: invalid operation: base (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/decode.go:1118:30: invalid operation: srt (variable of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/encoding/gob/decode.go:1019:31: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:948:22: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:825:35: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/gob/decode.go:656:21: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:625:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/text/template/exec.go:159:16: Error not declared by package runtime
/usr/local/go/src/encoding/gob/decode.go:562:21: MakeMapWithSize not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:564:19: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:568:18: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:568:27: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:569:18: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:569:28: invalid operation: mtyp (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:570:19: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:571:19: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:575:9: invalid operation: value (variable of type reflect.Value) has no field or method SetMapIndex
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
/usr/local/go/src/encoding/gob/decode.go:466:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
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
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
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
error: couldn't load packages due to errors: text/template/parse, encoding/gob, encoding/asn1 and 11 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_rpc:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
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
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/encoding/gob/type.go:496:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/type.go:514:37: invalid operation: t.Elem() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:531:17: invalid operation: typ (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/type.go:119:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/encoding/gob/type.go:126:9: invalid operation: rt (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/gob/type.go:142:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:643:70: invalid operation: f (variable of type reflect.StructField) has no field or method Index
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
/usr/local/go/src/encoding/gob/encode.go:603:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:562:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
/usr/local/go/src/encoding/gob/encoder.go:127:29: invalid operation: st (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/encode.go:320:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/encoding/gob/decode.go:1258:17: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:376:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/gob/decode.go:233:18: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:1202:17: invalid operation: base (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/decode.go:1118:30: invalid operation: srt (variable of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/encoding/gob/decode.go:1019:31: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
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
/usr/local/go/src/encoding/json/encode.go:364:11: WaitGroup not declared by package sync
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
/usr/local/go/src/text/template/exec.go:431:56: invalid operation: value.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/text/template/exec.go:375:20: invalid operation: val (variable of type reflect.Value) has no field or method Recv
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
error: couldn't load packages due to errors: net/rpc, net, encoding/json and 11 more
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
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
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
/usr/local/go/src/encoding/gob/type.go:119:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/encoding/gob/type.go:126:9: invalid operation: rt (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/funcs.go:278:14: invalid operation: v (variable of type reflect.Value) has no field or method Call
/usr/local/go/src/encoding/gob/type.go:142:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:643:70: invalid operation: f (variable of type reflect.StructField) has no field or method Index
/usr/local/go/src/text/template/exec.go:949:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:949:50: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/text/template/exec.go:950:30: PtrTo not declared by package reflect
/usr/local/go/src/text/template/exec.go:950:79: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:603:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:1102:24: invalid operation: sf (variable of type reflect.StructField) has no field or method PkgPath
/usr/local/go/src/encoding/json/encode.go:1103:11: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1118:15: invalid operation: sf (variable of type reflect.StructField) has no field or method Tag
/usr/local/go/src/encoding/json/encode.go:1131:11: invalid operation: ft (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/encode.go:1150:26: invalid operation: sf (variable of type reflect.StructField) has no field or method Anonymous
/usr/local/go/src/encoding/json/encode.go:1176:51: invalid operation: ft (variable of type reflect.Type) has no field or method Name
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
/usr/local/go/src/encoding/json/encode.go:749:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:698:11: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:703:9: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/encode.go:562:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/encode.go:391:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:395:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:400:7: invalid operation: t (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/json/encode.go:404:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/encode.go:364:11: WaitGroup not declared by package sync
/usr/local/go/src/text/template/exec.go:375:20: invalid operation: val (variable of type reflect.Value) has no field or method Recv
/usr/local/go/src/encoding/json/decode.go:459:41: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/encoder.go:127:29: invalid operation: st (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/encode.go:320:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/encoding/gob/decode.go:1258:17: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:376:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/json/decode.go:483:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:485:15: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/gob/decode.go:233:18: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:913:18: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:934:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:963:6: invalid operation: v (variable of type reflect.Value) has no field or method SetBytes
/usr/local/go/src/encoding/json/decode.go:967:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1001:9: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/json/decode.go:1009:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowInt
/usr/local/go/src/encoding/json/decode.go:1017:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowUint
/usr/local/go/src/encoding/json/decode.go:1025:23: invalid operation: v (variable of type reflect.Value) has no field or method OverflowFloat
/usr/local/go/src/encoding/json/decode.go:634:40: invalid operation: v (variable of type reflect.Value) has no field or method NumMethod
/usr/local/go/src/encoding/gob/decode.go:1202:17: invalid operation: base (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/decode.go:1118:30: invalid operation: srt (variable of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/encoding/gob/decode.go:1019:31: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:948:22: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:652:12: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:657:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:657:24: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:664:18: MakeMap not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:704:23: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:706:25: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:825:35: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/json/decode.go:741:25: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:748:38: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
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
/usr/local/go/src/encoding/asn1/marshal.go:537:47: invalid operation: v.Type() (value of type reflect.Type) has no field or method NumMethod
/usr/local/go/src/encoding/asn1/marshal.go:546:27: New not declared by package reflect
/usr/local/go/src/html/template/content.go:143:16: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
/usr/local/go/src/html/template/content.go:143:57: invalid operation: v.Type() (value of type reflect.Type) has no field or method Implements
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
error: couldn't load packages due to errors: net, encoding/asn1, net/rpc and 11 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/net_rpc_jsonrpc:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/internal/singleflight/singleflight.go:13:10: WaitGroup not declared by package sync
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/encoding/json/encode.go:345:23: Map not declared by package sync
/usr/local/go/src/encoding/json/encode.go:1249:21: Map not declared by package sync
/usr/local/go/src/text/template/parse/parse.go:196:26: Error not declared by package runtime
/usr/local/go/src/net/net_fake.go:178:17: Cond not declared by package sync
/usr/local/go/src/net/net_fake.go:179:17: Cond not declared by package sync
/usr/local/go/src/net/lookup.go:59:23: WaitGroup not declared by package sync
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
/usr/local/go/src/encoding/gob/type.go:119:12: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/encoding/gob/type.go:126:9: invalid operation: rt (variable of type reflect.Type) has no field or method Implements
/usr/local/go/src/encoding/gob/type.go:142:14: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:643:70: invalid operation: f (variable of type reflect.StructField) has no field or method Index
/usr/local/go/src/encoding/gob/encode.go:603:16: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/encode.go:562:34: invalid operation: t (variable of type reflect.Type) has no field or method Key
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
/usr/local/go/src/encoding/gob/encoder.go:127:29: invalid operation: st (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/encode.go:320:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/encoding/gob/decode.go:1258:17: New not declared by package reflect
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
/usr/local/go/src/encoding/gob/decode.go:376:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/encoding/gob/decode.go:233:18: New not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:1202:17: invalid operation: base (variable of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/gob/decode.go:1118:30: invalid operation: srt (variable of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/encoding/gob/decode.go:1019:31: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:948:22: PtrTo not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:825:35: invalid operation: t (variable of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:656:21: Zero not declared by package reflect
/usr/local/go/src/encoding/gob/decode.go:625:11: invalid operation: value (variable of type reflect.Value) has no field or method Cap
/usr/local/go/src/text/template/exec.go:591:33: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method FieldByName
/usr/local/go/src/text/template/exec.go:596:22: invalid operation: receiver (variable of type reflect.Value) has no field or method FieldByIndex
/usr/local/go/src/text/template/exec.go:612:21: invalid operation: nameVal.Type() (value of type reflect.Type) has no field or method AssignableTo
/usr/local/go/src/text/template/exec.go:612:50: invalid operation: receiver.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/text/template/exec.go:622:23: Zero not declared by package reflect
/usr/local/go/src/text/template/exec.go:529:25: cannot convert nil (untyped nil value) to reflect.Type
/usr/local/go/src/text/template/exec.go:431:56: invalid operation: value.Type() (value of type reflect.Type) has no field or method NumMethod
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
/usr/local/go/src/encoding/json/decode.go:704:23: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:706:25: Zero not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:741:25: New not declared by package reflect
/usr/local/go/src/encoding/json/decode.go:748:38: invalid operation: v.Type() (value of type reflect.Type) has no field or method Name
/usr/local/go/src/encoding/json/decode.go:789:19: invalid operation: v.Type() (value of type reflect.Type) has no field or method Key
/usr/local/go/src/encoding/gob/decode.go:466:18: invalid operation: value (variable of type reflect.Value) has no field or method FieldByIndex
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
/usr/local/go/src/mime/type.go:15:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:16:22: Map not declared by package sync
/usr/local/go/src/mime/type.go:21:20: Map not declared by package sync
/usr/local/go/src/mime/type.go:24:27: Map not declared by package sync
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
error: couldn't load packages due to errors: text/template, net, text/template/parse and 11 more
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
error: couldn't load packages due to errors: net, vendor/golang_org/x/crypto/cryptobyte, encoding/asn1 and 2 more
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
error: couldn't load packages due to errors: encoding/asn1, vendor/golang_org/x/crypto/cryptobyte, internal/singleflight and 2 more
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
error: couldn't load packages due to errors: net, context, internal/singleflight
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
error: couldn't load packages due to errors: internal/singleflight, context, net
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
error: couldn't load packages due to errors: runtime/pprof, context
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
error: couldn't load packages due to errors: runtime/pprof, context
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
error: couldn't load packages due to errors: context, runtime/trace
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
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/flag/flag.go:410:15: New not declared by package reflect
/usr/local/go/src/flag/flag.go:412:15: Zero not declared by package reflect
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
error: couldn't load packages due to errors: context, runtime/debug, runtime/trace and 2 more
Makefile:5: recipe for target 'build' failed
make: *** [build] Error 1

```

With Init Interp

```
docker run --rm -v /home/stephen/go/src/github.com/trashhalo/tinygo-import-report/tests/testing:/go/src/github.com/trashhalo/tinygo-import-report tinygo/tinygo \
build -initinterp -o /go/src/github.com/trashhalo/tinygo-import-report/wasm.wasm -target wasm github.com/trashhalo/tinygo-import-report
/usr/local/go/src/runtime/debug/stack.go:24:16: Stack not declared by package runtime
/usr/local/go/src/runtime/debug/garbage.go:37:34: MemStats not declared by package runtime
/usr/local/go/src/context/context.go:472:26: invalid operation: reflect.TypeOf(key) (value of type reflect.Type) has no field or method Comparable
/usr/local/go/src/flag/flag.go:410:15: New not declared by package reflect
/usr/local/go/src/flag/flag.go:412:15: Zero not declared by package reflect
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
error: couldn't load packages due to errors: runtime/debug, context, flag and 2 more
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
error: couldn't load packages due to errors: text/template/parse, text/template
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

