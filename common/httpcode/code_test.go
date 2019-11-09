package httpcode

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GetHTTPCodeTestSuite struct {
	suite.Suite
}

func TestGetHTTPCodeTestSuite(t *testing.T) {
	suite.Run(t, new(GetHTTPCodeTestSuite))
}

func (s *GetHTTPCodeTestSuite) TestGetHTTPCode() {
	type args struct {
		code int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TC01",
			args: args{
				code: 4010000,
			},
			want: 401,
		},
		{
			name: "TC02",
			args: args{
				code: 4040001,
			},
			want: 200,
		},
		{
			name: "TC03",
			args: args{
				code: 5000001,
			},
			want: 200,
		},
		{
			name: "TC04",
			args: args{
				code: 200,
			},
			want: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if got := GetHTTPCode(tt.args.code); got != tt.want {
				t.Errorf("GetHTTPCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *GetHTTPCodeTestSuite) TestParseHTTPStatus() {
	type args struct {
		code        int
		defaultCode []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TC01",
			args: args{
				code:        4000102,
				defaultCode: []int{SuccessfullyCode},
			},
			want: 400,
		},
		{
			name: "TC02",
			args: args{
				code:        -4000102,
				defaultCode: []int{SuccessfullyCode},
			},
			want: 200,
		},
		{
			name: "TC03",
			args: args{
				code:        5000102,
				defaultCode: []int{SuccessfullyCode},
			},
			want: 500,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if got := ParseHTTPStatus(tt.args.code, tt.args.defaultCode...); got != tt.want {
				t.Errorf("ParseHTTPStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
