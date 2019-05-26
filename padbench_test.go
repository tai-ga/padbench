package padbench

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/klauspost/cpuid"
)

type NoPad struct {
	a uint64
	b uint64
	c uint64
}

func (np *NoPad) Increase() {
	atomic.AddUint64(&np.a, 1)
	atomic.AddUint64(&np.b, 1)
	atomic.AddUint64(&np.c, 1)
}

type Pad struct {
	a uint64    // 8
	_ [7]uint64 // 8 * 7
	b uint64    // 8
	_ [7]uint64 // 8 * 7
	c uint64    // 8
	_ [7]uint64 // 8 * 7
}

func (p *Pad) Increase() {
	atomic.AddUint64(&p.a, 1)
	atomic.AddUint64(&p.b, 1)
	atomic.AddUint64(&p.c, 1)
}

func init() {
	/*
	   Print basic CPU information:
	*/
	fmt.Println("Name:", cpuid.CPU.BrandName)
	fmt.Println("PhysicalCores:", cpuid.CPU.PhysicalCores)
	fmt.Println("ThreadsPerCore:", cpuid.CPU.ThreadsPerCore)
	fmt.Println("LogicalCores:", cpuid.CPU.LogicalCores)
	fmt.Println("Family", cpuid.CPU.Family, "Model:", cpuid.CPU.Model)
	fmt.Println("Features:", cpuid.CPU.Features)
	fmt.Println("Cacheline bytes:", cpuid.CPU.CacheLine)
	fmt.Println("L1 Data Cache:", cpuid.CPU.Cache.L1D, "bytes")
	fmt.Println("L1 Instruction Cache:", cpuid.CPU.Cache.L1D, "bytes")
	fmt.Println("L2 Cache:", cpuid.CPU.Cache.L2, "bytes")
	fmt.Println("L3 Cache:", cpuid.CPU.Cache.L3, "bytes")

	/*
	   Test if we have a specific feature:
	*/
	if cpuid.CPU.SSE() {
		fmt.Println("We have Streaming SIMD Extensions")
	}
	fmt.Println("---")
}

func Benchmark_StructNoPad(b *testing.B) {
	nopad := &NoPad{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			nopad.Increase()
		}
	})
}

func Benchmark_StructPad(b *testing.B) {
	pad := &Pad{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pad.Increase()
		}
	})
}
