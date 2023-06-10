package hook

/*

// #include "event/hook_async.h"
*/
import "C"

import (
	"log"
	"time"

	"encoding/json"
)

//export go_send
func go_send(s *C.char) {
	str := []byte(C.GoString(s))
	out := Event{}

	err := json.Unmarshal(str, &out)
	if err != nil {
		log.Fatal("json.Unmarshal error is: ", err)
	}
	out.When = time.UnixMicro(int64(out.Time))
	//fmt.Println(fmt.Sprintf("now: %+v, out: %+v", time.Now(), out))
	if out.Keychar != CharUndefined {
		lck.Lock()
		raw2key[out.Rawcode] = string([]rune{out.Keychar})
		lck.Unlock()
	}

	// todo: maybe make non-bloking
	ev <- out
}
