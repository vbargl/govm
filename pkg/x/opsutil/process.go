package opsutil

import "barglvojtech.net/govm/pkg/embi/types"

func Process(ops ...types.Op) error {
	for _, op := range ops {
		if err := op.Process(); err != nil {
			return err
		}
	}
	return nil
}
