module github.com/telepresenceio/go-fuseftp

go 1.25

require (
	github.com/jlaffaye/ftp v0.2.0
	github.com/puzpuzpuz/xsync/v4 v4.2.0
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.11.1
	github.com/telepresenceio/go-ftpserver v0.0.0-20251115215116-e3eac8d99c06
	github.com/telepresenceio/go-fuseftp/rpc v0.6.7
	github.com/winfsp/cgofuse v1.6.0
	golang.org/x/sys v0.38.0
	google.golang.org/grpc v1.76.0
	google.golang.org/protobuf v1.36.10
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fclairamb/ftpserverlib v0.27.0 // indirect
	github.com/fclairamb/go-log v0.6.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	github.com/telepresenceio/dlib/v2 v2.0.1 // indirect
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251111163417-95abcf5c77ba // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/telepresenceio/go-fuseftp/rpc => ./rpc
