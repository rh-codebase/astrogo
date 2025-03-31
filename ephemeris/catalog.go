package ephemeris

// Read in brigh source catalog
import (
	"errors"
	"fmt"
	"strings"

	au "github.com/rh-codebase/astrogo/astrounit"
	gu "github.com/rh-codebase/genutilsgo"
)

type bSCstr struct {
	Epoch     string  `yaml:"Epoch"`
	RA_hms    string  `yaml:"RA_hms"`
	DEC_dms   string  `yaml:"DEC_dms"`
	PMRA_hms  string  `yaml:"PMRA_hms"`
	PMDEC_dms string  `yaml:"PMDEC_dms"`
	Magnitude float64 `yaml:"Magnitude"`
}

type BSCdata struct {
	Epoch     string
	RA        au.Angle
	DEC       au.Angle
	PMRA      au.Angle
	PMDEC     au.Angle
	Magnitude float64 `yaml:"Magnitude"`
}

type bSCs map[string]bSCstr

type BSC map[string]BSCdata

type CatalogSource struct {
	Name     string
	Filename string
}

type Catalogs []CatalogSource

var (
	cats Catalogs
)

func AddCatalog(cat CatalogSource) {
	cats = append(cats, cat)
}

func ClearCatalogs() {
	cats = nil
}

func (bsc *BSC) LoadCatalogs() error {
	var emsg string
	for _, c := range cats {
		err := bsc.ReadYaml(c.Filename)
		if err != nil {
			msg := fmt.Sprintf("Could not load catalog %s at %s, ", c.Name, c.Filename)
			emsg += msg
		}
	}
	if len(emsg) != 0 {
		return errors.New(emsg)
	}
	return nil
}

func (bsc *BSC) ReadYaml(fn string) error {
	bscs := make(bSCs)
	err := gu.ReadYaml(fn, &bscs)
	if err != nil {
		return err
	}

	// convert bscs to BSC
	for k, v := range bscs {
		var bscd BSCdata
		bscd.Epoch = v.Epoch
		valh, err := au.NewHMS(v.RA_hms)
		if err != nil {
			return err
		}
		bscd.RA = valh.Angle()

		vald, err := au.NewDMS(v.DEC_dms)
		if err != nil {
			return err
		}
		bscd.DEC = vald.Angle()

		valh, err = au.NewHMS(v.PMRA_hms)
		if err != nil {
			return err
		}
		bscd.PMRA = valh.Angle()

		vald, err = au.NewDMS(v.PMDEC_dms)
		if err != nil {
			return err
		}
		bscd.PMDEC = vald.Angle()

		bscd.Magnitude = v.Magnitude

		(*bsc)[strings.ToLower(k)] = bscd
	}
	return nil
}

func (bsc *BSC) GetSource(src string) (BSCdata, error) {
	src = strings.ToLower(src)
	if star, ok := (*bsc)[src]; !ok {
		emsg := fmt.Sprintf("Source %s not found in catalog", src)

		return star, errors.New(emsg)
	} else {
		return star, nil
	}
}
