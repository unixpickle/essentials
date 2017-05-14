package essentials

import "testing"

func TestRound(t *testing.T) {
	ins := []float64{
		2.33646,
		-3.25800,
		3.95544,
		4.41466,
		-9.20760,
		-1.76616,
		4.91430,
		-11.16655,
		-3.72920,
		3.41797,
		2.31483,
		-7.12645,
		0.45609,
		-15.44078,
		-12.33272,
		7.08801,
		-9.69377,
		-4.42842,
		-8.67027,
		12.98316,
		2.5,
		-2.5,
		3.5,
		-3.5,
	}
	outs := []float64{
		2,
		-3,
		4,
		4,
		-9,
		-2,
		5,
		-11,
		-4,
		3,
		2,
		-7,
		0,
		-15,
		-12,
		7,
		-10,
		-4,
		-9,
		13,
		3,
		-3,
		4,
		-4,
	}
	for i, in := range ins {
		actual := Round(in)
		expected := outs[i]
		if actual != expected {
			t.Errorf("round(%f) gave %f but should give %f", in, actual, expected)
		}
	}
}
