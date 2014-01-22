package pricing

type InstanceTypeConfig struct {
	Family             string
	Name               string
	Arch               string
	Cpus               int
	ECUs               float64
	ECUText            string
	Memory             float64
	Storage            string
	EbsOptimizable     bool
	NetworkPerformance string

	Turbo             bool
	AVX               bool
	AES               bool
	PhysicalProcessor string
}
