package rbac

import (
	"venus/service/casbin"
)

type Rbac struct {
	Sub  string
	Obj  string
	Act  string
	Auth *casbin.CasbinAuth
}

func NewRbac(sub, obj, act string) (*Rbac, error) {

	casbinAuth, err := casbin.NewCasbinAuth()
	if err != nil {
		return nil, err
	}
	success, err := casbinAuth.AddCasbinPolicies([][]string{{sub, obj, act}})
	if err != nil {
		return nil, err
	}

	if !success {
		return nil, nil
	}

	err = casbinAuth.Save()
	if err != nil {
		return nil, err
	}

	return &Rbac{
		Auth: casbinAuth,
		Sub:  sub,
		Obj:  obj,
		Act:  act,
	}, nil
}

func (r *Rbac) Close() error {
	_, err := r.Auth.DeleteCasbinPolicies([][]string{{r.Sub, r.Obj, r.Act}})
	return err
}
