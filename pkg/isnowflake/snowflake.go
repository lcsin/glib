package isnowflake

import "github.com/bwmarrin/snowflake"

func New(workerId int64) (*snowflake.Node, error) {
	node, err := snowflake.NewNode(workerId)
	return node, err
}
