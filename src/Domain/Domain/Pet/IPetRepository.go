package Pet

type IPetRepository interface {
	Save(pet *Pet) (error error)
}
