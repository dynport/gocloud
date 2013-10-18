package main

import (
	"fmt"
	"github.com/dynport/gocli"
	"github.com/dynport/gocloud/jiffybox"
	"os"
	"strconv"
	"strings"
)

func init() {
	router.Register("jb/backups/list", &gocli.Action{Handler: jiffyBoxListBackupsAction, Description: "List all running boxes"})
	router.Register("jb/backups/create", &gocli.Action{Handler: jiffyBoxCreateBackupAction, Usage: USAGE_CREATE_BACKUP, Description: "Create manual backup from box"})
}

const USAGE_CREATE_BACKUP = "id"

func jiffyBoxCreateBackupAction(args *gocli.Args) error {
	if len(args.Args) != 1 {
		return fmt.Errorf(USAGE_CREATE_BACKUP)
	}
	id, e := strconv.Atoi(args.Args[0])
	if e != nil {
		return fmt.Errorf("unable to use %s as server id", args.Args[0])
	}
	logger.Infof("creating backup for box %d", id)
	e = client().CreateBackup(id)
	if e != nil {
		return e
	}
	logger.Infof("created backup for box %d", id)
	return nil
}

const USAGE_START_SERVER = "id"

func init() {
	router.Register("jb/servers/shutdown", &gocli.Action{Handler: jiffyBoxStopServer, Description: "Shutdown Server", Usage: USAGE_START_SERVER})
}

func jiffyBoxStopServer(args *gocli.Args) error {
	id, e := serverFromArgs(args.Args)
	if e != nil {
		return e
	}
	s, e := client().ShutdownServer(id)
	if e != nil {
		return e
	}
	logger.Infof("started server %d", id)
	printServer(s)
	return nil
}

func init() {
	args := gocli.NewArgs(nil)
	args.RegisterInt(CLI_PLAN_ID, "plan_id", false, DEFAULT_PLAN_ID, "Plan id")
	router.Register("jb/servers/start", &gocli.Action{Handler: jiffyBoxStartServer, Description: "Start Server", Usage: USAGE_START_SERVER, Args: args})
}

func jiffyBoxStartServer(args *gocli.Args) error {
	id, e := serverFromArgs(args.Args)
	if e != nil {
		return e
	}
	planId := args.MustGetInt(CLI_PLAN_ID)
	s, e := client().StartServer(id, planId)
	if e != nil {
		return e
	}
	logger.Infof("started server %d", id)
	printServer(s)
	return nil
}

func init() {
	router.Register("jb/servers/list", &gocli.Action{Handler: jiffyBoxListServersAction})
	router.Register("jb/servers/show", &gocli.Action{Handler: jiffyBoxShowServersAction})

	args := gocli.NewArgs(nil)
	args.RegisterInt(CLI_PLAN_ID, "plan_id", false, DEFAULT_PLAN_ID, "Plan id")
	args.RegisterString(CLI_NAME, "name", false, "", "Name of the new box")
	router.Register("jb/servers/clone", &gocli.Action{Handler: jiffyBoxCloneServerAction, Usage: USAGE_CLONE_SERVER, Args: args})
}

func serverFromArgs(args []string) (id int, e error) {
	if len(args) != 1 {
		return id, fmt.Errorf(USAGE_CREATE_BACKUP)
	}
	id, e = strconv.Atoi(args[0])
	if e != nil {
		return id, fmt.Errorf("unable to use %s as server id", args[0])
	}
	return id, nil
}

func jiffyBoxCloneServerAction(args *gocli.Args) error {
	id, e := serverFromArgs(args.Args)
	if e != nil {
		return e
	}
	opts := &jiffybox.CreateOptions{
		PlanId:   args.MustGetInt(CLI_PLAN_ID),
		Name:     args.MustGetString(CLI_NAME),
		Password: os.Getenv("JIFFYBOX_DEFAULT_PASSWORD"),
	}
	logger.Infof("cloning server %d with %#v", id, opts)
	s, e := client().CloneServer(id, opts)
	if e != nil {
		return e
	}
	logger.Infof("cloned server %d", id)
	printServer(s)
	return nil

}

const USAGE_SHOW_SERVER = "id"

func jiffyBoxShowServersAction(args *gocli.Args) error {
	if len(args.Args) != 1 {
		return fmt.Errorf(USAGE_SHOW_SERVER)
	}
	id, e := strconv.Atoi(args.Args[0])
	if e != nil {
		return fmt.Errorf("unable to use %s as server id", args.Args[0])
	}
	server, e := client().JiffyBox(id)
	if e != nil {
		return e
	}
	printServer(server)
	return nil
}

func printServer(server *jiffybox.Server) {
	table := gocli.NewTable()
	table.Add("Id", server.Id)
	table.Add("Name", server.Name)
	table.Add("Status", server.Status)
	table.Add("Created", server.CreatedAt().Format(TIME_FORMAT))
	table.Add("Host", server.Host)
	table.Add("Running", server.Running)
	table.Add("RecoverymodeActive", server.RecoverymodeActive)
	table.Add("Plan", server.Plan.Id)
	table.Add("Cpu", server.Plan.Cpus)
	table.Add("RAM", server.Plan.RamInMB)
	table.Add("IsBeingCopied", server.IsBeingCopied)
	table.Add("ManualBackupRunning", server.ManualBackupRunning)
	if server.ActiveProfile != nil {
		table.Add("Profile Name", server.ActiveProfile.Name)
		table.Add("Profile Kernel", server.ActiveProfile.Kernel)
	}
	i := 0
	for k, v := range server.Ips {
		key := ""
		if i == 0 {
			key = "Ips"
			i++
		}
		table.Add(key, k+": "+strings.Join(v, ", "))
	}
	fmt.Println(table)
}

func init() {
	router.Register("jb/plans/list", &gocli.Action{Handler: jiffyBoxListPlansAction})
	router.Register("jb/distributions/list", &gocli.Action{Handler: jiffyBoxListDistributionsAction})
}

func init() {
	router.Register("jb/servers/delete", &gocli.Action{Handler: jiffyBoxDeleteAction})
}

func jiffyBoxListBackupsAction(args *gocli.Args) error {
	backups, e := client().Backups()
	if e != nil {
		return e
	}

	for _, backup := range backups {
		fmt.Println(backup.ServerId, backup.Key, backup.CreatedAt())
	}
	return nil
}

const USAGE_DELETE = "id"

func jiffyBoxDeleteAction(args *gocli.Args) error {
	if len(args.Args) != 1 {
		return fmt.Errorf(USAGE_DELETE)
	}
	id := args.Args[0]
	logger.Infof("deleting box with id %s", id)
	e := client().DeleteJiffyBox(id)
	if e != nil {
		return e
	}
	logger.Info("deleted box")
	return nil
}

//client().DeleteJiffyBox(id)

func client() *jiffybox.Client {
	return jiffybox.NewFromEnv()
}

func jiffyBoxListDistributionsAction(args *gocli.Args) error {
	distributions, e := client().Distributions()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	table.Add("Key", "Name", "Min Disk Size", "Default Kernel")
	for _, distribution := range distributions {
		table.Add(distribution.Key, distribution.Name, distribution.MinDiskSizeMB, distribution.DefaultKernel)
	}
	fmt.Println(table)
	return nil
}

const (
	CLI_NAME             = "--name"
	CLI_PLAN_ID          = "--plan-id"
	CLI_DISTRIBUTION     = "--distribution"
	DEFAULT_PLAN_ID      = 20
	DEFAULT_DISTRIBUTION = "ubuntu_12_4_lts_64bit"
	USAGE_CLONE_SERVER   = "id"
	HOURS_PER_MONTH      = 365 * 24.0 / 12.0
)

func init() {
	args := gocli.NewArgs(nil)
	args.RegisterString(CLI_NAME, "name", false, "", "Name of the new box")
	args.RegisterInt(CLI_PLAN_ID, "plan_id", false, DEFAULT_PLAN_ID, "Plan id")
	args.RegisterString(CLI_DISTRIBUTION, "distribution", false, DEFAULT_DISTRIBUTION, "Distribution")
	router.Register("jb/servers/create", &gocli.Action{Handler: jiffyBoxCreateAction, Args: args, Description: "Create new JiffyBox"})
}

func jiffyBoxListPlansAction(args *gocli.Args) error {
	plans, e := client().Plans()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	table.Add("Id", "Name", "Cpu", "Ram", "Disk", "Price/Hour", "Price/Month")
	for _, plan := range plans {
		table.Add(
			plan.Id, plan.Name, plan.Cpus, plan.RamInMB, plan.DiskSizeInMB,
			fmt.Sprintf("%.02f €", plan.PricePerHour),
			fmt.Sprintf("%.2f €", plan.PricePerHour*HOURS_PER_MONTH),
		)
	}
	fmt.Println(table)
	return nil
}

func jiffyBoxCreateAction(args *gocli.Args) error {
	logger.Infof("creating new jiffybox")
	opts := &jiffybox.CreateOptions{
		Name:         args.MustGetString(CLI_NAME),
		PlanId:       args.MustGetInt(CLI_PLAN_ID),
		Distribution: args.MustGetString(CLI_DISTRIBUTION),
		UseSshKey:    true,
		Password: os.Getenv("JIFFYBOX_DEFAULT_PASSWORD"),
	}
	s, e := client().CreateJiffyBox(opts)
	if e != nil {
		return e
	}
	fmt.Println("created server!")
	printServer(s)
	return nil
}

const TIME_FORMAT = "2006-01-02T15:04:05"

func jiffyBoxListServersAction(args *gocli.Args) error {
	servers, e := client().JiffyBoxes()
	if e != nil {
		return e
	}
	if len(servers) == 0 {
		fmt.Println("no boxes found")
		return nil
	}
	table := gocli.NewTable()
	table.Add("Created", "Id", "Status", "Running", "Name", "Cpu", "RAM", "Ip")
	for _, server := range servers {
		table.Add(server.CreatedAt().Format(TIME_FORMAT), server.Id, server.Status, server.Running, server.Name, server.Plan.Cpus, server.Plan.RamInMB, server.PublicIp())
	}
	fmt.Println(table)
	return nil
}
