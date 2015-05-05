package fs

import (
	"github.com/docker/libcontainer/cgroups"
	"github.com/docker/libcontainer/configs"
	"github.com/docker/libcontainer/utils"	

)

type NetClsGroup struct {
}

func (s *NetClsGroup) Apply(d *data) error {
	dir, err := d.join("net_cls")
	if err != nil {
		if cgroups.IsNotFound(err) {
			return nil
		} else {
			return err
		}
	}

	if err := s.Set(dir, d.c); err != nil {
		return err
	}

	return nil
}

func (s *NetClsGroup) Set(path string, cgroup *configs.Cgroup) error {
  classId := utils.GetNetClsClassId()	
  if err := writeFile(path, "net_cls.classid", classId); err != nil {
	  	return err
	}
	return nil
}

func (s *NetClsGroup) Remove(d *data) error {
	return removePath(d.path("devices"))
}

func (s *NetClsGroup) GetStats(path string, stats *cgroups.Stats) error {
	return nil
}
