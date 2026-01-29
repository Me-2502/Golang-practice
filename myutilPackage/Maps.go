package myutil

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

type emptyInterface struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

type hmap struct {
	count     int // number of elements
	flags     uint8
	B         uint8 // log_2 of # of buckets
	noverflow uint16
	hash0     uint32

	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr
	extra      unsafe.Pointer
}

func changeMapValue(m map[string]int, s string, i int) {
	m[s] = i
	m["Harry Hudson"] = 48
}

func Maps() {
	// MAPS
	domainToIp := map[string]string{}
	domainToIp["https://www.google.com"] = "8.8.8.8"
	domainToIp["https://www.amazon.com"] = "98.87.174.74"
	domainToIp["https://www.microsoft.com"] = "13.107.213.68"
	url := "https://www.amazon.com"
	fmt.Printf("Ip of %s is %s\n", url, domainToIp[url])
	fmt.Println(unsafe.Sizeof(domainToIp), domainToIp)
	fmt.Println()

	delete(domainToIp, "https://www.microsoft.com")
	url = "https://www.microsoft.com"
	v, ok := domainToIp[url]
	if ok {
		fmt.Printf("Ip of %s is %s\n", url, v)
	} else {
		fmt.Println("Record not found")
	}
	fmt.Println(unsafe.Sizeof(domainToIp), domainToIp)
	for k, v := range domainToIp {
		fmt.Printf("Ip of %s is %s\n", k, v)
	}
	for k, v := range domainToIp {
		fmt.Printf("Ip of %s is %s\n", k, v)
	}
	fmt.Println()

	nameToAge := map[string]int{
		"John Doe":  43,
		"Mary Jane": 28,
	}
	changeMapValue(nameToAge, "John Doe", 32)
	fmt.Println(nameToAge)
	fmt.Println()

	// var nill map[string]string
	// nill["Hello"] = "world!"
	// fmt.Println("Hello", nill["Hello"])
	var mapEmptyIface = any(domainToIp)
	mapEmptyIfacePtr := unsafe.Pointer(&mapEmptyIface)
	emptyIfaceDataPtr := (*emptyInterface)(mapEmptyIfacePtr)
	hm := (*hmap)(emptyIfaceDataPtr.data)
	fmt.Println(hm.count)
	fmt.Println(uint64(1) << (hm).B)
	fmt.Println()

	var startTime time.Time
	hinted := make(map[int]int, 500)
	startTime = time.Now()
	for i := 0; i < 500; i++ {
		hinted[i] = (i << 2)
	}
	fmt.Println(time.Since(startTime))
	unhinted := map[int]int{}
	startTime = time.Now()
	for i := 0; i < 500; i++ {
		unhinted[i] = (i << 2)
	}
	fmt.Println(time.Since(startTime))
	fmt.Println()

	var start time.Time
	var sMap sync.Map
	var nMap = make(map[int]int)
	var hMap = make(map[int]int, 300)
	start = time.Now()
	for i := 0; i < 500; i++ {
		sMap.Store(i, (i << 1))
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	for i := 0; i < 500; i++ {
		nMap[i] = (i << 1)
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	for i := 0; i < 500; i++ {
		hMap[i] = (i << 1)
	}
	fmt.Println(time.Since(start))
	fmt.Println()
}
