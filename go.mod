module github.com/kimpettersen/svc-voicepayments

go 1.12

require (
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/kimpettersen/svc-payments v0.0.0-20190304120234-76d7dffed7e9
	google.golang.org/grpc v1.19.0
)

// This is a hack to use a local version of svc-payments
replace github.com/kimpettersen/svc-payments => ../svc-payments
