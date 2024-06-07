package helpers

import (
	"bytes"
	"encoding/binary"
	"mongodb-go/src/models"
)

func EncodeBSON(data []models.User) ([]byte, error) {
	// Create a bytes.Buffer to hold the BSON data
	buf := bytes.Buffer{}

	// Iterate over each student and encode them to BSON
	for _, user := range data {
		// Write name length
		nameLen := int32(len(user.Name))
		err := binary.Write(&buf, binary.LittleEndian, nameLen)
		if err != nil {
			return nil, err
		}
		// Write name
		_, err = buf.WriteString(user.Name)
		if err != nil {
			return nil, err
		}
		// Write grade
		err = binary.Write(&buf, binary.LittleEndian, int32(user.Age))
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}
