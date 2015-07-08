package main

type dropletReboot struct {
	ID string `cli:"arg required"`
}

func (r *dropletReboot) Run() error {
	cl, err := client()
	if err != nil {
		return err
	}
	return cl.RebootDroplet(r.ID)
}
