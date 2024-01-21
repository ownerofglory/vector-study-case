package service

import "github.com/ownerofglory/vector-study-case/internal/logging"

type numService struct{}

func NewNumService() *numService {
	return &numService{}
}

func (n *numService) FindPairForTarget(a, b []int, target int) bool {
	logging.Info("Executing FindPairForTarget")
	logging.Debugf("a=%v, b=%v, target=%v", a, b, target)

	for _, outer := range a {
		for _, inner := range b {
			if outer+inner == target {
				logging.Info("Found pair for target number")
				logging.Debugf("Target %v: %v + %v", target, outer, inner)
				return true
			}
		}
	}

	logging.Info("No pair for target number found")
	return false
}
