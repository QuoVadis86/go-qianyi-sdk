package qianyi

import (
	"encoding/json"
	"fmt"
)

type BaseResponse struct {
	ErrorCode string `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
	State     string `json:"state"`
	BizContent string `json:"bizContent"`
	RequestID string `json:"requestId"`
}

func (r *BaseResponse) IsSuccess() bool {
	return r.State == "success" && r.ErrorCode == ""
}

type BizContent struct {
	State      string          `json:"state"`
	NotSuccess bool            `json:"notSuccess,omitempty"`
	Total      int             `json:"total,omitempty"`
	Result     json.RawMessage `json:"result,omitempty"`
}

func (r *BaseResponse) ParseBizContent() (*BizContent, error) {
	bc := &BizContent{}
	if r.BizContent == "" {
		return bc, nil
	}
	if err := json.Unmarshal([]byte(r.BizContent), bc); err != nil {
		return nil, fmt.Errorf("parse bizContent: %w", err)
	}
	return bc, nil
}

func (r *BaseResponse) HasError() bool {
	return r.ErrorCode != "" || r.State == "failure"
}

type APIError struct {
	ErrorCode string
	Message   string
	RequestID string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("qianyi api error [%s]: %s (request_id: %s)", e.ErrorCode, e.Message, e.RequestID)
}

func parseResponse(body []byte, result any) error {
	var base BaseResponse
	if err := json.Unmarshal(body, &base); err != nil {
		return fmt.Errorf("unmarshal response: %w (body: %s)", err, truncate(string(body), 500))
	}

	if result != nil {
		if r, ok := result.(*BaseResponse); ok {
			*r = base
			return nil
		}
		if w, ok := result.(*ResponseWrapper); ok {
			w.BaseResponse = base
			bc, err := base.ParseBizContent()
			if err != nil {
				return err
			}
			w.BizContent = bc
			if len(bc.Result) > 0 && w.Result != nil {
				if err := json.Unmarshal(bc.Result, w.Result); err != nil {
					return fmt.Errorf("unmarshal bizContent.result: %w", err)
				}
			}
			return nil
		}
	}

	if base.HasError() {
		return &APIError{ErrorCode: base.ErrorCode, Message: base.ErrorMsg, RequestID: base.RequestID}
	}

	return nil
}

type ResponseWrapper struct {
	BaseResponse
	BizContent *BizContent
	Result     any
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
