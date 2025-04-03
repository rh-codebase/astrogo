package ephemeris

import (
	"fmt"
	"strings"
	"testing"

	th "github.com/rh-codebase/genutilsgo"
)

func TestReadCatalog(t *testing.T) {

	bsc := make(BSC)

	err := bsc.ReadYaml("brightSourceCatalog.yml")
	if err != nil {
		fmt.Println("ReadYML returned err: ", err)
		t.Fail()
	}
	for k, v := range bsc {
		fmt.Printf("Source: %s, DEC: %s, RA: %s, Mag: %6.3f\n", k, v.DEC.SexagesimalDMS(), v.PMRA.SexagesimalHMS(), v.Magnitude)
	}
}

func TestAddCatalog(t *testing.T) {
	cs := CatalogSource{"BSC", "bsc.yml"}
	AddCatalog(cs)
}

func TestLoadCatalogs(t *testing.T) {
	cs := CatalogSource{"BSC1", "bsc1.yml"}
	AddCatalog(cs)
	cs2 := CatalogSource{"BSC2", "bsc2.yml"}
	AddCatalog(cs2)
	bsc := make(BSC)
	err := bsc.LoadCatalogs()
	if err == nil {
		t.Fail()
	}
	ClearCatalogs()
	cs = CatalogSource{"BSC", "brightSourceCatalog.yml"}
	AddCatalog(cs)
	err = bsc.LoadCatalogs()
	if err != nil {
		t.Fail()
	}
	fmt.Println("ZetPav: ", bsc[strings.ToLower("ZetPav")])
	fmt.Println("AlpBoo: ", bsc["alpboo"])
}

func TestClearCatalogs(t *testing.T) {
	cs := CatalogSource{"BSC1", "bsc1.yml"}
	AddCatalog(cs)
	ClearCatalogs()
	if cats != nil {
		t.Fail()
	}
}

func TestSourceNameCase(t *testing.T) {
	cs := CatalogSource{"BSC", "brightSourceCatalog.yml"}
	AddCatalog(cs)
	bsc := make(BSC)
	err := bsc.LoadCatalogs()
	if err != nil {
		fmt.Println("Failed to load Catalog")
		t.Fail()
	}
	_, err = bsc.GetSource("ZetPav")
	if err != nil {
		fmt.Println("ZetPav not found")
		t.Fail()
	}

	s, err := bsc.GetSource("zetpav")
	if err != nil {
		fmt.Println("zetpav not found")
		t.Fail()
	} else {
		th.CheckFT(t, s.RA.Hour().Value, 18.717252777777777, 1e-6, "Value Error")
		th.CheckFT(t, s.DEC.Degree().Value, -71.42811111111111, 1e-6, "Value Error")
		th.CheckFT(t, s.Magnitude, 4.01, 1e-6, "Value Error")
	}
}
