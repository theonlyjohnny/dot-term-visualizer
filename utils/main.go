package utils

import (
	"github.com/theonlyjohnny/dotCli/logger"
	terminaldimensions "github.com/wayneashleyberry/terminal-dimensions"
)

const (
	monitorWidthIn  = 32
	aspectRatio     = 21.0 / 9.0
	monitorHeightIn = float64(monitorWidthIn / aspectRatio)
	ptPerInch       = 1.0 / 72.0
)

var (
	log        = logger.Log
	maxRows    float64
	maxColumns float64
	colPerInch float64
	rowPerInch float64
)

func init() {
	maxHeightUint, heightError := terminaldimensions.Height()
	maxWidthUint, widthError := terminaldimensions.Width()

	if widthError != nil || heightError != nil {
		log.Error("Couldn't query terminal size:")
		log.Errorf("widthError: %s", widthError.Error())
		log.Errorf("heightError: %s", heightError.Error())
	}

	maxColumns = float64(maxWidthUint)
	maxRows = float64(maxHeightUint)

	colPerInch = maxColumns / monitorWidthIn
	rowPerInch = maxRows / monitorHeightIn
	log.Debugf("colPerInch: %v, rowPerInch: %v, ptPerInch: %v, maxColumns: %v, maxRows: %v, monitorHeightIn: %v, monitorWidthIn: %v", colPerInch, rowPerInch, ptPerInch, maxColumns, maxRows, monitorHeightIn, monitorWidthIn)
}

func GetRectFromCommaString(v string) [4]int {
	ptsRect := getPtRectFromCommaString(v)
	termRect := getTermRectFromPtRect(&ptsRect)
	scaleUp(&termRect)
	return termRect.getRoundedRectSlice()
}
