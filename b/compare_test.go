package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCompare(t *testing.T) {
	var i, X, K, s, f int
	var tasks []int
	var sd, fd, msd, mfd time.Duration
	const iterations = 1e6
	const maxN = 105
	const mx = 109
	const maxDuration = time.Second * 2

	defer func() {
		err := recover()
		fmt.Printf("i[%5d] X[%2d] K[%2d] s[%d] f[%d] msd[%v] mfd[%v] tasks%+v err->%v\n", i, X, K, s, f, msd, mfd, tasks, err)

		if i != iterations {
			t.FailNow()
		}
	}()

	baseTest(t, simpleCalsDays)
	baseTest(t, fastCalsDays)

	rsrc := rand.NewSource(666)
	rnd := rand.New(rsrc)

	for i = 1; i < iterations; i++ {
		N := rnd.Int31n(maxN-1) + 1
		X = int(rnd.Int31n(mx-1) + 1)
		K = int(rnd.Int31n(mx-1) + 1)
		tasks = make([]int, N)
		for i := 0; i < int(N); i++ {
			tasks[i] = int(rnd.Int31n(mx-1) + 1)
		}

		s, sd = makeCall(X, K, tasks, simpleCalsDays)
		f, fd = makeCall(X, K, tasks, fastCalsDays)

		if msd < sd {
			msd = sd
		}
		if mfd < fd {
			mfd = fd
		}

		require.Equalf(t, s, f, "i(%5d) s(%d)!=f(%d) tasks%+v", s, f, tasks)
		require.Lessf(t, sd, maxDuration, "sd(%+v) fd(%+v)", sd, fd)
		require.Lessf(t, fd, maxDuration, "sd(%+v) fd(%+v)", sd, fd)
	}
}

func makeCall(X, K int, tasks []int, days calcDays) (res int, duration time.Duration) {
	cpyTasks := make([]int, len(tasks))
	copy(cpyTasks, tasks)
	start := time.Now()
	res = days(X, K, cpyTasks)
	duration = time.Since(start)
	return
}
