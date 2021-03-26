package int_gen

import (
    "math/rand"
    "testing"
)

func TestIntGen_Gen(t *testing.T) {
    type fields struct {
        rand *rand.Rand
    }
    r := rand.New(rand.NewSource(1))
    i1 := r.Int()
    tests := []struct {
        name   string
        fields fields
        want   int
    }{
        {
            name:   "seed 1, 1",
            fields: fields{rand: rand.New(rand.NewSource(1))},
            want:   i1,
        },
        {
            name:   "seed 1, 1",
            fields: fields{rand: rand.New(rand.NewSource(1))},
            want:   i1,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            intGen := &IntGen{
                rand: tt.fields.rand,
            }
            if got := intGen.Gen(); got != tt.want {
                t.Errorf("Gen() = %v, want %v", got, tt.want)
            }
        })
    }
}
