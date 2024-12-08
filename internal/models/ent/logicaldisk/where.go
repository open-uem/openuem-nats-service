// Code generated by ent, DO NOT EDIT.

package logicaldisk

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/doncicuto/openuem_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldLTE(FieldID, id))
}

// Filesystem applies equality check predicate on the "filesystem" field. It's identical to FilesystemEQ.
func Filesystem(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEQ(FieldFilesystem, v))
}

// Usage applies equality check predicate on the "usage" field. It's identical to UsageEQ.
func Usage(v int8) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEQ(FieldUsage, v))
}

// SizeInUnits applies equality check predicate on the "size_in_units" field. It's identical to SizeInUnitsEQ.
func SizeInUnits(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEQ(FieldSizeInUnits, v))
}

// RemainingSpaceInUnits applies equality check predicate on the "remaining_space_in_units" field. It's identical to RemainingSpaceInUnitsEQ.
func RemainingSpaceInUnits(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEQ(FieldRemainingSpaceInUnits, v))
}

// VolumeName applies equality check predicate on the "volume_name" field. It's identical to VolumeNameEQ.
func VolumeName(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEQ(FieldVolumeName, v))
}

// BitlockerStatus applies equality check predicate on the "bitlocker_status" field. It's identical to BitlockerStatusEQ.
func BitlockerStatus(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEQ(FieldBitlockerStatus, v))
}

// LabelEQ applies the EQ predicate on the "label" field.
func LabelEQ(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEQ(FieldLabel, v))
}

// LabelNEQ applies the NEQ predicate on the "label" field.
func LabelNEQ(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNEQ(FieldLabel, v))
}

// LabelIn applies the In predicate on the "label" field.
func LabelIn(vs ...string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldIn(FieldLabel, vs...))
}

// LabelNotIn applies the NotIn predicate on the "label" field.
func LabelNotIn(vs ...string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNotIn(FieldLabel, vs...))
}

// LabelGT applies the GT predicate on the "label" field.
func LabelGT(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldGT(FieldLabel, v))
}

// LabelGTE applies the GTE predicate on the "label" field.
func LabelGTE(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldGTE(FieldLabel, v))
}

// LabelLT applies the LT predicate on the "label" field.
func LabelLT(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldLT(FieldLabel, v))
}

// LabelLTE applies the LTE predicate on the "label" field.
func LabelLTE(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldLTE(FieldLabel, v))
}

// LabelContains applies the Contains predicate on the "label" field.
func LabelContains(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldContains(FieldLabel, v))
}

// LabelHasPrefix applies the HasPrefix predicate on the "label" field.
func LabelHasPrefix(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldHasPrefix(FieldLabel, v))
}

// LabelHasSuffix applies the HasSuffix predicate on the "label" field.
func LabelHasSuffix(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldHasSuffix(FieldLabel, v))
}

// LabelEqualFold applies the EqualFold predicate on the "label" field.
func LabelEqualFold(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEqualFold(FieldLabel, v))
}

// LabelContainsFold applies the ContainsFold predicate on the "label" field.
func LabelContainsFold(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldContainsFold(FieldLabel, v))
}

// FilesystemEQ applies the EQ predicate on the "filesystem" field.
func FilesystemEQ(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEQ(FieldFilesystem, v))
}

// FilesystemNEQ applies the NEQ predicate on the "filesystem" field.
func FilesystemNEQ(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNEQ(FieldFilesystem, v))
}

// FilesystemIn applies the In predicate on the "filesystem" field.
func FilesystemIn(vs ...string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldIn(FieldFilesystem, vs...))
}

// FilesystemNotIn applies the NotIn predicate on the "filesystem" field.
func FilesystemNotIn(vs ...string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNotIn(FieldFilesystem, vs...))
}

// FilesystemGT applies the GT predicate on the "filesystem" field.
func FilesystemGT(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldGT(FieldFilesystem, v))
}

// FilesystemGTE applies the GTE predicate on the "filesystem" field.
func FilesystemGTE(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldGTE(FieldFilesystem, v))
}

// FilesystemLT applies the LT predicate on the "filesystem" field.
func FilesystemLT(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldLT(FieldFilesystem, v))
}

// FilesystemLTE applies the LTE predicate on the "filesystem" field.
func FilesystemLTE(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldLTE(FieldFilesystem, v))
}

// FilesystemContains applies the Contains predicate on the "filesystem" field.
func FilesystemContains(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldContains(FieldFilesystem, v))
}

// FilesystemHasPrefix applies the HasPrefix predicate on the "filesystem" field.
func FilesystemHasPrefix(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldHasPrefix(FieldFilesystem, v))
}

// FilesystemHasSuffix applies the HasSuffix predicate on the "filesystem" field.
func FilesystemHasSuffix(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldHasSuffix(FieldFilesystem, v))
}

// FilesystemIsNil applies the IsNil predicate on the "filesystem" field.
func FilesystemIsNil() predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldIsNull(FieldFilesystem))
}

// FilesystemNotNil applies the NotNil predicate on the "filesystem" field.
func FilesystemNotNil() predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNotNull(FieldFilesystem))
}

// FilesystemEqualFold applies the EqualFold predicate on the "filesystem" field.
func FilesystemEqualFold(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEqualFold(FieldFilesystem, v))
}

// FilesystemContainsFold applies the ContainsFold predicate on the "filesystem" field.
func FilesystemContainsFold(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldContainsFold(FieldFilesystem, v))
}

// UsageEQ applies the EQ predicate on the "usage" field.
func UsageEQ(v int8) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEQ(FieldUsage, v))
}

// UsageNEQ applies the NEQ predicate on the "usage" field.
func UsageNEQ(v int8) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNEQ(FieldUsage, v))
}

// UsageIn applies the In predicate on the "usage" field.
func UsageIn(vs ...int8) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldIn(FieldUsage, vs...))
}

// UsageNotIn applies the NotIn predicate on the "usage" field.
func UsageNotIn(vs ...int8) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNotIn(FieldUsage, vs...))
}

// UsageGT applies the GT predicate on the "usage" field.
func UsageGT(v int8) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldGT(FieldUsage, v))
}

// UsageGTE applies the GTE predicate on the "usage" field.
func UsageGTE(v int8) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldGTE(FieldUsage, v))
}

// UsageLT applies the LT predicate on the "usage" field.
func UsageLT(v int8) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldLT(FieldUsage, v))
}

// UsageLTE applies the LTE predicate on the "usage" field.
func UsageLTE(v int8) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldLTE(FieldUsage, v))
}

// SizeInUnitsEQ applies the EQ predicate on the "size_in_units" field.
func SizeInUnitsEQ(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEQ(FieldSizeInUnits, v))
}

// SizeInUnitsNEQ applies the NEQ predicate on the "size_in_units" field.
func SizeInUnitsNEQ(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNEQ(FieldSizeInUnits, v))
}

// SizeInUnitsIn applies the In predicate on the "size_in_units" field.
func SizeInUnitsIn(vs ...string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldIn(FieldSizeInUnits, vs...))
}

// SizeInUnitsNotIn applies the NotIn predicate on the "size_in_units" field.
func SizeInUnitsNotIn(vs ...string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNotIn(FieldSizeInUnits, vs...))
}

// SizeInUnitsGT applies the GT predicate on the "size_in_units" field.
func SizeInUnitsGT(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldGT(FieldSizeInUnits, v))
}

// SizeInUnitsGTE applies the GTE predicate on the "size_in_units" field.
func SizeInUnitsGTE(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldGTE(FieldSizeInUnits, v))
}

// SizeInUnitsLT applies the LT predicate on the "size_in_units" field.
func SizeInUnitsLT(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldLT(FieldSizeInUnits, v))
}

// SizeInUnitsLTE applies the LTE predicate on the "size_in_units" field.
func SizeInUnitsLTE(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldLTE(FieldSizeInUnits, v))
}

// SizeInUnitsContains applies the Contains predicate on the "size_in_units" field.
func SizeInUnitsContains(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldContains(FieldSizeInUnits, v))
}

// SizeInUnitsHasPrefix applies the HasPrefix predicate on the "size_in_units" field.
func SizeInUnitsHasPrefix(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldHasPrefix(FieldSizeInUnits, v))
}

// SizeInUnitsHasSuffix applies the HasSuffix predicate on the "size_in_units" field.
func SizeInUnitsHasSuffix(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldHasSuffix(FieldSizeInUnits, v))
}

// SizeInUnitsIsNil applies the IsNil predicate on the "size_in_units" field.
func SizeInUnitsIsNil() predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldIsNull(FieldSizeInUnits))
}

// SizeInUnitsNotNil applies the NotNil predicate on the "size_in_units" field.
func SizeInUnitsNotNil() predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNotNull(FieldSizeInUnits))
}

// SizeInUnitsEqualFold applies the EqualFold predicate on the "size_in_units" field.
func SizeInUnitsEqualFold(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEqualFold(FieldSizeInUnits, v))
}

// SizeInUnitsContainsFold applies the ContainsFold predicate on the "size_in_units" field.
func SizeInUnitsContainsFold(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldContainsFold(FieldSizeInUnits, v))
}

// RemainingSpaceInUnitsEQ applies the EQ predicate on the "remaining_space_in_units" field.
func RemainingSpaceInUnitsEQ(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEQ(FieldRemainingSpaceInUnits, v))
}

// RemainingSpaceInUnitsNEQ applies the NEQ predicate on the "remaining_space_in_units" field.
func RemainingSpaceInUnitsNEQ(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNEQ(FieldRemainingSpaceInUnits, v))
}

// RemainingSpaceInUnitsIn applies the In predicate on the "remaining_space_in_units" field.
func RemainingSpaceInUnitsIn(vs ...string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldIn(FieldRemainingSpaceInUnits, vs...))
}

// RemainingSpaceInUnitsNotIn applies the NotIn predicate on the "remaining_space_in_units" field.
func RemainingSpaceInUnitsNotIn(vs ...string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNotIn(FieldRemainingSpaceInUnits, vs...))
}

// RemainingSpaceInUnitsGT applies the GT predicate on the "remaining_space_in_units" field.
func RemainingSpaceInUnitsGT(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldGT(FieldRemainingSpaceInUnits, v))
}

// RemainingSpaceInUnitsGTE applies the GTE predicate on the "remaining_space_in_units" field.
func RemainingSpaceInUnitsGTE(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldGTE(FieldRemainingSpaceInUnits, v))
}

// RemainingSpaceInUnitsLT applies the LT predicate on the "remaining_space_in_units" field.
func RemainingSpaceInUnitsLT(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldLT(FieldRemainingSpaceInUnits, v))
}

// RemainingSpaceInUnitsLTE applies the LTE predicate on the "remaining_space_in_units" field.
func RemainingSpaceInUnitsLTE(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldLTE(FieldRemainingSpaceInUnits, v))
}

// RemainingSpaceInUnitsContains applies the Contains predicate on the "remaining_space_in_units" field.
func RemainingSpaceInUnitsContains(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldContains(FieldRemainingSpaceInUnits, v))
}

// RemainingSpaceInUnitsHasPrefix applies the HasPrefix predicate on the "remaining_space_in_units" field.
func RemainingSpaceInUnitsHasPrefix(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldHasPrefix(FieldRemainingSpaceInUnits, v))
}

// RemainingSpaceInUnitsHasSuffix applies the HasSuffix predicate on the "remaining_space_in_units" field.
func RemainingSpaceInUnitsHasSuffix(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldHasSuffix(FieldRemainingSpaceInUnits, v))
}

// RemainingSpaceInUnitsIsNil applies the IsNil predicate on the "remaining_space_in_units" field.
func RemainingSpaceInUnitsIsNil() predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldIsNull(FieldRemainingSpaceInUnits))
}

// RemainingSpaceInUnitsNotNil applies the NotNil predicate on the "remaining_space_in_units" field.
func RemainingSpaceInUnitsNotNil() predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNotNull(FieldRemainingSpaceInUnits))
}

// RemainingSpaceInUnitsEqualFold applies the EqualFold predicate on the "remaining_space_in_units" field.
func RemainingSpaceInUnitsEqualFold(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEqualFold(FieldRemainingSpaceInUnits, v))
}

// RemainingSpaceInUnitsContainsFold applies the ContainsFold predicate on the "remaining_space_in_units" field.
func RemainingSpaceInUnitsContainsFold(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldContainsFold(FieldRemainingSpaceInUnits, v))
}

// VolumeNameEQ applies the EQ predicate on the "volume_name" field.
func VolumeNameEQ(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEQ(FieldVolumeName, v))
}

// VolumeNameNEQ applies the NEQ predicate on the "volume_name" field.
func VolumeNameNEQ(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNEQ(FieldVolumeName, v))
}

// VolumeNameIn applies the In predicate on the "volume_name" field.
func VolumeNameIn(vs ...string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldIn(FieldVolumeName, vs...))
}

// VolumeNameNotIn applies the NotIn predicate on the "volume_name" field.
func VolumeNameNotIn(vs ...string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNotIn(FieldVolumeName, vs...))
}

// VolumeNameGT applies the GT predicate on the "volume_name" field.
func VolumeNameGT(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldGT(FieldVolumeName, v))
}

// VolumeNameGTE applies the GTE predicate on the "volume_name" field.
func VolumeNameGTE(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldGTE(FieldVolumeName, v))
}

// VolumeNameLT applies the LT predicate on the "volume_name" field.
func VolumeNameLT(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldLT(FieldVolumeName, v))
}

// VolumeNameLTE applies the LTE predicate on the "volume_name" field.
func VolumeNameLTE(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldLTE(FieldVolumeName, v))
}

// VolumeNameContains applies the Contains predicate on the "volume_name" field.
func VolumeNameContains(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldContains(FieldVolumeName, v))
}

// VolumeNameHasPrefix applies the HasPrefix predicate on the "volume_name" field.
func VolumeNameHasPrefix(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldHasPrefix(FieldVolumeName, v))
}

// VolumeNameHasSuffix applies the HasSuffix predicate on the "volume_name" field.
func VolumeNameHasSuffix(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldHasSuffix(FieldVolumeName, v))
}

// VolumeNameIsNil applies the IsNil predicate on the "volume_name" field.
func VolumeNameIsNil() predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldIsNull(FieldVolumeName))
}

// VolumeNameNotNil applies the NotNil predicate on the "volume_name" field.
func VolumeNameNotNil() predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNotNull(FieldVolumeName))
}

// VolumeNameEqualFold applies the EqualFold predicate on the "volume_name" field.
func VolumeNameEqualFold(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEqualFold(FieldVolumeName, v))
}

// VolumeNameContainsFold applies the ContainsFold predicate on the "volume_name" field.
func VolumeNameContainsFold(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldContainsFold(FieldVolumeName, v))
}

// BitlockerStatusEQ applies the EQ predicate on the "bitlocker_status" field.
func BitlockerStatusEQ(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEQ(FieldBitlockerStatus, v))
}

// BitlockerStatusNEQ applies the NEQ predicate on the "bitlocker_status" field.
func BitlockerStatusNEQ(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNEQ(FieldBitlockerStatus, v))
}

// BitlockerStatusIn applies the In predicate on the "bitlocker_status" field.
func BitlockerStatusIn(vs ...string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldIn(FieldBitlockerStatus, vs...))
}

// BitlockerStatusNotIn applies the NotIn predicate on the "bitlocker_status" field.
func BitlockerStatusNotIn(vs ...string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNotIn(FieldBitlockerStatus, vs...))
}

// BitlockerStatusGT applies the GT predicate on the "bitlocker_status" field.
func BitlockerStatusGT(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldGT(FieldBitlockerStatus, v))
}

// BitlockerStatusGTE applies the GTE predicate on the "bitlocker_status" field.
func BitlockerStatusGTE(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldGTE(FieldBitlockerStatus, v))
}

// BitlockerStatusLT applies the LT predicate on the "bitlocker_status" field.
func BitlockerStatusLT(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldLT(FieldBitlockerStatus, v))
}

// BitlockerStatusLTE applies the LTE predicate on the "bitlocker_status" field.
func BitlockerStatusLTE(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldLTE(FieldBitlockerStatus, v))
}

// BitlockerStatusContains applies the Contains predicate on the "bitlocker_status" field.
func BitlockerStatusContains(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldContains(FieldBitlockerStatus, v))
}

// BitlockerStatusHasPrefix applies the HasPrefix predicate on the "bitlocker_status" field.
func BitlockerStatusHasPrefix(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldHasPrefix(FieldBitlockerStatus, v))
}

// BitlockerStatusHasSuffix applies the HasSuffix predicate on the "bitlocker_status" field.
func BitlockerStatusHasSuffix(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldHasSuffix(FieldBitlockerStatus, v))
}

// BitlockerStatusIsNil applies the IsNil predicate on the "bitlocker_status" field.
func BitlockerStatusIsNil() predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldIsNull(FieldBitlockerStatus))
}

// BitlockerStatusNotNil applies the NotNil predicate on the "bitlocker_status" field.
func BitlockerStatusNotNil() predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldNotNull(FieldBitlockerStatus))
}

// BitlockerStatusEqualFold applies the EqualFold predicate on the "bitlocker_status" field.
func BitlockerStatusEqualFold(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldEqualFold(FieldBitlockerStatus, v))
}

// BitlockerStatusContainsFold applies the ContainsFold predicate on the "bitlocker_status" field.
func BitlockerStatusContainsFold(v string) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.FieldContainsFold(FieldBitlockerStatus, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.LogicalDisk {
	return predicate.LogicalDisk(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Agent) predicate.LogicalDisk {
	return predicate.LogicalDisk(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LogicalDisk) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LogicalDisk) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LogicalDisk) predicate.LogicalDisk {
	return predicate.LogicalDisk(sql.NotPredicates(p))
}
