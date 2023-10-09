package helmcharts

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/test-network-function/oct/pkg/http"
)

const (
	helmCatalogURL       = "https://charts.openshift.io/index.yaml"
	unofficialCatalogURL = ""
)

const (
	helmRelativePath = "%s/cmd/tnf/fetch/data/helm/helm.db"
)

// Returns all the URLs of the available oneline catalogs.
func GetHelmChartsCatalogs() []string {
	return []string{
		helmCatalogURL,
		unofficialCatalogURL,
	}
}

func DownloadHelmCatalogs() error {
	start := time.Now()
	err := removeHelmDB()
	if err != nil {
		return err
	}

	log.Infof("Getting helm charts catalog page, url: %s", helmCatalogURL)
	body, err := http.GetHTTPBody(helmCatalogURL)
	if err != nil {
		return err
	}

	path, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}
	filename := fmt.Sprintf(helmRelativePath, path)
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	_, err = f.Write(body)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filename, err)
	}

	log.Info("Time to process all the charts: ", time.Since(start))
	return nil
}

func removeHelmDB() error {
	path, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	filename := fmt.Sprintf(helmRelativePath, path)
	err = os.Remove(filename)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove file %s: %w", filename, err)
	}

	return nil
}
