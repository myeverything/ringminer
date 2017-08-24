
// TODO(fukun): add header
In loopring, we chose govendor to manager libs.

install govendor:
go get -u -v github.com/kardianos/govendor

initial project:
govendor init

add external libs:
govendor add +external
