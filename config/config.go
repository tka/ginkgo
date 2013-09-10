package config

import (
	"flag"
	"time"

	"fmt"
)

type GinkgoConfigType struct {
	RandomSeed        int64
	RandomizeAllSpecs bool
	FocusString       string
	ParallelNode      int
	ParallelTotal     int
	SkipBenchmarks    bool
}

var GinkgoConfig = GinkgoConfigType{}

type DefaultReporterConfigType struct {
	NoColor           bool
	SlowSpecThreshold float64
	NoisyPendings     bool
}

var DefaultReporterConfig = DefaultReporterConfigType{}

func processPrefix(prefix string) string {
	if prefix != "" {
		prefix = prefix + "."
	}
	return prefix
}

func Flags(prefix string, includeParallelFlags bool) {
	prefix = processPrefix(prefix)
	flag.Int64Var(&(GinkgoConfig.RandomSeed), prefix+"seed", time.Now().Unix(), "The seed used to randomize the spec suite.")
	flag.BoolVar(&(GinkgoConfig.RandomizeAllSpecs), prefix+"randomizeAllSpecs", false, "If set, ginkgo will randomize all specs together.  By default, ginkgo only randomizes the top level Describe/Context groups.")
	flag.BoolVar(&(GinkgoConfig.SkipBenchmarks), prefix+"skipBenchmarks", false, "If set, ginkgo will skip any benchmark specs.")
	flag.StringVar(&(GinkgoConfig.FocusString), prefix+"focus", "", "If set, ginkgo will only run specs that match this regular expression.")

	if includeParallelFlags {
		flag.IntVar(&(GinkgoConfig.ParallelNode), prefix+"parallel.node", 1, "This worker node's (one-indexed) node number.  For running specs in parallel.")
		flag.IntVar(&(GinkgoConfig.ParallelTotal), prefix+"parallel.total", 1, "The total number of worker nodes.  For running specs in parallel.")
	}

	flag.BoolVar(&(DefaultReporterConfig.NoColor), prefix+"noColor", false, "If set, suppress color output in default reporter.")
	flag.Float64Var(&(DefaultReporterConfig.SlowSpecThreshold), prefix+"slowSpecThreshold", 5.0, "(in seconds) Specs that take longer to run than this threshold are flagged as slow by the default reporter (default: 5 seconds).")
	flag.BoolVar(&(DefaultReporterConfig.NoisyPendings), prefix+"noisyPendings", true, "If set, default reporter will shout about pending tests.")
}

func BuildFlagArgs(prefix string, ginkgo GinkgoConfigType, reporter DefaultReporterConfigType) []string {
	prefix = processPrefix(prefix)
	result := make([]string, 0)

	if ginkgo.RandomSeed > 0 {
		result = append(result, fmt.Sprintf("--%sseed=%d", prefix, ginkgo.RandomSeed))
	}

	if ginkgo.RandomizeAllSpecs {
		result = append(result, fmt.Sprintf("--%srandomizeAllSpecs", prefix))
	}

	if ginkgo.SkipBenchmarks {
		result = append(result, fmt.Sprintf("--%sskipBenchmarks", prefix))
	}

	if ginkgo.FocusString != "" {
		result = append(result, fmt.Sprintf("--%sfocus=%s", prefix, ginkgo.FocusString))
	}

	if ginkgo.ParallelNode != 0 {
		result = append(result, fmt.Sprintf("--%sparallel.node=%d", prefix, ginkgo.ParallelNode))
	}

	if ginkgo.ParallelTotal != 0 {
		result = append(result, fmt.Sprintf("--%sparallel.total=%d", prefix, ginkgo.ParallelTotal))
	}

	if reporter.NoColor {
		result = append(result, fmt.Sprintf("--%snoColor", prefix))
	}

	if reporter.SlowSpecThreshold > 0 {
		result = append(result, fmt.Sprintf("--%sslowSpecThreshold=%.5f", prefix, reporter.SlowSpecThreshold))
	}

	if !reporter.NoisyPendings {
		result = append(result, fmt.Sprintf("--%snoisyPendings=false"))
	}

	return result
}