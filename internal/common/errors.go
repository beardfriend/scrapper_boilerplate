package common

import (
	"math"

	"github.com/pkg/errors"
)

var (
	ErrDatabaseError = errors.New("데이터베이스 연결을 확인해주세요")
	ErrServerError   = errors.New("서버 에러입니다. 관리자에게 문의해주세요.")

	ErrNotFound         = errors.New("검색 결과가 없습니다")
	ErrScrapperNetwork  = errors.New("크롤러 네트워크 에러입니다")
	ErrScrapperFunction = errors.New("크롤러 기능 고장입니다")
	ErrToomanyReqeuest  = errors.New("요청이 많습니다.")
)

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, ErrNotFound)
}

func IsToomanyRequest(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, ErrToomanyReqeuest)
}

func MatchScrapperError(statusCode int, inputErr error) (err error) {
	if inputErr != nil {
		err = ErrScrapperNetwork
		return
	}

	if statusCode == 429 {
		err = ErrToomanyReqeuest
		return
	}

	if statusCode != 404 && math.Abs(float64(statusCode/200)) > 1 {
		err = ErrScrapperFunction

		return
	}

	err = nil
	return
}
