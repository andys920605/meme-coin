package meme_coin

import (
	"strconv"

	"github.com/andys920605/meme-coin/pkg/errors"
	"github.com/andys920605/meme-coin/pkg/snowflake"
)

type ID int64

func NewID() ID {
	return ID(snowflake.New())
}

func ParseID(id string) (ID, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "convert to int64")
	}
	return ID(i), nil
}

func (i ID) Int64() int64 {
	return int64(i)
}

func (i ID) String() string {
	return strconv.FormatInt(i.Int64(), 10)
}
