package syscall
import (
	"grumpy"
	"reflect"
	mod "syscall"
)
func fun(f *grumpy.Frame, _ []*grumpy.Object) (*grumpy.Object, *grumpy.BaseException) {
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_APPLETALK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_APPLETALK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_CCITT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_CCITT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_CHAOS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_CHAOS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_CNT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_CNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_COIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_COIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_DATAKIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_DATAKIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_DECnet)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_DECnet", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_DLI)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_DLI", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_E164)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_E164", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_ECMA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_ECMA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_HYLINK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_HYLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_IEEE80211)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_IEEE80211", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_IMPLINK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_IMPLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_INET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_INET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_INET6)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_INET6", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_IPX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_IPX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_ISDN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_ISDN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_ISO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_ISO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_LAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_LAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_LINK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_LINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_LOCAL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_LOCAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_MAX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_NATM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_NATM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_NDRV)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_NDRV", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_NETBIOS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_NETBIOS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_NS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_NS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_OSI)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_OSI", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_PPP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_PPP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_PUP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_PUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_RESERVED_36)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_RESERVED_36", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_ROUTE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_ROUTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_SIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_SIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_SNA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_SNA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_SYSTEM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_SYSTEM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_UNIX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_UNIX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.AF_UNSPEC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "AF_UNSPEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Accept)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Accept", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Access)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Access", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Adjtime)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Adjtime", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B0)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B0", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B110)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B110", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B115200)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B115200", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B1200)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B1200", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B134)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B134", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B14400)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B14400", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B150)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B150", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B1800)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B1800", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B19200)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B19200", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B200)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B200", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B230400)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B230400", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B2400)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B2400", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B28800)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B28800", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B300)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B300", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B38400)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B38400", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B4800)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B4800", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B50)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B50", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B57600)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B57600", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B600)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B600", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B7200)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B7200", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B75)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B75", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B76800)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B76800", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.B9600)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "B9600", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCFLUSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCFLUSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCGBLEN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCGBLEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCGDLT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCGDLT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCGDLTLIST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCGDLTLIST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCGETIF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCGETIF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCGHDRCMPLT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCGHDRCMPLT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCGRSIG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCGRSIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCGRTIMEOUT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCGRTIMEOUT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCGSEESENT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCGSEESENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCGSTATS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCGSTATS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCIMMEDIATE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCIMMEDIATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCPROMISC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCPROMISC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCSBLEN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCSBLEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCSDLT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCSDLT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCSETF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCSETF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCSETIF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCSETIF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCSHDRCMPLT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCSHDRCMPLT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCSRSIG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCSRSIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCSRTIMEOUT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCSRTIMEOUT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCSSEESENT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCSSEESENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BIOCVERSION)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BIOCVERSION", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_A)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_A", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_ABS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_ABS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_ADD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_ADD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_ALIGNMENT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_ALIGNMENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_ALU)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_ALU", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_AND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_AND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_B)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_B", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_DIV)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_DIV", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_H)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_H", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_IMM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_IMM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_IND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_IND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_JA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_JA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_JEQ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_JEQ", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_JGE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_JGE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_JGT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_JGT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_JMP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_JMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_JSET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_JSET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_K)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_K", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_LD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_LD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_LDX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_LDX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_LEN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_LEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_LSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_LSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_MAJOR_VERSION)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MAJOR_VERSION", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_MAXBUFSIZE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MAXBUFSIZE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_MAXINSNS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MAXINSNS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_MEM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MEM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_MEMWORDS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MEMWORDS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_MINBUFSIZE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MINBUFSIZE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_MINOR_VERSION)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MINOR_VERSION", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_MISC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MISC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_MSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_MUL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_MUL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_NEG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_NEG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_OR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_OR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_RELEASE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_RELEASE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_RET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_RET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_RSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_RSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_ST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_ST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_STX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_STX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_SUB)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_SUB", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_TAX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_TAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_TXA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_TXA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_W)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_W", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BPF_X)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BPF_X", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BRKINT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BRKINT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Bind)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Bind", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BpfBuflen)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BpfBuflen", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BpfDatalink)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BpfDatalink", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.BpfHdr
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "BpfHdr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BpfHeadercmpl)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BpfHeadercmpl", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.BpfInsn
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "BpfInsn", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BpfInterface)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BpfInterface", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BpfJump)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BpfJump", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.BpfProgram
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "BpfProgram", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.BpfStat
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "BpfStat", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BpfStats)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BpfStats", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BpfStmt)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BpfStmt", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BpfTimeout)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BpfTimeout", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.BpfVersion
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "BpfVersion", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.BytePtrFromString)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "BytePtrFromString", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ByteSliceFromString)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ByteSliceFromString", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CFLUSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CFLUSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CLOCAL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CLOCAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CREAD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CREAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CS5)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CS5", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CS6)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CS6", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CS7)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CS7", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CS8)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CS8", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CSIZE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CSIZE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CSTART)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CSTART", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CSTATUS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CSTATUS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CSTOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CSTOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CSTOPB)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CSTOPB", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CSUSP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CSUSP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CTL_MAXNAME)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CTL_MAXNAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CTL_NET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CTL_NET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Chdir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Chdir", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CheckBpfVersion)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CheckBpfVersion", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Chflags)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Chflags", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Chmod)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Chmod", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Chown)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Chown", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Chroot)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Chroot", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Clearenv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Clearenv", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Close)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Close", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CloseOnExec)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CloseOnExec", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CmsgLen)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CmsgLen", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.CmsgSpace)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "CmsgSpace", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Cmsghdr
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Cmsghdr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Connect)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Connect", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Credential
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Credential", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_APPLE_IP_OVER_IEEE1394)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_APPLE_IP_OVER_IEEE1394", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_ARCNET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_ARCNET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_ATM_CLIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_ATM_CLIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_ATM_RFC1483)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_ATM_RFC1483", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_AX25)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_AX25", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_CHAOS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_CHAOS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_CHDLC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_CHDLC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_C_HDLC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_C_HDLC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_EN10MB)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_EN10MB", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_EN3MB)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_EN3MB", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_FDDI)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_FDDI", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_IEEE802)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_IEEE802", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_IEEE802_11)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_IEEE802_11", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_IEEE802_11_RADIO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_IEEE802_11_RADIO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_IEEE802_11_RADIO_AVS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_IEEE802_11_RADIO_AVS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_LINUX_SLL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_LINUX_SLL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_LOOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_LOOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_NULL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_NULL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_PFLOG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_PFLOG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_PFSYNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_PFSYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_PPP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_PPP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_PPP_BSDOS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_PPP_BSDOS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_PPP_SERIAL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_PPP_SERIAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_PRONET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_PRONET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_RAW)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_RAW", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_SLIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_SLIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DLT_SLIP_BSDOS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DLT_SLIP_BSDOS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DT_BLK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_BLK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DT_CHR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_CHR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DT_DIR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_DIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DT_FIFO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_FIFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DT_LNK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_LNK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DT_REG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_REG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DT_SOCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_SOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DT_UNKNOWN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_UNKNOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.DT_WHT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "DT_WHT", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Dirent
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Dirent", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Dup)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Dup", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Dup2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Dup2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.E2BIG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "E2BIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EACCES)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EACCES", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EADDRINUSE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EADDRINUSE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EADDRNOTAVAIL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EADDRNOTAVAIL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EAFNOSUPPORT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EAFNOSUPPORT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EAGAIN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EAGAIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EALREADY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EALREADY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EAUTH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EAUTH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EBADARCH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EBADARCH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EBADEXEC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EBADEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EBADF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EBADF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EBADMACHO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EBADMACHO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EBADMSG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EBADMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EBADRPC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EBADRPC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EBUSY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EBUSY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ECANCELED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECANCELED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ECHILD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHILD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ECHO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ECHOCTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHOCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ECHOE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHOE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ECHOK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHOK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ECHOKE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHOKE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ECHONL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHONL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ECHOPRT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECHOPRT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ECONNABORTED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECONNABORTED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ECONNREFUSED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECONNREFUSED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ECONNRESET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ECONNRESET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EDEADLK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EDEADLK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EDESTADDRREQ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EDESTADDRREQ", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EDEVERR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EDEVERR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EDOM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EDOM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EDQUOT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EDQUOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EEXIST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EEXIST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EFAULT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EFAULT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EFBIG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EFBIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EFTYPE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EFTYPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EHOSTDOWN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EHOSTDOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EHOSTUNREACH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EHOSTUNREACH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EIDRM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EIDRM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EILSEQ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EILSEQ", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EINPROGRESS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EINPROGRESS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EINTR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EINTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EINVAL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EINVAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EIO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EIO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EISCONN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EISCONN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EISDIR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EISDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ELAST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ELAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ELOOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ELOOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EMFILE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EMFILE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EMLINK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EMLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EMSGSIZE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EMSGSIZE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EMULTIHOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EMULTIHOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENAMETOOLONG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENAMETOOLONG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENEEDAUTH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENEEDAUTH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENETDOWN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENETDOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENETRESET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENETRESET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENETUNREACH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENETUNREACH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENFILE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENFILE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOATTR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOBUFS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOBUFS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENODATA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENODATA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENODEV)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENODEV", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOENT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOEXEC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOLCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOLCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOLINK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOMEM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOMEM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOMSG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOPOLICY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOPOLICY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOPROTOOPT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOPROTOOPT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOSPC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOSPC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOSR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOSR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOSTR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOSTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOSYS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOSYS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOTBLK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTBLK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOTCONN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTCONN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOTDIR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOTEMPTY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTEMPTY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOTRECOVERABLE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTRECOVERABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOTSOCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTSOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOTSUP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTSUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENOTTY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENOTTY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ENXIO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ENXIO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EOPNOTSUPP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EOPNOTSUPP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EOVERFLOW)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EOVERFLOW", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EOWNERDEAD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EOWNERDEAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EPERM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPERM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EPFNOSUPPORT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPFNOSUPPORT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EPIPE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPIPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EPROCLIM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPROCLIM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EPROCUNAVAIL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPROCUNAVAIL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EPROGMISMATCH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPROGMISMATCH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EPROGUNAVAIL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPROGUNAVAIL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EPROTO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPROTO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EPROTONOSUPPORT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPROTONOSUPPORT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EPROTOTYPE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPROTOTYPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EPWROFF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EPWROFF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ERANGE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ERANGE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EREMOTE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EREMOTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EROFS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EROFS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ERPCMISMATCH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ERPCMISMATCH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ESHLIBVERS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ESHLIBVERS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ESHUTDOWN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ESHUTDOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ESOCKTNOSUPPORT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ESOCKTNOSUPPORT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ESPIPE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ESPIPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ESRCH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ESRCH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ESTALE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ESTALE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ETIME)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ETIMEDOUT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETIMEDOUT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ETOOMANYREFS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETOOMANYREFS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ETXTBSY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ETXTBSY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EUSERS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EUSERS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EVFILT_AIO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EVFILT_AIO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EVFILT_FS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EVFILT_FS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EVFILT_MACHPORT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EVFILT_MACHPORT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EVFILT_PROC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EVFILT_PROC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EVFILT_READ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EVFILT_READ", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EVFILT_SIGNAL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EVFILT_SIGNAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EVFILT_SYSCOUNT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EVFILT_SYSCOUNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EVFILT_THREADMARKER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EVFILT_THREADMARKER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EVFILT_TIMER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EVFILT_TIMER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EVFILT_USER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EVFILT_USER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EVFILT_VM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EVFILT_VM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EVFILT_VNODE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EVFILT_VNODE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EVFILT_WRITE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EVFILT_WRITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EV_ADD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EV_ADD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EV_CLEAR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EV_CLEAR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EV_DELETE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EV_DELETE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EV_DISABLE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EV_DISABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EV_DISPATCH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EV_DISPATCH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EV_ENABLE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EV_ENABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EV_EOF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EV_EOF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EV_ERROR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EV_ERROR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EV_FLAG0)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EV_FLAG0", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EV_FLAG1)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EV_FLAG1", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EV_ONESHOT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EV_ONESHOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EV_OOBAND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EV_OOBAND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EV_POLL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EV_POLL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EV_RECEIPT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EV_RECEIPT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EV_SYSFLAGS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EV_SYSFLAGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EWOULDBLOCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EWOULDBLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EXDEV)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EXDEV", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EXTA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EXTA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EXTB)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EXTB", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.EXTPROC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EXTPROC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Environ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Environ", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Errno
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Errno", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Exchangedata)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Exchangedata", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Exec)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Exec", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Exit)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Exit", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FD_CLOEXEC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FD_CLOEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FD_SETSIZE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FD_SETSIZE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FLUSHO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FLUSHO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_ADDFILESIGS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_ADDFILESIGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_ADDSIGS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_ADDSIGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_ALLOCATEALL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_ALLOCATEALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_ALLOCATECONTIG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_ALLOCATECONTIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_CHKCLEAN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_CHKCLEAN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_DUPFD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_DUPFD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_DUPFD_CLOEXEC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_DUPFD_CLOEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_FLUSH_DATA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_FLUSH_DATA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_FREEZE_FS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_FREEZE_FS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_FULLFSYNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_FULLFSYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_GETFD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETFD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_GETFL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETFL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_GETLK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETLK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_GETLKPID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETLKPID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_GETNOSIGPIPE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETNOSIGPIPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_GETOWN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_GETPATH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETPATH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_GETPATH_MTMINFO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETPATH_MTMINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_GETPROTECTIONCLASS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GETPROTECTIONCLASS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_GLOBAL_NOCACHE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_GLOBAL_NOCACHE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_LOG2PHYS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_LOG2PHYS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_LOG2PHYS_EXT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_LOG2PHYS_EXT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_MARKDEPENDENCY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_MARKDEPENDENCY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_NOCACHE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_NOCACHE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_NODIRECT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_NODIRECT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_OK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_OK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_PATHPKG_CHECK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_PATHPKG_CHECK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_PEOFPOSMODE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_PEOFPOSMODE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_PREALLOCATE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_PREALLOCATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_RDADVISE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_RDADVISE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_RDAHEAD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_RDAHEAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_RDLCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_RDLCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_READBOOTSTRAP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_READBOOTSTRAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_SETBACKINGSTORE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETBACKINGSTORE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_SETFD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETFD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_SETFL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETFL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_SETLK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETLK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_SETLKW)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETLKW", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_SETNOSIGPIPE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETNOSIGPIPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_SETOWN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_SETPROTECTIONCLASS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETPROTECTIONCLASS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_SETSIZE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_SETSIZE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_THAW_FS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_THAW_FS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_UNLCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_UNLCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_VOLPOSMODE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_VOLPOSMODE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_WRITEBOOTSTRAP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_WRITEBOOTSTRAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.F_WRLCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "F_WRLCK", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Fbootstraptransfer_t
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Fbootstraptransfer_t", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Fchdir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fchdir", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Fchflags)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fchflags", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Fchmod)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fchmod", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Fchown)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fchown", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FcntlFlock)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FcntlFlock", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.FdSet
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "FdSet", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Flock)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Flock", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Flock_t
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Flock_t", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.FlushBpf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FlushBpf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ForkExec)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ForkExec", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ForkLock)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ForkLock", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Fpathconf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fpathconf", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Fsid
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Fsid", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Fstat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fstat", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Fstatfs)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fstatfs", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Fstore_t
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Fstore_t", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Fsync)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Fsync", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Ftruncate)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ftruncate", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Futimes)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Futimes", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getdirentries)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getdirentries", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getdtablesize)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getdtablesize", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getegid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getegid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getenv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getenv", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Geteuid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Geteuid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getfsstat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getfsstat", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getgid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getgid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getgroups)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getgroups", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getpagesize)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getpagesize", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getpeername)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getpeername", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getpgid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getpgid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getpgrp)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getpgrp", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getpid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getpid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getppid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getppid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getpriority)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getpriority", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getrlimit)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getrlimit", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getrusage)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getrusage", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getsid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getsid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getsockname)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getsockname", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GetsockoptByte)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetsockoptByte", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GetsockoptICMPv6Filter)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetsockoptICMPv6Filter", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GetsockoptIPMreq)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetsockoptIPMreq", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GetsockoptIPv6MTUInfo)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetsockoptIPv6MTUInfo", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GetsockoptIPv6Mreq)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetsockoptIPv6Mreq", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GetsockoptInet4Addr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetsockoptInet4Addr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.GetsockoptInt)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "GetsockoptInt", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Gettimeofday)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Gettimeofday", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getuid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getuid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Getwd)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Getwd", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.HUPCL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "HUPCL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ICANON)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ICANON", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ICMP6_FILTER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ICMP6_FILTER", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.ICMPv6Filter
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "ICMPv6Filter", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ICRNL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ICRNL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IEXTEN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IEXTEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_ALLMULTI)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_ALLMULTI", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_ALTPHYS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_ALTPHYS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_BROADCAST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_BROADCAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_DEBUG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_DEBUG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_LINK0)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_LINK0", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_LINK1)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_LINK1", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_LINK2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_LINK2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_LOOPBACK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_LOOPBACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_MULTICAST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_MULTICAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_NOARP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_NOARP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_NOTRAILERS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_NOTRAILERS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_OACTIVE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_OACTIVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_POINTOPOINT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_POINTOPOINT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_PROMISC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_PROMISC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_RUNNING)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_RUNNING", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_SIMPLEX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_SIMPLEX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFF_UP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFF_UP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFNAMSIZ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFNAMSIZ", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_1822)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_1822", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_AAL5)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_AAL5", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_ARCNET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_ARCNET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_ARCNETPLUS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_ARCNETPLUS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_ATM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_ATM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_BRIDGE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_BRIDGE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_CARP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_CARP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_CELLULAR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_CELLULAR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_CEPT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_CEPT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_DS3)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_DS3", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_ENC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_ENC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_EON)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_EON", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_ETHER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_ETHER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_FAITH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_FAITH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_FDDI)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_FDDI", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_FRELAY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_FRELAY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_FRELAYDCE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_FRELAYDCE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_GIF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_GIF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_HDH1822)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_HDH1822", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_HIPPI)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_HIPPI", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_HSSI)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_HSSI", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_HY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_HY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_IEEE1394)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_IEEE1394", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_IEEE8023ADLAG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_IEEE8023ADLAG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_ISDNBASIC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_ISDNBASIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_ISDNPRIMARY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_ISDNPRIMARY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_ISO88022LLC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_ISO88022LLC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_ISO88023)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_ISO88023", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_ISO88024)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_ISO88024", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_ISO88025)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_ISO88025", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_ISO88026)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_ISO88026", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_L2VLAN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_L2VLAN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_LAPB)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_LAPB", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_LOCALTALK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_LOCALTALK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_LOOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_LOOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_MIOX25)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_MIOX25", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_MODEM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_MODEM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_NSIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_NSIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_OTHER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_OTHER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_P10)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_P10", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_P80)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_P80", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_PARA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_PARA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_PDP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_PDP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_PFLOG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_PFLOG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_PFSYNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_PFSYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_PPP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_PPP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_PROPMUX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_PROPMUX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_PROPVIRTUAL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_PROPVIRTUAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_PTPSERIAL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_PTPSERIAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_RS232)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_RS232", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_SDLC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_SDLC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_SIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_SIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_SLIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_SLIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_SMDSDXI)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_SMDSDXI", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_SMDSICIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_SMDSICIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_SONET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_SONET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_SONETPATH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_SONETPATH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_SONETVT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_SONETVT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_STARLAN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_STARLAN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_STF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_STF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_T1)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_T1", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_ULTRA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_ULTRA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_V35)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_V35", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_X25)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_X25", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_X25DDN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_X25DDN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_X25PLE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_X25PLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IFT_XETHER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IFT_XETHER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IGNBRK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IGNBRK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IGNCR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IGNCR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IGNPAR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IGNPAR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IMAXBEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IMAXBEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.INLCR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "INLCR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.INPCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "INPCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IN_CLASSA_HOST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSA_HOST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IN_CLASSA_MAX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSA_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IN_CLASSA_NET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSA_NET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IN_CLASSA_NSHIFT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSA_NSHIFT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IN_CLASSB_HOST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSB_HOST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IN_CLASSB_MAX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSB_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IN_CLASSB_NET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSB_NET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IN_CLASSB_NSHIFT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSB_NSHIFT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IN_CLASSC_HOST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSC_HOST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IN_CLASSC_NET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSC_NET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IN_CLASSC_NSHIFT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSC_NSHIFT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IN_CLASSD_HOST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSD_HOST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IN_CLASSD_NET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSD_NET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IN_CLASSD_NSHIFT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_CLASSD_NSHIFT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IN_LINKLOCALNETNUM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_LINKLOCALNETNUM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IN_LOOPBACKNET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IN_LOOPBACKNET", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.IPMreq
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IPMreq", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_3PC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_3PC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_ADFS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_ADFS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_AH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_AH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_AHIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_AHIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_APES)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_APES", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_ARGUS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_ARGUS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_AX25)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_AX25", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_BHA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_BHA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_BLT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_BLT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_BRSATMON)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_BRSATMON", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_CFTP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_CFTP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_CHAOS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_CHAOS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_CMTP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_CMTP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_CPHB)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_CPHB", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_CPNX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_CPNX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_DDP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_DDP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_DGP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_DGP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_DIVERT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_DIVERT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_DONE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_DONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_DSTOPTS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_DSTOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_EGP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_EGP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_EMCON)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_EMCON", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_ENCAP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_ENCAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_EON)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_EON", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_ESP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_ESP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_ETHERIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_ETHERIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_FRAGMENT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_FRAGMENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_GGP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_GGP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_GMTP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_GMTP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_GRE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_GRE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_HELLO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_HELLO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_HMP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_HMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_HOPOPTS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_HOPOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_ICMP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_ICMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_ICMPV6)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_ICMPV6", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_IDP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IDP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_IDPR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IDPR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_IDRP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IDRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_IGMP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IGMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_IGP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IGP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_IGRP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_IL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_INLSP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_INLSP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_INP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_INP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_IP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_IPCOMP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IPCOMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_IPCV)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IPCV", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_IPEIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IPEIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_IPIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IPIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_IPPC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IPPC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_IPV4)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IPV4", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_IPV6)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IPV6", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_IRTP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_IRTP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_KRYPTOLAN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_KRYPTOLAN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_LARP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_LARP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_LEAF1)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_LEAF1", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_LEAF2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_LEAF2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_MAX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_MAXID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_MAXID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_MEAS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_MEAS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_MHRP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_MHRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_MICP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_MICP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_MTP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_MTP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_MUX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_MUX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_ND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_ND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_NHRP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_NHRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_NONE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_NONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_NSP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_NSP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_NVPII)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_NVPII", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_OSPFIGP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_OSPFIGP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_PGM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_PGM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_PIGP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_PIGP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_PIM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_PIM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_PRM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_PRM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_PUP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_PUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_PVP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_PVP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_RAW)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_RAW", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_RCCMON)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_RCCMON", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_RDP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_RDP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_ROUTING)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_ROUTING", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_RSVP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_RSVP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_RVD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_RVD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_SATEXPAK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_SATEXPAK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_SATMON)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_SATMON", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_SCCSP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_SCCSP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_SCTP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_SCTP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_SDRP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_SDRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_SEP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_SEP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_SRPC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_SRPC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_ST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_ST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_SVMTP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_SVMTP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_SWIPE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_SWIPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_TCF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_TCF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_TCP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_TCP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_TP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_TP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_TPXX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_TPXX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_TRUNK1)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_TRUNK1", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_TRUNK2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_TRUNK2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_TTP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_TTP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_UDP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_UDP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_VINES)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_VINES", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_VISA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_VISA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_VMTP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_VMTP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_WBEXPAK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_WBEXPAK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_WBMON)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_WBMON", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_WSN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_WSN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_XNET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_XNET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPPROTO_XTP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPPROTO_XTP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_2292DSTOPTS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_2292DSTOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_2292HOPLIMIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_2292HOPLIMIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_2292HOPOPTS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_2292HOPOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_2292NEXTHOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_2292NEXTHOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_2292PKTINFO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_2292PKTINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_2292PKTOPTIONS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_2292PKTOPTIONS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_2292RTHDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_2292RTHDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_BINDV6ONLY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_BINDV6ONLY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_BOUND_IF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_BOUND_IF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_CHECKSUM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_CHECKSUM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_DEFAULT_MULTICAST_HOPS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_DEFAULT_MULTICAST_HOPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_DEFAULT_MULTICAST_LOOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_DEFAULT_MULTICAST_LOOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_DEFHLIM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_DEFHLIM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_FAITH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_FAITH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_FLOWINFO_MASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_FLOWINFO_MASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_FLOWLABEL_MASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_FLOWLABEL_MASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_FRAGTTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_FRAGTTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_FW_ADD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_FW_ADD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_FW_DEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_FW_DEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_FW_FLUSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_FW_FLUSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_FW_GET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_FW_GET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_FW_ZERO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_FW_ZERO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_HLIMDEC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_HLIMDEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_IPSEC_POLICY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_IPSEC_POLICY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_JOIN_GROUP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_JOIN_GROUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_LEAVE_GROUP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_LEAVE_GROUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_MAXHLIM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_MAXHLIM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_MAXOPTHDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_MAXOPTHDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_MAXPACKET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_MAXPACKET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_MAX_GROUP_SRC_FILTER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_MAX_GROUP_SRC_FILTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_MAX_MEMBERSHIPS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_MAX_MEMBERSHIPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_MAX_SOCK_SRC_FILTER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_MAX_SOCK_SRC_FILTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_MIN_MEMBERSHIPS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_MIN_MEMBERSHIPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_MMTU)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_MMTU", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_MULTICAST_HOPS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_MULTICAST_HOPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_MULTICAST_IF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_MULTICAST_IF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_MULTICAST_LOOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_MULTICAST_LOOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_PORTRANGE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_PORTRANGE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_PORTRANGE_DEFAULT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_PORTRANGE_DEFAULT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_PORTRANGE_HIGH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_PORTRANGE_HIGH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_PORTRANGE_LOW)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_PORTRANGE_LOW", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_RECVTCLASS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RECVTCLASS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_RTHDR_LOOSE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RTHDR_LOOSE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_RTHDR_STRICT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RTHDR_STRICT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_RTHDR_TYPE_0)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_RTHDR_TYPE_0", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_SOCKOPT_RESERVED1)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_SOCKOPT_RESERVED1", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_TCLASS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_TCLASS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_UNICAST_HOPS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_UNICAST_HOPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_V6ONLY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_V6ONLY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_VERSION)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_VERSION", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IPV6_VERSION_MASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IPV6_VERSION_MASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_ADD_MEMBERSHIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_ADD_MEMBERSHIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_ADD_SOURCE_MEMBERSHIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_ADD_SOURCE_MEMBERSHIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_BLOCK_SOURCE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_BLOCK_SOURCE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_BOUND_IF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_BOUND_IF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_DEFAULT_MULTICAST_LOOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_DEFAULT_MULTICAST_LOOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_DEFAULT_MULTICAST_TTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_DEFAULT_MULTICAST_TTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_DF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_DF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_DROP_MEMBERSHIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_DROP_MEMBERSHIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_DROP_SOURCE_MEMBERSHIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_DROP_SOURCE_MEMBERSHIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_DUMMYNET_CONFIGURE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_DUMMYNET_CONFIGURE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_DUMMYNET_DEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_DUMMYNET_DEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_DUMMYNET_FLUSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_DUMMYNET_FLUSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_DUMMYNET_GET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_DUMMYNET_GET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_FAITH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_FAITH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_FW_ADD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_FW_ADD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_FW_DEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_FW_DEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_FW_FLUSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_FW_FLUSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_FW_GET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_FW_GET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_FW_RESETLOG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_FW_RESETLOG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_FW_ZERO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_FW_ZERO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_HDRINCL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_HDRINCL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_IPSEC_POLICY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_IPSEC_POLICY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_MAXPACKET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MAXPACKET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_MAX_GROUP_SRC_FILTER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MAX_GROUP_SRC_FILTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_MAX_MEMBERSHIPS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MAX_MEMBERSHIPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_MAX_SOCK_MUTE_FILTER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MAX_SOCK_MUTE_FILTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_MAX_SOCK_SRC_FILTER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MAX_SOCK_SRC_FILTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_MF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_MIN_MEMBERSHIPS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MIN_MEMBERSHIPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_MSFILTER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MSFILTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_MSS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MSS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_MULTICAST_IF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MULTICAST_IF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_MULTICAST_IFINDEX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MULTICAST_IFINDEX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_MULTICAST_LOOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MULTICAST_LOOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_MULTICAST_TTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MULTICAST_TTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_MULTICAST_VIF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_MULTICAST_VIF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_NAT__XXX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_NAT__XXX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_OFFMASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_OFFMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_OLD_FW_ADD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_OLD_FW_ADD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_OLD_FW_DEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_OLD_FW_DEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_OLD_FW_FLUSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_OLD_FW_FLUSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_OLD_FW_GET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_OLD_FW_GET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_OLD_FW_RESETLOG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_OLD_FW_RESETLOG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_OLD_FW_ZERO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_OLD_FW_ZERO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_OPTIONS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_OPTIONS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_PKTINFO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_PKTINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_PORTRANGE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_PORTRANGE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_PORTRANGE_DEFAULT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_PORTRANGE_DEFAULT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_PORTRANGE_HIGH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_PORTRANGE_HIGH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_PORTRANGE_LOW)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_PORTRANGE_LOW", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_RECVDSTADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RECVDSTADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_RECVIF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RECVIF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_RECVOPTS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RECVOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_RECVPKTINFO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RECVPKTINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_RECVRETOPTS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RECVRETOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_RECVTTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RECVTTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_RETOPTS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RETOPTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_RF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_RSVP_OFF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RSVP_OFF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_RSVP_ON)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RSVP_ON", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_RSVP_VIF_OFF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RSVP_VIF_OFF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_RSVP_VIF_ON)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_RSVP_VIF_ON", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_STRIPHDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_STRIPHDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_TOS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_TOS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_TRAFFIC_MGT_BACKGROUND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_TRAFFIC_MGT_BACKGROUND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_TTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_TTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IP_UNBLOCK_SOURCE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IP_UNBLOCK_SOURCE", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.IPv6MTUInfo
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IPv6MTUInfo", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.IPv6Mreq
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IPv6Mreq", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ISIG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ISIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ISTRIP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ISTRIP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IUTF8)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IUTF8", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IXANY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IXANY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IXOFF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IXOFF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.IXON)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IXON", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.IfData
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IfData", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.IfMsghdr
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IfMsghdr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.IfaMsghdr
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IfaMsghdr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.IfmaMsghdr
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IfmaMsghdr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.IfmaMsghdr2
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "IfmaMsghdr2", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ImplementsGetwd)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ImplementsGetwd", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Inet4Pktinfo
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Inet4Pktinfo", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Inet6Pktinfo
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Inet6Pktinfo", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.InterfaceAddrMessage
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "InterfaceAddrMessage", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.InterfaceMessage
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "InterfaceMessage", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.InterfaceMulticastAddrMessage
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "InterfaceMulticastAddrMessage", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Iovec
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Iovec", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Issetugid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Issetugid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Kevent)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Kevent", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Kevent_t
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Kevent_t", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Kill)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Kill", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Kqueue)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Kqueue", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.LOCK_EX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LOCK_EX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.LOCK_NB)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LOCK_NB", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.LOCK_SH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LOCK_SH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.LOCK_UN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "LOCK_UN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Lchown)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lchown", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Linger
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Linger", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Link)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Link", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Listen)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Listen", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Log2phys_t
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Log2phys_t", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Lstat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Lstat", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MADV_CAN_REUSE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_CAN_REUSE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MADV_DONTNEED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_DONTNEED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MADV_FREE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_FREE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MADV_FREE_REUSABLE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_FREE_REUSABLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MADV_FREE_REUSE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_FREE_REUSE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MADV_NORMAL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_NORMAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MADV_RANDOM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_RANDOM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MADV_SEQUENTIAL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_SEQUENTIAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MADV_WILLNEED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_WILLNEED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MADV_ZERO_WIRED_PAGES)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MADV_ZERO_WIRED_PAGES", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MAP_ANON)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_ANON", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MAP_COPY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_COPY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MAP_FILE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_FILE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MAP_FIXED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_FIXED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MAP_HASSEMAPHORE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_HASSEMAPHORE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MAP_JIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_JIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MAP_NOCACHE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_NOCACHE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MAP_NOEXTEND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_NOEXTEND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MAP_NORESERVE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_NORESERVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MAP_PRIVATE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_PRIVATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MAP_RENAME)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_RENAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MAP_RESERVED0080)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_RESERVED0080", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MAP_SHARED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MAP_SHARED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MCL_CURRENT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MCL_CURRENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MCL_FUTURE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MCL_FUTURE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MSG_CTRUNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_CTRUNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MSG_DONTROUTE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_DONTROUTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MSG_DONTWAIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_DONTWAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MSG_EOF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_EOF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MSG_EOR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_EOR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MSG_FLUSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_FLUSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MSG_HAVEMORE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_HAVEMORE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MSG_HOLD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_HOLD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MSG_NEEDSA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_NEEDSA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MSG_OOB)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_OOB", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MSG_PEEK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_PEEK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MSG_RCVMORE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_RCVMORE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MSG_SEND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_SEND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MSG_TRUNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_TRUNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MSG_WAITALL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_WAITALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MSG_WAITSTREAM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MSG_WAITSTREAM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MS_ASYNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_ASYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MS_DEACTIVATE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_DEACTIVATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MS_INVALIDATE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_INVALIDATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MS_KILLPAGES)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_KILLPAGES", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.MS_SYNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "MS_SYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mkdir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mkdir", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mkfifo)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mkfifo", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mknod)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mknod", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mlock)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mlock", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mlockall)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mlockall", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mmap)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mmap", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Mprotect)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Mprotect", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Msghdr
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Msghdr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Munlock)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Munlock", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Munlockall)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Munlockall", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Munmap)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Munmap", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NAME_MAX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NAME_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NET_RT_DUMP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NET_RT_DUMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NET_RT_DUMP2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NET_RT_DUMP2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NET_RT_FLAGS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NET_RT_FLAGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NET_RT_IFLIST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NET_RT_IFLIST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NET_RT_IFLIST2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NET_RT_IFLIST2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NET_RT_MAXID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NET_RT_MAXID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NET_RT_STAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NET_RT_STAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NET_RT_TRASH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NET_RT_TRASH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOFLSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOFLSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_ABSOLUTE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_ABSOLUTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_ATTRIB)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_ATTRIB", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_CHILD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_CHILD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_DELETE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_DELETE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_EXEC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_EXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_EXIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_EXIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_EXITSTATUS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_EXITSTATUS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_EXTEND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_EXTEND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_FFAND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_FFAND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_FFCOPY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_FFCOPY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_FFCTRLMASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_FFCTRLMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_FFLAGSMASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_FFLAGSMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_FFNOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_FFNOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_FFOR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_FFOR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_FORK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_FORK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_LINK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_LINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_LOWAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_LOWAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_NONE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_NONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_NSECONDS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_NSECONDS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_PCTRLMASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_PCTRLMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_PDATAMASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_PDATAMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_REAP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_REAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_RENAME)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_RENAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_RESOURCEEND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_RESOURCEEND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_REVOKE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_REVOKE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_SECONDS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_SECONDS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_SIGNAL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_SIGNAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_TRACK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_TRACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_TRACKERR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_TRACKERR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_TRIGGER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_TRIGGER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_USECONDS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_USECONDS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_VM_ERROR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_VM_ERROR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_VM_PRESSURE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_VM_PRESSURE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_VM_PRESSURE_SUDDEN_TERMINATE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_VM_PRESSURE_SUDDEN_TERMINATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_VM_PRESSURE_TERMINATE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_VM_PRESSURE_TERMINATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NOTE_WRITE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NOTE_WRITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NsecToTimespec)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NsecToTimespec", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.NsecToTimeval)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NsecToTimeval", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.OCRNL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "OCRNL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.OFDEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "OFDEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.OFILL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "OFILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ONLCR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ONLCR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ONLRET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ONLRET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ONOCR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ONOCR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ONOEOT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ONOEOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.OPOST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "OPOST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_ACCMODE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_ACCMODE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_ALERT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_ALERT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_APPEND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_APPEND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_ASYNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_ASYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_CLOEXEC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_CLOEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_CREAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_CREAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_DIRECTORY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_DIRECTORY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_DSYNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_DSYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_EVTONLY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_EVTONLY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_EXCL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_EXCL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_EXLOCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_EXLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_FSYNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_FSYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_NDELAY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_NDELAY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_NOCTTY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_NOCTTY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_NOFOLLOW)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_NOFOLLOW", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_NONBLOCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_NONBLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_POPUP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_POPUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_RDONLY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_RDONLY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_RDWR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_RDWR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_SHLOCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_SHLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_SYMLINK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_SYMLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_SYNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_SYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_TRUNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_TRUNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.O_WRONLY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "O_WRONLY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Open)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Open", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PARENB)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PARENB", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PARMRK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PARMRK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PARODD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PARODD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PENDIN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PENDIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PRIO_PGRP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PRIO_PGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PRIO_PROCESS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PRIO_PROCESS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PRIO_USER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PRIO_USER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PROT_EXEC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PROT_EXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PROT_NONE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PROT_NONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PROT_READ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PROT_READ", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PROT_WRITE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PROT_WRITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PTRACE_CONT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_CONT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PTRACE_KILL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_KILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PTRACE_TRACEME)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PTRACE_TRACEME", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_ATTACH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_ATTACH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_ATTACHEXC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_ATTACHEXC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_CONTINUE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_CONTINUE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_DENY_ATTACH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_DENY_ATTACH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_DETACH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_DETACH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_FIRSTMACH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_FIRSTMACH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_FORCEQUOTA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_FORCEQUOTA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_KILL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_KILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_READ_D)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_READ_D", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_READ_I)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_READ_I", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_READ_U)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_READ_U", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_SIGEXC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_SIGEXC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_STEP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_STEP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_THUPDATE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_THUPDATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_TRACE_ME)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_TRACE_ME", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_WRITE_D)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_WRITE_D", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_WRITE_I)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_WRITE_I", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PT_WRITE_U)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PT_WRITE_U", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ParseDirent)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ParseDirent", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ParseRoutingMessage)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ParseRoutingMessage", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ParseRoutingSockaddr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ParseRoutingSockaddr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ParseSocketControlMessage)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ParseSocketControlMessage", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ParseUnixRights)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ParseUnixRights", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pathconf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pathconf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pipe)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pipe", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pread)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pread", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.ProcAttr
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "ProcAttr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PtraceAttach)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PtraceAttach", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.PtraceDetach)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "PtraceDetach", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Pwrite)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Pwrite", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RLIMIT_AS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RLIMIT_AS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RLIMIT_CORE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RLIMIT_CORE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RLIMIT_CPU)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RLIMIT_CPU", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RLIMIT_DATA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RLIMIT_DATA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RLIMIT_FSIZE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RLIMIT_FSIZE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RLIMIT_NOFILE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RLIMIT_NOFILE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RLIMIT_STACK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RLIMIT_STACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RLIM_INFINITY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RLIM_INFINITY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTAX_AUTHOR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_AUTHOR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTAX_BRD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_BRD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTAX_DST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_DST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTAX_GATEWAY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_GATEWAY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTAX_GENMASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_GENMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTAX_IFA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_IFA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTAX_IFP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_IFP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTAX_MAX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_MAX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTAX_NETMASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTAX_NETMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTA_AUTHOR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_AUTHOR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTA_BRD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_BRD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTA_DST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_DST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTA_GATEWAY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_GATEWAY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTA_GENMASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_GENMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTA_IFA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_IFA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTA_IFP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_IFP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTA_NETMASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTA_NETMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_BLACKHOLE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_BLACKHOLE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_BROADCAST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_BROADCAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_CLONING)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_CLONING", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_CONDEMNED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_CONDEMNED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_DELCLONE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_DELCLONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_DONE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_DONE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_DYNAMIC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_DYNAMIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_GATEWAY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_GATEWAY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_HOST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_HOST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_IFREF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_IFREF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_IFSCOPE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_IFSCOPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_LLINFO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_LLINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_LOCAL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_LOCAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_MODIFIED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_MODIFIED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_MULTICAST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_MULTICAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_PINNED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_PINNED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_PRCLONING)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_PRCLONING", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_PROTO1)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_PROTO1", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_PROTO2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_PROTO2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_PROTO3)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_PROTO3", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_REJECT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_REJECT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_STATIC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_STATIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_UP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_UP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_WASCLONED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_WASCLONED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTF_XRESOLVE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTF_XRESOLVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_ADD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_ADD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_CHANGE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_CHANGE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_DELADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_DELADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_DELETE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_DELETE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_DELMADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_DELMADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_GET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_GET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_GET2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_GET2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_IFINFO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_IFINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_IFINFO2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_IFINFO2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_LOCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_LOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_LOSING)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_LOSING", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_MISS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_MISS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_NEWADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NEWADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_NEWMADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NEWMADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_NEWMADDR2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_NEWMADDR2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_OLDADD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_OLDADD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_OLDDEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_OLDDEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_REDIRECT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_REDIRECT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_RESOLVE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_RESOLVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_RTTUNIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_RTTUNIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTM_VERSION)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTM_VERSION", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTV_EXPIRE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTV_EXPIRE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTV_HOPCOUNT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTV_HOPCOUNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTV_MTU)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTV_MTU", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTV_RPIPE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTV_RPIPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTV_RTT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTV_RTT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTV_RTTVAR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTV_RTTVAR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTV_SPIPE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTV_SPIPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RTV_SSTHRESH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RTV_SSTHRESH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RUSAGE_CHILDREN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RUSAGE_CHILDREN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RUSAGE_SELF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RUSAGE_SELF", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Radvisory_t
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Radvisory_t", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RawSockaddr
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RawSockaddr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RawSockaddrAny
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RawSockaddrAny", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RawSockaddrDatalink
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RawSockaddrDatalink", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RawSockaddrInet4
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RawSockaddrInet4", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RawSockaddrInet6
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RawSockaddrInet6", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RawSockaddrUnix
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RawSockaddrUnix", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RawSyscall)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RawSyscall", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RawSyscall6)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RawSyscall6", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Read)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Read", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.ReadDirent)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ReadDirent", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Readlink)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Readlink", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Recvfrom)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Recvfrom", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Recvmsg)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Recvmsg", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Rename)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Rename", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Revoke)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Revoke", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Rlimit
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Rlimit", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Rmdir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Rmdir", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.RouteMessage
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RouteMessage", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.RouteRIB)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "RouteRIB", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.RtMetrics
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RtMetrics", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RtMsghdr
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RtMsghdr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Rusage
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Rusage", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SCM_CREDS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SCM_CREDS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SCM_RIGHTS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SCM_RIGHTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SCM_TIMESTAMP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SCM_TIMESTAMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SCM_TIMESTAMP_MONOTONIC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SCM_TIMESTAMP_MONOTONIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SHUT_RD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SHUT_RD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SHUT_RDWR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SHUT_RDWR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SHUT_WR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SHUT_WR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGABRT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGABRT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGALRM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGALRM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGBUS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGBUS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGCHLD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGCHLD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGCONT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGCONT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGEMT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGEMT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGFPE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGFPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGHUP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGHUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGILL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGINFO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGINFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGINT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGINT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGIO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGIO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGIOT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGIOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGKILL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGKILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGPIPE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGPIPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGPROF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGPROF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGQUIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGQUIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGSEGV)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGSEGV", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGSTOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGSTOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGSYS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGSYS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGTERM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGTERM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGTRAP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGTRAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGTSTP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGTSTP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGTTIN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGTTIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGTTOU)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGTTOU", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGURG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGURG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGUSR1)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGUSR1", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGUSR2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGUSR2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGVTALRM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGVTALRM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGWINCH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGWINCH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGXCPU)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGXCPU", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIGXFSZ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIGXFSZ", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCADDMULTI)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCADDMULTI", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCAIFADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCAIFADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCALIFADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCALIFADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCARPIPLL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCARPIPLL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCATMARK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCATMARK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCAUTOADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCAUTOADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCAUTONETMASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCAUTONETMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCDELMULTI)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCDELMULTI", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCDIFADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCDIFADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCDIFPHYADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCDIFPHYADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCDLIFADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCDLIFADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGDRVSPEC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGDRVSPEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGETSGCNT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGETSGCNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGETVIFCNT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGETVIFCNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGETVLAN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGETVLAN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGHIWAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGHIWAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFALTMTU)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFALTMTU", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFASYNCMAP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFASYNCMAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFBOND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFBOND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFBRDADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFBRDADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFCAP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFCAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFCONF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFCONF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFDEVMTU)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFDEVMTU", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFDSTADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFDSTADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFFLAGS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFFLAGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFGENERIC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFGENERIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFKPI)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFKPI", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFMAC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFMAC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFMEDIA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFMEDIA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFMETRIC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFMETRIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFMTU)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFMTU", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFNETMASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFNETMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFPDSTADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFPDSTADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFPHYS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFPHYS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFPSRCADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFPSRCADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFSTATUS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFSTATUS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFVLAN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFVLAN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGIFWAKEFLAGS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGIFWAKEFLAGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGLIFADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGLIFADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGLIFPHYADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGLIFPHYADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGLOWAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGLOWAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCGPGRP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCGPGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCIFCREATE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCIFCREATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCIFCREATE2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCIFCREATE2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCIFDESTROY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCIFDESTROY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCRSLVMULTI)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCRSLVMULTI", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSDRVSPEC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSDRVSPEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSETVLAN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSETVLAN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSHIWAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSHIWAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFALTMTU)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFALTMTU", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFASYNCMAP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFASYNCMAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFBOND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFBOND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFBRDADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFBRDADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFCAP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFCAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFDSTADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFDSTADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFFLAGS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFFLAGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFGENERIC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFGENERIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFKPI)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFKPI", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFLLADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFLLADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFMAC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFMAC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFMEDIA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFMEDIA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFMETRIC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFMETRIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFMTU)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFMTU", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFNETMASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFNETMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFPHYADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFPHYADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFPHYS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFPHYS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSIFVLAN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSIFVLAN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSLIFPHYADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSLIFPHYADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSLOWAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSLOWAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SIOCSPGRP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SIOCSPGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SOCK_DGRAM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOCK_DGRAM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SOCK_MAXADDRLEN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOCK_MAXADDRLEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SOCK_RAW)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOCK_RAW", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SOCK_RDM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOCK_RDM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SOCK_SEQPACKET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOCK_SEQPACKET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SOCK_STREAM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOCK_STREAM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SOL_SOCKET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOL_SOCKET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SOMAXCONN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SOMAXCONN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_ACCEPTCONN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_ACCEPTCONN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_BROADCAST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_BROADCAST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_DEBUG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_DEBUG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_DONTROUTE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_DONTROUTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_DONTTRUNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_DONTTRUNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_ERROR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_ERROR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_KEEPALIVE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_KEEPALIVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_LABEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_LABEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_LINGER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_LINGER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_LINGER_SEC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_LINGER_SEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_NKE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_NKE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_NOADDRERR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_NOADDRERR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_NOSIGPIPE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_NOSIGPIPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_NOTIFYCONFLICT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_NOTIFYCONFLICT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_NP_EXTENSIONS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_NP_EXTENSIONS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_NREAD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_NREAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_NWRITE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_NWRITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_OOBINLINE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_OOBINLINE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_PEERLABEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_PEERLABEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_RANDOMPORT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_RANDOMPORT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_RCVBUF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_RCVBUF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_RCVLOWAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_RCVLOWAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_RCVTIMEO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_RCVTIMEO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_RESTRICTIONS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_RESTRICTIONS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_RESTRICT_DENYIN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_RESTRICT_DENYIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_RESTRICT_DENYOUT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_RESTRICT_DENYOUT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_RESTRICT_DENYSET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_RESTRICT_DENYSET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_REUSEADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_REUSEADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_REUSEPORT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_REUSEPORT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_REUSESHAREUID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_REUSESHAREUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_SNDBUF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_SNDBUF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_SNDLOWAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_SNDLOWAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_SNDTIMEO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_SNDTIMEO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_TIMESTAMP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_TIMESTAMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_TIMESTAMP_MONOTONIC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_TIMESTAMP_MONOTONIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_TYPE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_TYPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_UPCALLCLOSEWAIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_UPCALLCLOSEWAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_USELOOPBACK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_USELOOPBACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_WANTMORE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_WANTMORE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SO_WANTOOBFLAG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SO_WANTOOBFLAG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_ACCEPT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ACCEPT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_ACCEPT_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ACCEPT_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_ACCESS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ACCESS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_ACCESS_EXTENDED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ACCESS_EXTENDED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_ACCT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ACCT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_ADD_PROFIL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ADD_PROFIL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_ADJTIME)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ADJTIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_AIO_CANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_AIO_CANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_AIO_ERROR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_AIO_ERROR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_AIO_FSYNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_AIO_FSYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_AIO_READ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_AIO_READ", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_AIO_RETURN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_AIO_RETURN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_AIO_SUSPEND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_AIO_SUSPEND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_AIO_SUSPEND_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_AIO_SUSPEND_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_AIO_WRITE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_AIO_WRITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_ATGETMSG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ATGETMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_ATPGETREQ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ATPGETREQ", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_ATPGETRSP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ATPGETRSP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_ATPSNDREQ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ATPSNDREQ", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_ATPSNDRSP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ATPSNDRSP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_ATPUTMSG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ATPUTMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_ATSOCKET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ATSOCKET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_AUDIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_AUDIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_AUDITCTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_AUDITCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_AUDITON)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_AUDITON", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_AUDIT_SESSION_JOIN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_AUDIT_SESSION_JOIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_AUDIT_SESSION_PORT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_AUDIT_SESSION_PORT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_AUDIT_SESSION_SELF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_AUDIT_SESSION_SELF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_BIND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_BIND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_BSDTHREAD_CREATE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_BSDTHREAD_CREATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_BSDTHREAD_REGISTER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_BSDTHREAD_REGISTER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_BSDTHREAD_TERMINATE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_BSDTHREAD_TERMINATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_CHDIR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CHDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_CHFLAGS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CHFLAGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_CHMOD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CHMOD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_CHMOD_EXTENDED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CHMOD_EXTENDED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_CHOWN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CHOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_CHROOT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CHROOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_CHUD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CHUD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_CLOSE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CLOSE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_CLOSE_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CLOSE_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_CONNECT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CONNECT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_CONNECT_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CONNECT_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_COPYFILE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_COPYFILE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_CSOPS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_CSOPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_DELETE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_DELETE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_DUP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_DUP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_DUP2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_DUP2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_EXCHANGEDATA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_EXCHANGEDATA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_EXECVE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_EXECVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_EXIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_EXIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FCHDIR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FCHDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FCHFLAGS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FCHFLAGS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FCHMOD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FCHMOD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FCHMOD_EXTENDED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FCHMOD_EXTENDED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FCHOWN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FCHOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FCNTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FCNTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FCNTL_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FCNTL_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FDATASYNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FDATASYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FFSCTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FFSCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FGETATTRLIST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FGETATTRLIST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FGETXATTR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FGETXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FHOPEN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FHOPEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FILEPORT_MAKEFD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FILEPORT_MAKEFD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FILEPORT_MAKEPORT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FILEPORT_MAKEPORT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FLISTXATTR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FLISTXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FLOCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FORK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FORK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FPATHCONF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FPATHCONF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FREMOVEXATTR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FREMOVEXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FSCTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FSETATTRLIST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSETATTRLIST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FSETXATTR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSETXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FSGETPATH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSGETPATH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FSTAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSTAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FSTAT64)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSTAT64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FSTAT64_EXTENDED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSTAT64_EXTENDED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FSTATFS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSTATFS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FSTATFS64)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSTATFS64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FSTATV)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSTATV", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FSTAT_EXTENDED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSTAT_EXTENDED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FSYNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FSYNC_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FSYNC_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FTRUNCATE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FTRUNCATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_FUTIMES)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_FUTIMES", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETATTRLIST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETATTRLIST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETAUDIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETAUDIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETAUDIT_ADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETAUDIT_ADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETAUID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETAUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETDIRENTRIES)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETDIRENTRIES", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETDIRENTRIES64)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETDIRENTRIES64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETDIRENTRIESATTR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETDIRENTRIESATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETDTABLESIZE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETDTABLESIZE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETEGID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETEGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETEUID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETEUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETFH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETFH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETFSSTAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETFSSTAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETFSSTAT64)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETFSSTAT64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETGID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETGROUPS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETGROUPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETHOSTUUID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETHOSTUUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETITIMER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETITIMER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETLCID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETLCID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETLOGIN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETLOGIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETPEERNAME)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETPEERNAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETPGID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETPGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETPGRP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETPGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETPID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETPID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETPPID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETPPID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETPRIORITY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETPRIORITY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETRLIMIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETRLIMIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETRUSAGE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETRUSAGE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETSGROUPS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETSGROUPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETSID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETSID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETSOCKNAME)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETSOCKNAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETSOCKOPT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETSOCKOPT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETTID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETTID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETTIMEOFDAY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETTIMEOFDAY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETUID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETWGROUPS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETWGROUPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_GETXATTR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_GETXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_IDENTITYSVC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_IDENTITYSVC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_INITGROUPS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_INITGROUPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_IOCTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_IOCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_IOPOLICYSYS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_IOPOLICYSYS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_ISSETUGID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_ISSETUGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_KDEBUG_TRACE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_KDEBUG_TRACE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_KEVENT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_KEVENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_KEVENT64)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_KEVENT64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_KILL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_KILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_KQUEUE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_KQUEUE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_LCHOWN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LCHOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_LINK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_LIO_LISTIO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LIO_LISTIO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_LISTEN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LISTEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_LISTXATTR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LISTXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_LSEEK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LSEEK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_LSTAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LSTAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_LSTAT64)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LSTAT64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_LSTAT64_EXTENDED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LSTAT64_EXTENDED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_LSTATV)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LSTATV", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_LSTAT_EXTENDED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_LSTAT_EXTENDED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MADVISE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MADVISE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MAXSYSCALL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MAXSYSCALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MINCORE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MINCORE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MINHERIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MINHERIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MKCOMPLEX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MKCOMPLEX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MKDIR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MKDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MKDIR_EXTENDED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MKDIR_EXTENDED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MKFIFO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MKFIFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MKFIFO_EXTENDED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MKFIFO_EXTENDED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MKNOD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MKNOD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MLOCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MLOCKALL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MLOCKALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MMAP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MMAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MODWATCH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MODWATCH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MOUNT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MOUNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MPROTECT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MPROTECT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MSGCTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MSGCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MSGGET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MSGGET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MSGRCV)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MSGRCV", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MSGRCV_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MSGRCV_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MSGSND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MSGSND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MSGSND_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MSGSND_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MSGSYS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MSGSYS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MSYNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MSYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MSYNC_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MSYNC_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MUNLOCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MUNLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MUNLOCKALL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MUNLOCKALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_MUNMAP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_MUNMAP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_NFSCLNT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_NFSCLNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_NFSSVC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_NFSSVC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_OPEN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_OPEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_OPEN_EXTENDED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_OPEN_EXTENDED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_OPEN_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_OPEN_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PATHCONF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PATHCONF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PID_HIBERNATE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PID_HIBERNATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PID_RESUME)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PID_RESUME", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PID_SHUTDOWN_SOCKETS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PID_SHUTDOWN_SOCKETS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PID_SUSPEND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PID_SUSPEND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PIPE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PIPE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_POLL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_POLL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_POLL_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_POLL_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_POSIX_SPAWN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_POSIX_SPAWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PREAD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PREAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PREAD_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PREAD_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PROCESS_POLICY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PROCESS_POLICY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PROC_INFO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PROC_INFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PROFIL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PROFIL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PSYNCH_CVBROAD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PSYNCH_CVBROAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PSYNCH_CVCLRPREPOST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PSYNCH_CVCLRPREPOST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PSYNCH_CVSIGNAL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PSYNCH_CVSIGNAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PSYNCH_CVWAIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PSYNCH_CVWAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PSYNCH_MUTEXDROP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PSYNCH_MUTEXDROP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PSYNCH_MUTEXWAIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PSYNCH_MUTEXWAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PSYNCH_RW_DOWNGRADE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PSYNCH_RW_DOWNGRADE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PSYNCH_RW_LONGRDLOCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PSYNCH_RW_LONGRDLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PSYNCH_RW_RDLOCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PSYNCH_RW_RDLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PSYNCH_RW_UNLOCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PSYNCH_RW_UNLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PSYNCH_RW_UNLOCK2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PSYNCH_RW_UNLOCK2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PSYNCH_RW_UPGRADE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PSYNCH_RW_UPGRADE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PSYNCH_RW_WRLOCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PSYNCH_RW_WRLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PSYNCH_RW_YIELDWRLOCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PSYNCH_RW_YIELDWRLOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PTRACE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PTRACE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PWRITE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PWRITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_PWRITE_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_PWRITE_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_QUOTACTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_QUOTACTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_READ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_READ", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_READLINK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_READLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_READV)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_READV", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_READV_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_READV_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_READ_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_READ_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_REBOOT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_REBOOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_RECVFROM)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RECVFROM", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_RECVFROM_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RECVFROM_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_RECVMSG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RECVMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_RECVMSG_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RECVMSG_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_REMOVEXATTR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_REMOVEXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_RENAME)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RENAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_REVOKE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_REVOKE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_RMDIR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_RMDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SEARCHFS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEARCHFS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SELECT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SELECT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SELECT_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SELECT_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SEMCTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEMCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SEMGET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEMGET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SEMOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEMOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SEMSYS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEMSYS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SEM_CLOSE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEM_CLOSE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SEM_DESTROY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEM_DESTROY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SEM_GETVALUE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEM_GETVALUE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SEM_INIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEM_INIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SEM_OPEN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEM_OPEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SEM_POST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEM_POST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SEM_TRYWAIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEM_TRYWAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SEM_UNLINK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEM_UNLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SEM_WAIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEM_WAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SEM_WAIT_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SEM_WAIT_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SENDFILE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SENDFILE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SENDMSG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SENDMSG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SENDMSG_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SENDMSG_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SENDTO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SENDTO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SENDTO_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SENDTO_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETATTRLIST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETATTRLIST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETAUDIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETAUDIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETAUDIT_ADDR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETAUDIT_ADDR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETAUID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETAUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETEGID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETEGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETEUID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETEUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETGID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETGROUPS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETGROUPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETITIMER)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETITIMER", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETLCID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETLCID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETLOGIN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETLOGIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETPGID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETPGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETPRIORITY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETPRIORITY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETPRIVEXEC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETPRIVEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETREGID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETREGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETREUID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETREUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETRLIMIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETRLIMIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETSGROUPS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETSGROUPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETSID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETSID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETSOCKOPT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETSOCKOPT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETTID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETTID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETTID_WITH_PID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETTID_WITH_PID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETTIMEOFDAY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETTIMEOFDAY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETUID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETWGROUPS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETWGROUPS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SETXATTR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SETXATTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SHARED_REGION_CHECK_NP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SHARED_REGION_CHECK_NP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SHARED_REGION_MAP_AND_SLIDE_NP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SHARED_REGION_MAP_AND_SLIDE_NP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SHMAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SHMAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SHMCTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SHMCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SHMDT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SHMDT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SHMGET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SHMGET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SHMSYS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SHMSYS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SHM_OPEN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SHM_OPEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SHM_UNLINK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SHM_UNLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SHUTDOWN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SHUTDOWN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SIGACTION)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SIGACTION", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SIGALTSTACK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SIGALTSTACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SIGPENDING)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SIGPENDING", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SIGPROCMASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SIGPROCMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SIGRETURN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SIGRETURN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SIGSUSPEND)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SIGSUSPEND", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SIGSUSPEND_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SIGSUSPEND_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SOCKET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SOCKET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SOCKETPAIR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SOCKETPAIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_STACK_SNAPSHOT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_STACK_SNAPSHOT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_STAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_STAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_STAT64)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_STAT64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_STAT64_EXTENDED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_STAT64_EXTENDED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_STATFS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_STATFS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_STATFS64)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_STATFS64", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_STATV)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_STATV", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_STAT_EXTENDED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_STAT_EXTENDED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SWAPON)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SWAPON", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SYMLINK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SYMLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SYNC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SYNC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_SYSCALL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_SYSCALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_THREAD_SELFID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_THREAD_SELFID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_TRUNCATE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_TRUNCATE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_UMASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_UMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_UMASK_EXTENDED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_UMASK_EXTENDED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_UNDELETE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_UNDELETE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_UNLINK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_UNLINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_UNMOUNT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_UNMOUNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_UTIMES)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_UTIMES", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_VFORK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_VFORK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_VM_PRESSURE_MONITOR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_VM_PRESSURE_MONITOR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_WAIT4)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_WAIT4", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_WAIT4_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_WAIT4_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_WAITEVENT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_WAITEVENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_WAITID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_WAITID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_WAITID_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_WAITID_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_WATCHEVENT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_WATCHEVENT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_WORKQ_KERNRETURN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_WORKQ_KERNRETURN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_WORKQ_OPEN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_WORKQ_OPEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_WRITE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_WRITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_WRITEV)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_WRITEV", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_WRITEV_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_WRITEV_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS_WRITE_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS_WRITE_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___DISABLE_THREADSIGNAL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___DISABLE_THREADSIGNAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_EXECVE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_EXECVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_GETFSSTAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_GETFSSTAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_GET_FD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_GET_FD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_GET_FILE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_GET_FILE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_GET_LCID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_GET_LCID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_GET_LCTX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_GET_LCTX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_GET_LINK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_GET_LINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_GET_MOUNT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_GET_MOUNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_GET_PID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_GET_PID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_GET_PROC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_GET_PROC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_MOUNT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_MOUNT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_SET_FD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_SET_FD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_SET_FILE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_SET_FILE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_SET_LCTX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_SET_LCTX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_SET_LINK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_SET_LINK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_SET_PROC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_SET_PROC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___MAC_SYSCALL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___MAC_SYSCALL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___OLD_SEMWAIT_SIGNAL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___OLD_SEMWAIT_SIGNAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___OLD_SEMWAIT_SIGNAL_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___OLD_SEMWAIT_SIGNAL_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___PTHREAD_CANCELED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___PTHREAD_CANCELED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___PTHREAD_CHDIR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___PTHREAD_CHDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___PTHREAD_FCHDIR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___PTHREAD_FCHDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___PTHREAD_KILL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___PTHREAD_KILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___PTHREAD_MARKCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___PTHREAD_MARKCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___PTHREAD_SIGMASK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___PTHREAD_SIGMASK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___SEMWAIT_SIGNAL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___SEMWAIT_SIGNAL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___SEMWAIT_SIGNAL_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___SEMWAIT_SIGNAL_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___SIGWAIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___SIGWAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___SIGWAIT_NOCANCEL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___SIGWAIT_NOCANCEL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SYS___SYSCTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SYS___SYSCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IEXEC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IEXEC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IFBLK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFBLK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IFCHR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFCHR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IFDIR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFDIR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IFIFO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFIFO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IFLNK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFLNK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IFMT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFMT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IFREG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFREG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IFSOCK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFSOCK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IFWHT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IFWHT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IREAD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IREAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IRGRP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IRGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IROTH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IROTH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IRUSR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IRUSR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IRWXG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IRWXG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IRWXO)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IRWXO", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IRWXU)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IRWXU", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_ISGID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_ISGID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_ISTXT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_ISTXT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_ISUID)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_ISUID", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_ISVTX)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_ISVTX", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IWGRP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IWGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IWOTH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IWOTH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IWRITE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IWRITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IWUSR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IWUSR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IXGRP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IXGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IXOTH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IXOTH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.S_IXUSR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "S_IXUSR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Seek)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Seek", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Select)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Select", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sendfile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sendfile", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sendmsg)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sendmsg", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SendmsgN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SendmsgN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sendto)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sendto", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetBpf)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetBpf", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetBpfBuflen)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetBpfBuflen", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetBpfDatalink)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetBpfDatalink", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetBpfHeadercmpl)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetBpfHeadercmpl", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetBpfImmediate)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetBpfImmediate", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetBpfInterface)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetBpfInterface", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetBpfPromisc)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetBpfPromisc", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetBpfTimeout)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetBpfTimeout", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetKevent)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetKevent", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetNonblock)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetNonblock", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Setegid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setegid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Setenv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setenv", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Seteuid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Seteuid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Setgid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setgid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Setgroups)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setgroups", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Setlogin)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setlogin", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Setpgid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setpgid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Setpriority)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setpriority", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Setprivexec)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setprivexec", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Setregid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setregid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Setreuid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setreuid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Setrlimit)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setrlimit", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Setsid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setsid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetsockoptByte)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptByte", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetsockoptICMPv6Filter)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptICMPv6Filter", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetsockoptIPMreq)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptIPMreq", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetsockoptIPv6Mreq)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptIPv6Mreq", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetsockoptInet4Addr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptInet4Addr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetsockoptInt)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptInt", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetsockoptLinger)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptLinger", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetsockoptString)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptString", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SetsockoptTimeval)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SetsockoptTimeval", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Settimeofday)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Settimeofday", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Setuid)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Setuid", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Shutdown)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Shutdown", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Signal
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Signal", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofBpfHdr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofBpfHdr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofBpfInsn)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofBpfInsn", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofBpfProgram)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofBpfProgram", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofBpfStat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofBpfStat", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofBpfVersion)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofBpfVersion", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofCmsghdr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofCmsghdr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofICMPv6Filter)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofICMPv6Filter", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofIPMreq)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofIPMreq", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofIPv6MTUInfo)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofIPv6MTUInfo", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofIPv6Mreq)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofIPv6Mreq", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofIfData)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofIfData", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofIfMsghdr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofIfMsghdr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofIfaMsghdr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofIfaMsghdr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofIfmaMsghdr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofIfmaMsghdr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofIfmaMsghdr2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofIfmaMsghdr2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofInet4Pktinfo)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofInet4Pktinfo", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofInet6Pktinfo)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofInet6Pktinfo", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofLinger)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofLinger", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofMsghdr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofMsghdr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofRtMetrics)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofRtMetrics", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofRtMsghdr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofRtMsghdr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofSockaddrAny)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofSockaddrAny", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofSockaddrDatalink)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofSockaddrDatalink", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofSockaddrInet4)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofSockaddrInet4", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofSockaddrInet6)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofSockaddrInet6", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SizeofSockaddrUnix)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SizeofSockaddrUnix", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SlicePtrFromStrings)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SlicePtrFromStrings", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.SockaddrDatalink
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SockaddrDatalink", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.SockaddrInet4
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SockaddrInet4", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.SockaddrInet6
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SockaddrInet6", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.SockaddrUnix
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SockaddrUnix", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Socket)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Socket", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.SocketControlMessage
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SocketControlMessage", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SocketDisableIPv6)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SocketDisableIPv6", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Socketpair)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Socketpair", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.StartProcess)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StartProcess", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Stat)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Stat", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Stat_t
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Stat_t", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Statfs)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Statfs", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Statfs_t
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Statfs_t", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Stderr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Stderr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Stdin)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Stdin", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Stdout)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Stdout", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.StringBytePtr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StringBytePtr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.StringByteSlice)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StringByteSlice", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.StringSlicePtr)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "StringSlicePtr", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Symlink)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Symlink", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sync)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sync", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.SysProcAttr
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "SysProcAttr", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Syscall)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Syscall", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Syscall6)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Syscall6", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Syscall9)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Syscall9", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Sysctl)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Sysctl", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.SysctlUint32)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SysctlUint32", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCIFLUSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCIFLUSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCIOFLUSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCIOFLUSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCOFLUSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCOFLUSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCP_CONNECTIONTIMEOUT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_CONNECTIONTIMEOUT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCP_KEEPALIVE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_KEEPALIVE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCP_MAXHLEN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_MAXHLEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCP_MAXOLEN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_MAXOLEN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCP_MAXSEG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_MAXSEG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCP_MAXWIN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_MAXWIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCP_MAX_SACK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_MAX_SACK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCP_MAX_WINSHIFT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_MAX_WINSHIFT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCP_MINMSS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_MINMSS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCP_MINMSSOVERLOAD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_MINMSSOVERLOAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCP_MSS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_MSS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCP_NODELAY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_NODELAY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCP_NOOPT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_NOOPT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCP_NOPUSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_NOPUSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCP_RXT_CONNDROPTIME)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_RXT_CONNDROPTIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCP_RXT_FINDROP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCP_RXT_FINDROP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TCSAFLUSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TCSAFLUSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCCBRK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCCBRK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCCDTR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCCDTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCCONS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCCONS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCDCDTIMESTAMP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCDCDTIMESTAMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCDRAIN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCDRAIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCDSIMICROCODE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCDSIMICROCODE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCEXCL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCEXCL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCEXT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCEXT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCFLUSH)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCFLUSH", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCGDRAINWAIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCGDRAINWAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCGETA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCGETA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCGETD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCGETD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCGPGRP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCGPGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCGWINSZ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCGWINSZ", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCIXOFF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCIXOFF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCIXON)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCIXON", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCMBIC)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCMBIC", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCMBIS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCMBIS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCMGDTRWAIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCMGDTRWAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCMGET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCMGET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCMODG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCMODG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCMODS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCMODS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCMSDTRWAIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCMSDTRWAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCMSET)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCMSET", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCM_CAR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_CAR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCM_CD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_CD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCM_CTS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_CTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCM_DSR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_DSR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCM_DTR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_DTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCM_LE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_LE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCM_RI)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_RI", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCM_RNG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_RNG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCM_RTS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_RTS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCM_SR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_SR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCM_ST)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCM_ST", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCNOTTY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCNOTTY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCNXCL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCNXCL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCOUTQ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCOUTQ", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCPKT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCPKT_DATA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT_DATA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCPKT_DOSTOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT_DOSTOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCPKT_FLUSHREAD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT_FLUSHREAD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCPKT_FLUSHWRITE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT_FLUSHWRITE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCPKT_IOCTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT_IOCTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCPKT_NOSTOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT_NOSTOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCPKT_START)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT_START", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCPKT_STOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPKT_STOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCPTYGNAME)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPTYGNAME", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCPTYGRANT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPTYGRANT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCPTYUNLK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCPTYUNLK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCREMOTE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCREMOTE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCSBRK)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSBRK", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCSCONS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSCONS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCSCTTY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSCTTY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCSDRAINWAIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSDRAINWAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCSDTR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSDTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCSETA)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSETA", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCSETAF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSETAF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCSETAW)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSETAW", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCSETD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSETD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCSIG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSIG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCSPGRP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSPGRP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCSTART)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSTART", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCSTAT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSTAT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCSTI)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSTI", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCSTOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSTOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCSWINSZ)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCSWINSZ", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCTIMESTAMP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCTIMESTAMP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TIOCUCNTL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TIOCUCNTL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TOSTOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TOSTOP", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Termios
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Termios", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Timespec
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Timespec", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TimespecToNsec)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TimespecToNsec", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Timeval
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Timeval", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Timeval32
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Timeval32", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.TimevalToNsec)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TimevalToNsec", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Truncate)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Truncate", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Umask)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Umask", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Undelete)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Undelete", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.UnixRights)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UnixRights", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Unlink)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Unlink", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Unmount)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Unmount", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Unsetenv)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Unsetenv", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Utimes)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Utimes", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.UtimesNano)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "UtimesNano", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VDISCARD)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VDISCARD", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VDSUSP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VDSUSP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VEOF)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VEOF", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VEOL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VEOL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VEOL2)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VEOL2", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VERASE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VERASE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VINTR)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VINTR", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VKILL)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VKILL", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VLNEXT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VLNEXT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VMIN)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VMIN", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VQUIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VQUIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VREPRINT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VREPRINT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VSTART)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VSTART", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VSTATUS)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VSTATUS", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VSTOP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VSTOP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VSUSP)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VSUSP", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VT0)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VT0", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VT1)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VT1", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VTDLY)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VTDLY", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VTIME)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VTIME", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.VWERASE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VWERASE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.WCONTINUED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WCONTINUED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.WCOREFLAG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WCOREFLAG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.WEXITED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WEXITED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.WNOHANG)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WNOHANG", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.WNOWAIT)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WNOWAIT", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.WORDSIZE)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WORDSIZE", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.WSTOPPED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WSTOPPED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.WUNTRACED)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WUNTRACED", o); raised != nil {
		return nil, raised
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Wait4)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Wait4", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.WaitStatus
		if o, raised := grumpy.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "WaitStatus", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := grumpy.WrapNative(f, reflect.ValueOf(mod.Write)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Write", o); raised != nil {
		return nil, raised
	}

	return nil, nil
}
var Code = grumpy.NewCode("<module>", "syscall", nil, 0, fun)
func init() {
	grumpy.RegisterModule("__go__/syscall", Code)
}
