package internal

import (
	"github.com/goexl/gox/field"
	"github.com/goexl/id/internal/core"
	"github.com/goexl/log"

	"sync"
	"time"
)

var (
	// epoch 设置为2010年11月4日UTC，Twitter雪花算法开始时间（以毫秒为单位）
	// 可以对其进行自定义，以为您的应用程序设置另一个开始时间
	epoch int64 = 1288834974657
	// nodeBits 表示节点数量所占用的位数
	// 10位的话，就表示总共有2^10（1024个节点数量）个节点数可使用
	nodeBits uint8 = 10
	// stepBits 表示生成编号时，可以自增的位数（大小）
	// 12位的话，就表示总共有2^12（4096个自增数量）个自增数可使用
	stepBits uint8 = 12
)

type Snowflake struct {
	mutex sync.Mutex
	epoch time.Time
	time  int64
	node  uint16
	step  int64

	nodeMax   int16
	nodeMask  uint16
	stepMask  int64
	timeShift uint8
	nodeShift uint8

	logger log.Logger
}

func NewSnowflake(node uint16, logger log.Logger) (snowflake *Snowflake) {
	snowflake = new(Snowflake)
	snowflake.node = node
	snowflake.nodeMax = -1 ^ (-1 << nodeBits)
	snowflake.nodeMask = uint16(snowflake.nodeMax << stepBits)
	snowflake.stepMask = -1 ^ (-1 << stepBits)
	snowflake.timeShift = nodeBits + stepBits
	snowflake.nodeShift = stepBits
	snowflake.logger = logger

	if snowflake.node < 0 || snowflake.node > uint16(snowflake.nodeMax) {
		snowflake.node = 1
		logger.Error("节点编号出错", field.New("node.old", node), field.New("node.new", 1))
	}

	now := time.Now()
	// 确保单一时钟可用（减少雪花算法里面因为时间回调问题而引发的编号冲突）
	snowflake.epoch = now.Add(time.Unix(epoch/1000, (epoch%1000)*1000000).Sub(now))

	return
}

func (s *Snowflake) Next() (id core.Id) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	now := time.Since(s.epoch).Nanoseconds() / 1000000
	if now == s.time {
		s.step = (s.step + 1) & s.stepMask

		if 0 == s.step {
			for now <= s.time {
				now = time.Since(s.epoch).Nanoseconds() / 1000000
			}
		}
	} else {
		s.step = 0
	}
	s.time = now
	id = core.Id(now<<s.timeShift | int64(s.node<<s.nodeShift) | s.step)

	return
}
