package plugins

import(
    "strconv"
    "testing"
)

func CheckData(data map[string]*MetricValue, key string, val string) bool {
    metVal := (*data[key]).Val
    return metVal == val
}

func TestCollectData_ValidInput(t *testing.T){
    testVals := map[string]string{
        "hw.physicalcpu":"4",
        "hw.logicalcpu":"8",
        "hw.cpufrequency":"2300000000",
        "machdep.cpu.vendor":"GenuineIntel\n",
        "machdep.cpu.brand_string":"Intel(R) Core(TM) i7-3615QM CPU @ 2.30GHz\n",
        "machdep.cpu.model":"58",
        "machdep.cpu.family":"6",
        "machdep.cpu.stepping":"9",
        "machdep.cpu.features":"FPU VME DE PSE TSC MSR PAE MCE CX8 APIC SEP MTRR PGE MCA CMOV PAT PSE36 CLFSH DS ACPI MMX FXSR SSE SSE2 SS HTT TM PBE SSE3 PCLMULQDQ DTES64 MON DSCPL VMX EST TM2 SSSE3 CX16 TPR PDCM SSE4.1 SSE4.2 x2APIC POPCNT AES PCID XSAVE OSXSAVE TSCTMR AVX1.0 RDRAND F16C",
    }

	identifiers := map[string]string{
		"real":       "hw.physicalcpu",
		"total":      "hw.logicalcpu",
		"mhz":        "hw.cpufrequency",
        "vendor_id": "machdep.cpu.vendor",
		"model_name":  "machdep.cpu.brand_string",
		"model": "machdep.cpu.model",
		"family":      "machdep.cpu.family",
		"stepping":     "machdep.cpu.stepping",
		"flags":   "machdep.cpu.features",
	}

    SetLookup := func (testVals map[string]string) (func(string) (string, error)){
        return func(metric string) (string, error){
            return testVals[metric], nil
        }
    }
    GetCpuMetric = SetLookup(testVals)

    testCpu := &Cpu{"test"}
    _, data, err := testCpu.CollectData()

    if err != nil {
        t.Fail()
    }

	mhz, err := strconv.Atoi(testVals["hw.cpufrequency"])
	testVals["hw.cpufrequency"] = strconv.Itoa(mhz / 1000000)


    for metric, sysMetric := range identifiers {
        if !CheckData(data, metric, testVals[sysMetric]) {
            t.Fail()
        }
    }

}
