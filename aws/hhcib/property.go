package main

import (
	"log"
	"strings"
)

type Property struct {
	Description string
	Default     string
	Name        string
	Required    bool
	Type        string
	List        bool
	ValidValues []string
}

var null = struct{}{}

var uncountable = map[string]struct{}{
	"PricingDetails":             null,
	"InstanceCounts":             null,
	"Details":                    null,
	"ProductCodes":               null,
	"Instances":                  null,
	"PrivateIpAddresses":         null,
	"IpRanges":                   null,
	"Events":                     null,
	"PriceSchedules":             null,
	"ModificationResults":        null,
	"RecurringCharges":           null,
	"ReservedInstancesOfferings": null,
	"Volumes":                    null,
	"PrivateIpAddress":           null,
	"IpPermissions":              null,
}

func pluralize(name string) string {
	name = strings.TrimSuffix(name, "Set")
	_, uncountable := uncountable[name]
	if uncountable {
		return name
	}
	if strings.HasSuffix(name, "s") {
		log.Print("ERROR: " + name)
	}
	return name + "s"
}

func normalizeCustomType(t string) string {
	for _, suffix := range []string{"Type"} {
		t = strings.TrimSuffix(t, suffix)
	}
	return t
}

type Comments map[string]string

func (property *Property) toTypeDefinition() *TypeField {
	typ := strings.TrimSuffix(property.Type, ".")
	name := strings.TrimSuffix(property.Name, ".item")
	if strings.HasSuffix(typ, "SetType") {
		return newTypeField(pluralize(normalizeName(name)), "[]*"+normalizeCustomType(typ), Comments{"xml": name + ",omitempty"})
	} else if strings.HasSuffix(typ, "ItemType") {
		return newTypeField(pluralize(normalizeName(name)), "[]*"+normalizeCustomType(typ), Comments{"xml": name + ">item,omitempty"})
	}
	return newTypeField(normalizeName(property.Name), convertType(typ), Comments{"xml": name + ",omitempty"})
}

func normalizeName(name string) string {
	name = strings.TrimSuffix(name, ".item")
	name = strings.Replace(name, ".n", "", -1)
	name = strings.Replace(name, ".", "", -1)
	name = strings.Title(name)
	return name
}

func (p *Property) toRequestTypeDefinition() string {
	if strings.HasSuffix(p.Name, ".n") {
		name := strings.TrimSuffix(p.Name, ".n")
		return newTypeField(pluralize(name), "[]"+convertType(p.Type), Comments{"aws": name}).String()
	}
	return newTypeField(normalizeName(p.Name), convertType(p.Type), Comments{"aws": p.Name}).String()
}
