package account

type EmailRepo interface {
	Create(id string) (ret bool)
	A(i string)
}

type Email struct {
	emailRepo EmailRepo
}

func NewEmailService(emailRepo EmailRepo) *Email {
	return &Email{emailRepo: emailRepo}
}

func (e *Email) Create(emailID string) bool {
	return e.emailRepo.Create(emailID)
}
