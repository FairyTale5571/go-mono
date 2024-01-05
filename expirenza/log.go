package expirenza

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func msglog(caller int, format string, a ...interface{}) {

	pc, file, line, _ := runtime.Caller(caller)

	files := strings.Split(file, "/")
	file = files[len(files)-1]

	name := runtime.FuncForPC(pc).Name()
	fns := strings.Split(name, ".")
	name = fns[len(fns)-1]

	msg := fmt.Sprintf(format, a...)

	log.Printf("[EXPIRENZA] %s:%d:%s() %s\n", file, line, name, msg)
}

func (e *Expirenza) log(format string, a ...interface{}) {
	if !e.debug {
		return
	}
	msglog(2, format, a...)
}
