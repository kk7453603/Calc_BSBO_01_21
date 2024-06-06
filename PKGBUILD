pkgname=calculator
pkgver=1.0
pkgrel=1
pkgdesc="A simple calculator written in Go"
arch=('x86_64')
url="https://github.com/kk7453603/Calc_BSBO_01_21"
license=('MIT')
depends=('go')

# prepare() {
    # echo 'penis 2'
    # tar -czvf "$pkgname-$pkgver.tar.gz" ./
    # source=("file://./$pkgname-$pkgver.tar.gz")
# }


build() {
    go build ../cmd/main.go
}

package() {
    install -Dm755 "./main" "../output/main"
}