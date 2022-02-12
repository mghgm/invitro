package trace

import (
	"fmt"
	"math"

	util "github.com/eth-easl/loader/pkg"
)

//* A bit of a heck to get around cyclic import.
type FunctionSpecsGen func(Function) (int, int)

type FunctionConcurrencyStats struct {
	Average float64
	Count   float64
	Median  float64
	Minimum float64
	Maximum float64
	data    []float64
}

type FunctionInvocationStats struct {
	Average int
	Count   int
	Median  int
	Minimum int
	Maximum int
	data    []int
}
type FunctionRuntimeStats struct {
	Average       int
	Count         int
	Minimum       int
	Maximum       int
	Percentile0   int
	Percentile1   int
	Percentile25  int
	Percentile50  int
	Percentile75  int
	Percentile99  int
	Percentile100 int
}

type FunctionMemoryStats struct {
	Average       int
	Count         int
	Percentile1   int
	Percentile5   int
	Percentile25  int
	Percentile50  int
	Percentile75  int
	Percentile95  int
	Percentile99  int
	Percentile100 int
}

type Function struct {
	Mame            string
	Url             string
	AppHash         string
	Hash            string
	Deployed        bool
	ConcurrencySats FunctionConcurrencyStats
	InvocationStats FunctionInvocationStats
	RuntimeStats    FunctionRuntimeStats
	MemoryStats     FunctionMemoryStats
}

type FunctionTraces struct {
	Path                      string
	Functions                 []Function
	WarmupScales              []int
	InvocationsEachMinute     [][]int
	TotalInvocationsPerMinute []int
}

func (f *Function) SetHash(hash int) {
	f.Hash = fmt.Sprintf("%015d", hash)
}

func (f *Function) SetName(name string) {
	f.Mame = name
}

func (f *Function) SetStatus(b bool) {
	f.Deployed = b
}

func (f *Function) GetStatus() bool {
	return f.Deployed
}

func (f *Function) GetName() string {
	return f.Mame
}

func (f *Function) GetUrl() string {
	return f.Url
}

func (f *Function) SetUrl(url string) {
	f.Url = url
}

func (f *Function) GetExpectedConcurrency() int {
	expectedRps := f.InvocationStats.Median / 60
	expectedFinishingRatePerSec := float64(f.RuntimeStats.Percentile100) / 1000
	expectedConcurrency := float64(expectedRps) * expectedFinishingRatePerSec

	// log.Info(expectedRps, expectedFinishingRatePerSec, expectedConcurrency)

	return util.MaxOf(
		MIN_CONCURRENCY,
		util.MinOf(
			MAX_CONCURRENCY,
			int(math.Ceil(expectedConcurrency)),
		),
	)
}
