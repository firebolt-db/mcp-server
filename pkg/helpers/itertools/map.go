package itertools

// Map transforms each element of a slice by applying a converter function to it.
func Map[I, O any](inputs []I, converter func(I) O) []O {
	outputs := make([]O, len(inputs))
	for i, input := range inputs {
		outputs[i] = converter(input)
	}
	return outputs
}

// MapWithFailure is the same as Map but takes a converter function which can return an error.
// In the case that conversion returns an error the conversion process will be short circuited
// and the error and an empty slice will be returned
func MapWithFailure[I, O any](inputs []I, converter func(I) (O, error)) ([]O, error) {
	outputs := make([]O, 0, len(inputs))
	for _, input := range inputs {
		o, err := converter(input)
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, o)
	}
	return outputs, nil
}
