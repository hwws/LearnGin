package messages

type Person struct {
	ID      string `xml:"id"`
	Name    string `xml:"name"`
	Email   string `xml:"email"`
	Phone   string `xml:"phone"`
	Address string `xml:"address"`
}
