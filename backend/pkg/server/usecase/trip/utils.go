package trip

import (
	"fmt"
	"time"

	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
)

// GPS情報を抜き取る
func ExtractExifContent(img []byte) (map[string]interface{}, error) {
	// ---
	// exif抽出
	// ---
	rawExif, err := exif.SearchAndExtractExif(img)
	if err != nil {
		return nil, err
	}
	im, err := exifcommon.NewIfdMappingWithStandard()
	if err != nil {
		return nil, err
	}
	ti := exif.NewTagIndex()
	_, index, err := exif.Collect(im, ti, rawExif)
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})

	// ---
	// gps
	// ---
	ifd, err := index.RootIfd.ChildWithIfdPath(exifcommon.IfdGpsInfoStandardIfdIdentity)
	if err != nil {
		return nil, err
	}
	gi, err := ifd.GpsInfo()
	if err != nil {
		return nil, err
	}
	data["gps_latitude"] = gi.Latitude.Decimal()
	data["gps_longitude"] = gi.Longitude.Decimal()

	// ---
	// time
	// ---
	results, err := index.Tree[0].FindTagWithName("DateTime")
	if err != nil {
		return nil, err
	}
	if len(results) != 1 {
		errMsg := fmt.Errorf("there wasn't exactly one result. len=%d", len(results))
		return nil, errMsg
	}
	ite := results[0]
	time, err := ite.Value()
	if err != nil {
		return nil, err
	}
	// 型アサーション
	t, ok := time.(string)
	if !ok {
		return nil, fmt.Errorf("extracted type is wrong. alt=%s", "time")
	}

	timeTime, err := stringToTime(t)
	if err != nil {
		return nil, err
	}
	data["time"] = *timeTime

	return data, nil
}

func stringToTime(str string) (*time.Time, error) {
	t, err := time.Parse(imgLayout, str)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
