package util

import (
	"github.com/cheggaaa/pb"
	"github.com/inancgumus/screen"
)

type PBar struct {
	pool              *pb.Pool
	progress_bar_list []*pb.ProgressBar
}

func (pPool *PBar) reset() {
	pPool.pool.Stop()
	screen.Clear()
	pPool.pool = pb.NewPool()
	for _, n_value := range pPool.progress_bar_list {
		pPool.pool.Add(n_value)
	}
	pPool.pool.Start()
}

func NewPBar(pbar *pb.ProgressBar) *PBar {
	var pp PBar
	pp.pool = pb.NewPool()
	pp.pool.Add(pbar)
	pp.progress_bar_list = append(pp.progress_bar_list, pbar)
	screen.Clear()
	pp.pool.Start()
	return &pp
}

func (pPool *PBar) Add(pBar *pb.ProgressBar) {
	var new_array []*pb.ProgressBar
	new_array = append(new_array, pBar)
	for _, n_value := range pPool.progress_bar_list {
		new_array = append(new_array, n_value)
	}
	pPool.progress_bar_list = new_array
	pPool.reset()
}

func (pPool *PBar) UpdateFinished() {
	var new_progress_bar []*pb.ProgressBar

	for _, n_value := range pPool.progress_bar_list {
		if n_value.IsFinished() == false {
			new_progress_bar = append(new_progress_bar, n_value)
		}
	}
	pPool.progress_bar_list = new_progress_bar
	pPool.reset()
}

func (pPool *PBar) End() {
	pPool.pool.Stop()
}
