package snowflake

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/andys920605/meme-coin/pkg/logging"
)

const (
	epoch             = int64(1672531200000) // 設置起始時間(時間戳/毫秒)：2023-01-01 00:00:00，有效期17年
	timestampBits     = uint(39)             // 時間戳占用位數
	datacenteridBits  = uint(8)              // 數據中心id所占位數
	workeridBits      = uint(8)              // 機器id所占位數
	sequenceBits      = uint(8)              // 序列所占的位數
	timestampMax      = int64(-1 ^ (-1 << timestampBits))
	datacenteridMax   = int64(-1 ^ (-1 << datacenteridBits))
	workeridMax       = int64(-1 ^ (-1 << workeridBits))
	sequenceMask      = int64(-1 ^ (-1 << sequenceBits))
	workeridShift     = sequenceBits
	datacenteridShift = sequenceBits + workeridBits
	timestampShift    = sequenceBits + workeridBits + datacenteridBits
)

type Snowflake struct {
	sync.Mutex
	timestamp    int64
	workerId     int64
	datacenterId int64
	sequence     int64
	logger       *logging.Logging
}

var (
	instance *Snowflake
	once     sync.Once
)

func Init(logger *logging.Logging) {
	dataCenterId, workerId := getDeviceID(logger)
	once.Do(func() {
		var err error
		instance, err = newSnowflake(dataCenterId, workerId, logger)
		if err != nil {
			logger.Emergencyf("failed to initialize snowflake: %v", err)
		}
	})
}

func New() int64 {
	if instance == nil {
		panic("Snowflake not initialized. Please call Init first.")
	}
	return instance.NextVal()
}

func newSnowflake(datacenterId, workerId int64, logger *logging.Logging) (*Snowflake, error) {
	if datacenterId < 0 || datacenterId > datacenteridMax {
		return nil, fmt.Errorf("datacenterId must be between 0 and %d", datacenteridMax-1)
	}
	if workerId < 0 || workerId > workeridMax {
		return nil, fmt.Errorf("workerId must be between 0 and %d", workeridMax-1)
	}
	return &Snowflake{
		timestamp:    0,
		datacenterId: datacenterId,
		workerId:     workerId,
		sequence:     0,
		logger:       logger,
	}, nil
}

func (s *Snowflake) NextVal() int64 {
	s.Lock()
	defer s.Unlock()
	now := time.Now().UnixNano() / 1000000 // 轉毫秒
	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.sequence = 0
	}
	t := now - epoch
	if t > timestampMax {
		s.logger.Errorf("timestamp exceeds maximum, current: %d", now)
		return 0
	}
	s.timestamp = now
	r := (t << timestampShift) | (s.datacenterId << datacenteridShift) | (s.workerId << workeridShift) | s.sequence
	return r
}

func getDeviceID(logger *logging.Logging) (datacenterid, workerid int64) {
	ip := getExternalIP(logger)
	datacenterid = int64(ip[2]) & datacenteridMax
	workerid = int64(ip[3]) & workeridMax
	return
}

func getExternalIP(logger *logging.Logging) net.IP {
	ifaces, err := net.Interfaces()
	if err != nil {
		logger.Emergencyf("failed to initial snowflake getExternalIP: %v", err)
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			logger.Emergencyf("failed to initial snowflake getExternalIP: %v", err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip
		}
	}
	logger.Emergency("failed to initial snowflake: no ipv4 address")
	return net.IP{}
}
