package casbin_casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"testing"
)

func TestCasbin(t *testing.T) {

	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	fmt.Println(e)
	fmt.Println(err)

	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	//act := "read"  // the operation that the user performs on the resource.
	act := "read1" // the operation that the user performs on the resource.

	if res, _ := e.Enforce(sub, obj, act); res {
		// permit alice to read data1
		println("allow")
	} else {
		println("deny")
		// deny the request, show an error
	}

}
