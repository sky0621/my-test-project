package converter

func ToPtr[T any](v T) *T {
	return &v
}

func ToVal[T any](p *T) T {
	return *p
}

func ToVals[T any](ps []*T) []T {
	results := make([]T, len(ps))
	for i, v := range ps {
		results[i] = *v
	}
	return results
}

func PtrIfNotNil[S any, T any](value *S, converter func(S) T) *T {
	if value == nil {
		return nil
	}
	converted := converter(*value)
	return &converted
}

func PtrIfNotEmpty[T any](s []T) *[]T {
	if len(s) == 0 {
		return nil
	}
	return &s
}
