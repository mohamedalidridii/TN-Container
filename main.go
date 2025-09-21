package main
import "os" 
import "os/exec"
import "syscall"
import "strconv"
import "path/filepath"
import "fmt" 
import "io/ioutil"

func main () {
	switch os.Args[1]{
	case "run":
	run()
	case "child":
		child()
	default:
		panic ("bad  command")
	}
}
func run (){
	fmt.Printf("Running %v \n", os.Args[2:], os.Getpid())
	cmd := exec.Command("/proc/self/exe",append([]string{"child"},os.Args[2:]...)...)	
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID| syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}	
	cmd.Run()

}

func child (){
	fmt.Printf("Running %v \n", os.Args[2:], os.Getpid())
	cg()
	syscall.Sethostname([]byte("TunisianContainer"))
	syscall.Chroot("/home/medalydridi/archfs")
	syscall.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")
	cmd := exec.Command(os.Args[2], os.Args[3:]...)	
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	cmd.Run()

	syscall.Unmount("/proc", 0)

}

func cg() {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")
	os.MkdirAll(filepath.Join(pids, "medalydridi"), 0755)
	ioutil.WriteFile(filepath.Join(pids, "medalydridi/pids.max"), []byte("20"), 0700)
	// Removes the new cgroup in place after the container exits
	ioutil.WriteFile(filepath.Join(pids, "medalydridi/notify_on_release"), []byte("1"), 0700)
	ioutil.WriteFile(filepath.Join(pids, "medalydridi/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700)
}

func must(err error) {
	if err!= nil {
		panic (err)
	}
}
