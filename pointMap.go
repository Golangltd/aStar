package main

import (
	"math/rand"
	"time"
)

type PointMap struct {
	Size          int // 大小
	Obstacle      int //障碍物
	ObstaclePoint []*Point
}

func NewPointMap(size int) *PointMap {
	//设置障碍物的数量为地图大小除以8
	randomMap := &PointMap{Size: size, Obstacle: size / 8}
	//调用GenerateObstacle生成随机障碍物；
	randomMap.GenerateObstacle()
	return randomMap
}

func (pm *PointMap) GenerateObstacle() {
	pm.ObstaclePoint = append(pm.ObstaclePoint, NewPoint(pm.Size/2, pm.Size/2))
	pm.ObstaclePoint = append(pm.ObstaclePoint, NewPoint(pm.Size/2, pm.Size/2-1))
	//在地图的中间生成一个斜着的障碍物；
	x := make([]int, 4)
	for i := range x {
		v := pm.Size/2 - 4 + i
		pm.ObstaclePoint = append(pm.ObstaclePoint, NewPoint(v, pm.Size-v))
		pm.ObstaclePoint = append(pm.ObstaclePoint, NewPoint(v, pm.Size-v-1))
		pm.ObstaclePoint = append(pm.ObstaclePoint, NewPoint(pm.Size-v, v))
		pm.ObstaclePoint = append(pm.ObstaclePoint, NewPoint(pm.Size-v, v-1))
	}
	//随机生成其他几个障碍物；
	x = make([]int, pm.Obstacle-1)
	for range x {
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(pm.Size)
		y := rand.Intn(pm.Size)
		pm.ObstaclePoint = append(pm.ObstaclePoint, NewPoint(x, y))
		//障碍物的方向也是随机的；
		xl := make([]int, pm.Size/4)
		randNum := rand.Intn(100)
		for l := range xl {
			if randNum > 50 {
				pm.ObstaclePoint = append(pm.ObstaclePoint, NewPoint(x, y+l))
			} else {
				pm.ObstaclePoint = append(pm.ObstaclePoint, NewPoint(x+l, y))
			}
		}
	}
}
//定义一个方法来判断某个节点是否是障碍物
func (pm *PointMap) IsObstacle(i, j int) bool {
	for _, p := range pm.ObstaclePoint {
		if i == p.X && j == p.Y {
			return true
		}
	}
	return false
}
