package main
//Hi my soul! How much you will to sleep now? I need you very much


import (
        "log"
        "net/http"
        _ "net/http/pprof"
        "runtime"
        "sync"
)

func bigBytes() *[]byte {
        s := make([]byte, 100000000)
        return &s
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("That is reaction to /"))	
}

func main() {

//"Во дела 180701" for mn.go


	http.HandleFunc("/", rootHandler)

        var wg sync.WaitGroup

        go func() {
                log.Println(http.ListenAndServe(":6060", nil))
        }()

        var mem runtime.MemStats
        runtime.ReadMemStats(&mem)
        log.Println(mem.Alloc)
        log.Println(mem.TotalAlloc)
        log.Println(mem.HeapAlloc)
        log.Println(mem.HeapSys)

        for i := 0; i < 10; i++ {
                s := bigBytes()
                if s == nil {
                        log.Println("oh noes")
                }
        }

        runtime.ReadMemStats(&mem)
        log.Println(mem.Alloc)
        log.Println(mem.TotalAlloc)
        log.Println(mem.HeapAlloc)
        log.Println(mem.HeapSys)

        wg.Add(1)
        wg.Wait()

}
