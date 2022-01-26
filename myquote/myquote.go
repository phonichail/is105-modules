package myquote

import myPKG "rsc.io/quote"

func GetGlassQuote() string {
	return myPKG.Glass()
}

func GetGoQuote() string {
	return myPKG.Go()
}
func GetHelloQuote() string {
	return myPKG.Hello()
}
func GetOptQuote() string {
	return myPKG.Opt()
}
