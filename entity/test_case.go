package entity

type TestCase struct {
	Number      int
	Name        string `json:"name"`
	Description string `json:"description"`
}
