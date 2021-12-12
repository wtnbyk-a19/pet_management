package Create

type IPetCreateUseCase interface {
	Handle(inputData PetCreateInputData) (outputData PetCreateOutputData)
}
