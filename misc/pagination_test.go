package misc

import (
	"fmt"
	"testing"
)

func TestPagination_CountSql(t *testing.T) {
	p := NewPagination(10,20)
	fmt.Println(p.CountSql("Select user,name,*  From   users t1 join (select * from payment) t2 on t1.uid = t2.uid"))
}
