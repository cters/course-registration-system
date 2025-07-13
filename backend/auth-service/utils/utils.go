package utils

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GenerateTokenUUID(userId int) string {
	newUUID := uuid.New()

	// convert to string, remove "-"
	uuidString := strings.ReplaceAll((newUUID).String(), "", "")
	return strconv.Itoa(userId) + "token" + uuidString
}