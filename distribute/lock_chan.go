package distribute

//chan锁

type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	var lock Lock
	lock.c = make(chan struct{}, 1)
	lock.c <- struct{}{}
	return lock
}

func (lock Lock) Lock() bool {
	var result bool = false
	select {
	case <- lock.c:
		result = true
	default:
	}
	return result
}

func (lock Lock) UnLock() {
	lock.c <- struct{}{}
}
