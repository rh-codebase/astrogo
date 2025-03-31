package astrounit

import (
	"fmt"
	"testing"
	"time"
)

func TestGetDataDirNow(t *testing.T) {
	var df DataFilename
	s := df.GetDataDirNow()
	fmt.Println(s)
}

func TestGetDataDir(t *testing.T) {
	var df DataFilename
	s := df.GetDataDir(time.Now())
	fmt.Println(s)

	df.UseDayFmt()
	s = df.GetDataDirNow()
	fmt.Println(s)
}

func TestGetDataPath(t *testing.T) {
	var df DataFilename
	s := df.GetDataPath(time.Now())
	fmt.Println(s)

	df.DataDirRoot = "/dataroot"
	s = df.GetDataPath(time.Now())
	fmt.Println(s)
}

func TestGetDataFilename(t *testing.T) {
	var df DataFilename
	ti := time.Now()
	s := df.GetDataFilename(ti, "test")
	fmt.Println(s)
	df.DataDirRoot = "test"
	df.FilenamePrefix = "pre_"
	df.FilenameSuffix = ".log"
	s = df.GetDataFilename(ti, "test")
	fmt.Println(s)

	df.SetFilenamePrefixHour(ti)
	df.UseDayFmt()
	s = df.GetDataFilename(ti, "test")
	fmt.Println(s)

}

func TestCreateDataDir(t *testing.T) {
	var df DataFilename
	ti := time.Now()
	df.UseDayFmt()
	df.DataDirRoot = "test"
	fmt.Println("dataPath: ", df.GetDataPath(ti))
	err := df.CreateDataDir(ti, 0755)
	if err != nil {
		t.Fail()
	}
}

func TestGetDataFilenamesStr(t *testing.T) {
	var df DataFilename
	//ti := time.Now()
	//s := df.GetDataFilename(ti, "test")
	//fmt.Println(s)
	df.DataDirRoot = "test"
	df.UseDayFmt()
	st := "2024-10-31T23:05:00-07:00"
	et := "2024-11-01T02:05:00-07:00"

	as, err := df.GetDataFilenamesStr(st, et)
	if err != nil {
		fmt.Println("Error GetDataFilenameStr: ", err)
		t.Fail()
	}
	for _, s := range as {
		fmt.Println(s)
	}

}
