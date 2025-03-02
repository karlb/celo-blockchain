package main

// The following directive is necessary to make the package coherent:

// This program generates contributors.go. It can be invoked by running
// go generate

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"
	"time"
)

func readABI(truffleJsonFile string) (string, error) {
	jsonData, err := ioutil.ReadFile(truffleJsonFile)
	if err != nil {
		return "", fmt.Errorf("Can't read build fild for %s: %w", truffleJsonFile, err)
	}

	var data struct {
		ABI json.RawMessage `json:"abi"`
	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return "", fmt.Errorf("Can't read ABI for %s: %w", truffleJsonFile, err)
	}

	return string(data.ABI), nil
}

var contractNames = []string{
	"Proxy",
	"Registry",
	"Freezer",
	"TransferWhitelist",
	"FeeCurrencyWhitelist",
	"GoldToken",
	"SortedOracles",
	"GasPriceMinimum",
	"ReserveSpenderMultiSig",
	"Reserve",
	"StableToken",
	"StableTokenEUR",
	"StableTokenBRL",
	"Exchange",
	"ExchangeEUR",
	"ExchangeBRL",
	"Accounts",
	"LockedGold",
	"Validators",
	"Election",
	"EpochRewards",
	"Random",
	"Attestations",
	"Escrow",
	"BlockchainParameters",
	"GovernanceSlasher",
	"DoubleSigningSlasher",
	"DowntimeSlasher",
	"GovernanceApproverMultiSig",
	"Governance",
}

var buildPath = flag.String("buildpath", "", "the folder where truffle contract build live (on monorepo ./packages/protocol/build/contracts )")
var outPath = flag.String("outpath", "./mycelo/contract", "relative path to mycelo/contract package")

func main() {
	flag.Parse()

	if buildPath == nil || *buildPath == "" {
		fmt.Println("Missing --buildpath variable")
		flag.PrintDefaults()
		os.Exit(1)
	}

	outfile := path.Join(*outPath, "gen_abis.go")

	fmt.Printf("Generating abi mappings on %s.\nReading contracts from %s\n", outfile, *buildPath)

	abis := make(map[string]string)
	for _, name := range contractNames {
		abi, err := readABI(path.Join(*buildPath, name+".json"))
		die(err)
		abis[name] = abi
	}

	f, err := os.Create(outfile)
	die(err)
	defer f.Close()

	fileTemplate.Execute(f, struct {
		Timestamp time.Time
		ABIs      map[string]string
	}{
		Timestamp: time.Now(),
		ABIs:      abis,
	})
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var fileTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
// {{ .Timestamp }}
package contract

import "github.com/celo-org/celo-blockchain/accounts/abi"

var abis map[string]*abi.ABI

func init() {
	abis = make(map[string]*abi.ABI)

	{{range $name, $abi := .ABIs -}}
	// {{$name}} ABI
	abis["{{$name}}"] = mustParseABI(` + "`{{$abi}}`" + `)


	{{- end }}
}
`))
