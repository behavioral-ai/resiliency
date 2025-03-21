package limit

import "fmt"

func ExampleNewRateLimiter() {
	l := NewRateLimiter(100, 10)

	fmt.Printf("test: NewRateLimiter() -> [%v]\n", l)

	//Output:
	//test: NewRateLimiter() -> [rate: 100 burst: 10]

}

/*
func rateLimiterSetValues(limit rate.Limit, burst int) url.Values {
	v := make(url.Values)
	if limit != -2 {
		v.Add(RateLimitKey, fmt.Sprintf("%v", limit))
	}
	if burst != -2 {
		v.Add(RateBurstKey, strconv.Itoa(burst))
	}
	return v
}

func Example_newRateLimiter() {
	t := newRateLimiter("test-routing", newTable(true), NewRateLimiterConfig(true, 503, 1, 100))
	fmt.Printf("test: newRateLimiter() -> [name:%v] [limit:%v] [burst:%v] [statusCode:%v]\n", t.name, t.config.Limit, t.config.Burst, t.StatusCode())

	t = newRateLimiter("test-route2", newTable(true), NewRateLimiterConfig(true, 429, rate.Inf, DefaultBurst))
	fmt.Printf("test: newRateLimiter() -> [name:%v] [limit:%v] [burst:%v] [statusCode:%v]\n", t.name, t.config.Limit, t.config.Burst, t.StatusCode())

	t2 := cloneRateLimiter(t)
	t2.config.Limit = 123
	fmt.Printf("test: cloneRateLimiter() -> [prev-limit:%v] [prev-name:%v] [curr-limit:%v] [curr-name:%v]\n", t.config.Limit, t.name, t2.config.Limit, t2.name)

	//Output:
	//test: newRateLimiter() -> [name:test-routing] [limit:1] [burst:100] [statusCode:503]
	//test: newRateLimiter() -> [name:test-route2] [limit:1.7976931348623157e+308] [burst:1] [statusCode:429]
	//test: cloneRateLimiter() -> [prev-limit:1.7976931348623157e+308] [prev-name:test-route2] [curr-limit:123] [curr-name:test-route2]

}

func ExampleRateLimiter_State() {
	tbl := newTable(true)
	t := newRateLimiter("test-routing", tbl, NewRateLimiterConfig(true, 503, rate.Inf, 5))
	fmt.Printf("test: newRateLimiter() -> [name:%v] [limit:%v] [burst:%v] [statusCode:%v]\n", t.name, t.config.Limit, t.config.Burst, t.StatusCode())

	limit, burst := t.state()
	fmt.Printf("test: rateLimiterState(t) -> [enabled:%v] [limit:%v] [burst:%v]\n", t.IsEnabled(), limit, burst)

	t.config.Enabled = false

	limit, burst = t.state()
	fmt.Printf("test: rateLimiterState(t) -> [enabled:%v] [limit:%v] [burst:%v]\n", t.IsEnabled(), limit, burst)

	//Output:
	//test: newRateLimiter() -> [name:test-routing] [limit:1.7976931348623157e+308] [burst:5] [statusCode:503]
	//test: rateLimiterState(t) -> [enabled:true] [limit:99999] [burst:5]
	//test: rateLimiterState(t) -> [enabled:false] [limit:-1] [burst:-1]

}

func ExampleRateLimiter_Toggle() {
	name := "test-routing"
	config := NewRateLimiterConfig(true, 503, 100, 10)
	t := newTable(true)
	err := t.add(newConfig(name, config))
	fmt.Printf("test: Add() -> [%v] [count:%v]\n", err, t.count())

	ctrl := t.lookup(name)
	fmt.Printf("test: IsEnabled() -> [%v]\n", ctrl.RateLimiter().IsEnabled())
	prevEnabled := ctrl.RateLimiter().IsEnabled()

	ctrl.RateLimiter().Signal(enableValues(false))
	ctrl1 := t.lookup(name)
	fmt.Printf("test: Disable() -> [prev-enabled:%v] [curr-enabled:%v]\n", prevEnabled, ctrl1.RateLimiter().IsEnabled())
	prevEnabled = ctrl1.RateLimiter().IsEnabled()

	ctrl1.RateLimiter().Signal(enableValues(true))
	ctrl = t.lookup(name)
	fmt.Printf("test: Enable() -> [prev-enabled:%v] [curr-enabled:%v]\n", prevEnabled, ctrl.RateLimiter().IsEnabled())

	//Output:
	//test: Add() -> [[]] [count:1]
	//test: IsEnabled() -> [true]
	//test: Disable() -> [prev-enabled:true] [curr-enabled:false]
	//test: Enable() -> [prev-enabled:false] [curr-enabled:true]

}

func ExampleRateLimiter_Signal() {
	name := "test-routing"
	burst := 1
	var limit rate.Limit
	config := NewRateLimiterConfig(true, 503, 100, 10)
	t := newTable(true)
	errs := t.add(newConfig(name, config))
	fmt.Printf("test: Add() -> [%v] [count:%v]\n", errs, t.count())

	ctrl := t.lookup(name)
	if c2, ok := any(ctrl).(*controller); ok {
		limit, burst = c2.rateLimiter.state()
	}
	fmt.Printf("test: rateLimiterState(100,10) -> [limit:%v] [burst:%v]\n", limit, burst)

	fmt.Printf("test: Signal(nil) -> [nil:%v]\n", ctrl.RateLimiter().Signal(nil))
	fmt.Printf("test: Signal([]) ->  [empty:%v]\n", ctrl.RateLimiter().Signal(make(url.Values)))

	err := ctrl.RateLimiter().Signal(rateLimiterSetValues(100, 0))
	ctrl1 := t.lookup(name)
	if c2, ok := any(ctrl1).(*controller); ok {
		limit, burst = c2.rateLimiter.state()
	}
	fmt.Printf("test: Signal(100,0) -> [error:%v] [limit:%v] [burst:%v]\n", err, limit, burst)

	err = ctrl1.RateLimiter().Signal(rateLimiterSetValues(-1, 10))
	ctrl = t.lookup(name)
	if c2, ok := any(ctrl).(*controller); ok {
		limit, burst = c2.rateLimiter.state()
	}
	fmt.Printf("test: Signal(-1,10) -> [error:%v] [limit:%v] [burst:%v]\n", err, limit, burst)

	err = ctrl.RateLimiter().Signal(rateLimiterSetValues(100, 10))
	ctrl1 = t.lookup(name)
	if c2, ok := any(ctrl1).(*controller); ok {
		limit, burst = c2.rateLimiter.state()
	}
	fmt.Printf("test: Signal(100,10) -> [error:%v] [limit:%v] [burst:%v]\n", err, limit, burst)

	err = ctrl1.RateLimiter().Signal(rateLimiterSetValues(100, 8))
	ctrl = t.lookup(name)
	if c2, ok := any(ctrl).(*controller); ok {
		limit, burst = c2.rateLimiter.state()
	}
	fmt.Printf("test: Signal(100,8) -> [error:%v] [limit:%v] [burst:%v]\n", err, limit, burst)

	err = ctrl.RateLimiter().Signal(rateLimiterSetValues(99, 8))
	ctrl1 = t.lookup(name)
	if c2, ok := any(ctrl1).(*controller); ok {
		limit, burst = c2.rateLimiter.state()
	}
	fmt.Printf("test: Signal(99,8) -> [error:%v] [limit:%v] [burst:%v]\n", err, limit, burst)

	err = ctrl1.RateLimiter().Signal(rateLimiterSetValues(-2, 5))
	ctrl = t.lookup(name)
	if c2, ok := any(ctrl).(*controller); ok {
		limit, burst = c2.rateLimiter.state()
	}
	fmt.Printf("test: Signal(99,5) -> [error:%v] [limit:%v] [burst:%v]\n", err, limit, burst)

	err = ctrl.RateLimiter().Signal(rateLimiterSetValues(88, -2))
	ctrl1 = t.lookup(name)
	if c2, ok := any(ctrl1).(*controller); ok {
		limit, burst = c2.rateLimiter.state()
	}
	fmt.Printf("test: Signal(88,5) -> [error:%v] [limit:%v] [burst:%v]\n", err, limit, burst)

	//Output:
	//test: Add() -> [[]] [count:1]
	//test: rateLimiterState(100,10) -> [limit:100] [burst:10]
	//test: Signal(nil) -> [nil:invalid argument: values are nil for rate limiter signal]
	//test: Signal([]) ->  [empty:<nil>]
	//test: Signal(100,0) -> [error:invalid argument: burst value is <= 0 [0]] [limit:100] [burst:10]
	//test: Signal(-1,10) -> [error:invalid argument: limit value is <= 0 [-1]] [limit:100] [burst:10]
	//test: Signal(100,10) -> [error:<nil>] [limit:100] [burst:10]
	//test: Signal(100,8) -> [error:<nil>] [limit:100] [burst:8]
	//test: Signal(99,8) -> [error:<nil>] [limit:99] [burst:8]
	//test: Signal(99,5) -> [error:<nil>] [limit:99] [burst:5]
	//test: Signal(88,5) -> [error:<nil>] [limit:88] [burst:5]

}


*/
