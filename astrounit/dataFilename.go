package astrounit

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	at "github.com/rh-codebase/astrogo/astrotime"
)

/*
   provides a simple interface for handling data filenames
of the format:
<data_dir_root>/YYYY-MM/<filename_prefix><filename>.<filename_suffix>
as well as:
<data_dir_root>/YYYY-MM-DD/<filename_prefix><filename>.<filename_suffix>

where <data_dir_root>/YYYY-MM(-DD)/<filename_prefix><filename>.<filename_suffix>
is determined from the current time on the machine running
the code. 4-digit year, 2-digit month, 2-digit day.

<data_dir_root>/YYYY-MM(-DD) will be created if it does not exist.

For example:

If

  data_dir_root = '/opt/testbed_data'
  filename_prefix = 'PacketDump'
  filename = 1903
  filename_suffix = 'h5'

Then

  the filename will be: /opt/testbed_data/YYYY-MM/PacketDump1903.h5
*/

type DataFilename struct {
	DataDirRoot    string
	FilenamePrefix string
	FilenameSuffix string
}

const (
	YM_FMT   = "2006-01"
	YMD_FMT  = "2006-01-02"
	HOUR_FMT = "2006-01-02T15_"
)

var (
	Fmt         string
	currentPath *string
)

func init() {
	Fmt = YM_FMT
}

func (df *DataFilename) UseDayFmt() {
	Fmt = YMD_FMT
}

func (df *DataFilename) UseMonthFmt() {
	Fmt = YM_FMT
}

func (df *DataFilename) GetDataDirNow() string {
	return time.Now().Format(Fmt)
}

func (df *DataFilename) GetDataDir(t time.Time) string {
	return t.Format(Fmt)
}

func (df *DataFilename) GetDataPath(t time.Time) string {
	return filepath.Join(df.DataDirRoot, df.GetDataDir(t))
}

func (df *DataFilename) GetDataFilename(t time.Time, fn string) string {
	p := df.GetDataPath(t)
	return filepath.Join(p, fmt.Sprintf("%s%s%s", df.FilenamePrefix, fn, df.FilenameSuffix))
}

func (df *DataFilename) SetFilenamePrefixHour(t time.Time) {
	df.FilenamePrefix = t.Format(HOUR_FMT)
}

// idempotent.
func (df *DataFilename) CreateDataDir(t time.Time, m os.FileMode) error {
	p := df.GetDataPath(t)
	return os.MkdirAll(p, m)
}

func (df *DataFilename) GetDataFilenamesStr(s, e string) ([]string, error) {
	var as []string
	sti, err := at.ParseLocalTimeToUTC(s)

	if err != nil {
		return as, err
	}
	eti, err := at.ParseLocalTimeToUTC(e)
	if err != nil {
		return as, err
	}

	return df.GetDataFilenames(sti, eti)
}

func (df *DataFilename) GetDataFilenames(st, et time.Time) ([]string, error) {
	var res []string
	tt := st
	for tt.Compare(et) != 1 { // tt <= et
		df.SetFilenamePrefixHour(tt)
		fn := df.GetDataFilename(tt, "monitordata")
		res = append(res, fn)
		tt = tt.Add(1.0 * time.Hour)
	}
	return res, nil
}
