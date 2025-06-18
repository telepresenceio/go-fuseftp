module github.com/telepresenceio/go-fuseftp

go 1.24

toolchain go1.24.4

require (
	github.com/datawire/go-ftpserver v0.1.3
	github.com/jlaffaye/ftp v0.2.0
	github.com/puzpuzpuz/xsync/v3 v3.5.1
	github.com/puzpuzpuz/xsync/v4 v4.1.0
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.8.3
	github.com/telepresenceio/go-fuseftp/rpc v0.6.6
	github.com/winfsp/cgofuse v1.6.0
	golang.org/x/sys v0.33.0
	google.golang.org/grpc v1.73.0
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/datawire/dlib v1.3.1-0.20220715022530-b09ab2e017e1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fclairamb/ftpserverlib v0.21.0 // indirect
	github.com/fclairamb/go-log v0.4.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250603155806-513f23925822 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/telepresenceio/go-fuseftp/rpc => ./rpc
