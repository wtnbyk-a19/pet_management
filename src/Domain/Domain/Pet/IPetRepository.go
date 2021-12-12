package Pet

type IPetRepository interface {
	Save(pet *Pet) (result *Pet, error error)
}
