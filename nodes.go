package proxmox

import "fmt"

func (c *Client) Nodes() (ns NodeStatuses, err error) {
	return ns, c.Get("/nodes", &ns)
}

func (c *Client) Node(name string) (*Node, error) {
	var node Node
	if err := c.Get(fmt.Sprintf("/nodes/%s/status", name), &node); err != nil {
		return nil, err
	}
	node.Name = name
	node.client = c

	return &node, nil
}

func (n *Node) Version() (version *Version, err error) {
	return version, n.client.Get("/nodes/%s/version", &version)
}

func (n *Node) VirtualMachines() (vms VirtualMachines, err error) {
	return vms, n.client.Get(fmt.Sprintf("/nodes/%s/qemu", n.Name), &vms)
}

func (n *Node) VirtualMachine(vmid int) (vm *VirtualMachine, err error) {
	return vm, n.client.Get(fmt.Sprintf("/nodes/%s/qemu/%d/status/current", n.Name, vmid), &vm)
}

func (n *Node) Containers() (c Containers, err error) {
	if err := n.client.Get(fmt.Sprintf("/nodes/%s/lxc", n.Name), &c); err != nil {
		return nil, err
	}

	for _, container := range c {
		container.client = n.client
		container.Node = n.Name
	}

	return c, nil
}

func (n *Node) Container(vmid int) (*Container, error) {
	var c Container
	if err := n.client.Get(fmt.Sprintf("/nodes/%s/lxc/%d/status/current", n.Name, vmid), &c); err != nil {
		return nil, err
	}
	c.client = n.client
	c.Node = n.Name

	return &c, nil
}

func (n *Node) NewContainer(t NewContainer) (string, error) {
	var s string
	if err := n.client.Post(fmt.Sprintf("/nodes/%s/lxc", n.Name), t, &s); err != nil {
		return "", err
	}

	return s, nil
}

func (n *Node) Appliances() (appliances Appliances, err error) {
	err = n.client.Get(fmt.Sprintf("/nodes/%s/aplinfo", n.Name), &appliances)
	if err != nil {
		return appliances, err
	}

	for _, t := range appliances {
		t.client = n.client
		t.Node = n.Name
	}

	return appliances, nil
}

func (n *Node) DownloadAppliance(template, storage string) (ret string, err error) {
	return ret, n.client.Post(fmt.Sprintf("/nodes/%s/aplinfo", n.Name), map[string]string{
		"template": template,
		"storage":  storage,
	}, &ret)
}

func (n *Node) VzTmpls(storage string) (templates VzTmpls, err error) {
	return templates, n.client.Get(fmt.Sprintf("/nodes/%s/storage/%s/content?content=vztmpl", n.Name, storage), &templates)
}

func (n *Node) VzTmpl(template, storage string) (*VzTmpl, error) {
	templates, err := n.VzTmpls(storage)
	if err != nil {
		return nil, err
	}

	volid := fmt.Sprintf("%s:vztmpl/%s", storage, template)
	for _, t := range templates {
		if t.VolID == volid {
			return t, nil
		}
	}

	return nil, fmt.Errorf("could not find vztmpl: %s", template)
}

func (n *Node) Storages() (storages Storages, err error) {
	err = n.client.Get(fmt.Sprintf("/nodes/%s/storage", n.Name), &storages)
	if err != nil {
		return
	}

	for _, s := range storages {
		s.Node = n.Name
		s.client = n.client
	}

	return
}

func (n *Node) Storage(name string) (storage *Storage, err error) {
	err = n.client.Get(fmt.Sprintf("/nodes/%s/storage/%s/status", n.Name, name), &storage)
	if err != nil {
		return
	}

	storage.Node = n.Name
	storage.client = n.client
	storage.Name = name

	return
}
