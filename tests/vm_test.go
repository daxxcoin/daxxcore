// Copyright 2014 The daxxcoreAuthors
// This file is part of the daxxcore library.
//
// The daxxcore library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The daxxcore library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the daxxcore library. If not, see <http://www.gnu.org/licenses/>.

package tests

import (
	"os"
	"path/filepath"
	"testing"
)

func BenchmarkVmAckermann32Tests(b *testing.B) {
	fn := filepath.Join(vmTestDir, "vmPerformanceTest.json")
	if err := BenchVmTest(fn, bconf{"ackermann32", os.Getenv("JITFORCE") == "true", os.Getenv("JITVM") == "true"}, b); err != nil {
		b.Error(err)
	}
}

func BenchmarkVmFibonacci16Tests(b *testing.B) {
	fn := filepath.Join(vmTestDir, "vmPerformanceTest.json")
	if err := BenchVmTest(fn, bconf{"fibonacci16", os.Getenv("JITFORCE") == "true", os.Getenv("JITVM") == "true"}, b); err != nil {
		b.Error(err)
	}
}

// I've created a new function for each tests so it's easier to identify where the problem lies if any of them fail.
func TestVmVMArithmetic(t *testing.T) {
	fn := filepath.Join(vmTestDir, "vmArithmeticTest.json")
	if err := RunVmTest(fn, VmSkipTests); err != nil {
		t.Error(err)
	}
}

func TestVmBitwiseLogicOperation(t *testing.T) {
	fn := filepath.Join(vmTestDir, "vmBitwiseLogicOperationTest.json")
	if err := RunVmTest(fn, VmSkipTests); err != nil {
		t.Error(err)
	}
}

func TestVmBlockInfo(t *testing.T) {
	fn := filepath.Join(vmTestDir, "vmBlockInfoTest.json")
	if err := RunVmTest(fn, VmSkipTests); err != nil {
		t.Error(err)
	}
}

func TestVmEnvironmentalInfo(t *testing.T) {
	fn := filepath.Join(vmTestDir, "vmEnvironmentalInfoTest.json")
	if err := RunVmTest(fn, VmSkipTests); err != nil {
		t.Error(err)
	}
}

func TestVmFlowOperation(t *testing.T) {
	fn := filepath.Join(vmTestDir, "vmIOandFlowOperationsTest.json")
	if err := RunVmTest(fn, VmSkipTests); err != nil {
		t.Error(err)
	}
}

func TestVmLogTest(t *testing.T) {
	fn := filepath.Join(vmTestDir, "vmLogTest.json")
	if err := RunVmTest(fn, VmSkipTests); err != nil {
		t.Error(err)
	}
}

func TestVmPerformance(t *testing.T) {
	fn := filepath.Join(vmTestDir, "vmPerformanceTest.json")
	if err := RunVmTest(fn, VmSkipTests); err != nil {
		t.Error(err)
	}
}

func TestVmPushDupSwap(t *testing.T) {
	fn := filepath.Join(vmTestDir, "vmPushDupSwapTest.json")
	if err := RunVmTest(fn, VmSkipTests); err != nil {
		t.Error(err)
	}
}

func TestVmVMSha3(t *testing.T) {
	fn := filepath.Join(vmTestDir, "vmSha3Test.json")
	if err := RunVmTest(fn, VmSkipTests); err != nil {
		t.Error(err)
	}
}

func TestVm(t *testing.T) {
	fn := filepath.Join(vmTestDir, "vmtests.json")
	if err := RunVmTest(fn, VmSkipTests); err != nil {
		t.Error(err)
	}
}

func TestVmLog(t *testing.T) {
	fn := filepath.Join(vmTestDir, "vmLogTest.json")
	if err := RunVmTest(fn, VmSkipTests); err != nil {
		t.Error(err)
	}
}

func TestVmInputLimits(t *testing.T) {
	fn := filepath.Join(vmTestDir, "vmInputLimits.json")
	if err := RunVmTest(fn, VmSkipTests); err != nil {
		t.Error(err)
	}
}

func TestVmInputLimitsLight(t *testing.T) {
	fn := filepath.Join(vmTestDir, "vmInputLimitsLight.json")
	if err := RunVmTest(fn, VmSkipTests); err != nil {
		t.Error(err)
	}
}

func TestVmVMRandom(t *testing.T) {
	fns, _ := filepath.Glob(filepath.Join(baseDir, "RandomTests", "*"))
	for _, fn := range fns {
		if err := RunVmTest(fn, VmSkipTests); err != nil {
			t.Error(err)
		}
	}
}