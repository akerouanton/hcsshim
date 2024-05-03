//go:build windows

package hcsshim

import (
	"context"
	"crypto/sha1"
	"path/filepath"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/Microsoft/hcsshim/internal/wclayer"
)

func layerPath(info *DriverInfo, id string) string {
	return filepath.Join(info.HomeDir, id)
}

func ActivateLayer(ctx context.Context, info DriverInfo, id string) error {
	return wclayer.ActivateLayer(ctx, layerPath(&info, id))
}
func CreateLayer(ctx context.Context, info DriverInfo, id, parent string) error {
	return wclayer.CreateLayer(ctx, layerPath(&info, id), parent)
}

// New clients should use CreateScratchLayer instead. Kept in to preserve API compatibility.
func CreateSandboxLayer(ctx context.Context, info DriverInfo, layerId, parentId string, parentLayerPaths []string) error {
	return wclayer.CreateScratchLayer(ctx, layerPath(&info, layerId), parentLayerPaths)
}
func CreateScratchLayer(ctx context.Context, info DriverInfo, layerId, parentId string, parentLayerPaths []string) error {
	return wclayer.CreateScratchLayer(ctx, layerPath(&info, layerId), parentLayerPaths)
}
func DeactivateLayer(ctx context.Context, info DriverInfo, id string) error {
	return wclayer.DeactivateLayer(ctx, layerPath(&info, id))
}

func DestroyLayer(ctx context.Context, info DriverInfo, id string) error {
	return wclayer.DestroyLayer(ctx, layerPath(&info, id))
}

// New clients should use ExpandScratchSize instead. Kept in to preserve API compatibility.
func ExpandSandboxSize(ctx context.Context, info DriverInfo, layerId string, size uint64) error {
	return wclayer.ExpandScratchSize(ctx, layerPath(&info, layerId), size)
}
func ExpandScratchSize(ctx context.Context, info DriverInfo, layerId string, size uint64) error {
	return wclayer.ExpandScratchSize(ctx, layerPath(&info, layerId), size)
}
func ExportLayer(ctx context.Context, info DriverInfo, layerId string, exportFolderPath string, parentLayerPaths []string) error {
	return wclayer.ExportLayer(ctx, layerPath(&info, layerId), exportFolderPath, parentLayerPaths)
}
func GetLayerMountPath(ctx context.Context, info DriverInfo, id string) (string, error) {
	return wclayer.GetLayerMountPath(ctx, layerPath(&info, id))
}
func GetSharedBaseImages(ctx context.Context) (imageData string, err error) {
	return wclayer.GetSharedBaseImages(ctx)
}
func ImportLayer(ctx context.Context, info DriverInfo, layerID string, importFolderPath string, parentLayerPaths []string) error {
	return wclayer.ImportLayer(ctx, layerPath(&info, layerID), importFolderPath, parentLayerPaths)
}
func LayerExists(ctx context.Context, info DriverInfo, id string) (bool, error) {
	return wclayer.LayerExists(ctx, layerPath(&info, id))
}
func PrepareLayer(ctx context.Context, info DriverInfo, layerId string, parentLayerPaths []string) error {
	return wclayer.PrepareLayer(ctx, layerPath(&info, layerId), parentLayerPaths)
}
func ProcessBaseLayer(ctx context.Context, path string) error {
	return wclayer.ProcessBaseLayer(ctx, path)
}
func ProcessUtilityVMImage(ctx context.Context, path string) error {
	return wclayer.ProcessUtilityVMImage(ctx, path)
}
func UnprepareLayer(ctx context.Context, info DriverInfo, layerId string) error {
	return wclayer.UnprepareLayer(ctx, layerPath(&info, layerId))
}
func ConvertToBaseLayer(ctx context.Context, path string) error {
	return wclayer.ConvertToBaseLayer(ctx, path)
}

type DriverInfo struct {
	Flavour int
	HomeDir string
}

type GUID [16]byte

func NameToGuid(ctx context.Context, name string) (id GUID, err error) {
	g, err := wclayer.NameToGuid(ctx, name)
	return g.ToWindowsArray(), err
}

func NewGUID(source string) *GUID {
	h := sha1.Sum([]byte(source))
	var g GUID
	copy(g[0:], h[0:16])
	return &g
}

func (g *GUID) ToString() string {
	return guid.FromWindowsArray(*g).String()
}

type LayerReader = wclayer.LayerReader

func NewLayerReader(ctx context.Context, info DriverInfo, layerID string, parentLayerPaths []string) (LayerReader, error) {
	return wclayer.NewLayerReader(ctx, layerPath(&info, layerID), parentLayerPaths)
}

type LayerWriter = wclayer.LayerWriter

func NewLayerWriter(ctx context.Context, info DriverInfo, layerID string, parentLayerPaths []string) (LayerWriter, error) {
	return wclayer.NewLayerWriter(ctx, layerPath(&info, layerID), parentLayerPaths)
}

type WC_LAYER_DESCRIPTOR = wclayer.WC_LAYER_DESCRIPTOR
