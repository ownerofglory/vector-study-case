package port

//go:generate mockgen -package=port -source=./num_port.go -destination=./num_port_mock.go
type NumService interface {
	FindPairForTarget(a, b []int, target int) bool
}
