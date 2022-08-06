
Pdf library I made for personal use

![](examples/example.png)

Requirements<br>
https://go.dev/<br>
https://github.com/libvips/libvips<br>

Cloning<br>
This project uses pdf.js and bulma css
```
git clone --recurse-submodules https://github.com/lenivyyluitel/emperor
cd emperor
make build
make
```
Usage<br>
```go run cmd/*```
or
```./emperor 8080 false```

Flags<br>
-log (disable or enable logging)<br>
-port