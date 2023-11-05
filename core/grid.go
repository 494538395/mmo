package core

import (
	"fmt"
	"sync"
)

type Grid struct {
	GID       int          // 格子 ID
	MinX      int          // 格子左边界
	MinY      int          // 格子上边界
	MaxX      int          // 格子右边界
	MaxY      int          // 格子下边界
	playerIDs map[int]bool // 格子内玩家集合
	lock      sync.RWMutex //
}

func NewGrid(gid int, minX, minY, maxX, maxY int) *Grid {
	return &Grid{
		GID:       gid,
		MinX:      minX,
		MinY:      minY,
		MaxX:      maxX,
		MaxY:      maxY,
		playerIDs: make(map[int]bool),
	}
}

// AddPlayer 添加玩家到网格中
func (g *Grid) AddPlayer(pid int) {
	g.lock.Lock()
	defer g.lock.Unlock()

	if _, ok := g.playerIDs[g.GID]; !ok {
		g.playerIDs[g.GID] = true
	}
}

// RemovePayer 删除格子中的玩家
func (g *Grid) RemovePayer(pid int) {
	g.lock.Lock()
	defer g.lock.Unlock()

	delete(g.playerIDs, pid)
}

// GetPlayerIDs 获取格子中的所有玩家ID
func (g *Grid) GetPlayerIDs() (pids []int) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	for pid, _ := range g.playerIDs {
		pids = append(pids, pid)
	}

	return
}

func (g *Grid) String() string {
	return fmt.Sprintf("Grid GID: %d, MinX: %d, MaxX: %d, MinY: %d, MaxY:%d playerIDs: %v,",
		g.GID, g.MinX, g.MaxX, g.MinY, g.MaxY, g.playerIDs)
}
