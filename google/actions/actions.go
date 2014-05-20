package actions

import (
	"fmt"
	"github.com/dynport/dgtk/log"
	"github.com/dynport/gocli"
	"github.com/dynport/gocloud/google"
	"os"
)

const (
	CLI_PROJECT       = "--project-id"
	CLI_CLIENT_ID     = "--client-id"
	CLI_CLIENT_SECRET = "--client-secret"
	CLI_ZONE          = "--zone-id"
)

type Arg struct {
	Cli      string
	Json     string
	Required bool
	Default  string
}

type Flags struct {
	ProjectId    string
	ZoneId       string
	ClientId     string
	ClientSecret string
}

func (flags *Flags) ProjectIdRequired() bool {
	return flags.ProjectId == ""
}

func (flags *Flags) ClientIdrequired() bool {
	return flags.ClientId == ""
}

func (flags *Flags) ClientSecretRequired() bool {
	return flags.ClientSecret == ""
}

func (flags *Flags) ZoneIdRequired() bool {
	return flags.ZoneId == ""
}

var flags = &Flags{
	ProjectId:    os.Getenv("GOOGLE_COMPUTE_PROJECT_ID"),
	ClientId:     os.Getenv("GOOGLE_COMPUTE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_COMPUTE_CLIENT_SECRET"),
	ZoneId:       os.Getenv("GOOGLE_COMPUTE_ZONE_ID"),
}

func projectArgs() *gocli.Args {
	args := gocli.NewArgs(nil)
	args.RegisterString(CLI_PROJECT, "project_id", flags.ProjectIdRequired(), flags.ProjectId, "ProjectId")
	args.RegisterString(CLI_CLIENT_ID, "client_id", flags.ClientIdrequired(), flags.ClientId, "ClientId")
	args.RegisterString(CLI_CLIENT_SECRET, "client_secret", flags.ClientSecretRequired(), flags.ClientSecret, "ClientSecret")
	return args
}

func computeFromArgs() (c *google.Compute, e error) {
	client := &google.Client{
		ProjectId:    projectId(),
		ClientId:     os.Getenv(CLI_CLIENT_ID),
		ClientSecret: os.Getenv(CLI_CLIENT_SECRET),
	}
	return client.Compute()
}

func projectId() string {
	return os.Getenv(CLI_PROJECT)
}

func ListImagesHandler() error {
	compute, e := computeFromArgs()
	if e != nil {
		return e
	}
	rsp, e := compute.NewImagesService().List(projectId()).Do()
	if e != nil {
		return e
	}
	log.Info("%#v", rsp)
	table := gocli.NewTable()
	table.Add("Status", "Name", "Status", "Created", "Description")
	for _, zone := range rsp.Items {
		table.Add(zone.Status, zone.Name, zone.Status, zone.CreationTimestamp, zone.Description)
	}
	fmt.Println(table)
	return nil
}

func ListZonesHandler() error {
	compute, e := computeFromArgs()
	if e != nil {
		return e
	}
	rsp, e := compute.NewZonesService().List(projectId()).Do()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	table.Add("Status", "Name", "Created", "Description")
	for _, zone := range rsp.Items {
		table.Add(zone.Status, zone.Name, zone.CreationTimestamp, zone.Description)
	}
	fmt.Println(table)
	return nil
}

type ListMachineTypes struct {
	ProjectId string `cli:"type=arg required=true"`
	ZoneId    string `cli:"type=arg required=true"`
}

func (a *ListMachineTypes) Run() error {
	compute, e := computeFromArgs()
	if e != nil {
		return e
	}
	rsp, e := compute.NewMachineTypesService().List(a.ProjectId, a.ZoneId).Do()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	table.Add("Zone", "Name", "CPUs", "Memory", "ImageSpace", "Created", "Description")
	for _, t := range rsp.Items {
		table.Add(t.Zone, t.Name, t.GuestCpus, t.MemoryMb, t.ImageSpaceGb, t.CreationTimestamp, t.Description)
	}
	fmt.Println(table)
	return nil
}
