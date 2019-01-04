package hlcuptester

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"strings"
)

type ReqResp struct {
	URI            string
	RequestBody    string
	ResponseStatus int
	ResponseBody   string
	Err            error
}

func skip(r *bufio.Scanner, n int) bool {
	for i := 0; i < n; i++ {
		if !r.Scan() {
			return false
		}
	}
	return true
}

func Load(datapath string, phase int, uriFilter ... Filter) (ch chan *ReqResp, err error) {
	var ammo, answ *os.File
	var method string

	if phase == 2 {
		method = "post"
	} else if phase == 1 || phase == 3 {
		method = "get"
	} else {
		return nil, errors.New("wrong phase number")
	}

	ammo, err = os.Open(fmt.Sprintf("%sammo/phase_%d_%s.ammo", datapath, phase, method))
	if err != nil {
		return nil, errors.Wrap(err, "error opening Ammo file")
	}

	answ, err = os.Open(fmt.Sprintf("%sanswers/phase_%d_%s.answ", datapath, phase, method))
	if err != nil {
		return nil, errors.Wrap(err, "error opening Answers file")
	}

	ch = make(chan *ReqResp, 10)

	ammoScan := bufio.NewScanner(ammo)
	answScan := bufio.NewScanner(answ)
	// запросы
	go func() {
		defer ammo.Close()
		defer answ.Close()

		for ammoScan.Scan() {
			skip(ammoScan, 1)
			rs := &ReqResp{}
			s := strings.Split(ammoScan.Text(), " ")
			rs.URI = s[1]

			if phase == 2 {
				skip(ammoScan, 5)

				cls := strings.Split(ammoScan.Text(), " ")
				if len(cls) < 2 {
					ch <- &ReqResp{Err: errors.New("wrong content length")}
					return
				}
				cl, err := strconv.Atoi(cls[1])

				if err != nil {
					ch <- &ReqResp{Err: errors.Wrap(err, "can't parse content length")}
					return
				}

				skip(ammoScan, 2)
				// grap body
				if cl > 0 {
					skip(ammoScan, 1)
					rs.RequestBody = ammoScan.Text()
				}
			} else {
				skip(ammoScan, 5)
			}

			if !answScan.Scan() {
				ch <- &ReqResp{Err: errors.New("error getting answer")}
				return
			}

			a := strings.Split(answScan.Text(), "\t")
			rs.ResponseStatus, err = strconv.Atoi(a[2])
			if err != nil {
				ch <- &ReqResp{Err: errors.Wrap(err, "error parsing return status")}
				return
			}

			if len(a) == 4 {
				rs.ResponseBody = a[3]
			}

			if len(uriFilter) > 0 {
				for _, f := range uriFilter {
					if f(rs.URI) {
						ch <- rs
						break
					}
				}
			} else {
				ch <- rs
			}

		}

		close(ch)
	}()

	return ch, nil
}
