package v1

//go:generate oapi-codegen -config=../deepmap.cfg.yaml -package "${GOPACKAGE}" "${APIFILE}"

//nolint:revive //needed for generating types.
import _ "github.com/deepmap/oapi-codegen/pkg/types"
