package schema

// UploaderResponse is the response of uploader.
type UploaderResponse struct {
	TotalCount      uint
	SuccessCount    uint
	SuccessContacts []string
	Errors          []UploaderError
}

// UploaderError is the error occured in uploader process
type UploaderError struct {
	Line    uint
	Contact string
	Err     error
}

// AddError appends the error in uploader response.
func (u *UploaderResponse) AddError(line uint, contact string, err error) {
	u.Errors = append(u.Errors, UploaderError{
		Line:    line,
		Contact: contact,
		Err:     err,
	})
	u.TotalCount++
}

// Successfull appends the successful contact.
func (u *UploaderResponse) Successfull(contact string) {
	u.SuccessContacts = append(u.SuccessContacts, contact)
	u.SuccessCount++
	u.TotalCount++
}
