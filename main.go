// Copyright 2019 The crowdcompute:cc-go-sdk Authors
// This file is part of the crowdcompute:cc-go-sdk library.
//
// The crowdcompute:cc-go-sdk library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The crowdcompute:cc-go-sdk library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the crowdcompute:cc-go-sdk library. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// InfoStat represents the information of a CPU
type InfoStat struct {
	CPU       int32   `json:"cpu"`
	VendorID  string  `json:"vendorId"`
	CoreID    string  `json:"coreId"`
	Cores     int32   `json:"cores"`
	ModelName string  `json:"modelName"`
	Mhz       float64 `json:"mhz"`
	CacheSize int32   `json:"cacheSize"`
}

// InfoCPU All the info of the CPU
type InfoCPU struct {
	PhysicalCPU int `json:"physical_cpu"`
	InfoStat    []InfoStat
}

func printCPUInfo() {

	cpus, _ := cpu.Counts(true)
	info, _ := cpu.Info()

	infoCPU := &InfoCPU{
		PhysicalCPU: cpus,
	}

	infoStat := make([]InfoStat, 0)
	for _, i := range info {
		iStat := InfoStat{
			CPU:       i.CPU,
			VendorID:  i.VendorID,
			CoreID:    i.CoreID,
			Cores:     i.Cores,
			ModelName: i.ModelName,
			Mhz:       i.Mhz,
			CacheSize: i.CacheSize,
		}
		infoStat = append(infoStat, iStat)
	}
	infoCPU.InfoStat = infoStat
	b, _ := json.Marshal(infoCPU)
	fmt.Print(string(b))
}

func main() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			printCPUInfo()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
