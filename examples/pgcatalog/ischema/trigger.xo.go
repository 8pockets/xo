// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/8pockets/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// Trigger represents a row from 'information_schema.triggers'.
type Trigger struct {
	TriggerCatalog          pgtypes.SQLIdentifier  `json:"trigger_catalog"`            // trigger_catalog
	TriggerSchema           pgtypes.SQLIdentifier  `json:"trigger_schema"`             // trigger_schema
	TriggerName             pgtypes.SQLIdentifier  `json:"trigger_name"`               // trigger_name
	EventManipulation       pgtypes.CharacterData  `json:"event_manipulation"`         // event_manipulation
	EventObjectCatalog      pgtypes.SQLIdentifier  `json:"event_object_catalog"`       // event_object_catalog
	EventObjectSchema       pgtypes.SQLIdentifier  `json:"event_object_schema"`        // event_object_schema
	EventObjectTable        pgtypes.SQLIdentifier  `json:"event_object_table"`         // event_object_table
	ActionOrder             pgtypes.CardinalNumber `json:"action_order"`               // action_order
	ActionCondition         pgtypes.CharacterData  `json:"action_condition"`           // action_condition
	ActionStatement         pgtypes.CharacterData  `json:"action_statement"`           // action_statement
	ActionOrientation       pgtypes.CharacterData  `json:"action_orientation"`         // action_orientation
	ActionTiming            pgtypes.CharacterData  `json:"action_timing"`              // action_timing
	ActionReferenceOldTable pgtypes.SQLIdentifier  `json:"action_reference_old_table"` // action_reference_old_table
	ActionReferenceNewTable pgtypes.SQLIdentifier  `json:"action_reference_new_table"` // action_reference_new_table
	ActionReferenceOldRow   pgtypes.SQLIdentifier  `json:"action_reference_old_row"`   // action_reference_old_row
	ActionReferenceNewRow   pgtypes.SQLIdentifier  `json:"action_reference_new_row"`   // action_reference_new_row
	Created                 pgtypes.TimeStamp      `json:"created"`                    // created
}
