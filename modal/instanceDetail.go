package modal

type InstanceDetail struct {
	InstanceName string
	PublicIp string
}

func NewInstanceDetail(name, publicId string) InstanceDetail {
	instDetail:= &InstanceDetail{}
	instDetail.InstanceName = name
	instDetail.PublicIp = publicId;
	return *instDetail;
}

func (id InstanceDetail) IsValid() bool {
	if id.PublicIp != "" {
		return true
	}
	return false
}

func (id InstanceDetail) ToString() string {
	return "{'InstanceName': "+id.InstanceName+", 'PublicIp': "+id.PublicIp+"}"
}