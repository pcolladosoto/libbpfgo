package libbpfgo

/*
#include "libbpfgo.h"
*/
import "C"

const (
	// Maximum number of channels (RingBuffers + PerfBuffers) supported
	maxEventChannels = 512
)

var (
	eventChannels = newRWArray(maxEventChannels)
)
