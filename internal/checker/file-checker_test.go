package Checker

import (
    "bytes"
    "io"
    "reflect"
    "testing"
)

func Test_readline(t *testing.T) {
    str := "hello, world\nI am zhaopeng.\n\n"
    buff := new(bytes.Buffer)
    buff.WriteString(str)

    type args struct {
        reader io.Reader
    }
    tests := []struct {
        name      string
        args      args
        wantCount int
        wantLine  []byte
        wantErr   error
    }{
        {
            name:      "1st line",
            args:      args{reader: buff},
            wantCount: 12,
            wantLine:  []byte("hello, world"),
            wantErr:   nil,
        },
        {
            name:      "2nd line",
            args:      args{reader: buff},
            wantCount: 14,
            wantLine:  []byte("I am zhaopeng."),
            wantErr:   nil,
        },
        {
            name:      "3rd line",
            args:      args{reader: buff},
            wantCount: 0,
            wantLine:  []byte(""),
            wantErr:   nil,
        },
        {
            name:      "4st line",
            args:      args{reader: buff},
            wantCount: 0,
            wantLine:  []byte(""),
            wantErr:   io.EOF,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotCount, gotLine, err := readline(tt.args.reader)
            if (err != nil) && err != tt.wantErr {
                t.Errorf("readline() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if gotCount != tt.wantCount {
                t.Errorf("readline() gotCount = %v, want %v", gotCount, tt.wantCount)
            }
            if !reflect.DeepEqual(gotLine, tt.wantLine) {
                t.Errorf("readline() gotLine = %v, want %v", gotLine, tt.wantLine)
            }
        })
    }
}

func TestFileChecker_check(t *testing.T) {
    testCases := []struct {
        name        string
        path        string
        want        bool
        wantedCount int
    }{
        {
            name:        "case1: 空文件",
            path:        "./testdata/int-0-asc.txt",
            want:        true,
            wantedCount: 0,
        },
        {
            name:        "case2:3items,asc",
            path:        "./testdata/int-3-asc.txt",
            want:        true,
            wantedCount: 3,
        },
        {
            name:        "case3:3items,desc",
            path:        "./testdata/int-3-desc.txt",
            want:        true,
            wantedCount: 3,
        },
        {
            name:        "case4:4items,desc,false",
            path:        "./testdata/int-4-asc-false.txt",
            want:        false,
            wantedCount: 3,
        },
        {
            name:        "case5:5items,desc,false",
            path:        "./testdata/int-5-desc-false.txt",
            want:        false,
            wantedCount: 4,
        },
    }
    for _, testCase := range testCases {
        t.Run(testCase.name, func(t *testing.T) {
            fc := NewFileChecker(testCase.path)
            if got, count := fc.Check(); got != testCase.want || count != testCase.wantedCount {
                t.Errorf("Check() = %v, count= %d, want %v, wantedCount: %d", got, count, testCase.want, testCase.wantedCount)
            }
        })
    }
}
