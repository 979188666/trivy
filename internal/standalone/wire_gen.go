// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package standalone

import (
	"github.com/aquasecurity/trivy-db/pkg/db"
	"github.com/aquasecurity/trivy/pkg/detector/library"
	"github.com/aquasecurity/trivy/pkg/detector/ospkg"
	"github.com/aquasecurity/trivy/pkg/scanner"
	library2 "github.com/aquasecurity/trivy/pkg/scanner/library"
	ospkg2 "github.com/aquasecurity/trivy/pkg/scanner/ospkg"
	"github.com/aquasecurity/trivy/pkg/vulnerability"
)

// Injectors from inject.go:

func initializeScanner() scanner.Scanner {
	detector := ospkg.Detector{}
	ospkgScanner := ospkg2.NewScanner(detector)
	driverFactory := library.DriverFactory{}
	libraryDetector := library.NewDetector(driverFactory)
	libraryScanner := library2.NewScanner(libraryDetector)
	scannerScanner := scanner.NewScanner(ospkgScanner, libraryScanner)
	return scannerScanner
}

func initializeVulnerabilityClient() vulnerability.Client {
	config := db.Config{}
	client := vulnerability.NewClient(config)
	return client
}
