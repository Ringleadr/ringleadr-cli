package Format

import (
	"fmt"
	"github.com/GodlikePenguin/agogos-datatypes"
	"github.com/dustin/go-humanize"
	"time"
)

func PrintNodeStats(stats *Datatypes.NodeStatsEntry) error {
	if len(stats.Stats) == 0 {
		fmt.Printf("No stats recorded for %s", stats.Name)
		return nil
	}
	fmt.Printf("Most recent statistics for %s", stats.Name)
	mostRecent := stats.Stats[len(stats.Stats)-1]
	fmt.Printf("Recorded at:\t\t\t%s\n", humanize.Time(time.Unix(mostRecent.Timestamp, 0)))
	fmt.Printf("Number of CPUs:\t\t\t%d\n", mostRecent.Cpus)
	fmt.Printf("CPU Usage:\t\t\t%f%%\n", mostRecent.CpuPercent)
	fmt.Printf("Total Memory:\t\t\t%s\n", humanize.Bytes(mostRecent.TotalMem))
	fmt.Printf("Used Memory:\t\t\t%s\n", humanize.Bytes(mostRecent.UsedMem))
	fmt.Printf("Used Memory percent:\t\t%f%%\n", mostRecent.UsedMemPercent)
	fmt.Printf("Number of containers on host:\t%d\n", mostRecent.NumContainers)
	return nil
}
