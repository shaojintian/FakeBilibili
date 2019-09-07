module FakeBilibili/src

go 1.12

require (
	github.com/go-chassis/go-chassis v1.6.1
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jinzhu/gorm v1.9.10
	github.com/kr/pretty v0.1.0 // indirect
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/ratelimit v0.1.0 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)

replace (
	cloud.google.com/go v0.37.4 => github.com/googleapis/google-cloud-go v0.37.4
	go.uber.org/ratelimit v0.0.0-20180316092928-c15da0234277 => github.com/uber-go/ratelimit v0.0.0-20180316092928-c15da0234277
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c => github.com/golang/crypto v0.0.0-20190325154230-a5d413f7728c
	golang.org/x/sys v0.0.0-20181122145206-62eef0e2fa9b => github.com/golang/sys v0.0.0-20181122145206-62eef0e2fa9b
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0

)
