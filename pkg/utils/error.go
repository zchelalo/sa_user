package utils

import "google.golang.org/grpc/codes"

func IsErrorType(err error, errArray []error) bool {
	for _, e := range errArray {
		if err == e {
			return true
		}
	}
	return false
}

func ConvertStatusCodeToProtoCode(code int32) codes.Code {
	switch code {
	case 400:
		return codes.InvalidArgument
	case 404:
		return codes.NotFound
	case 409:
		return codes.AlreadyExists
	default:
		return codes.Internal
	}
}
