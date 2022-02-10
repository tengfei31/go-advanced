package assembly

//go:nosplit
func AsmCallCAdd(cfun uintptr, a, b int64) int64

