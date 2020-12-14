package env

import (
	"fmt"

	"venus/env/global"
	"venus/service/venus"
)

func services() error {
	inits := []func() error{
		func() error {
			venusService, err := venus.NewVenusService(global.Config.VenusServiceHostAddr, defaultTimeout(global.Config.DefaultRPCTimeout), defaultKeepAliveSec(global.Config.DefaultRPCKeepAliveSec))
			if err != nil {
				return err
			}
			global.VenusService = venusService
			return nil
		},
	}

	var (
		errs = make(chan error)
		err  error
	)

	for _, initFunc := range inits {
		go func(f func() error, e chan error) {
			e <- f()
		}(initFunc, errs)
	}

	for i := 0; i < len(inits); i++ {
		if ferr := <-errs; ferr != nil {
			if err != nil {
				err = fmt.Errorf("%s %s", err.Error(), ferr.Error())
			} else {
				err = ferr
			}
		}
	}

	return err
}

func defaultTimeout(sec int) int {
	if sec > 0 {
		return sec
	}

	return global.Config.DefaultRPCTimeout
}

func defaultKeepAliveSec(sec int) int {
	if sec > 0 {
		return sec
	}

	return global.Config.DefaultRPCKeepAliveSec
}
