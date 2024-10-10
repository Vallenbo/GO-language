package engine

import (
	"sync"
	"time"
)

// LockResourceManager 是全局锁资源管理器结构体
type LockResourceManager struct {
	locks           map[string]*Locker // locks 存储不同资源的互斥锁
	mu              sync.Mutex         // 添加互斥锁
	deadlockChecker *DeadlockChecker   // deadlockChecker 用于死锁检测
}

type Locker struct {
	resourceName   string
	isLocked       bool
	lock           *sync.Mutex
	timeout        time.Duration
	timeoutHandler func()
}

// DeadlockChecker 用于检测死锁的结构体
type DeadlockChecker struct {
	lockOrder     []string      // lockOrder 记录锁的获取顺序
	checkInterval time.Duration // checkInterval 死锁检测的时间间隔
	mu            sync.Mutex
}

// NewLockResourceManager 创建一个新的全局锁资源管理器实例
func NewLockResourceManager(deadlocktime time.Duration) *LockResourceManager {
	return &LockResourceManager{
		locks:           make(map[string]*Locker),         // 初始化锁资源池
		deadlockChecker: NewDeadlockChecker(deadlocktime), // 创建死锁检测器并传入检测时间间隔
	}
}

func NewLocker(resourceName string, duration time.Duration, fencer func()) *Locker {
	return &Locker{
		resourceName:   resourceName,
		isLocked:       false,
		lock:           &sync.Mutex{},
		timeout:        duration,
		timeoutHandler: fencer,
	}
}

// AcquireLock 获取指定资源的锁
func (m *LockResourceManager) AcquireLock(resourceName string, duration time.Duration, fencer func()) {
	m.mu.Lock() // 添加锁以保护对 locks 的并发访问
	defer m.mu.Unlock()

	if _, ok := m.locks[resourceName]; !ok { // 如果资源对应的锁不存在，则创建一个新的互斥锁
		m.locks[resourceName] = NewLocker(resourceName, duration, fencer) // 使用 NewLocker 创建新的锁
	}
	m.locks[resourceName].lock.Lock()                     // 锁定资源
	m.deadlockChecker.recordLockAcquisition(resourceName) // 记录锁的获取顺序到死锁检测器
}

// ReleaseLock 释放指定资源的锁
func (m *LockResourceManager) ReleaseLock(resource string) {
	m.mu.Lock() // 添加锁以保护对 locks 的并发访问
	defer m.mu.Unlock()

	if locker, ok := m.locks[resource]; ok { // 如果资源存在对应的锁
		locker.lock.Unlock()
		m.deadlockChecker.recordLockRelease(resource) // 记录锁的释放顺序到死锁检测器
	}
}

// NewDeadlockChecker 创建一个新的死锁检测器实例
func NewDeadlockChecker(checkInterval time.Duration) *DeadlockChecker {
	return &DeadlockChecker{
		checkInterval: checkInterval,
	}
}

// recordLockAcquisition 记录锁的获取顺序
func (d *DeadlockChecker) recordLockAcquisition(resource string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.lockOrder = append(d.lockOrder, resource) // 将资源名称添加到锁顺序列表中
}

// recordLockRelease 记录锁的释放顺序
func (d *DeadlockChecker) recordLockRelease(resource string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	for i, r := range d.lockOrder {
		if r == resource {
			d.lockOrder = append(d.lockOrder[:i], d.lockOrder[i+1:]...) // 从锁顺序列表中删除相应的资源名称
			break
		}
	}
}

// StartChecking 启动死锁检测
func (d *DeadlockChecker) StartChecking() {
	go func() { // 使用 goroutine 启动死锁检测
		for {
			time.Sleep(d.checkInterval) // 按照指定时间间隔进行死锁检测
			d.checkForDeadlock()
		}
	}()
}

// checkForDeadlock 检查是否存在死锁
func (d *DeadlockChecker) checkForDeadlock() {
	d.mu.Lock()
	defer d.mu.Unlock()
	if len(d.lockOrder) > 1 {
		// 这里可以进行更复杂的死锁检测逻辑，目前只是一个占位符
		println("Potential deadlock detected!")
	}
}
