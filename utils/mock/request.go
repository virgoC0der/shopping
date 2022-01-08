package mock

import (
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"

	. "shopping/utils/log"
	"shopping/utils/webbase"
)

func RequestApi(url string, method string, body []byte) (*webbase.CommonResp, error) {
	reqData := bytes.NewReader(body)
	request, err := http.NewRequest(method, url, reqData)
	if err != nil {
		Logger.Warn("new request err", zap.Error(err))
		return nil, err
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		Logger.Warn("do request err", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Logger.Warn("read response err", zap.Error(err))
		return nil, err
	}

	response := &webbase.CommonResp{}
	err = json.Unmarshal(b, response)
	if err != nil {
		Logger.Warn("json unmarshal err", zap.Error(err))
		return nil, err
	}

	if response.Code != 0 {
		Logger.Warn("request api err", zap.Error(err))
		return response, err
	}

	return response, nil
}
