package stypes

import (
	"encoding/xml"
	"errors"
)

type PTabLeader string

const (
	PTabLeaderNone       PTabLeader = "none"
	PTabLeaderDot        PTabLeader = "dot"
	PTabLeaderHyphen     PTabLeader = "hyphen"
	PTabLeaderUnderscore PTabLeader = "underscore"
	PTabLeaderMiddleDot  PTabLeader = "middleDot"
)

func PTabLeaderFromStr(value string) (PTabLeader, error) {
	switch value {
	case "none":
		return PTabLeaderNone, nil
	case "dot":
		return PTabLeaderDot, nil
	case "hyphen":
		return PTabLeaderHyphen, nil
	case "underscore":
		return PTabLeaderUnderscore, nil
	case "middleDot":
		return PTabLeaderMiddleDot, nil
	default:
		return "", errors.New("invalid PTabLeader value")
	}
}

func (l *PTabLeader) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := PTabLeaderFromStr(attr.Value)
	if err != nil {
		return err
	}

	*l = val
	return nil
}

type PTabRelativeTo string

const (
	PTabRelativeToMargin PTabRelativeTo = "margin"
	PTabRelativeToIndent PTabRelativeTo = "indent"
)

func PTabRelativeToFromStr(value string) (PTabRelativeTo, error) {
	switch value {
	case "margin":
		return PTabRelativeToMargin, nil
	case "indent":
		return PTabRelativeToIndent, nil
	default:
		return "", errors.New("invalid PTabRelativeTo value")
	}
}

func (r *PTabRelativeTo) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := PTabRelativeToFromStr(attr.Value)
	if err != nil {
		return err
	}

	*r = val
	return nil
}

type PTabAlignment string

const (
	PTabAlignmentLeft   PTabAlignment = "left"
	PTabAlignmentCenter PTabAlignment = "center"
	PTabAlignmentRight  PTabAlignment = "right"
)

func PTabAlignmentFromStr(value string) (PTabAlignment, error) {
	switch value {
	case "left":
		return PTabAlignmentLeft, nil
	case "center":
		return PTabAlignmentCenter, nil
	case "right":
		return PTabAlignmentRight, nil
	default:
		return "", errors.New("invalid PTabAlignment value")
	}
}

func (a *PTabAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := PTabAlignmentFromStr(attr.Value)
	if err != nil {
		return err
	}

	*a = val
	return nil
}
