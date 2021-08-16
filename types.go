package proxmox

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Otp      string `json:"otp,omitempty"`
	Path     string `json:"path,omitempty"`
	Privs    string `json:"privs,omitempty"`
	Realm    string `json:"realm,omitempty"`
}

type Session struct {
	Username            string `json:"username"`
	CsrfPreventionToken string `json:"CSRFPreventionToken,omitempty"`
	ClusterName         string `json:"clustername,omitempty"`
	Ticket              string `json:"ticket,omitempty"`
}

type Version struct {
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
	Version string `json:"version"`
}

type NodeStatuses []*NodeStatus
type NodeStatus struct {
	Status         string  `json:",omitempty"`
	Level          string  `json:",omitempty"`
	CPU            float64 `json:",omitempty"`
	Node           string  `json:",omitempty"`
	MaxMem         uint64  `json:",omitempty"`
	Disk           uint64  `json:",omitempty"`
	SslFingerprint string  `json:"ssl_fingerprint,omitempty"`
	MaxDisk        uint64  `json:",omitempty"`
	MaxCPU         int     `json:",omitempty"`
	ID             string  `json:",omitempty"`
	Mem            uint64  `json:",omitempty"`
}

type Node struct {
	Name       string
	client     *Client
	Kversion   string
	LoadAvg    []string
	CPU        float64
	RootFS     RootFS
	PVEVersion string
	CPUInfo    CPUInfo
	Swap       Memory
	Idle       int
	Memory     Memory
	Ksm        Ksm
	Uptime     uint64
	Wait       float64
}

type IsTemplate bool

func (it *IsTemplate) UnmarshalJSON(b []byte) error {
	*it = true
	if string(b) == "\"\"" {
		*it = false
	}

	return nil
}

type VirtualMachines []*VirtualMachine
type VirtualMachine struct {
	Name      string
	NetIn     uint64
	CPUs      int
	DiskWrite uint64
	Status    string
	VMID      string
	PID       string
	Netout    uint64
	Disk      uint64
	Uptime    uint64
	Mem       uint64
	CPU       float64
	MaxMem    uint64
	MaxDisk   uint64
	DiskRead  uint64
	Template  IsTemplate // empty str if a vm, int 1 if a template
	HA        HA         `json:",omitempty"`
}

type HA struct {
	Managed int
}

type RootFS struct {
	Avail uint64
	Total uint64
	Free  uint64
	Used  uint64
}

type CPUInfo struct {
	UserHz  int `json:"user_hz"`
	MHZ     string
	Mode    string
	Cores   int
	Sockets int
	Flags   string
	CPUs    int
	HVM     string
}

type Memory struct {
	Used  uint64
	Free  uint64
	Total uint64
}

type Ksm struct {
	Shared int
}

type Time struct {
	Timezone  string
	Time      uint64
	Localtime uint64
}

type Containers []*Container
type Container struct {
	Name       string
	Node       string
	client     *Client
	CPUs       int
	Status     string
	VMID       int
	OSTemplate string `json:"ostemplate"`
	Uptime     uint64
	MaxMem     uint64
	MaxDisk    uint64
	MaxSwap    uint64
}

type NewContainer struct {
	Node       string `json:"node"`
	VMID       int    `json:"vmid"`
	OSTemplate string `json:"ostemplate"`
	Storage    string `json:"storage"`
}

type ContainerStatuses []*ContainerStatus
type ContainerStatus struct {
	Data string `json:",omitempty"`
}

type Appliances []*Appliance
type Appliance struct {
	client       *Client
	Node         string `json:",omitempty"`
	Os           string
	Source       string
	Type         string
	SHA512Sum    string
	Package      string
	Template     string
	Architecture string
	InfoPage     string
	Description  string
	ManageURL    string
	Version      string
	Section      string
	Headline     string
}

type Storages []*Storage
type Storage struct {
	client       *Client
	Node         string
	Name         string `json:"storage"`
	Enabled      int
	UsedFraction float64 `json:"used_fraction"`
	Active       int
	Content      string
	Shared       int
	Avail        uint64
	Type         string
	Used         uint64
	Total        uint64
	Storage      string
}

type Volume interface {
	Delete() error
}

type ISOs []*ISO
type ISO struct{ Content }

type VzTmpls []*VzTmpl
type VzTmpl struct{ Content }

type Backups []*Backup
type Backup struct{ Content }

type Content struct {
	client  *Client
	URL     string
	Node    string
	Storage string
	Content string
	VolID   string
	CTime   uint64
	Format  string
	Size    uint64
	Used    uint64 `json:",omitempty"`
	Path    string `json:",omitempty"`
	Notes   string `json:",omitempty"`
}
