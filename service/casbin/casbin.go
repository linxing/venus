package casbin

import (
	"path"
	"runtime"

	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"

	"venus/model"
	"venus/service/casbin/xormadapter"
)

type CasbinAuth struct {
	enforcer *casbin.Enforcer
}

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)

	return path.Dir(filename)
}

func (ca *CasbinAuth) GetEnforcer() *casbin.Enforcer {
	return ca.enforcer
}

func NewCasbinAuth() (*CasbinAuth, error) {

	a, err := xormadapter.NewAdapterByEngine(model.GetEngine())
	if err != nil {
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer(getCurrentPath()+"/rbac.conf", a)
	if err != nil {
		return nil, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}

	return &CasbinAuth{
		enforcer: enforcer,
	}, nil
}

func (ca *CasbinAuth) Save() error {
	return ca.enforcer.SavePolicy()
}

func (ca *CasbinAuth) AddCasbinPolicies(rules [][]string) (bool, error) {

	success, err := ca.enforcer.AddPolicies(rules)
	if err != nil {
		return false, err
	}

	return success, nil
}

func (ca *CasbinAuth) DeleteCasbinPolicies(rules [][]string) (bool, error) {

	success, err := ca.enforcer.RemovePolicies(rules)
	if err != nil {
		return false, err
	}

	return success, nil
}

func (ca *CasbinAuth) CheckPermission(sub, obj, act string) (bool, error) {

	isAllow, err := ca.enforcer.Enforce(sub, obj, act)
	if err != nil {
		return false, err
	}

	return isAllow, nil
}
