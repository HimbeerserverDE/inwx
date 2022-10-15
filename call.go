package inwx

type Call interface {
	expectedStatus() []Status
}
