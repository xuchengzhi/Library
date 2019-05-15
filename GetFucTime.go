package GetFuncTime

import "time"
import "fmt"
import "runtime"
import "reflect"

//获取方法运行时间
func TC(f func()) time.Duration {
	time1 := time.Now()
	fname := fmt.Sprintf("func name is %s", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
	fmt.Println(fname)
	f()
	time2 := time.Now()
	return time2.Sub(time1)
}
