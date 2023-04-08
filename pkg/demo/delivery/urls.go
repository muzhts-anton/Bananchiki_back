package demodel

const (
	urlViewJoin = "/presentation/view/join/{code}"
	urlView     = "/presentation/view/{hash}"
	urlShowGo   = "/presentation/{presId}/show/go/{idx}" // обнулить все vote относящиеся к prezId
	urlShowStop = "/presentation/{presId}/show/stop" // сделать все vote равными единице относящиеся к prezId
)
