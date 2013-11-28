package pricing

type InstanceTypeConfig struct {
	Name               string
	Arch               string
	Cpus               int
	ECUs               float64
	Memory             float64
	Storage            string
	EbsOptimizable     bool
	NetworkPerformance string
}

var InstanceTypeConfigs = []*InstanceTypeConfig{
	{Name: "m3.xlarge", Arch: "64-bit", Cpus: 4, ECUs: 13, Memory: 15, Storage: "EBS only", EbsOptimizable: true, NetworkPerformance: "Moderate"},
	{Name: "m3.2xlarge", Arch: "64-bit", Cpus: 8, ECUs: 26, Memory: 30, Storage: "EBS only", EbsOptimizable: true, NetworkPerformance: "High"},
	{Name: "m1.small", Arch: "32-bit/64-bit", Cpus: 1, ECUs: 1, Memory: 1.7, Storage: "1 x 160", NetworkPerformance: "Low"},
	{Name: "m1.medium", Arch: "32-bit/64-bit", Cpus: 1, ECUs: 2, Memory: 3.75, Storage: "1 x 410", NetworkPerformance: "Moderate"},
	{Name: "m1.large", Arch: "64-bit", Cpus: 2, ECUs: 4, Memory: 7.5, Storage: "2 x 420", EbsOptimizable: true, NetworkPerformance: "Moderate"},
	{Name: "m1.xlarge", Arch: "64-bit", Cpus: 4, ECUs: 8, Memory: 15, Storage: "4 x 420", EbsOptimizable: true, NetworkPerformance: "High"},
	{Name: "c3.large", Arch: "64-bit", Cpus: 2, ECUs: 7, Memory: 3.75, Storage: "2 x 16 SSD", NetworkPerformance: "Moderate"},
	{Name: "c3.xlarge", Arch: "64-bit", Cpus: 4, ECUs: 14, Memory: 7, Storage: "2 x 40 SSD", EbsOptimizable: true, NetworkPerformance: "High"},
	{Name: "c3.2xlarge", Arch: "64-bit", Cpus: 8, ECUs: 28, Memory: 15, Storage: "2 x 80 SSD", EbsOptimizable: true, NetworkPerformance: "High"},
	{Name: "c3.4xlarge", Arch: "64-bit", Cpus: 16, ECUs: 55, Memory: 30, Storage: "2 x 160 SSD", EbsOptimizable: true, NetworkPerformance: "High"},
	{Name: "c3.8xlarge", Arch: "64-bit", Cpus: 32, ECUs: 108, Memory: 60, Storage: "2 x 320 SSD", NetworkPerformance: "High"},
	{Name: "c1.medium", Arch: "32-bit/64-bit", Cpus: 2, ECUs: 5, Memory: 1.7, Storage: "1 x 350", NetworkPerformance: "Moderate"},
	{Name: "c1.xlarge", Arch: "64-bit", Cpus: 8, ECUs: 20, Memory: 7, Storage: "4 x 420", EbsOptimizable: true, NetworkPerformance: "High"},
	{Name: "cc2.8xlarge", Arch: "64-bit", Cpus: 32, ECUs: 88, Memory: 60.5, Storage: "4 x 840", NetworkPerformance: "10 Gigabit"},
	{Name: "g2.2xlarge", Arch: "64-bit", Cpus: 8, ECUs: 26, Memory: 15, Storage: "1 x 60 SSD", EbsOptimizable: true, NetworkPerformance: "High"},
	{Name: "cg1.4xlarge", Arch: "64-bit", Cpus: 16, ECUs: 33.5, Memory: 22.5, Storage: "2 x 840", NetworkPerformance: "10 Gigabit"},
	{Name: "m2.xlarge", Arch: "64-bit", Cpus: 2, ECUs: 6.5, Memory: 17.1, Storage: "1 x 420", NetworkPerformance: "Moderate"},
	{Name: "m2.2xlarge", Arch: "64-bit", Cpus: 4, ECUs: 13, Memory: 34.2, Storage: "1 x 850", EbsOptimizable: true, NetworkPerformance: "Moderate"},
	{Name: "m2.4xlarge", Arch: "64-bit", Cpus: 8, ECUs: 26, Memory: 68.4, Storage: "2 x 840", EbsOptimizable: true, NetworkPerformance: "High"},
	{Name: "cr1.8xlarge", Arch: "64-bit", Cpus: 32, ECUs: 88, Memory: 244, Storage: "2 x 120 SSD", NetworkPerformance: "10 Gigabit"},
	{Name: "hi1.4xlarge", Arch: "64-bit", Cpus: 16, ECUs: 35, Memory: 60.5, Storage: "2 x 1,024 SSD", NetworkPerformance: "10 Gigabit"},
	{Name: "hs1.8xlarge", Arch: "64-bit", Cpus: 16, ECUs: 35, Memory: 117, Storage: "24 x 2,048", NetworkPerformance: "10 Gigabit"},
	{Name: "t1.micro", Arch: "32-bit/64-bit", Cpus: 1, ECUs: 1, Memory: 0.615, NetworkPerformance: "Very Low"},
}
