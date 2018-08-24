package galpine

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	_ "github.com/ieee0824/galpine/statik"
	"github.com/rakyll/statik/fs"
)

// sorce "http://www.gsi.go.jp/KOKUJYOHO/MOUNTAIN/data.js"

type Coordinate struct {
	B int `json:"b"` // 緯度
	L int `json:"l"` // 経度
	H int `json:"h"` // 高度
}

type Name struct {
	Kanji    string `json:"kanji"`
	Hiragana string `json:"hiragana"`
}

type Place struct {
	Whereabouts     string    `json:"whereabouts"`      // 所在
	PrefectureNames [3]string `json:"prefecture_names"` // 県名
}

type Data struct {
	ID           int `json:"id"`
	Coordinate   `json:"coordinate"`
	MountainName Name   `json:"mountain_name"`
	TopName      Name   `json:"top_name"`
	PointName    string `json:"point_name"`
	Remarks      string `json:"remarks"`
	Place        `json:"place"`
}

func NewData(s string) (*Data, error) {
	s = strings.TrimPrefix(s, " ")
	s = strings.TrimSuffix(s, ";")
	if s == "" {
		return nil, errors.New("empty string")
	}
	if strings.HasPrefix(s, "//") {
		return nil, errors.New("comment line")
	}
	ss := strings.Split(s, "=")
	if len(ss) < 2 {
		return nil, errors.New("illegal format")
	}
	s = ss[1]
	s = strings.Replace(s, "'", "", -1)

	tokens := strings.Split(s, ",")
	if len(tokens) < 14 {
		return nil, errors.New("illegal format")
	}

	ret := new(Data)

	{
		id, err := strconv.Atoi(tokens[0])
		if err != nil {
			return nil, err
		}
		ret.ID = id
	}

	{
		b, err := strconv.Atoi(tokens[1])
		if err != nil {
			return nil, err
		}
		l, err := strconv.Atoi(tokens[2])
		if err != nil {
			return nil, err
		}
		h, err := strconv.Atoi(tokens[3])
		if err != nil {
			return nil, err
		}
		ret.Coordinate.B = b
		ret.Coordinate.L = l
		ret.Coordinate.H = h
	}

	ret.MountainName.Kanji = tokens[4]
	ret.MountainName.Hiragana = tokens[5]
	ret.TopName.Kanji = tokens[6]
	ret.TopName.Hiragana = tokens[7]
	ret.PointName = tokens[8]
	ret.Remarks = tokens[9]
	ret.Place.Whereabouts = tokens[10]
	ret.Place.PrefectureNames[0] = tokens[11]
	ret.Place.PrefectureNames[1] = tokens[12]
	ret.Place.PrefectureNames[2] = tokens[13]

	return ret, nil
}

func NewDatas() []*Data {
	fs, err := fs.New()
	if err != nil {
		return nil
	}
	f, err := fs.Open("/data.js")
	if err != nil {
		return nil
	}
	ret := []*Data{}
	bin, err := ioutil.ReadAll(f)
	if err != nil {
		return nil
	}
	lines := strings.Split(string(bin), "\n")

	for _, line := range lines {
		d, err := NewData(line)
		if err != nil {
			continue
		}
		ret = append(ret, d)
	}
	return ret
}
