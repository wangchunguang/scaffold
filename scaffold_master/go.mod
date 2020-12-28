module master/scaffold_master

go 1.15

require (
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.3 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/prometheus/common v0.15.0
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.7.0
	github.com/tebeka/strftime v0.1.5 // indirect
	github.com/x-cray/logrus-prefixed-formatter v0.5.2
	google.golang.org/grpc v1.34.0
	go.etcd.io/etcd v3.3.25+incompatible
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
