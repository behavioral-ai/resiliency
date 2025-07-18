package operations

import (
	"fmt"
	"github.com/behavioral-ai/collective/exchange"
	"github.com/behavioral-ai/core/rest"
	"reflect"
)

func ExampleAuthorization_Chain() {
	name := "agent/authorization"
	chain := rest.BuildExchangeChain([]any{Authorization})
	fmt.Printf("test: BuildExchangeChain() -> %v\n", chain != nil)

	exchange.RegisterExchangeHandler(name, Authorization)
	l := exchange.ExchangeHandler(name)
	fmt.Printf("test: ExchangeLink() -> %v %v\n", reflect.TypeOf(Authorization), reflect.TypeOf(l))
	chain = rest.BuildExchangeChain([]any{l})
	fmt.Printf("test: repository.ExchangeLink() -> %v\n", chain != nil)

	//Output:
	//test: BuildExchangeChain() -> true
	//test: ExchangeLink() -> func(rest.Exchange) rest.Exchange func(rest.Exchange) rest.Exchange
	//test: repository.ExchangeLink() -> true

}
