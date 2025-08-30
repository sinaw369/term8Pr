package SearchLogsEng

import (
	"github.com/sinaw369/term8Pr/Tpage"
	"github.com/sinaw369/term8Pr/handler/SearchLogsEng/SearchLogEngine"
	"github.com/sinaw369/term8Pr/handler/SearchLogsEng/SearchLogEngine/adminHandler"
	"github.com/sinaw369/term8Pr/handler/SearchLogsEng/search"
	"github.com/sinaw369/term8Pr/handler/SearchLogsEng/service/adminService"
)

func SetUpSearchLogEngine(Path string) *Tpage.Deliver {

	SearchLog := search.New(Path)
	AdminSvc := adminService.New(SearchLog)
	AdminHand := adminHandler.New(*AdminSvc)
	SL := SearchLogEngine.New(*AdminHand)
	return Tpage.New(SL)
}

// PrintMemUsage
// Alloc -> uint64
// Alloc is bytes of allocated heap objects.
// "Allocated" heap objects include all reachable objects, as well as unreachable objects that the garbage collector has not yet freed.
// Specifically, Alloc increases as heap objects are allocated and decreases as the heap is swept and unreachable objects are freed.
// Sweeping occurs incrementally between GC cycles, so these two processes occur simultaneously, and as a result Alloc tends to change smoothly (in contrast with the sawtooth that is typical of stop-the-world garbage collectors).
// TotalAlloc -> uint64
// TotalAlloc is cumulative bytes allocated for heap objects.
// TotalAlloc increases as heap objects are allocated, but unlike Alloc and HeapAlloc, it does not decrease when objects are freed.
// Sys -> uint64
// Sys is the total bytes of memory obtained from the OS.
// Sys is the sum of the XSys fields below. Sys measures the virtual address space reserved by the Go runtime for the heap, stacks, and other internal data structures. It's likely that not all of the virtual address space is backed by physical memory at any given moment, though in general it all was at some point.
// NumGC -> uint32
// NumGC is the number of completed GC cycles./*
/*
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024) //Print only the current memory usage
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	//runtime.GC()
}
*/
