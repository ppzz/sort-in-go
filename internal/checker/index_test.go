package Checker

import "testing"

func TestChecker_IsOrdered(t *testing.T) {
    type fields struct {
        lastVal int
        asc     bool
    }
    type args struct {
        item int
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   bool
    }{
        {
            name: "asc, last 100, next: 100",
            fields: fields{
                lastVal: 100,
                asc:     true,
            },
            args: args{
                item: 100,
            },
            want: true,
        },
        {
            name: "desc, last 100, next: 100",
            fields: fields{
                lastVal: 100,
                asc:     false,
            },
            args: args{
                item: 100,
            },
            want: true,
        },
        {
            name: "asc, last 100, next: 99",
            fields: fields{
                lastVal: 100,
                asc:     true,
            },
            args: args{
                item: 99,
            },
            want: false,
        },
        {
            name: "asc, last 100, next: 101",
            fields: fields{
                lastVal: 100,
                asc:     true,
            },
            args: args{
                item: 101,
            },
            want: true,
        },
        {
            name: "desc, last 100, next: 99",
            fields: fields{
                lastVal: 100,
                asc:     false,
            },
            args: args{
                item: 99,
            },
            want: true,
        },
        {
            name: "desc, last 100, next: 101",
            fields: fields{
                lastVal: 100,
                asc:     false,
            },
            args: args{
                item: 101,
            },
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            c := &Checker{
                lastVal: tt.fields.lastVal,
                asc:     tt.fields.asc,
            }
            if got := c.IsOrdered(tt.args.item); got != tt.want {
                t.Errorf("IsOrdered() = %v, want %v", got, tt.want)
            }
        })
    }
}
