package id

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
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

	mu        sync.Mutex
	nodeMax   int64 = -1 ^ (-1 << nodeBits)
	nodeMask        = nodeMax << stepBits
	stepMask  int64 = -1 ^ (-1 << stepBits)
	timeShift       = nodeBits + stepBits
	nodeShift       = stepBits
)

// snowflake 雪花生成器
type snowflake struct {
	mu    sync.Mutex
	epoch time.Time
	time  int64
	node  int64
	step  int64

	nodeMax   int64
	nodeMask  int64
	stepMask  int64
	timeShift uint8
	nodeShift uint8
}

func newSnowflake(node int) (_snowflake *snowflake, err error) {
	mu.Lock()
	nodeMax = -1 ^ (-1 << nodeBits)
	nodeMask = nodeMax << stepBits
	stepMask = -1 ^ (-1 << stepBits)
	timeShift = nodeBits + stepBits
	nodeShift = stepBits
	mu.Unlock()

	_snowflake = &snowflake{}
	_snowflake.node = int64(node)
	_snowflake.nodeMax = -1 ^ (-1 << nodeBits)
	_snowflake.nodeMask = _snowflake.nodeMax << stepBits
	_snowflake.stepMask = -1 ^ (-1 << stepBits)
	_snowflake.timeShift = nodeBits + stepBits
	_snowflake.nodeShift = stepBits

	if _snowflake.node < 0 || _snowflake.node > _snowflake.nodeMax {
		err = errors.New(fmt.Sprintf("Node必须在0和%d之间", _snowflake.nodeMax))
	}

	now := time.Now()
	// 确保单一时钟可用（减少雪花算法里面因为时间回调问题而引发的编号冲突）
	_snowflake.epoch = now.Add(time.Unix(epoch/1000, (epoch%1000)*1000000).Sub(now))

	return
}

// Next 返回下一个编号
func (s *snowflake) Next() (id Id) {
	s.mu.Lock()

	now := time.Since(s.epoch).Nanoseconds() / 1000000
	if now == s.time {
		s.step = (s.step + 1) & s.stepMask

		if s.step == 0 {
			for now <= s.time {
				now = time.Since(s.epoch).Nanoseconds() / 1000000
			}
		}
	} else {
		s.step = 0
	}
	s.time = now
	id = Id((now)<<s.timeShift | (s.node << s.nodeShift) | (s.step))

	s.mu.Unlock()

	return
}

// NextId 下一个长整形Id
func (s *snowflake) NextId() int64 {
	return s.Next().Int64()
}

// NextString 下一个字符串形式的Id
func (s *snowflake) NextString() string {
	return s.Next().String()
}

// Int64 返回整形形式
func (i Id) Int64() int64 {
	return int64(i)
}

// ParseInt64 从整形获得编号
func ParseInt64(id int64) Id {
	return Id(id)
}

// String 返回编号的字符串形式
func (i Id) String() string {
	return strconv.FormatInt(int64(i), 10)
}

// ParseString 从十进制字符串解析编号
func ParseString(str string) (id Id, err error) {
	var original int64

	if original, err = strconv.ParseInt(str, 10, 64); nil != err {
		return
	}

	id = Id(original)

	return
}

// Base64 返回编号的Base64字符串
func (i Id) Base64() string {
	return base64.StdEncoding.EncodeToString(i.Bytes())
}

// ParseBase64 从Base64字符串中解析编号
func ParseBase64(str string) (id Id, err error) {
	var bytes []byte

	if bytes, err = base64.StdEncoding.DecodeString(str); nil != err {
		return
	}

	id, err = ParseBytes(bytes)

	return
}

// Capacity 返回二进制数组表示形式
func (i Id) Bytes() []byte {
	return []byte(i.String())
}

// ParseBytes 从二进制数组里面获得编号
func ParseBytes(bytes []byte) (id Id, err error) {
	var original int64

	if original, err = strconv.ParseInt(string(bytes), 10, 64); nil != err {
		return
	}

	id = Id(original)

	return
}

// IntBytes 返回8位二进制表示形式
func (i Id) IntBytes() [8]byte {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], uint64(i))

	return b
}

// ParseIntBytes 从二进制数组获得编号
func ParseIntBytes(id [8]byte) Id {
	return Id(int64(binary.BigEndian.Uint64(id[:])))
}

// Time 返回时间
func (i Id) Time() int64 {
	return (int64(i) >> timeShift) + epoch
}

// Node 返回节点数
func (i Id) Node() int64 {
	return int64(i) & nodeMask >> nodeShift
}

// Step 返回当前自增数量
func (i Id) Step() int64 {
	return int64(i) & stepMask
}

func (i Id) MarshalJSON() ([]byte, error) {
	buff := make([]byte, 0, 22)
	buff = append(buff, '"')
	buff = strconv.AppendInt(buff, int64(i), 10)
	buff = append(buff, '"')

	return buff, nil
}

func (i *Id) UnmarshalJSON(data []byte) (err error) {
	var original int64

	if original, err = strconv.ParseInt(string(data[1:len(data)-1]), 10, 64); nil != err {
		return
	}

	*i = Id(original)

	return
}
