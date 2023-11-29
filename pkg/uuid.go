package pkg

import uuid "github.com/satori/go.uuid"

func CheckUUID(ids ...string) error {
	for _, id := range ids {
		_, err := uuid.FromString(id)
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateUUID() uuid.UUID {
	return uuid.Must(uuid.NewV4(), nil)
}
