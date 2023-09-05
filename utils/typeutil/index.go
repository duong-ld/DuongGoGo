package typeutil

import "encoding/json"

func TypeConverter[R any, T any](data T) (*R, error) {
	var result R
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func ArrayTypeConverter[R any, T any](data []T) ([]*R, error) {
	var err error = nil
	result := make([]*R, len(data))

	for i := range result {
		result[i], err = TypeConverter[R, T](data[i])

		if err != nil {
			return result, err
		}
	}
	return result, err
}
